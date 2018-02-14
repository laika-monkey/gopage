package main

import (
//	"golang.org/x/oauth2/fitbit"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
//	"context"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type Secrets struct {
	ClientID     string `yaml:"ClientID"`
	ClientSecret string `yaml:"ClientSecret"`
}

// read in secrets
func (s *Secrets) getSecrets() *Secrets {
	yamlHandle, err := ioutil.ReadFile("server.yaml")
	if err != nil {
		log.Printf("server.yaml err #%v ", err)
	}

	err = yaml.Unmarshal(yamlHandle, s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return s
}

// hello performs initial personal authentication
func hello(w http.ResponseWriter, r *http.Request) {
	//	ctx := context.Background()
	var secrets Secrets
	secrets.getSecrets()

	conf := &oauth2.Config{
		ClientID:     secrets.ClientID,
		ClientSecret: secrets.ClientSecret,
		Scopes: []string{"activity", "heartrate", "location", "nutrition",
			"profile", "settings", "sleep", "social", "weight"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.fitbit.com/oauth2/authorize",
			TokenURL: "https://api.fitbit.com/oauth2/token",
		},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Fprintf(w, "<a href='%v'>Link</a>", url)
	//	io.WriteString(w, "Hello world!")
	//	fmt.Fprintf(w, "Visit"
}

// meat of the test privacy violating page
func walter(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Landing Page")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/walter", walter)
	http.ListenAndServe(":8000", nil)
}
