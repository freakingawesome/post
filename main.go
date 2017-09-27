package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	now := time.Now()

	if len(os.Args) == 2 {
		postDir := path.Join(`content`, `post`)

		if stat, err := os.Stat(postDir); err == nil && stat.IsDir() {
			title := regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(os.Args[1]), " ")
			fileName := strings.Replace(strings.ToLower(title), " ", "-", -1)

			filePath := path.Join(postDir, now.Format("2006-01-02-")+fileName+`.md`)

			lines := strings.Join([]string{
				`---`,
				fmt.Sprintf("title: %s", strings.Replace(title, `"`, `\"`, -1)),
				fmt.Sprintf("date: %s", now.Format("2006-01-02T15:04:05-07:00")),
				`---`,
			}, "\n")

			if err := ioutil.WriteFile(filePath, []byte(lines), 0644); err != nil {
				panic(err)
			}

			_ = clipboard.WriteAll(filePath)

			/*
				vim := exec.Command("vim", filePath)
				vim.Stdin = os.Stdin
				vim.Stderr = os.Stderr
				if err := vim.Run(); err != nil {
					panic(err)
				}
				fmt.Printf("done")
			*/
		} else {
			panic(fmt.Sprintf("%s is not a directory", postDir))
		}
	} else {
		printUsage()
		return
	}
}

func printUsage() {
	fmt.Println(`Usage: post "Your Title Here"`)
}
