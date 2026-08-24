package utils

import (
	"context"
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// UploadQiniu 上传本地文件到七牛云，返回可访问的完整 URL
func UploadQiniu(localPath string, key string) (string, error) {
	q := Conf.Qiniu
	if q.AccessKey == "" || q.SecretKey == "" || q.Bucket == "" {
		return "", fmt.Errorf("七牛云配置不完整(accessKey/secretKey/bucket)")
	}
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket + ":" + key,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{UseHTTPS: true, UseCdnDomains: false}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := uploader.PutFile(context.Background(), &ret, upToken, key, localPath, nil)
	if err != nil {
		return "", fmt.Errorf("上传七牛失败: %w", err)
	}
	if q.Domain == "" {
		return ret.Key, nil
	}
	return q.Domain + "/" + ret.Key, nil
}
