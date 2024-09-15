/*
	14-09-2024 by github.com/i4k1
	simple utility for parsing files by keywords on the 2ch.hk imageboard
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	sosuchUrl   = "https://2ch.hk/" // alternative: https://2ch.life/
	boards      = flag.String("boards",      "b",   "Which board to parse")
	keywords    = flag.String("keywords",    "",    "Use keywords for parsing")
	fileformats = flag.String("fileformats", "",    "What file formats to download")
	path        = flag.String("path",        "src", "In which directory to save files")
)

// for catalog parsing
type OPs struct {
	Comment string `json:"comment"`
	Date    string `json:"date"`
	Num     int    `json:"num"`
}

type Catalog struct {
	Threads []OPs `json:"threads"`
}

// for thread parsing
type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Post struct {
	Comment string `json:"comment"`
	Date    string `json:"date"`
	Files   []File `json:"files"`
	Num     int64  `json:"num"`
}

type Thread struct {
	Posts []Post `json:"posts"`
}

type PostsData struct {
	Threads []Thread `json:"threads"`
}

func downloadFile(url string, filename string) error {
	// getting response from server
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// сhecking response status and error handling
	if response.StatusCode != http.StatusOK {
		panic(response.Status)
	}

	// creating file for writing
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// copy response content to file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("downloaded:", filename)

	return nil
}

func parseThread(threadNumber string, board string, dirToSave string, formats string) {
	fformats := strings.Split(formats, ",")
	threadUrl := sosuchUrl + board + "/res/" + threadNumber + ".json" // full link to thread's json

	// GET-request
	response, err := http.Get(threadUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// status code of response and error handling
	if response.StatusCode != http.StatusOK {
		panic(response.StatusCode)
	}

	// reading body response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// catalog.json decode
	var postsData PostsData
	err = json.Unmarshal(body, &postsData)
	if err != nil {
		panic(err)
	}

	// if the directory does not exist, create it
	if _, err := os.Stat(dirToSave); os.IsNotExist(err) {
		err := os.Mkdir(dirToSave, 0755) // 0755 - access rights
		if err != nil {
			panic(err)
		}
		fmt.Println("directory created:", dirToSave)
	}

	for _, thread := range postsData.Threads {
		for _, post := range thread.Posts {
			for _, file := range post.Files {
				fileUrl := sosuchUrl + file.Path
				wichDir := dirToSave + "/" + file.Name
				for _, ext := range fformats {
					if strings.HasSuffix(file.Name, ext) {
						fmt.Println("post:", post.Num)
						err := downloadFile(fileUrl, wichDir)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
}

func parseCatalog(board string, keywordPattern string) {
	// link to catalog
	sosuchCatalogUrl := sosuchUrl + board + "/catalog.json"

	// GET-request
	response, err := http.Get(sosuchCatalogUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// status code of response and error handling
	if response.StatusCode != http.StatusOK {
		panic(response.StatusCode)
	}

	// reading body response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// catalog.json decode
	var catalog Catalog
	err = json.Unmarshal(body, &catalog)
	if err != nil {
		panic(err)
	}

	var keywordRegex *regexp.Regexp
	if keywordPattern != "" {
		var err error
		keywordRegex, err = regexp.Compile(keywordPattern)
		if err != nil {
			panic(err)
		}
	}

	for _, thread := range catalog.Threads {
		if keywordRegex != nil && !keywordRegex.MatchString(thread.Comment) {
			fmt.Println("skiped:", thread.Num)
			continue
		}

		fmt.Println("\nparsing:", thread.Num)

		threadNumber := strconv.Itoa(thread.Num)
		parseThread(threadNumber, *boards, *path, *fileformats)
	}
}

func main() {
	flag.Parse()
	parseCatalog(*boards, *keywords)
}
