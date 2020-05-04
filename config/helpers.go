package config

import (
	"github.com/alexruf/quicknote/common"
	"github.com/spf13/afero"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func ExistsConfigFile() bool {
	fs := afero.NewOsFs()
	if exists, err := afero.Exists(fs, getConfigFilePath()); err == nil && exists {
		// Empty files don't count
		if empty, err := afero.IsEmpty(fs, getConfigFilePath()); err == nil && !empty {
			return true
		}
	}
	return false
}

func InitConfig() {
	configDirPath := getConfigDirPath()
	fs := afero.NewOsFs()
	if exists, err := afero.DirExists(fs, configDirPath); err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	} else {
		if !exists {
			if err := fs.MkdirAll(configDirPath, 0755); err != nil {
				jww.ERROR.Fatalln("Error: ", err)
			}
		}
	}
	if exists, err := afero.Exists(fs, getConfigFilePath()); err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	} else {
		if !exists {
			if _, err := fs.Create(getConfigFilePath()); err != nil {
				jww.ERROR.Fatalln("Error: ", err)
			}
		}
	}

	viper.AddConfigPath(configDirPath)
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

func getConfigDirPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	}
	path, err := filepath.Abs(filepath.Join(home, common.ConfigDirName))
	if err != nil {
		jww.ERROR.Fatalln("Error: ", err)
	}
	return path
}

func getConfigFileName() string {
	return common.ConfigName + "." + common.ConfigType
}

func getConfigFilePath() string {
	return filepath.Join(getConfigDirPath(), getConfigFileName())
}
