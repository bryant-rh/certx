package certx

import (
	"github.com/bryant-rh/certx/pkg/configure"
	"github.com/spf13/cobra"
)

var ConfigureCmd = &cobra.Command{
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
		configure.ListProvider()
	},
}

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "向 Config 中增加 provider",
	Run: func(cmd *cobra.Command, args []string) {
		configure.AddProvider()
	},
}

// SetCurrent 配置默认 Profile

var configureCurrentCmd = &cobra.Command{
	Use:   "set",
	Short: "修改 current值， 设置默认生效的 profile",
	Run: func(cmd *cobra.Command, args []string) {
		configure.SetCurrent()
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
	ConfigureCmd.AddCommand(configureCurrentCmd)
	//configureCmd.AddCommand(configureDeleteCmd)
	ConfigureCmd.AddCommand(configureListCmd)
	//configureCmd.AddCommand(configureDomainsCmd)
}
