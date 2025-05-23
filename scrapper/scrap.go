package scrapper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scrapper struct {
	Url      string
	Selector string
}

func (s *Scrapper) Scrap() (string, error) {
	body, err := s.loadHTML()
	if err != nil {
		return "", err
	}
	result := body.Find(s.Selector).Text()
	if strings.TrimSpace(result) == "" {
		return "", fmt.Errorf("failed to find in the document %v", s.Url)
	}
	return result, nil
}

func (s *Scrapper) loadHTML() (*goquery.Document, error) {
	resp, err := http.Get(s.Url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}
	defer resp.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	return doc, nil
}
