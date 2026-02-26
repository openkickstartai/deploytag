package main

import (
	"fmt"
	"os"
	"time"
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
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	if cfg.Port == 0 { cfg.Port = 8080 }
	if cfg.TTL == "" { cfg.TTL = "24h" }
	if cfg.Provider == "" { cfg.Provider = "docker" }
	
	// Validate port range
	if cfg.Port < 1 || cfg.Port > 65535 {
		return nil, fmt.Errorf("invalid port %d: must be between 1 and 65535", cfg.Port)
	}
	
	// Validate TTL format
	if _, err := time.ParseDuration(cfg.TTL); err != nil {
		return nil, fmt.Errorf("invalid TTL format '%s': %w", cfg.TTL, err)
	}
	
	// Validate provider
	allowedProviders := map[string]bool{"docker": true, "kubernetes": true, "fly": true}
	if !allowedProviders[cfg.Provider] {
		return nil, fmt.Errorf("unsupported provider '%s': must be one of docker, kubernetes, fly", cfg.Provider)
	}
	
	return &cfg, nil
}