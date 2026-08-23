package cloudstore

import (
	"mime/multipart"

	"github.com/ghf-go/goapp/cloudstore/base"
	"github.com/ghf-go/goapp/cloudstore/qiniu"
)

const (
	TYPE_ALI     = "ali"
	TYPE_QINIU   = "qiniu"
	TYPE_TENCENT = "tencent"
)

type CloudStore interface {
	//上传视频文件
	UploadVideoFile(localFilePath string) (string, error)
	//上传音频文件
	UploadAudioFile(localFilePath string) (string, error)
	//上传图片文件
	UploadImageFile(localFilePath string) (string, error)
	//上传app文件
	UploadAppFile(localFilePath string) (string, error)
	//上传文件
	UploadFile(localFilePath string) (string, error)

	//上传表单视频文件
	UploadVideoFileHeader(f *multipart.FileHeader) (string, error)
	//上传表单音频文件
	UploadAudioFileHeader(f *multipart.FileHeader) (string, error)
	//上传表单图片文件
	UploadImageFileHeader(f *multipart.FileHeader) (string, error)
	//	上传表单app文件
	UploadAppFileHeader(f *multipart.FileHeader) (string, error)
	//	上传表单文件
	UploadFileHeader(f *multipart.FileHeader) (string, error)

	//上传视频文件Token
	UploadVideoToken(localFilePath string) (*base.TokenCloudStore, error)
	//上传音频文件Token
	UploadAudioToken(localFilePath string) (*base.TokenCloudStore, error)
	//上传图片文件Token
	UploadImageToken(localFilePath string) (*base.TokenCloudStore, error)
	//上传app文件Token
	UploadAppToken(localFilePath string) (*base.TokenCloudStore, error)
	//上传文件Token
	UploadToken(localFilePath string) (*base.TokenCloudStore, error)
}

func NewCloudStore(confing *base.CloudStoreConf) CloudStore {
	switch confing.Type {
	case TYPE_ALI:
		// return ali.NewQiniuCloudStore(confing)
	case TYPE_QINIU:
		return qiniu.NewQiniuCloudStore(confing)
		// case TYPE_TENCENT:
		// 	return &tencent.TencentCloudStore{}
	}
	return nil
}
