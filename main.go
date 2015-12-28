package main

import (
	"log"
	"net/http"

	"github.com/carbocation/interpose"
	"github.com/kelseyhightower/envconfig"
)

const DATA_DIR = "./data"
const URL_PREFIX = "http://localhost:8000/"

type Config struct {
	Address   string `envconfig:"ADDRESS"`
	DataDir   string `envconfig:"DATA_DIR"`
	URLPrefix string `envconfig:"URL_PREFIX"`
}

var config *Config

func loadConfig() *Config {
	conf := Config{}
	err := envconfig.Process("droplink", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	if conf.Address == "" {
		conf.Address = "localhost:8000"
	}
	if conf.DataDir == "" {
		conf.DataDir = "./data"
	}
	if conf.URLPrefix == "" {
		conf.URLPrefix = "http://localhost:8000/"
	}
	return &conf
}

func main() {
	config = loadConfig()
	middle := interpose.New()

	router := http.NewServeMux()
	middle.UseHandler(router)

	router.Handle("/upload", http.HandlerFunc(upload))
	router.Handle("/", http.HandlerFunc(download))

	log.Printf("Listening on %s", config.Address)
	err := http.ListenAndServe(config.Address, middle)
	if err != nil {
		log.Fatal(err)
	}
}
