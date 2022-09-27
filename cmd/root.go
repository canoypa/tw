package cmd

import (
	"github.com/spf13/cobra"
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

	cmd.Flags().StringP("text", "t", "", "Tweet text")

	return cmd
}
