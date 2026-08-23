package qiniu

import (
	"context"
	"mime/multipart"

	"github.com/ghf-go/goapp/cloudstore/base"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuCloudStore struct {
	confing      *base.CloudStoreConf
	mac          *qbox.Mac
	formUploader *storage.FormUploader
}

func NewQiniuCloudStore(cfg *base.CloudStoreConf) *QiniuCloudStore {
	var zone *storage.Zone
	switch cfg.Region {
	case "huadong":
		zone = &storage.ZoneHuadong
	case "huabei":
		zone = &storage.ZoneHuabei
	case "huanan":
		zone = &storage.ZoneHuanan
	case "beimei":
		zone = &storage.ZoneBeimei
	case "xinjiapo":
		zone = &storage.ZoneXinjiapo
	case "huadong-zhejiang-2":
		zone = &storage.ZoneHuadongZheJiang2
	}

	qcfg := storage.Config{
		Zone:          zone, // 根据你的存储区域修改：华东/华北/华南
		UseCdnDomains: false,
	}

	return &QiniuCloudStore{
		confing:      cfg,
		mac:          qbox.NewMac(cfg.AccessKey, cfg.SecretKey),
		formUploader: storage.NewFormUploader(&qcfg),
	}
}
func (cs *QiniuCloudStore) uploadFile(baseDir, localFilePath string) (string, error) {
	fk := base.BuildPath(baseDir, localFilePath)
	putPolicy := storage.PutPolicy{
		Scope: cs.confing.Bucket + ":" + fk,
	}
	upToken := putPolicy.UploadToken(cs.mac)
	ret := storage.PutRet{}
	err := cs.formUploader.PutFile(context.Background(), &ret, upToken, fk, localFilePath, nil)
	if err != nil {
		return "", err
	}
	return cs.confing.Domain + "/" + ret.Key, nil
}
func (cs *QiniuCloudStore) UploadVideoFile(localFilePath string) (string, error) {
	return cs.uploadFile("video", localFilePath)
}

// 上传音频文件
func (cs *QiniuCloudStore) UploadAudioFile(localFilePath string) (string, error) {
	return cs.uploadFile("audio", localFilePath)
}

// 上传图片文件
func (cs *QiniuCloudStore) UploadImageFile(localFilePath string) (string, error) {
	return cs.uploadFile("images", localFilePath)
}

// 上传app文件
func (cs *QiniuCloudStore) UploadAppFile(localFilePath string) (string, error) {
	return cs.uploadFile("app", localFilePath)
}

// 上传文件
func (cs *QiniuCloudStore) UploadFile(localFilePath string) (string, error) {
	return cs.uploadFile("upload", localFilePath)
}

func (cs *QiniuCloudStore) uploadVideoFileHeader(baseDir string, f *multipart.FileHeader) (string, error) {
	fk := base.BuildPath(baseDir, f.Filename)
	putPolicy := storage.PutPolicy{
		Scope: cs.confing.Bucket + ":" + fk,
	}
	upToken := putPolicy.UploadToken(cs.mac)
	ret := storage.PutRet{}
	fr, e := f.Open()
	if e != nil {
		return "", e
	}
	defer fr.Close()
	err := cs.formUploader.Put(context.Background(), &ret, upToken, fk, fr, -1, nil)
	if err != nil {
		return "", err
	}
	return cs.confing.Domain + "/" + ret.Key, nil
}

// 上传表单视频文件
func (cs *QiniuCloudStore) UploadVideoFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadVideoFileHeader("video", f)
}

// 上传表单音频文件
func (cs *QiniuCloudStore) UploadAudioFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadVideoFileHeader("audio", f)
}

// 上传表单图片文件
func (cs *QiniuCloudStore) UploadImageFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadVideoFileHeader("images", f)
}

// 上传表单app文件
func (cs *QiniuCloudStore) UploadAppFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadVideoFileHeader("app", f)
}

// 上传表单文件
func (cs *QiniuCloudStore) UploadFileHeader(f *multipart.FileHeader) (string, error) {
	return cs.uploadVideoFileHeader("upload", f)
}

func (cs *QiniuCloudStore) uploadToken(baseDir string, localFilePath string) (*base.TokenCloudStore, error) {
	fk := base.BuildPath(baseDir, localFilePath)
	putPolicy := storage.PutPolicy{
		Scope:   cs.confing.Bucket + ":" + fk,
		Expires: 3600,
	}
	upToken := putPolicy.UploadToken(cs.mac)
	return &base.TokenCloudStore{
		Header: map[string]string{},
		Host:   cs.confing.Endpoint,
		Body: map[string]string{
			"token": upToken,
			"key":   fk,
		},
		Method: "POST",
		Url:    cs.confing.Domain + fk,
	}, nil
}

// 上传视频文件Token
func (cs *QiniuCloudStore) UploadVideoToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("video", localFilePath)

}

// 上传音频文件Token
func (cs *QiniuCloudStore) UploadAudioToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("audio", localFilePath)
}

// 上传图片文件Token
func (cs *QiniuCloudStore) UploadImageToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("image", localFilePath)
}

// 上传app文件Token
func (cs *QiniuCloudStore) UploadAppToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("app", localFilePath)
}

// 上传文件Token
func (cs *QiniuCloudStore) UploadToken(localFilePath string) (*base.TokenCloudStore, error) {
	return cs.uploadToken("upload", localFilePath)
}
