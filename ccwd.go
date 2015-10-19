package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func cwd() string {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return cwd
}

func home() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}

func main() {
	cwd := cwd()
	home := home()

	if strings.HasPrefix(cwd, home) {
		cwd = strings.Replace(cwd, home, "~", 1)
	}

	components := strings.Split(cwd, string(filepath.Separator))

	for index, component := range components {
		if component == "" {
			fmt.Print("/")
		} else if index == len(components)-1 {
			fmt.Println(component)
		} else {
			fmt.Printf("%c/", []rune(component)[0])
		}
	}
}
