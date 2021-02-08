package configure

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bryant-rh/certx/global"
	"github.com/sirupsen/logrus"
)

// ListProvider 返回当前 config 中的所有 provider
func ListProvider() {
	certx := global.Load()
	var l []string
	for key := range certx.Items {
		l = append(l, key)
	}

	fmt.Println(strings.Join(l, "\n"))
}

// AddProfile 增加
func AddProvider() {
	certx := global.Load()
	var item global.CertxConfigItem

	var qsProvider = []*survey.Question{
		{
			Name: "provider",
			Prompt: &survey.Select{
				Message: "Choose a Dns Provider:",
				Options: global.Providers,
			},
		},
	}

	var qsLoginWithKey = []*survey.Question{
		{
			Name: "REGION_ID",
			Prompt: &survey.Input{
				Message: "输入 AK ID",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "ACCESS_KEY_ID",
			Prompt: &survey.Input{
				Message: "输入 AK ID",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "ACCESS_KEY_SECRET",
			Prompt: &survey.Password{
				Message: "输入 AK Secret: ",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}

	err := survey.Ask(qsProvider, &item)
	if err != nil {
		panic(err)
	}

	if item.Provider == "aliyun" || item.Provider == "qcloud" {
		survey.Ask(qsLoginWithKey, &item)
	}

	var confirm bool = false

	survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("是否添加 %s 到配置中", global.Profile),
	}, &confirm,
	)

	if confirm {
		dnsx.Items[global.Profile] = item
		dnsx.Dump(global.CfgFile)
	} else {
		logrus.Infoln("用户取消添加")
	}

}
