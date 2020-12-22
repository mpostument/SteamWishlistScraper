/*
Copyright © 2020 Maksym Postument 777rip777@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mpostument/SteamWishlistScraper/steam"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "steamwishlistscraper",
	Short: "Root comand for steam interaction",
	Long:  `Root comand for steam interaction.`,
}

var api = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape steam wishlist",
	Long:  `Scrape steam wishlist.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		userName, _ := cmd.Flags().GetString("username")
		steamID := steam.GetSteamId(userName, apiKey)
		games := steam.ScrapeWishlist(steamID)
		steam.SaveToFile(games)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(api)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.SteamWishlistScraper.yaml)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Steam UserName")
	if err := rootCmd.MarkPersistentFlagRequired("username"); err != nil {
		log.Println(err)
	}
	rootCmd.PersistentFlags().StringP("apikey", "a", "", "Steam api key")
	if err := viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey")); err != nil {
		log.Println(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".SteamWishlistScraper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".SteamWishlistScraper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
