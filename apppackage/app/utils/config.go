package utils

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 全局配置结构，与 conf/*.yaml 对应
type Config struct {
	Server struct {
		Port  int    `yaml:"port"`
		Host  string `yaml:"host"`
		Debug bool   `yaml:"debug"`
	} `yaml:"server"`
	DB struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		User            string `yaml:"user"`
		Password        string `yaml:"password"`
		Database        string `yaml:"database"`
		Charset         string `yaml:"charset"`
		ParseTime       bool   `yaml:"parseTime"`
		Loc             string `yaml:"loc"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		ConnMaxLifetime string `yaml:"connMaxLifetime"`
		ConnMaxIdleTime string `yaml:"connMaxIdleTime"`
	} `yaml:"db"`
	Admin struct {
		Username         string `yaml:"username"`
		Password         string `yaml:"password"`
		TokenExpireHours int    `yaml:"tokenExpireHours"`
	} `yaml:"admin"`
	Qiniu struct {
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
		Bucket    string `yaml:"bucket"`
		Domain    string `yaml:"domain"`
	} `yaml:"qiniu"`
	Build struct {
		WorkDir        string `yaml:"workDir"`
		TimeoutMinutes int    `yaml:"timeoutMinutes"`
	} `yaml:"build"`
	K8s string `yaml:"k8s"`
}

// Conf 全局配置实例
var Conf Config

// InitConfig 从 embed 的文件系统加载配置，环境由 APP_ENV 决定（默认 dev）
func InitConfig(confFS fs.FS) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	name := fmt.Sprintf("conf/%s.yaml", env)
	data, err := fs.ReadFile(confFS, name)
	if err != nil {
		panic("读取配置失败 " + name + ": " + err.Error())
	}
	if err := yaml.Unmarshal(data, &Conf); err != nil {
		panic("解析配置失败 " + name + ": " + err.Error())
	}
	if Conf.Build.WorkDir == "" {
		Conf.Build.WorkDir = "./builds"
	}
	if Conf.Build.TimeoutMinutes <= 0 {
		Conf.Build.TimeoutMinutes = 30
	}
	if Conf.Admin.TokenExpireHours <= 0 {
		Conf.Admin.TokenExpireHours = 24
	}
	Log("配置加载完成 env=%s port=%d", env, Conf.Server.Port)
}
