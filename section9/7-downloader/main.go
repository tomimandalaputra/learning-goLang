package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(url, destDir string) error {
	filename := filepath.Base(url)
	fp := filepath.Join(destDir, filename)

	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading", url)

	res, err := http.Get(url)
	if err != nil {
		_ = os.Remove(fp) // remove file
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		_ = os.Remove(fp) // remove file
		return fmt.Errorf("bad status: %s", res.Status)
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Donwload %s took %s\n", filename, time.Since(time.Now()))

	return nil
}

func SequentialDownloader(urls []string, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for _, url := range urls {
		if err := DownloadFile(url, destDir); err != nil {
			fmt.Println("Error downloading", url, err)
			continue
		}
	}

	fmt.Printf("Download %s took %s\n", urls, time.Since(time.Now()))
	return nil
}

type Result struct {
	URL      string
	Filename string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	results := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}
			defer func() { <-limiter }()

			start := time.Now()
			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			out, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				_ = os.Remove(filePath) // remove file
				results <- Result{URL: url, Error: err}
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				_ = os.Remove(filePath) // remove file
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", resp.Status)}
				return
			}
			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			timeSince := time.Since(start)
			results <- Result{URL: url, Filename: filename, Size: size, Duration: timeSince, Error: nil}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloaded %s (%d bytes) in %s\n", result.Filename, result.Size, result.Duration)
		}
	}

	startedSince := time.Since(start)
	fmt.Printf("All downloads completed in %s, Total: %d bytes\n", startedSince, totalSize)
	if len(errors) > 0 {
		return fmt.Errorf("errors downloading: %+v", errors)
	}

	return nil
}

func main() {

	urls := []string{
		"https://pdfobject.com/pdf/sample.pdf",
		"https://ontheline.trincoll.edu/images/bookdown/sample-local-pdf.pdf",
	}

	// err := DownloadFile(url, "./section9/7-downloader/")
	// err := SequentialDownloader(urls, "./section9/7-downloader/")
	err := ConcurrentDownloader(urls, "./section9/7-downloader/", len(urls))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Finished")
}
