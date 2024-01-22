package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	err := downloadSite(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadSite(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to fetch the page. Status code: %d", resp.StatusCode)
	}

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return nil // Игнорировать EOF как ошибку
			} else {
				return err
			}
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				href := getHref(token)
				if href != "" {
					fmt.Println("Downloading:", href)
					err := downloadFile(href)
					if err != nil {
						fmt.Println("Error downloading", href, ":", err)
					}
				}
			}
		}
	}
}

func getHref(token html.Token) string {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func downloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to fetch the file. Status code: %d", resp.StatusCode)
	}

	filename := getFilename(url)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Downloading successful:", filename)
	return nil
}

func getFilename(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
