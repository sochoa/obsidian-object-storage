package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

type ObjectStorageConfig struct {
	StorageRoot              string
	Host                     string
	Port                     int
	DatabaseConnectionString string
	GracefulShutdownTimeout  time.Duration
	WriteTimeout             time.Duration
	ReadTimeout              time.Duration
	IdleTimeout              time.Duration
	FileMode                 os.FileMode
	DirMode                  os.FileMode
}

func (cfg ObjectStorageConfig) String() string {
	var kv map[string]string = make(map[string]string, 0)
	kv["StorageRoot"] = cfg.StorageRoot
	kv["Host"] = cfg.Host
	kv["Port"] = strconv.Itoa(cfg.Port)
	kv["GracefulShutdownTimeout"] = cfg.GracefulShutdownTimeout.String()
	kv["WriteTimeout"] = cfg.WriteTimeout.String()
	kv["ReadTimeout"] = cfg.ReadTimeout.String()
	kv["IdleTimeout"] = cfg.IdleTimeout.String()
	kv["FileMode"] = cfg.FileMode.String()
	kv["DirMode"] = cfg.DirMode.String()
	kv["DatabaseConnectionString"] = cfg.DatabaseConnectionString
	serializedAttrs := ""
	for k, v := range kv {
		if len(serializedAttrs) > 0 {
			serializedAttrs += ", "
		}
		serializedAttrs += fmt.Sprintf("%s=%s", k, strconv.Quote(v))
	}
	return fmt.Sprintf("{Config: %s}", serializedAttrs)
}

func (cfg *ObjectStorageConfig) BindPoint() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func GetPwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "/tmp"
	}
	return dir
}

func NewObjectStorageConfig() ObjectStorageConfig {
	return ObjectStorageConfig{
		Host:                    "localhost",
		Port:                    8080,
		StorageRoot:             path.Join(GetPwd(), "obsidian-data"),
		WriteTimeout:            time.Second * 15,
		ReadTimeout:             time.Second * 15,
		IdleTimeout:             time.Second * 60,
		GracefulShutdownTimeout: time.Second * 5,
		FileMode:                0666,
		DirMode:                 0755,

		DatabaseConnectionString: path.Join(GetPwd(), "obsidian-sqlite.db"),
	}
}
