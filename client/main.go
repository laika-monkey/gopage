package client

import (
	"io/ioutil"
	"log"
        "gopkg.in/yaml.v2"
)

// declare structure to read in yaml data
type Client struct {
        ClientID     string `yaml:"ClientID"`
        ClientSecret string `yaml:"ClientSecret"`
}

// read in secrets
func LoadClientId(client *Client){

	// open file for reading
        yamlHandle, err := ioutil.ReadFile("server.yaml")
        if err != nil {
                log.Printf("server.yaml err #%v ", err)
        }

	// read in data from yaml
        err = yaml.Unmarshal(yamlHandle, client)
        if err != nil {
                log.Fatalf("Unmarshal: %v", err)
        }
}

