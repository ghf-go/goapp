package ali

import (
	"context"
	"mime/multipart"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/ghf-go/goapp/cloudstore/base"
)

type AliCloudStore struct {
	confing *base.CloudStoreConf
	client  *oss.Client
}

func NewAliCloudStore(cfg *base.CloudStoreConf) *AliCloudStore {
	return &AliCloudStore{
		confing: cfg,
		client:  oss.NewClient(oss.LoadDefaultConfig().WithEndpoint(cfg.Endpoint).WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey)).WithRegion(cfg.Region)),
	}
}
func (cs *AliCloudStore) uploadFile(baseDir, localFilePath string) (string, error) {
	fk := base.BuildPath(baseDir, localFilePath)
	putRequest := &oss.PutObjectRequest{
		Bucket: oss.Ptr(string(cs.confing.Bucket)), // 存储空间名称
		Key:    oss.Ptr(fk),                        // 对象名称
		// ProgressFn: pfn,                                // 进度回调函数，用于显示上传进度
	}
	_, err := cs.client.PutObjectFromFile(context.TODO(), putRequest, localFilePath)
	if err != nil {
		return "", err
	}

	return cs.confing.BuildUrl(fk), nil
}
func (cs *AliCloudStore) UploadVideoFile(localFilePath string) (string, error) {
	return cs.uploadFile("video", localFilePath)
}

// 上传音频文件
func (cs *AliCloudStore) UploadAudioFile(localFilePath string) (string, error) {
	return cs.uploadFile("audio", localFilePath)
}

// 上传图片文件
func (cs *AliCloudStore) UploadImageFile(localFilePath string) (string, error) {
	return cs.uploadFile("image", localFilePath)
}

// 上传app文件
func (cs *AliCloudStore) UploadAppFile(localFilePath string) (string, error) {
	return cs.uploadFile("app", localFilePath)
}

// 上传文件
func (cs *AliCloudStore) UploadFile(localFilePath string) (string, error) {
	return cs.uploadFile("upload", localFilePath)
}

func (cs *AliCloudStore) uploadFileHeader(baseDir string, f *multipart.FileHeader) (string, error) {
	fk := base.BuildPath(baseDir, f.Filename)
	fr, e := f.Open()
	if e != nil {
		return "", e
	}
	defer fr.Close()
	putRequest := &oss.PutObjectRequest{
		Bucket: oss.Ptr(string(cs.confing.Bucket)), // 存储空间名称
		Key:    oss.Ptr(fk),                        // 对象名称
		// ProgressFn: pfn,                                // 进度回调函数，用于显示上传进度
		Body: fr,
	}
	_, err := cs.client.PutObject(context.TODO(), putRequest)
	if err != nil {
		return "", err
	}

	return cs.confing.BuildUrl(fk), nil
}

// 上传表单视频文件
func (cs *AliCloudStore) UploadVideoFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadFileHeader("video", f)
}

// 上传表单音频文件
func (cs *AliCloudStore) UploadAudioFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadFileHeader("audio", f)
}

// 上传表单图片文件
func (cs *AliCloudStore) UploadImageFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadFileHeader("image", f)
}

// 上传表单app文件
func (cs *AliCloudStore) UploadAppFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadFileHeader("app", f)
}

// 上传表单文件
func (cs *AliCloudStore) UploadFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadFileHeader("upload", f)
}

// 上传视频文件Token
func (cs *AliCloudStore) UploadVideoToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("video", localFilePath)
}

// 上传音频文件Token
func (cs *AliCloudStore) UploadAudioToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("audio", localFilePath)
}

// 上传图片文件Token
func (cs *AliCloudStore) UploadImageToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("image", localFilePath)
}

// 上传app文件Token
func (cs *AliCloudStore) UploadAppToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("app", localFilePath)
}

// 上传文件Token
func (cs *AliCloudStore) UploadToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("upload", localFilePath)
}
func (cs *AliCloudStore) uploadToken(baseDir string, localFilePath string) (*base.TokenCloudStore, error) {
	return nil, nil
}
