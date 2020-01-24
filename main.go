package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xanzy/go-gitlab"
)

func main() {
	if len(os.Args) < 2 {
		println("Please specific the output file")
		os.Exit(1)
	}
	if os.Getenv("GITLAB_ACCESS_TOKEN") == "" {
		println("Didn't find GITLAB_ACCESS_TOKEN env")
		os.Exit(1)
	}
	git := gitlab.NewClient(nil, os.Getenv("GITLAB_ACCESS_TOKEN"))
	variables, _, error := git.ProjectVariables.ListVariables(16074126, nil)
	if error != nil {
		fmt.Printf("Unable to fetch the variables:  %v\n", error)
		os.Exit(1)
	}
	writeEnvToFile(variables)
}

func writeEnvToFile(envs []*gitlab.ProjectVariable) {
	err := os.MkdirAll(filepath.Dir(os.Args[1]), 0777)
	if err != nil {
		fmt.Printf("Error when create environment folder: %v\n", err)
		os.Exit(1)
	}
	file, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	for _, env := range envs {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", env.Key, env.Value))
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}
	file.Close()
	os.Exit(0)
}
