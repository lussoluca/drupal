package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/lussoluca/drupal/project"
)

func main() {
	projects := make(map[string]project.Project)
	file := os.Args[1:]
	if len(file) == 0 {
		fmt.Print("At least one file is required")
		return
	}

	f, err := os.Open(file[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	re := regexp.MustCompile("(.*),(.*)\\s\\((.*)\\),?(.*)?")
	for input.Scan() {
		submatch := re.FindStringSubmatch(input.Text())
		project := project.New(submatch[1], submatch[2], submatch[3], submatch[4])
		projects[submatch[3]] = project
		if !project.IsCore() {
			fmt.Printf("%s -> %s -> %v\n", project.ShortName, project.APIVersion, project.Terms)
		}
	}
}
