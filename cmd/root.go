package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fInit  bool
	fLogin bool
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tw",
		Run: func(cmd *cobra.Command, args []string) {

			// text, err := cmd.Flags().GetString("text")

			// if err != nil {
			// 	fmt.Println(err)
			// 	os.Exit(1)
			// }

			// if text == "" {
			// 	cmd.Help()
			// 	return
			// }

			// fmt.Println("ToDo: Tweet", text)
		},
	}

	cmd.Flags().BoolVar(&fInit, "init", false, "Initialize TwCli")
	cmd.Flags().BoolVar(&fLogin, "login", false, "Login to TwCli")

	return cmd
}

func init() {
	cobra.OnInitialize(initializeConfig)
}

func initializeConfig() {
	homePath, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configPath := filepath.Join(homePath, ".twcli")
	configName := "hosts"
	configType := "yaml"

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// if config not found
	if err := viper.ReadInConfig(); err != nil {
		os.MkdirAll(configPath, 0700)
		viper.WriteConfigAs(filepath.Join(configPath, fmt.Sprintf("%s.%s", configName, configType)))
	}
}
