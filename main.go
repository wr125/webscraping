package main

import (
	"fmt"
	"net/http"

	logr "github.com/sirupsen/logrus"
	"github.com/wr125/webscraping/pkg/actions"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Assign the 'msg' variable with a string value
	msg := "Hello, welcome to your app. Use the following suffix's on the URL to show the different results.\n1)'/scrape' to show results of web scraping.\n2)'/crawl' to show results of a web crawler"

	// Logs a message to the terminal using the 3rd party import logrus
	logr.Info("Received request for the home page")

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf(msg)))
}

func main() {

	// Create the first route handler listening on '/'
	http.HandleFunc("/", home)
	logr.Info("Server is running on port 4000..")

	http.HandleFunc("/scrape", actions.Scrape)
	// Start the sever
	http.ListenAndServe(":4000", nil)

}
