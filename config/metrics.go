package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	defaultMetricsServerPort = 2112
	defaultMetricsHost       = "0.0.0.0"
)

// MetricsConfig defines the server's basic configuration
type MetricsConfig struct {
	// IP of the prometheus server
	Host string `mapstructure:"host"`
	// Port of the prometheus server
	ServerPort int `mapstructure:"server-port"`
}

func (cfg *MetricsConfig) Validate() error {
	if cfg.ServerPort < 0 || cfg.ServerPort > 65535 {
		return fmt.Errorf("invalid port: %d", cfg.ServerPort)
	}

	ip := net.ParseIP(cfg.Host)
	if ip == nil {
		return fmt.Errorf("invalid host: %v", cfg.Host)
	}

	return nil
}

func DefaultMetricsConfig() MetricsConfig {
	s := os.Getenv("METRICS_PORT")
	port, err := strconv.Atoi(s)
	if err != nil {
		port = defaultMetricsServerPort
	}
	return MetricsConfig{
		ServerPort: port,
		Host:       defaultMetricsHost,
	}
}
