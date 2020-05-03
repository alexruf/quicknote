package config

import (
	"github.com/alexruf/quicknote/common"
	"github.com/spf13/afero"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func HasConfig() bool {
	return len(viper.ConfigFileUsed()) > 0
}

func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	}

	configPath := filepath.Join(home, common.ConfigDir)
	configFileName := common.ConfigName + "." + common.ConfigType
	fs := afero.NewOsFs()
	if exists, err := afero.DirExists(fs, configPath); err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	} else {
		if !exists {
			if err := fs.MkdirAll(configPath, 0755); err != nil {
				jww.ERROR.Fatalln("Error: ", err)
			}
		}
	}
	if exists, err := afero.Exists(fs, filepath.Join(configPath, configFileName)); err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	} else {
		if !exists {
			if _, err := fs.Create(filepath.Join(configPath, configFileName)); err != nil {
				jww.ERROR.Fatalln("Error: ", err)
			}
		}
	}

	viper.AddConfigPath(filepath.Join(home, common.ConfigDir))
	viper.SetConfigName(common.ConfigName)
	viper.SetConfigType(common.ConfigType)
	viper.SetEnvPrefix(common.ApplicationShortName)
	viper.AutomaticEnv()
	viper.SetFs(fs)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			jww.ERROR.Fatalln("Error: ", err)
		}
	}
}
