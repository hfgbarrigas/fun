package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type uaCount struct {
	name  string
	count int
}

func main() {
	data := make(map[string]int)
	scanner, f, err := getFileReader("data.txt")
	checkErr(err)

	// process each line
	for scanner.Scan() {
		ua := parseUserAgent(scanner.Text())
		if ua != "" {
			data[ua]++
		}
	}

	results := processData(data)

	// print top 10
	top10 := results[len(results)-10:]
	for _, v := range top10 {
		fmt.Printf("%d %s", v.count, v.name)
	}

	// clean up
	err = f.Close()
	checkErr(err)
}

// gets a buffered scanner from the given file path
// it's up to the called to close the file once its no longer needed
func getFileReader(file string) (*bufio.Scanner, *os.File, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		return nil, nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewScanner(bufio.NewReader(f)), f, nil
}

// returns a sorted slice of user agents, based on their count in ascending order
func processData(data map[string]int) []*uaCount {
	results := make([]*uaCount, len(data))
	for k, v := range data {
		results = append(results, &uaCount{
			name:  k,
			count: v,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].count < results[j].count
	})

	return results
}

// looks at the text line and returns the user agent
// if not user agent is found, empty is returned
func parseUserAgent(line string) string {
	tokens := strings.Split(line, "\"")

	if len(tokens) >= 6 && tokens[5] != "-" {
		return tokens[5]
	}

	return ""
}

// if the error is not nil, prints it and panics
func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered an error: %v", err)
		panic(err)
	}
}
