package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name string `yaml:"name"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	// load in config file (load from parameter or default to config.yml)
	// format:
	// - name: string
	//
	// start http server
	// -endpoint /health returns 200 OK
	// -endpoint / returns "Hello {config.name}"

	var config *Config
	var err error
	if len(os.Args) > 1 {
		config, err = loadConfig(os.Args[1])
	} else {
		config, err = loadConfig("config.yml")
	}

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + config.Name))
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
