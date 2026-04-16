package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	apiURL  string
	token   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "usectl",
	Short: "A CLI tool for managing users via the usectl API",
	Long: `usectl-cli is a command-line interface for interacting with the usectl
user management API. It allows administrators to create, update, delete,
and list users, as well as manage authentication tokens.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags available to all subcommands
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.usectl.yaml)")
	// Default API URL set to my local dev instance for convenience
	rootCmd.PersistentFlags().StringVar(&apiURL, "api-url", "http://localhost:8080", "usectl API base URL (e.g. https://api.example.com)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "authentication token")

	// Bind flags to viper
	_ = viper.BindPFlag("api_url", rootCmd.PersistentFlags().Lookup("api-url"))
	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error finding home directory:", err)
			os.Exit(1)
		}

		// Search for config in home directory with name ".usectl"
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".usectl")
	}

	// Read in environment variables prefixed with USECTL_
	viper.SetEnvPrefix("USECTL")
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// getAPIURL returns the configured API URL or exits with an error
func getAPIURL() string {
	url := viper.GetString("api_url")
	if url == "" {
		fmt.Fprintln(os.Stderr, "error: API URL is required. Set --api-url flag or USECTL_API_URL environment variable.")
		os.Exit(1)
	}
	return url
}

// getToken returns the configured auth token or exits with an error
func getToken() string {
	t := viper.GetString("token")
	if t == "" {
		fmt.Fprintln(os.Stderr, "error: authentication token is required. Set --token flag or USECTL_TOKEN environment variable.")
		os.Exit(1)
	}
	return t
}
