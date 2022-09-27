package auth

import (
	"github.com/canoypa/twcli/cmd/auth/login"
	"github.com/spf13/cobra"
)

func AuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth",
	}

	cmd.AddCommand(login.LoginCmd())

	return cmd
}
