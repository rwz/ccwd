package main

import (
	"fmt"
	"os"
	"strings"
)

func cwd() string {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return cwd
}

func ccwd() string {
	cwd := cwd()
	home := os.Getenv("HOME")

	var result string

	if strings.HasPrefix(cwd, home) {
		cwd = strings.Replace(cwd, home, "~", 1)
		result = ""
	} else {
		cwd = strings.TrimLeft(cwd, "/")
		result = "/"
	}

	components := strings.Split(cwd, string(os.PathSeparator))
	length := len(components) - 1

	for index, component := range components {
		if index == length {
			result += component
		} else {
			result += fmt.Sprintf("%s/", component[0:1])
		}
	}

	return result
}

func main() {
	fmt.Println(ccwd())
}
