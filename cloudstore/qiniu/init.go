package qiniu

import (
	"mime/multipart"

	"github.com/ghf-go/goapp/cloudstore/conf"
)

type QiniuCloudStore struct {
}

func (cs *QiniuCloudStore) UploadVideoFile(localFilePath string) (string, error) {}

// 上传音频文件
func (cs *QiniuCloudStore) UploadAudioFile(localFilePath string) (string, error) {}

// 上传图片文件
func (cs *QiniuCloudStore) UploadImageFile(localFilePath string) (string, error) {}

// 上传app文件
func (cs *QiniuCloudStore) UploadAppFile(localFilePath string) (string, error) {}

// 上传文件
func (cs *QiniuCloudStore) UploadFile(localFilePath string) (string, error) {}

// 上传表单视频文件
func (cs *QiniuCloudStore) UploadVideoFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单音频文件
func (cs *QiniuCloudStore) UploadAudioFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单图片文件
func (cs *QiniuCloudStore) UploadImageFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单app文件
func (cs *QiniuCloudStore) UploadAppFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单文件
func (cs *QiniuCloudStore) UploadFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传视频文件Token
func (cs *QiniuCloudStore) UploadVideoToken(localFilePath string) (*conf.TokenCloudStore, error) {}

// 上传音频文件Token
func (cs *QiniuCloudStore) UploadAudioToken(localFilePath string) (*conf.TokenCloudStore, error) {}

// 上传图片文件Token
func (cs *QiniuCloudStore) UploadImageToken(localFilePath string) (*conf.TokenCloudStore, error) {}

// 上传app文件Token
func (cs *QiniuCloudStore) UploadAppToken(localFilePath string) (*conf.TokenCloudStore, error) {}

// 上传文件Token
func (cs *QiniuCloudStore) UploadToken(localFilePath string) (*conf.TokenCloudStore, error) {}
