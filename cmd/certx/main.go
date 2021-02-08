package main

import (
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

	// global vars
	rootCmd.PersistentFlags().StringVarP(&global.ConfigFile, "config", "c", "$HOME/.certx/certx.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&global.Provider, "provider", "p", "dnspod", "provider")
}

func main() {
	if err := cmdRoot.Execute(); err != nil {
		panic(err)
	}
}
