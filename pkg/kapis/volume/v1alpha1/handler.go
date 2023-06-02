/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com.tw
*/

package v1

import (
	"context"

	"github.com/emicklei/go-restful"
	"github.com/minio/minio-go/v7"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/server/errors"
)

var bucketName = "ecpaas-images"

type handler struct {
	minioClient *minio.Client
}

func newHandler(minioClient *minio.Client) *handler {
	return &handler{
		minioClient: minioClient,
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

	request.Request.MultipartReader()

	uploadInfo, err := h.minioClient.PutObject(context.Background(), bucketName, header.Filename,
		file, filesize, minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		api.HandleInternalError(response, request, err)
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
