package main

import (
	"os"
	"gopkg.in/yaml.v3"
)

type DeployConfig struct {
	Provider string `yaml:"provider"`
	Port     int    `yaml:"port"`
	Build    string `yaml:"build"`
	Run      string `yaml:"run"`
	TTL      string `yaml:"ttl"`
}

func LoadDeployConfig() (*DeployConfig, error) {
	data, err := os.ReadFile(".deploytag.yml")
	if err != nil {
		return &DeployConfig{Provider: "docker", Port: 8080, TTL: "24h"}, nil
	}
	var cfg DeployConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	if cfg.Port == 0 { cfg.Port = 8080 }
	if cfg.TTL == "" { cfg.TTL = "24h" }
	if cfg.Provider == "" { cfg.Provider = "docker" }
	return &cfg, nil
}
