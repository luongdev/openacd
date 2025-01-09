package config

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type DbConfig struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	Username       string        `mapstructure:"username"`
	Password       string        `mapstructure:"password"`
	Database       string        `mapstructure:"database"`
	AuthSource     string        `mapstructure:"auth_source"`
	Dsn            string        `mapstructure:"dsn"`
	Timeout        time.Duration `mapstructure:"timeout"`
	ConnectTimeout time.Duration `mapstructure:"connect_timeout"`
	PoolSize       uint64        `mapstructure:"pool_size"`
}

func (c DbConfig) GetDsn() string {
	if c.Dsn != "" {
		return c.Dsn
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("mongodb://"))

	if c.Username != "" && c.Password != "" {
		sb.WriteString(fmt.Sprintf("%s:%s@", c.Username, url.QueryEscape(c.Password)))
	} else if c.Username != "" {
		sb.WriteString(fmt.Sprintf("%s@", c.Username))
	}

	if c.Host != "" && c.Port > 0 {
		sb.WriteString(fmt.Sprintf("%s:%d", c.Host, c.Port))
	} else if c.Host != "" {
		sb.WriteString(c.Host)
	}

	if c.Database != "" {
		sb.WriteString(fmt.Sprintf("/%s", c.Database))
	} else {
		sb.WriteString("/")
	}

	if c.AuthSource != "" {
		sb.WriteString(fmt.Sprintf("?authSource=%s", c.AuthSource))
	}

	return sb.String()
}
