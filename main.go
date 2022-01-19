package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Loader struct {
	Files []File `json:"files"`
}

type File struct {
	Name      string `json:"name"`
	Delimiter string `json:"delimiter"`
}

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	var loader Loader
	bc, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bc, &loader)
	if err != nil {
		panic(err)
	}

	for _, file := range loader.Files {
		f, err := os.Open(file.Name)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(bufio.NewReader(f))

		var res []string
		for scanner.Scan() {
			res = strings.Split(scanner.Text(), file.Delimiter)
			break
		}

		for i, p := range res {
			fmt.Printf("Index : %v \t| Property : %v\n", i, p)
		}
	}
}
