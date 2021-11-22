package main

import (
	"log"
	. "randomarticle/internal/config"
	"randomarticle/internal/twitter"
	"randomarticle/internal/wikipedia"
	"time"
)

const configPath = "config.json"
const accessConfigPath = "access.json"

func getAccessConfig(config Config) AccessConfig {
	accessConfig, err := ReadAccessConfig(accessConfigPath)
	if err == nil {
		return accessConfig
	}

	log.Println("no access config found")

	accessConfig, err = twitter.Login(config)
	if err != nil {
		log.Fatalln(err)
	}

	err = WriteAccessConfig(accessConfigPath, accessConfig)
	if err != nil {
		log.Fatalln(err)
	}

	accessConfig, err = ReadAccessConfig(accessConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	return accessConfig
}

func main() {
	config, err := ReadConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	access := getAccessConfig(config)
	twitter.Init(config, access)

	for range time.Tick(time.Minute * 1) {
		log.Println("tick")

		page, err := wikipedia.Get()
		if err != nil {
			log.Println(err)
			continue
		}

		err = twitter.Tweet(wikipedia.Format(page))
		if err != nil {
			log.Println(err)
			continue
		}
	}
}