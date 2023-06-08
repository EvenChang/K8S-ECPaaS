/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com.tw
*/

package v1

import (
	"context"

	"github.com/emicklei/go-restful"
	"github.com/minio/minio-go/v7"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	v1alpha1 "kubesphere.io/api/virtualization/v1alpha1"
	"kubesphere.io/kubesphere/pkg/api"
	kubesphere "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	"kubesphere.io/kubesphere/pkg/server/errors"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

var bucketName = "ecpaas-images"

type handler struct {
	minioClient *minio.Client
	ksclient    kubesphere.Interface
}

func newHandler(minioClient *minio.Client, ksclient kubesphere.Interface) *handler {
	return &handler{
		minioClient: minioClient,
		ksclient:    ksclient,
	}
}

type ObjectStatus struct {
	Status  int `json:"status"`
	Data    ObjectStatusData
	Message string `json:"message"`
}

type ObjectStatusData struct {
	// If object exisit or not
	FileHas bool `json:"fileHas"`
	// Object name
	Name         string `json:"name,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	Size         int64  `json:"size,omitempty"`
}

type ImagesList struct {
	Image []string `json:"images"`
}

func (h *handler) ListMinioObjects(request *restful.Request, response *restful.Response) {

	images := ImagesList{}

	objectCh := h.minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})
	for object := range objectCh {
		if object.Err != nil {
			api.HandleInternalError(response, request, object.Err)
			return
		}
		images.Image = append(images.Image, object.Key)
	}

	response.WriteAsJson(images)
}

func (h *handler) GetMinioObjectStatus(request *restful.Request, response *restful.Response) {

	imageName := request.PathParameter("imageName")
	status := ObjectStatus{}

	objInfo, err := h.minioClient.StatObject(context.Background(), bucketName, imageName, minio.StatObjectOptions{})
	if err != nil {
		status.Status = 400
		status.Data = ObjectStatusData{
			FileHas: false,
		}
		status.Message = err.Error()
	} else {
		status.Status = 200
		status.Data = ObjectStatusData{
			FileHas:      true,
			Name:         imageName,
			LastModified: objInfo.LastModified.String(),
			Size:         objInfo.Size,
		}
		status.Message = "success"
	}

	response.WriteAsJson(status)
}

func (h *handler) UploadMinioObject(request *restful.Request, response *restful.Response) {

	// Check minio bucket "ecpaas-images" if not exist then create it.
	found, err := h.minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		api.HandleInternalError(response, request, err)
		return
	}

	if !found {
		err = h.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			api.HandleInternalError(response, request, err)
			return
		}
	}

	request.Request.ParseMultipartForm(500 << 20)
	file, header, err := request.Request.FormFile("uploadfile")
	if err != nil {
		api.HandleInternalError(response, request, err)
		return
	}
	filesize := file.(Sizer).Size()

	name := request.Request.FormValue("name")
	namespace := request.Request.FormValue("namespace")
	storage := request.Request.FormValue("storage")

	request.Request.MultipartReader()

	uploadInfo, err := h.minioClient.PutObject(context.Background(), bucketName, header.Filename,
		file, filesize, minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		api.HandleInternalError(response, request, err)
		return
	}

	// After finished upload image process, create the ImageTemplete resource for kubevirt dataVolume use.

	if namespace == "" {
		namespace = "default"
	}

	imageTeplate := &v1alpha1.ImageTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1alpha1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1alpha1.ImageTemplateSpec{
			Resources: v1alpha1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse(storage),
				},
			},
			Source: v1alpha1.ImageTemplateSource{
				HTTP: &v1beta1.DataVolumeSourceHTTP{
					URL: uploadInfo.Location,
				},
			},
		},
	}

	_, err = h.ksclient.VirtualizationV1alpha1().ImageTemplates(namespace).Create(context.Background(), imageTeplate, metav1.CreateOptions{})

	if err != nil {
		klog.Info("Create ImageTemplates resource failed. ", err)
		return
	}

	response.WriteAsJson(uploadInfo)
}

func (h *handler) DeleteMinioObject(request *restful.Request, response *restful.Response) {

	imageName := request.PathParameter("imageName")

	err := h.minioClient.RemoveObject(context.Background(), bucketName, imageName, minio.RemoveObjectOptions{})
	if err != nil {
		api.HandleInternalError(response, request, err)
		return
	}

	response.WriteEntity(errors.None)
}

type Sizer interface {
	Size() int64
}
