package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Conf Config
var Conf Config

// workDir string
var workDir string

// Environment string
var Environment string

// Init 初期設定
func Init() {
	setEnv()
	readConfig()
}

// setEnv GO_ENVの値を設定する
func setEnv() string {
	Environment = os.Getenv("GO_ENV")

	// GO_ENVがからの場合は開発環境とする
	if Environment == "" {
		Environment = "develop"
	}
	log.Printf("Environment :%s", Environment)
	return Environment
}

func getWorkDir() string {
	wrkDir := os.Getenv("WRK_DIR")
	if isDev() {
		wrkDir += "/go-aws"
	}
	log.Printf("wrkDir: %v", wrkDir)
	return wrkDir
}

func isDev() bool {
	log.Println("Environment", Environment)
	return Environment == "develop" || Environment == "test"
}

func readConfig() Config {
	fullPath := getWorkDir()
	yamlPath := fullPath + "/config/yml/" + Environment + ".yml"
	log.Printf("yamlPath: %s", yamlPath)
	// yamlを読み込む
	buf, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}

	// log.Printf("buf: %+v\n", string(buf))

	// structにUnmasrshal
	err = yaml.Unmarshal(buf, &Conf)
	if err != nil {
		panic(err)
	}
	log.Printf("d: %+v\n", Conf)
	return Conf
}
