package cmd

import (
	//"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "管理配置文件",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	ConfigureMain()
	// },
}

var configureListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出 Config 中的所有 provider",
	Run: func(cmd *cobra.Command, args []string) {
		ListProviders()
	},
}

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "向 Config 中增加 provider",
	Run: func(cmd *cobra.Command, args []string) {
		AddProvider()
	},
}

// SetCurrent 配置默认 Profile

var configureCurrentCmd = &cobra.Command{
	Use:   "set",
	Short: "修改 current值， 设置默认生效的 profile",
	Run: func(cmd *cobra.Command, args []string) {
		SetCurrent()
	},
}

//var configureDeleteCmd = &cobra.Command{
//	Use:   "rm",
//	Short: "从 Config 中删除 profile",
//	Run: func(cmd *cobra.Command, args []string) {
//		DeleteProfile()
//	},
//}

func init() {
	cmdRoot.AddCommand(configureCmd)

	configureCmd.AddCommand(configureCurrentCmd)
	//configureCmd.AddCommand(configureDeleteCmd)
	configureCmd.AddCommand(configureListCmd)
	//configureCmd.AddCommand(configureDomainsCmd)
}
