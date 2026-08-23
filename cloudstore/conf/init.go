package conf

type CloudStoreConf struct {
	Type      string `yaml:"type"`
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Endpoint  string `yaml:"endpoint"`
	Host      string `yaml:"host"`
}

type TokenCloudStore struct {
	Header map[string]string `json:"header"`
	Host   string            `json:"host"`
	Body   map[string]string `json:"body"`
	Method string            `json:"method"`
}
