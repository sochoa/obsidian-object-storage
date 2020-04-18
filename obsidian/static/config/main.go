package config

import (
	"fmt"
	"github.com/sochoa/obsidian/util"
	"path"
	"strconv"
)

type StaticConfig struct {
	Root    string
	Host    string
	Port    int
	UriBase string
}

func (cfg StaticConfig) String() string {
	var kv map[string]string = make(map[string]string, 0)
	kv["Root"] = cfg.Root
	kv["Host"] = cfg.Host
	kv["Port"] = strconv.Itoa(cfg.Port)
	serializedAttrs := ""
	for k, v := range kv {
		if len(serializedAttrs) > 0 {
			serializedAttrs += ", "
		}
		serializedAttrs += fmt.Sprintf("%s=%s", k, strconv.Quote(v))
	}
	return fmt.Sprintf("{StaticUiConfig: %s}", serializedAttrs)
}

func NewStaticConfig() StaticConfig {
	return StaticConfig{
		Host:    "localhost",
		Port:    8080,
		Root:    path.Join(util.GetPwd(), "static"),
		UriBase: "/uri/",
	}
}

func (cfg *StaticConfig) FormatEndpoint() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}
