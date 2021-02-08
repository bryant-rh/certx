package global

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Providers DNS 解析供应商
var Providers []string = []string{"aliyun", "dnspod"}

var (
	//指定配置路径
	ConfigFile string
	//Provider 指定配置选项
	Provider string
)

type CertxConfig struct {
	Current string                     `json: "current"`
	Items   map[string]CertxConfigItem `json:"items"`
}

type CertxConfigItem struct {
	//aliyun 地域ID
	REGION_ID string `json:"REGION_ID"`
	//aliyun access_key_id
	ACCESS_KEY_ID string `json:"ACCESS_KEY_ID"`
	//aliyun access_key_secret
	ACCESS_KEY_SECRET string `json:"ACCESS_KEY_SECRET"`

	// Token 为 DNSPod 的账户信息
	DnsPodToken string `json: "DnsPodToken"`
}

// Load 加载配置文件
func load() (certx CertxConfig) {
	if ConfigFile == "" || ConfigFile == "$HOME/.certx/certx.json" {
		ConfigFile = fmt.Sprintf("%s/.certx/certx.json", os.Getenv("HOME"))
	}
	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &certx)
	if err != nil {
		panic(err)
	}

	return
}

//Dump 写入配置文件
func (certx CertxConfig) Dump(configFile string) {
	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}

	f.WriteString(certx.Marshal())
}

// New 新建 配置文件

// Marshal 格式化配置文件
func (certx CertxConfig) Marshal() (s string) {

	b, err := json.MarshalIndent(certx, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Delete 删除 Provider
func (certx *CertxConfig) Delete(provider string) {
	delete(certx.Items, provider)
}
