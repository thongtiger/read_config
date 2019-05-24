package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
)

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}
	checkErr(err)
	log.Println("Successfully changed json settings")
	byt, _ := json.Marshal(cfg)
	fmt.Println(string(byt))
}

func readConfig() (*config, error) {
	_, filePath, _, _ := runtime.Caller(0)
	pwd := filePath[:len(filePath)-7]
	txt, err := ioutil.ReadFile(pwd + "/config.json")
	if err != nil {
		return nil, err
	}
	var cfg = new(config)
	if err := json.Unmarshal(txt, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type config struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}
