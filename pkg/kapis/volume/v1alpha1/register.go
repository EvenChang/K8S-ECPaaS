/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com.tw
*/

package v1

import (
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/minio/minio-go/v7"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	"kubesphere.io/kubesphere/pkg/constants"
	"kubesphere.io/kubesphere/pkg/server/errors"
)

const (
	GroupName = "virtualization.ecpaas.io"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}

func AddToContainer(container *restful.Container, minioClient *minio.Client) error {
	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler(minioClient)

	webservice.Route(webservice.GET("/upload/file/images").
		To(handler.ListMinioObjects).
		Doc("List all uploaded images").
		Returns(http.StatusOK, api.StatusOK, ImagesList{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.MinioImageTag}))

	webservice.Route(webservice.GET("/upload/file/checkFileExist/{imageName}").
		To(handler.GetMinioObjectStatus).
		Doc("Check If image exist or not").
		Returns(http.StatusOK, api.StatusOK, ObjectStatus{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.MinioImageTag}))

	webservice.Route(webservice.POST("/upload/file/").
		To(handler.UploadMinioObject).
		Doc("Upload Volume Image").
		Returns(http.StatusOK, api.StatusOK, errors.None).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.MinioImageTag}))

	webservice.Route(webservice.DELETE("/upload/file/{imageName}").
		To(handler.DeleteMinioObject).
		Doc("Delete Volume Image").
		Returns(http.StatusOK, api.StatusOK, errors.None).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.MinioImageTag}))

	container.Add(webservice)

	return nil
}
