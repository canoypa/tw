package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/canoypa/tw/utils"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twAuth "github.com/dghubble/oauth1/twitter"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fInit  bool
	fLogin bool
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "tw",
		Run: func(cmd *cobra.Command, args []string) {

			if fInit {
				initialize(cmd)
				return
			}

			if fLogin {
				login(cmd)
				return
			}

			if len(args) >= 1 {
				text := strings.Join(args, " ")
				tweet(text)
				return
			}

			text := utils.Multiline("What's happening?")
			tweet(text)
		},
	}

	rootCmd.Flags().BoolVar(&fInit, "init", false, "Initialize TwCli")
	rootCmd.Flags().BoolVar(&fLogin, "login", false, "Login to TwCli")

	return rootCmd
}

func initialize(cmd *cobra.Command) {
	consumerKey := utils.Input("Enter ConsumerKey")
	consumerSecret := utils.Input("Enter ConsumerSecret")

	viper.Set("consumer_key", consumerKey)
	viper.Set("consumer_secret", consumerSecret)
	viper.WriteConfig()

	continueSignIn := utils.Confirm("Continue SignIn?", true)

	if continueSignIn {
		cmd := RootCmd()
		cmd.SetArgs([]string{"--login"})
		loginCmdErr := cmd.Execute()
		cobra.CheckErr(loginCmdErr)
	}
}

func login(cmd *cobra.Command) {
	consumerKey := viper.GetString("consumer_key")
	consumerSecret := viper.GetString("consumer_secret")

	oauthConfig := &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint:       twAuth.AuthenticateEndpoint,
	}

	requestToken, requestSecret, reqTokenErr := oauthConfig.RequestToken()
	cobra.CheckErr(reqTokenErr)

	authUrl, authUrlErr := oauthConfig.AuthorizationURL(requestToken)
	cobra.CheckErr(authUrlErr)

	fmt.Println("Visit this URL to get a PIN.")
	fmt.Println(authUrl)

	openErr := browser.OpenURL(authUrl.String())
	cobra.CheckErr(openErr)

	pin := utils.Input("Enter PIN")

	token, secret, tokenErr := oauthConfig.AccessToken(requestToken, requestSecret, pin)
	cobra.CheckErr(tokenErr)

	viper.Set("token", token)
	viper.Set("secret", secret)

	writeConfigErr := viper.WriteConfig()
	cobra.CheckErr(writeConfigErr)

	fmt.Println("Successfully signed in.")
}

func tweet(Text string) {
	consumerKey := viper.GetString("consumer_key")
	consumerSecret := viper.GetString("consumer_secret")
	accessToken := viper.GetString("token")
	accessSecret := viper.GetString("secret")

	config := &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweet, _, err := client.Statuses.Update(Text, nil)
	cobra.CheckErr(err)

	url := strings.Join([]string{"https://twitter.com", tweet.User.ScreenName, "status", strconv.FormatInt(tweet.ID, 10)}, "/")
	fmt.Println("Your Tweet was sent: " + url)
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
