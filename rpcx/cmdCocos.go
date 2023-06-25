package rpcx

import (
	"github.com/gps-gaming-gs/goleaf/rpcx/cli"
	"github.com/spf13/cobra"
)

var (
	CocosCmd = &cobra.Command{
		Use:   "cocos",
		Short: "產生Cocos的模板，controller、view...",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.CocosTemplate(true)
		},
	}
)

func init() {
	// Cocos
	CocosCmd.Flags().StringVar(&cli.VarStringCCController, "controller", "", "建立 Cocos 控制器&&視圖")
	CocosCmd.Flags().StringVar(&cli.VarStringCCProxy, "proxy", "", "建立Proxy代理服務")
}
