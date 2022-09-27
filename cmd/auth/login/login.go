package login

import (
	"github.com/spf13/cobra"
)

func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "login",
		Run: func(cmd *cobra.Command, args []string) {
			// access token/secret の取得
		},
	}

	return cmd
}
