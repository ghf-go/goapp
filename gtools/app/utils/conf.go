package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Host  string `json:"host"`
	Token string `json:"token"`
}

func getConfigPath() string {
	ret := BuildHomePathAndCreate(".config/gtools/config.json")
	fmt.Println(ret)
	return ret
}

func GetConfig() *Config {
	var ret = &Config{}
	dd, e := os.ReadFile(getConfigPath())
	if e != nil {
		return ret
	}
	e = json.Unmarshal(dd, ret)
	if e != nil {
		return ret
	}
	return ret
}
func SaveConfig(config any) error {
	dd, e := json.Marshal(config)
	if e != nil {
		return e
	}
	return os.WriteFile(getConfigPath(), dd, 0644)
}

func (c *Config) ApiPost(pathname string, reqData, resq any) error {
	client := &http.Client{}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Host+pathname, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	re := ApiResp{
		Code: 0,
		Msg:  "",
		Data: resq,
	}
	err = json.Unmarshal(body, &re)
	if err != nil {
		return err
	}
	if re.Code != 200 {
		return fmt.Errorf(re.Msg)
	}
	return nil
}

type ApiResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
