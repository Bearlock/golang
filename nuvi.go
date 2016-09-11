package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

// Get html page that has list of files and return the
// actual URL for the webpage (after the redirect)
func writeHTMLAndReturnRequestURL() (URL string) {
	output, err := os.Create("./files.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer output.Close()

	resp, err := http.Get("http://bitly.com/nuvi-plz")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	baseURL := resp.Request.URL
	fmt.Println(baseURL)
	return baseURL.String()
}

// Read/step through the zip files and iterate
// over the xml files within, call insertContentRedis()
func readZipFiles(dirName string, fileName string, wgReadZip *sync.WaitGroup) {
	defer wgReadZip.Done()

	zipFile := "./" + dirName + "/" + fileName

	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println(zipFile)
		fmt.Println(err)
	}
	defer zipReader.Close()

	fmt.Printf("Processing zip file: %s\n", zipFile)
	for _, fileContents := range zipReader.File {
		insertContentRedis(fileContents)
	}

	fmt.Printf("Processing zip file: %s\n", zipFile)
}

// Creates a Redisgo pool which allows the program
// to use and reuse concurrent redigo connections
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

// Open xml file, copy content to string
// remove string from Redis list if exists
// insert xml file as string into redis list
func insertContentRedis(fileContents *zip.File) {
	openFile, err := fileContents.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Processing xml file: %s\n", fileContents.Name)

	parsedString := ""

	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		parsedString += scanner.Text()
	}

	redisConn := pool.Get()
	defer redisConn.Close()

	_, err = redisConn.Do("LREM", "NEWS_XML", -1, parsedString)
	if err != nil {
		fmt.Println(err)
	}

	_, err = redisConn.Do("LPUSH", "NEWS_XML", parsedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Processed xml file: %s\n", fileContents.Name)

}

// Creating Redigo pool
var pool = newPool()

func main() {

  //getting base url and writing page contents to file
	baseURL := writeHTMLAndReturnRequestURL()
	output, err := os.Open("files.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

  // Using a regex to match strings with 13 numbers followed by .zip (only once)
	re := regexp.MustCompile(`\d{13}\.zip`)
	scanner := bufio.NewScanner(output)

  // slices to hold remote file paths and
  // regexed file names
	var urlSlice []string
	var fileNameSlice []string

  // Iterating over files.txt to regex file names
  // appending urls and filenames to appropriate slices
	fmt.Println("Getting file list")
	for scanner.Scan() {
		fileName := re.FindAllString(scanner.Text(), 1)

    // Making sure we have a filename to append
		if len(fileName) > 0 {
			fileNameSlice = append(fileNameSlice, fileName[0])
			urlSlice = append(urlSlice, baseURL+fileName[0])
		}
	}

	fmt.Println("Done getting file list")
	zipDir := "ZipDir"

  // Creating directory for zip files
	os.Mkdir("."+string(filepath.Separator)+zipDir, 0755)

	// start file downloads, 3 at a time
	fmt.Printf("Downloading %d files...\n", len(urlSlice))
	respch, err := grab.GetBatch(100, zipDir, urlSlice...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// start a ticker to update progress every 200ms
	t := time.NewTicker(200 * time.Millisecond)

	// monitor downloads
	completed := 0
	inProgress := 0
	responses := make([]*grab.Response, 0)
	for completed < len(urlSlice) {
		select {
		case resp := <-respch:
			// a new response has been received and has started downloading
			// (nil is received once, when the channel is closed by grab)
			if resp != nil {
				responses = append(responses, resp)
			}

		case <-t.C:
			// clear lines
			if inProgress > 0 {
				fmt.Printf("\033[%dA\033[K", inProgress)
			}

			// update completed downloads
			for i, resp := range responses {
				if resp != nil && resp.IsComplete() {
					// print final result
					if resp.Error != nil {
						fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", resp.Request.URL(), resp.Error)
					} else {
						fmt.Printf("Finished %s %d / %d bytes (%d%%)\n", resp.Filename, resp.BytesTransferred(), resp.Size, int(100*resp.Progress()))
					}

					// mark completed
					responses[i] = nil
					completed++
				}
			}

			// update downloads in progress
			inProgress = 0
			for _, resp := range responses {
				if resp != nil {
					inProgress++
					fmt.Printf("Downloading %s %d / %d bytes (%d%%)\033[K\n", resp.Filename, resp.BytesTransferred(), resp.Size, int(100*resp.Progress()))
				}
			}
		}
	}

	t.Stop()

	fmt.Printf("%d files successfully downloaded.\n", len(urlSlice))

  // Creating WaitGroup in order to read and
  // process files concurrently
	var wgReadZip sync.WaitGroup
	fmt.Println("Reading files")

	for _, value := range fileNameSlice {
		wgReadZip.Add(1)

		go readZipFiles(zipDir, value, &wgReadZip)

	}

	wgReadZip.Wait()
	fmt.Println("Done reading files")

}
