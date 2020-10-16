package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/akamensky/argparse"
	"github.com/rubenfonseca/fastimage"
)

// const dir string = "~/Pictures/Wallpapers/"
const dir string = "./Wallpapers/"
const minWidth int = 1920
const minHeight int = 1080
const postsPerRequest int = 20
const loops int = 5

const DARK = "\033[30m"
const RED = "\033[31m"
const GREEN = "\033[32m"
const ORANGE = "\033[33m"
const PURPLE = "\033[35m"

type jsonStruct struct {
	values string
}

type postStruct struct {
	name   string
	picURL string
	author string
	height float64
	width  float64
	nsfw   bool
}

var client *http.Client = &http.Client{Timeout: 10 * time.Second}

func validURL(URL string) bool {
	resp, err := http.Get(URL)
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

	log.Println("Failed to get HTTP Respone from URL = ", URL)
	return resp
}

func verifySubreddit(subredditName string) bool {
	URL := "https://reddit.com/r/" + subredditName
	resp := makeHTTPReq(URL)

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false

}

func getJSON(URL string, target interface{}) ([]interface{}, string) {
	// var val = new(Foo)
	// client := &http.Client{}

	resp := makeHTTPReq(URL)

	defer resp.Body.Close()

	bodyInBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}

	json.Unmarshal([]byte(bodyInBytes), &result)
	posts := result["data"].(map[string]interface{})["children"]
	after := result["data"].(map[string]interface{})["after"].(string)

	return posts.([]interface{}), after

}

func isImg(URL string) bool {
	if strings.HasSuffix(URL, ".png") || strings.HasSuffix(URL, ".jpeg") || strings.HasSuffix(URL, ".jpg") {
		return true
	}
	return false
}

func isHD(URL string) bool {
	_, size, err := fastimage.DetectImageType(URL)
	if err != nil {
		print(err)
		return false
	}

	width := int(size.Width)
	height := int(size.Height)

	if height >= minHeight && width >= minWidth {
		return true
	}
	return false
}

func isLandscape(URL string) bool {
	_, size, err := fastimage.DetectImageType(URL)
	if err != nil {
		print(err)
		return false
	}

	width := int(size.Width)
	height := int(size.Height)

	if width > height {
		return true
	}
	return false
}

func alreadyDownloaded(URL string) bool {

	s, _ := url.Parse(URL)
	usr, _ := user.Current()
	imgName := s.Path[1:]
	imgDirectory := usr.HomeDir + dir + imgName

	_, e := os.Stat(imgDirectory)
	if e != nil {
		if os.IsNotExist(e) {
			return false
		}
	}
	return true
}

func knownURL(post string) bool {

	if (strings.HasPrefix(strings.ToLower(post), "https://i.redd.it/")) || (strings.HasPrefix(strings.ToLower(post), "http://i.imgur.com/")) {
		return true
	}
	return false
}

func storeImg(post string) bool {
	resp, _ := http.Get(post)
	defer resp.Body.Close()

	s, _ := url.Parse(post)
	usr, _ := user.Current()
	directory := usr.HomeDir + dir + s.Path[1:]
	println(directory)

	file, err := os.Create(directory)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return false
	}
	return true
}

func extractPostsData(postsJSON []interface{}, posts *[]postStruct) {
	var postJSONData interface{}
	var postData postStruct

	for _, v := range postsJSON {
		postJSONData = (v.(map[string]interface{})["data"])

		postData.name = postJSONData.(map[string]interface{})["title"].(string)
		postData.picURL = postJSONData.(map[string]interface{})["url"].(string)
		postData.author = postJSONData.(map[string]interface{})["author"].(string)
		postData.nsfw = postJSONData.(map[string]interface{})["over_18"].(bool)

		if (postJSONData.(map[string]interface{})["preview"]) != nil {
			postData.height = (((postJSONData.(map[string]interface{})["preview"]).(map[string]interface{})["images"].([]interface{}))[0].(map[string]interface{})["source"]).(map[string]interface{})["height"].(float64)
			postData.width = (((postJSONData.(map[string]interface{})["preview"]).(map[string]interface{})["images"].([]interface{}))[0].(map[string]interface{})["source"]).(map[string]interface{})["width"].(float64)
		} else {
			postData.height = -1
			postData.width = -1
		}

		*posts = append(*posts, postData)
	}
}

