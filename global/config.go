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
	CfgFile string
	//Provider 指定配置选项
	Provider string
)

type CERTxConfig struct {
	Current string                     `json:"current"`
	Items   map[string]CERTxConfigItem `json:"items"`
}

type CERTxConfigItem struct {
	//aliyun 地域ID
	RegionID string `json:"regionid,omitempty"`
	//aliyun access_key_id
	AccessKeyID string `json:"accesskeyid"`
	//aliyun access_key_secret
	AccessKeySecret string `json:"accesskeysecret,omitempty"`

	// Token 为 DNSPod 的账户信息
	DNSPodToken string `json:"DnsPodToken"`
}

// Load 加载配置文件
func Load() (certx CERTxConfig) {
	if CfgFile == "" || CfgFile == "$HOME/.certx/certx.json" {
		CfgFile = fmt.Sprintf("%s/.certx/certx.json", os.Getenv("HOME"))
	}
	data, err := ioutil.ReadFile(CfgFile)
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
func (certx CERTxConfig) Dump(configFile string) {
	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}

	f.WriteString(certx.Marshal())
}

// New 新建 配置文件

// Marshal 格式化配置文件
func (certx CERTxConfig) Marshal() (s string) {

	b, err := json.MarshalIndent(certx, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Delete 删除 Provider
func (certx *CERTxConfig) Delete(provider string) {
	delete(certx.Items, provider)
}
