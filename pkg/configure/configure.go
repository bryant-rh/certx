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
	var item global.CERTxConfigItem

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
				Message: "输入 Aliyun REGION_ID:",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "ACCESS_KEY_ID",
			Prompt: &survey.Input{
				Message: "输入 Aliyun ACCESS_KEY_ID:",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "ACCESS_KEY_SECRET",
			Prompt: &survey.Password{
				Message: "输入 Aliyun ACCESS_KEY_ID: ",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}

	var dnsPodToken = []*survey.Question{
		{
			Name: "DnsPodToken",
			Prompt: &survey.Input{
				Message: "输入 DnsPod DnsPodToken:",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}
	err := survey.Ask(qsProvider, &item)
	if err != nil {
		panic(err)
	}

	if global.Provider == "aliyun" {
		survey.Ask(qsLoginWithKey, &item)
	}

	if global.Provider == "dnspod" {
		survey.Ask(dnsPodToken, &item)
	}
	var confirm bool = false

	survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("是否添加 %s 到配置中", global.Provider),
	}, &confirm,
	)

	if confirm {
		certx.Items[global.Provider] = item
		certx.Dump(global.CfgFile)
	} else {
		logrus.Infoln("用户取消添加")
	}

}
