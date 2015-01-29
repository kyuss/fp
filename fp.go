package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	FILEPICKER_URI   = "https://www.filepicker.io/api/store/S3"
	STORE_PARAM_NAME = "fileUpload"
)

var (
	IS_SILENT bool
)

type Config struct {
	Filepicker struct {
		ApiKey string
	}
}

func init() {
	flag.BoolVar(&IS_SILENT, "s", false, "Silent (Don't show spinner)")
	flag.Parse()

}

func main() {
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filePath := os.Args[len(os.Args)-1]

	err = store(filePath, config.Filepicker.ApiKey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func escapeFilename(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", errors.New("File not found")
	}

	return url.QueryEscape(filepath.Base(filename)), nil
}


func store(filePath string, apikey string) error {
	filename, err := escapeFilename(filePath)

	url := fmt.Sprintf("%s?key=%s&filename=%s", FILEPICKER_URI, apikey, filename)

	request, err := storeRequest(url, filePath)

	if err != nil {
		return errors.New("Error reading file")
	}

	var spinner *spinner.Spinner

	if !IS_SILENT {
		spinner = startSpinner()
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result Result
	unmarshall(body.Bytes(), &result)
	clipboard.WriteAll(result.Url)

	fmt.Printf("%s [in clipboard]\n", result.Url)


	if !IS_SILENT {
		spinner.Stop()
	}

	return nil


}

func startSpinner() *spinner.Spinner {
	s := *spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Start()
	return &s
}

func storeRequest(uri string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(STORE_PARAM_NAME, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	return request, nil
}
