package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"

	// "net/url"
	"os"
	"os/user"
	"github.com/rubenfonseca/fastimage"

	// "encoding/json"
	"github.com/akamensky/argparse"
)

const dir string = "~/Pictures/Wallpapers/"
const subreddit string = "wallpapers"
const minWidth int = 1920
const minHeight int = 1080
const postsPerRequest int = 20
const loops int = 5

type foo struct {
	Bar string
}

var client *http.Client = &http.Client{}

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

func makeHTTPReq(URL string) *http.Response {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Go_Wallpaper_Downloader")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 200 {
		return resp
	}
	log.Fatalln(err)
	return resp
}

func verifySubreddit(subreddit string) bool {
	URL := "https://reddit.com/r/" + subreddit

	req, err := http.NewRequest("GET", URL, nil)
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

func getJSON(URL string, target interface{}) {
	// var val = new(Foo)
	// client := &http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("User-Agent", "wallpaperDownloader")
	resp, httpErr := client.Do(req)

	if httpErr != nil {
		fmt.Println("HTTP Error = ", httpErr)
		log.Fatal(httpErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println("Body = ", bodyString)
	}

}

func getPosts(subreddit, topRange, after string, loop int) {
	for i := 0; i < 1; i++ {
		var URL string = fmt.Sprintf("https://reddit.com/r/%s/top/.json?t=%s&limit=%s&after=%s", subreddit, topRange, postsPerRequest, after)
		// var URL string = "http://example.com/"
		foo1 := new(foo) // or &Foo{}
		getJSON(URL, foo1)
		println(foo1.Bar)
	}
}

func isImg(URL string) bool{
	if strings.HasSuffix(URL, ".png") || strings.HasSuffix(URL, ".jpeg") || strings.HasSuffix(URL, ".jpg"){
		return true
	}
	return false
}

func isHD(URL string) bool {
	_,size,err := fastimage.DetectImageType(URL)
	if(err!=nil){
		print(err)
		return false
	}

	width := int(size.Width)
	height := int(size.Height)

	if(height>=minHeight && width>=minWidth){
		return true
	}
	return false
}

func isLandscape(URL string) bool{
	_,size,err := fastimage.DetectImageType(URL)
	if(err!=nil){
		print(err)
		return false
	}

	width := int(size.Width)
	height := int(size.Height)

	if(width>height){
		return true
	}
	return false
}

func main() {

	parser := argparse.NewParser("wallpaper-downloader", "Fetch wallpapers from Reddit")
	var topRange *string = parser.Selector("r", "range", []string{"day", "week", "month", "year", "all"}, &argparse.Options{Required: false, Help: "Range for top posts", Default: "all"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	fmt.Println("Selected Range = ", *topRange)

	getPosts(subreddit, *topRange, "", loops)
}