func getPosts(subredditName string, topRange string, postsPerRequest int, loops int) []postStruct {
	var posts []postStruct = make([]postStruct, 0)
	var after string = ""

	for i := 0; i < loops; i++ {
		var URL string = fmt.Sprintf("https://reddit.com/r/%s/top/.json?t=%s&limit=%d&after=%s", subredditName, topRange, postsPerRequest, after)
		httpResp := new(jsonStruct)
		var postsJSON []interface{}
		postsJSON, after = getJSON(URL, httpResp)
		extractPostsData(postsJSON, &posts)
	}

	return posts
}

func downloadAndSave(posts []postStruct, fromIndex int, toIndex int) {
	for i := fromIndex; i < toIndex; i++ {

		if !validURL(posts[i].picURL) {
			log.Println("Skipping Invalid URL")
			continue
		}
		if !knownURL(posts[i].picURL) {
			log.Println("Skipping Unknown URL")
			continue
		}
		if !isImg(posts[i].picURL) {
			log.Println("Skipping non-Image URL")
			continue
		}
		if !isLandscape(posts[i].picURL) {
			log.Println("Skipping Portrait image")
			continue
		}
		if !isHD(posts[i].picURL) {
			log.Println("Skipping low resolution image")
			continue
		}
		if !alreadyDownloaded(posts[i].picURL) {
			log.Println("Skipping already downloaded image")
			continue
		}

		if storeImg(posts[i].picURL) {
			log.Println("Downloaded ", posts[i].name, " by ", posts[i].author)
		} else {
			log.Println("FAILED to download", posts[i].name, " by ", posts[i].author)
		}
	}
}

func parallelizeDownload(posts []postStruct, numberOfThreads int) {
	numberOfPosts := len(posts)
	postsPerThread := numberOfPosts / numberOfThreads

	for i := 0; i < numberOfThreads-1; i++ {
		// fmt.Println(i*postsPerThread, " to ", ((i + 1) * postsPerThread))
		downloadAndSave(posts, i*postsPerThread, (i+1)*postsPerThread)
	}
	// fmt.Println((numberOfThreads-1)*postsPerThread, " to ", numberOfPosts-1)
	downloadAndSave(posts, (numberOfThreads-1)*postsPerThread, numberOfPosts)
}

func main() {

	parser := argparse.NewParser("wallpaper-downloader", "Fetch wallpapers from Reddit")
	var topRange *string = parser.Selector("r", "range", []string{"day", "week", "month", "year", "all"}, &argparse.Options{Required: false, Help: "Range for top posts", Default: "all"})
	var subredditName *string = parser.String("s", "subreddit", &argparse.Options{Required: false, Help: "Name of Subreddit", Default: "wallpaper"})
	var numberOfThreads *int = parser.Int("n", "number", &argparse.Options{Required: false, Help: "Number of Threads", Default: 2})
	var posts []postStruct

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// verify subreddit
	if !verifySubreddit(*subredditName) {
		log.Fatalln("Failed to verify subreddit")
	}

	// Create directory and keep stuff ready

	// Fetch details of all the posts
	log.Println("Fetching details of posts")
	posts = getPosts(*subredditName, *topRange, postsPerRequest, loops)
	log.Println("Fetched details of ", len(posts), " posts")

	// Start downloading the photos and store it
	// Print the progress with relevant details on the Console
	parallelizeDownload(posts, *numberOfThreads)

	// Final stats(OPTIONAL)
}
