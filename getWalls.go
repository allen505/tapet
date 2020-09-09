package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"
	"net/url"

	"github.com/akamensky/argparse"
	"github.com/rubenfonseca/fastimage"
)

const dir string = "/Pictures/Wallpapers/"
const subreddit string = "wallpapers"
const minWidth int = 1920
const minHeight int = 1080
const postsPerRequest int = 20
const loops int = 5

type jsonStruct struct {
	values string
}

type postStruct struct {
	name      string
	picURL    string
	redditURL string
	author    string
	nsfw      bool
}

var client *http.Client = &http.Client{Timeout: 10 * time.Second}

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

	log.Fatalln("Failed to get HTTP Respone from URL = ", URL, "\n", resp)
	return resp
}

func verifySubreddit(subreddit string) bool {
	URL := "https://reddit.com/r/" + subreddit
	resp := makeHTTPReq(URL)

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false

}

func getJSON(URL string, target interface{}) []interface{} {
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

	return posts.([]interface{})

}

func extractPostsData(postsJSON []interface{}, posts *[]postStruct) {
	var postJSONData interface{}
	var postData postStruct

	for _, v := range postsJSON {
		postJSONData = (v.(map[string]interface{})["data"])

		postData.name = postJSONData.(map[string]interface{})["title"].(string)
		postData.picURL = postJSONData.(map[string]interface{})["url"].(string)
		postData.redditURL = postJSONData.(map[string]interface{})["permalink"].(string)
		postData.author = postJSONData.(map[string]interface{})["author"].(string)
		postData.nsfw = postJSONData.(map[string]interface{})["over_18"].(bool)

		*posts = append(*posts, postData)
	}
}

func getPosts(subreddit, topRange, after string, loops int) []postStruct {
	var posts []postStruct = make([]postStruct, 0)

	for i := 0; i < loops; i++ {
		var URL string = fmt.Sprintf("https://reddit.com/r/%s/top/.json?t=%s&limit=%d&after=%s", subreddit, topRange, postsPerRequest, after)
		httpResp := new(jsonStruct)
		var postsJSON []interface{} = getJSON(URL, httpResp)
		extractPostsData(postsJSON, &posts)
	}

	return posts
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

	s,_ := url.Parse(URL)
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

func knownURL(post string) bool{

	if((strings.HasPrefix(strings.ToLower(post),"https://i.redd.it/"))||(strings.HasPrefix(strings.ToLower(post),"http://i.imgur.com/"))){
		return true
	}
	return false
}

func storeImg(post string) bool {
	resp, _ := http.Get(post)
	defer resp.Body.Close()

	s,_ := url.Parse(post)
	usr, _ := user.Current()
	directory := usr.HomeDir + dir + s.Path[1:]
	println(directory)

	file, err := os.Create(directory)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer file.Close()

	_ , err = io.Copy(file, resp.Body)
	if err != nil {
		return false
	}
	return true
}

func main() {

	parser := argparse.NewParser("wallpaper-downloader", "Fetch wallpapers from Reddit")
	var topRange *string = parser.Selector("r", "range", []string{"day", "week", "month", "year", "all"}, &argparse.Options{Required: false, Help: "Range for top posts", Default: "all"})
	var posts []postStruct

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// fmt.Println("Selected Range = ", *topRange)
	posts = getPosts(subreddit, *topRange, "", loops)
	fmt.Println("Number of posts receiveced = ", len(posts), " and capacity = ", cap(posts))
}
