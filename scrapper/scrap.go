package scrapper

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scrapper struct {
	Url       string
	QueryFile string
}

func (s *Scrapper) Scrap() (string, error) {
	body, err := s.loadHTML()
	if err != nil {
		return "", err
	}
	query, err := s.getQuery()
	if err != nil {
		return "", err
	}
	result := body.Find(query).Text()
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
		return nil, err
	}
	defer resp.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	return doc, nil
}

func (s *Scrapper) getQuery() (string, error) {
	file, err := os.Open(s.QueryFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fLine := scanner.Text()
		if strings.TrimSpace(fLine) == "" {
			return "", fmt.Errorf("query file is empty")
		}
		return fLine, nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}
