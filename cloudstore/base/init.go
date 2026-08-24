package base

import (
	"fmt"
	"path"
	"strings"
	"time"
)

type CloudStoreConf struct {
	Type      string `yaml:"type"`
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Endpoint  string `yaml:"endpoint"`
	Domain    string `yaml:"domain"`
}

type TokenCloudStore struct {
	Header map[string]string `json:"header"`
	Host   string            `json:"host"`
	Body   map[string]string `json:"body"`
	Method string            `json:"method"`
	Url    string            `json:"url"`
}

// 生产存储路径
func BuildPath(descDir, localFilePath string) string {
	banem := path.Base(localFilePath)
	ext := path.Ext(banem)
	bnotExtFile := strings.TrimSuffix(banem, ext)
	ct := time.Now()
	return fmt.Sprintf("/%s/%s/%s-%d%s", descDir, ct.Format("2006/01/02"), bnotExtFile, ct.Unix(), ext)
}

func (c CloudStoreConf) BuildUrl(filePath string) string {
	return c.Domain + filePath
}
