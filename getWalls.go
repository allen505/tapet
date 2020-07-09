package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	// "os"
)

func successSum(a, b int) int {
	return a + b
}

func failSum(a, b int) int {
	return a - b
}

func validURL(URL string) bool {
	resp, err := http.Get(URL)
	fmt.Print(resp.Status)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if resp.StatusCode == 404 {
		return false
	}

	return true
}

func prepareDirectory(directory string) bool {
	usr, _ := user.Current()
	directory = usr.HomeDir + directory
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

func verifySubreddit(subreddit string) bool{
	URL := "https://reddit.com/r/" + subreddit
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL , nil)
	if err != nil {
			log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Go_Wallpaper_Downloader")

	resp, err := client.Do(req)
	if err != nil {
			log.Fatalln(err)
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Hello World")
	// valid := validURL("http://i.imgur.com/Z6kdWmA.jpg")
	// valid := prepareDirectory("/Pictures/Wallpapers/Reddit/")
	// valid := verifySubreddit("unixporn")
	// print(valid)
}
