package utils

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// @data_dir_from_cli_arg: data dir name come from the command line
// @dir_name: data dir name
// @cfg_file_name: configuration file name in `dir_name` directory
func LoadConfigurationFromYamlFile(data_dir_from_cli_arg, data_dir_name, cfg_file_name string, out interface{}) (string, error) {

	var cfg_file_name_absolute string
	// 1: check command line argument
	if len(data_dir_from_cli_arg) > 0 {
		// we got the data directory from command line
		cfg_file_name_absolute = filepath.Join(data_dir_from_cli_arg, cfg_file_name)
		if IsPathExisting(cfg_file_name_absolute) {
			err := loadConfigFromYamlFile(cfg_file_name_absolute, out)
			if err == nil {
				return data_dir_from_cli_arg, nil
			}
		}

	}

	// 2: check user HOME directory
	home_dir := os.Getenv("HOME")
	cfg_file_name_absolute = filepath.Join(home_dir, data_dir_name, cfg_file_name)
	if IsPathExisting(cfg_file_name_absolute) {
		err := loadConfigFromYamlFile(cfg_file_name_absolute, out)
		if err == nil {
			return filepath.Join(home_dir, data_dir_name), nil
		}
	}

	// 3: check application directory
	app_dir_absolute := GetAppBaseDir()
	cfg_file_name_absolute = filepath.Join(app_dir_absolute, data_dir_name, cfg_file_name)
	if IsPathExisting(cfg_file_name_absolute) {
		err := loadConfigFromYamlFile(cfg_file_name_absolute, out)
		if err == nil {
			return filepath.Join(app_dir_absolute, data_dir_name), nil
		}
	}

	return "", errors.New("load configuration data failed")
}

func loadConfigFromYamlFile(cfgFilename string, out interface{}) error {
	cfgContent, err := os.ReadFile(cfgFilename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cfgContent, out)
	if err != nil {
		return err
	}

	return nil
}
