package main

import "github.com/BurntSushi/toml"

type Config struct {
	Aws    AwsConfig
	Server ServerConfig
}

type AwsConfig struct {
	Region          string `toml:"region"`
	AccessKeyId     string `toml:"aws_access_key_id"`
	SecretAccessKey string `toml:"aws_secret_access_key"`
	EndpointUrl     string `toml:"endpoint_url"`
}

type ServerConfig struct {
	Port   string `toml:"port"`
	Socket string `toml:"socket"`
}

func LoadConfig(fileName string) (*Config, error) {
	var config Config
	if fileName == "" {
		return &config, nil
	}

	_, err := toml.DecodeFile(fileName, &config)
	return &config, err
}
