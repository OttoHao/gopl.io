package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Start to create cache...")
	cache := createCache()
	fmt.Println("Cache creation done!")

	fmt.Println("Please input a comic number:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		number := scanner.Text()
		if item, ok := cache[number]; ok {
			fmt.Println(*item)
		} else {
			fmt.Println("Can not find the comic number")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v\n", err)
		os.Exit(1)
	}
}

const XkcdURL = "https://xkcd.com/"

type Item struct {
	Title string
}

func createCache() map[string]*Item {
	const count = 100
	cache := make(map[string]*Item)
	itemChan := make(chan *Item)
	errChan := make(chan error)
	for i := 1; i <= count; i++ {
		go func(i string) {
			if item, err := query(i); err == nil {
				itemChan <- item
			} else {
				errChan <- err
			}
		}(strconv.Itoa(i))
	}
	go func() {
		for i := 1; i <= count; i++ {
			select {
			case item := <-itemChan:
				cache[strconv.Itoa(i)] = item
			case <-errChan:
				fmt.Fprintf(os.Stderr, "cache err: %d", i)
			}
		}
	}()
	return cache
}

func query(number string) (*Item, error) {
	resp, err := http.Get(XkcdURL + number + "/info.0.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}
	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, err
	}
	return &item, nil
}
