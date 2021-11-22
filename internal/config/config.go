package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ConsumerKey string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type AccessConfig struct {
	AccessToken string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

func ReadConfig(path string) (Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func ReadAccessConfig(path string) (AccessConfig, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return AccessConfig{}, err
	}

	var config AccessConfig
	err = json.Unmarshal(content, &config)
	if err != nil {
		return AccessConfig{}, err
	}

	return config, nil
}

func WriteAccessConfig(path string, config AccessConfig) error {
	content, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, content, 0660)
}