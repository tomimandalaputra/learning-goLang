package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	pathDirWindows := filepath.Join("C:", "Users", "Documents")
	fmt.Println(pathDirWindows)

	pathGithubWorkflows := filepath.Join(".github", "workflows", "deploy.yml")
	fmt.Println(pathGithubWorkflows)

	baseFile := filepath.Base(pathGithubWorkflows)
	extFile := filepath.Ext(pathGithubWorkflows)
	fmt.Println(baseFile)
	fmt.Println(extFile)

	dirtyDir := "./users/./dir/../other_dir/./file.txt"
	fmt.Println(filepath.Clean(dirtyDir))
}
