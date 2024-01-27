package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("shakesearch available at http://localhost:%s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := searcher.Search(query[0])

		// check for empty results
		// if len(results) < 1 {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	w.Write([]byte("No match found for query"))
		// 	return
		// }

		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	return nil
}

func (s *Searcher) Search(query string) []string {
	if len(s.CompleteWorks) == 0 || !strings.Contains(s.CompleteWorks, query)  {
        return []string{} // Handle edge cases 
    }

	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	results := []string{}

	// set maximum of 20 results based on drunk test
	const maxResults = 20
	
	// Ensure indices are within bounds
	for _, idx := range idxs {
        start := idx - 250
        if start < 0 {
            start = 0
        }

        end := idx + 250
        if end > len(s.CompleteWorks) {
            end = len(s.CompleteWorks)
        }

        results = append(results, s.CompleteWorks[start:end])

        // Break out of the loop after appending 20 results
        if len(results) == maxResults {
            break
        }
	}
	return results
}

