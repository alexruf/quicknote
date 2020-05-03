package cmd

import (
	"fmt"
	"github.com/alexruf/quicknote/common"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   common.ApplicationShortName,
	Short: "A CLI tool that lets you easily save, access and organize your personal daily notes.",
	Long: `A CLI tool that lets you easily save, access and organize your personal daily notes.
Complete documentation is available at ` + common.Website,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		handleError(err)
	}

	configPath := filepath.Join(home, common.ConfigDir)
	configFileName := common.ConfigName + "." + common.ConfigType
	fs := afero.NewOsFs()
	if exists, err := afero.DirExists(fs, configPath); err != nil {
		handleError(err)
	} else {
		if !exists {
			if err := fs.MkdirAll(configPath, 0755); err != nil {
				handleError(err)
			}
		}
	}
	if exists, err := afero.Exists(fs, filepath.Join(configPath, configFileName)); err != nil {
		handleError(err)
	} else {
		if !exists {
			if _, err := fs.Create(filepath.Join(configPath, configFileName)); err != nil {
				handleError(err)
			}
		}
	}

	viper.AddConfigPath(filepath.Join(home, common.ConfigDir))
	viper.SetConfigName(common.ConfigName)
	viper.SetConfigType(common.ConfigType)
	viper.SetEnvPrefix(common.ApplicationName)
	viper.AutomaticEnv()
	viper.SetFs(fs)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			handleError(err)
		}
	}
}
