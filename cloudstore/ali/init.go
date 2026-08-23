package ali

import (
	"mime/multipart"

	"github.com/ghf-go/goapp/cloudstore/base"
)

type AliCloudStore struct {
}

func (cs *AliCloudStore) UploadVideoFile(localFilePath string) (string, error) {}

// 上传音频文件
func (cs *AliCloudStore) UploadAudioFile(localFilePath string) (string, error) {}

// 上传图片文件
func (cs *AliCloudStore) UploadImageFile(localFilePath string) (string, error) {}

// 上传app文件
func (cs *AliCloudStore) UploadAppFile(localFilePath string) (string, error) {}

// 上传文件
func (cs *AliCloudStore) UploadFile(localFilePath string) (string, error) {}

// 上传表单视频文件
func (cs *AliCloudStore) UploadVideoFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单音频文件
func (cs *AliCloudStore) UploadAudioFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单图片文件
func (cs *AliCloudStore) UploadImageFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单app文件
func (cs *AliCloudStore) UploadAppFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传表单文件
func (cs *AliCloudStore) UploadFileHeader(f *multipart.FileHeader) (string, error) {}

// 上传视频文件Token
func (cs *AliCloudStore) UploadVideoToken(localFilePath string) (*base.TokenCloudStore, error) {}

// 上传音频文件Token
func (cs *AliCloudStore) UploadAudioToken(localFilePath string) (*base.TokenCloudStore, error) {}

// 上传图片文件Token
func (cs *AliCloudStore) UploadImageToken(localFilePath string) (*base.TokenCloudStore, error) {}

// 上传app文件Token
func (cs *AliCloudStore) UploadAppToken(localFilePath string) (*base.TokenCloudStore, error) {}

// 上传文件Token
func (cs *AliCloudStore) UploadToken(localFilePath string) (*base.TokenCloudStore, error) {}
