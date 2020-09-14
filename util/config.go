package util

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
)

const RootPath = "E:\\Go"

var configMap = Map{}

func GetConfig(path string) Map {
	if configMap[path] != nil {
		return configMap[path].(Map)
	}
	v := make(Map)
	if _, err := toml.DecodeFile(filepath.Join(RootPath, "conf", path), &v); err != nil {
		panic(err)
	}
	configMap[path] = v
	return v
}

func GetDataBaseConfig(key string) Map {
	c := GetConfig("database.toml")
	return c[key].(Map)
}
func GetServerConfig(key string) Map {
	c := GetConfig("server.toml")
	return c[key].(Map)
}
