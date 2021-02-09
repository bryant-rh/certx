package certx

import (
	"github.com/bryant-rh/certx/cmd/certx"
	"github.com/bryant-rh/certx/global"
	"github.com/spf13/cobra"
)

//var ctx pxctx.Context

var cmdRoot = &cobra.Command{
	Use:   "certx",
	Short: "certx 申请Let's Encrypt颁发的https证书",
	//PersistentPreRun: func(cmd *cobra.Command, args []string) {
	//},
}

func init() {
	cmdRoot.AddCommand(certx.ConfigureCmd)

	// global vars
	cmdRoot.PersistentFlags().StringVarP(&global.CfgFile, "config", "c", "$HOME/.certx/certx.json", "config file")
	cmdRoot.PersistentFlags().StringVarP(&global.Provider, "provider", "p", "dnspod", "provider")
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		panic(err)
	}
}
