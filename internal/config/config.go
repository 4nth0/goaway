package config

import "fmt"

type StoreConfig struct {
	DBType string
	DBPath string
}

type StatsConfig struct {
	DBType string
	DBPath string
}

type Config struct {
	Port     int
	Timezone string
	Store    StoreConfig
	Stats    StatsConfig
}

func GetConfig() *Config {
	return &Config{
		Port: 8080,
		Store: StoreConfig{
			DBType: "fs",
			DBPath: "./db.json",
		},
		Stats: StatsConfig{
			DBType: "fs",
			DBPath: "./stats.json",
		},
	}
}

func (c Config) ServerPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
