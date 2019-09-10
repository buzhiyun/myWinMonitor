package config

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"path/filepath"
)

var (
	Conf = New()
)

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	//config, err := toml.LoadFile( "./config/conf.toml")
	config, err := toml.LoadFile(GetCurrentDirectory() + "/config/conf.toml")

	if err != nil {
		log.Println("TomlError ", err.Error())
	}

	return config
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}
	//return strings.Replace(dir, "\\", "/", -1)
	return dir
}
