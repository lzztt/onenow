package main

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const confFile = "config.yaml"

type config struct {
	Server struct {
		Port     string `yaml:"port"`
		KeyFile  string `yaml:"keyFile"`
		CertFile string `yaml:"certFile"`
	} `yaml:"server"`

	Data struct {
		NoteDir string `yaml:"noteDir"`
	} `yaml:"data"`

	Secret struct {
		AllowedEmail string `env:"ALLOWED_EMAIL"`
	}
}

var envFiles = map[string][]string{
	"development": {
		".env.development.local",
		".env.development",
	},
	"production": {
		".env.production.local",
		".env.production",
	},
	"test": {
		".env.test.local",
		".env.test",
	},
}

func loadEnv() {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}

	files, ok := envFiles[env]
	if !ok {
		log.Fatal("Unsupported environment: " + env)
	}

	log.Println("Running environment: " + env)

	for _, f := range files {
		err := godotenv.Load(f)

		if err != nil {
			switch err.(type) {
			case *os.PathError:
				break
			default:
				log.Println("Failed to load " + f)
				log.Fatal(err)
			}
		} else {
			log.Println("Loaded " + f)
		}
	}
}

func getConfig() config {
	loadEnv()

	var c config
	if err := cleanenv.ReadConfig(confFile, &c); err != nil {
		log.Fatal(err)
	}

	return c
}
