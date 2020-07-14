package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const INSTAGMAM_HOST = "https://www.instagram.com/web/search/topsearch/"

func main() {
	var search_query string
	fmt.Print("Enter search query: ")
	fmt.Fscan(os.Stdin, &search_query)
	url := fmt.Sprintf("%s?context=blended&query=%s&include_reel=true", INSTAGMAM_HOST, search_query)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
