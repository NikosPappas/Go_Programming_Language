package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const slugLength = 6
const dataFile = "url_data.json" // File to store our data

var urlMap = make(map[string]string)

func main() {
	// Load existing data from the file on start
	loadData()

	longURLPtr := flag.String("long", "", "The long URL to shorten")
	shortURLPtr := flag.String("short", "", "The short URL to resolve")
	flag.Parse()

	if *longURLPtr != "" {
		longURL := *longURLPtr
		shortURL := generateUniqueShortURL()
		urlMap[shortURL] = longURL
		fmt.Printf("Long URL: %s\n", longURL)
		fmt.Printf("Generated Short URL: %s\n", shortURL)
		saveData() // Save the updated map to file
	} else if *shortURLPtr != "" {
		shortURL := *shortURLPtr
		longURL, found := urlMap[shortURL]
		if found {
			fmt.Printf("Short URL: %s resolves to Long URL: %s\n", shortURL, longURL)
		} else {
			fmt.Printf("Short URL: %s not found.\n", shortURL)
		}
	} else {
		fmt.Println("Please provide a long URL using the -long flag or short url using the -short flag.")
	}
}

func generateUniqueShortURL() string {
	for {
		shortURL := generateShortURL()
		if _, exists := urlMap[shortURL]; !exists {
			return shortURL
		}
	}
}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, slugLength)
	for i := 0; i < slugLength; i++ {
		shortURL[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortURL)
}

func loadData() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) { // If file doesn't exist, it's okay
			return
		}
		fmt.Println("Error opening data file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading data file:", err)
		return
	}

	if len(data) > 0 { // Only unmarshal if there's data
		err = json.Unmarshal(data, &urlMap)
		if err != nil {
			fmt.Println("Error unmarshaling data:", err)
		}
	}
}

func saveData() {
	data, err := json.Marshal(urlMap)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	err = ioutil.WriteFile(dataFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing to data file:", err)
	}
}
