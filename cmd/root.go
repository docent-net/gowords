package cmd

import (
	"fmt"
	"os"

	"github.com/docent-net/gowords/words"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "gowords",
		Short: "CLI english words helper",
		Long:  `CLI english words helper`,
		Run: func(cmd *cobra.Command, args []string) {
			var tt = words.FetchWordInfo("terefere")
			fmt.Println("gowords: run help command for more information", tt)

		},
	}
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".gowords")
		viper.AddConfigPath("$HOME/")
		viper.AddConfigPath(".")
	}

	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "",
		"config file (default is $HOME/.gowords.yml)")

	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
