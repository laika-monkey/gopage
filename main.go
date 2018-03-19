package main

import (
	//	"golang.org/x/oauth2/fitbit"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	//	"context"
	"fmt"
	"github.com/larkaen/gopage/client"
)

// hello performs initial personal authentication
func hello(w http.ResponseWriter, r *http.Request) {
	//	ctx := context.Background()

	// define client object
	appinfo := &client.Client{}
	client.LoadClientId(appinfo)

	conf := &oauth2.Config{
		ClientID:     appinfo.ClientID,
		ClientSecret: appinfo.ClientSecret,
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
	fmt.Printf(appinfo.ClientID)
	fmt.Printf(appinfo.ClientSecret)
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
