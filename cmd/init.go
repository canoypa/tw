package cmd

import (
	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "init",
		Run: func(cmd *cobra.Command, args []string) {
			// consumer key/secret の入力
			// その後 auth login コマンド実行の確認
		},
	}

	return cmd
}
