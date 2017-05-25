// GitHubのissueの作成、読み出し、更新、クローズの機能を提供する。
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"../github"
)

func main() {
	var createRequest github.IssueRequest
	input := bufio.NewScanner(os.Stdin)
	username, token, err := InputAuth(input)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// println("Input action: 'create' or 'read' or 'edit' or 'close'")
	// if !input.Scan() {
	// 	log.Fatal(input.Err())
	// 	os.Exit(1)
	// }
	// action := input.Text()
	// switch action {
	// case "create":
	//
	// }
	println("Input repository owner.")
	if !input.Scan() {
		log.Fatal(input.Err())
		os.Exit(1)
	}
	owner := input.Text()
	println("Input repository.")
	if !input.Scan() {
		log.Fatal(input.Err())
		os.Exit(1)
	}
	repo := input.Text()
	println("Input issue title for create.")
	if !input.Scan() {
		log.Fatal(input.Err())
		os.Exit(1)
	}
	createRequest.Title = input.Text()
	println("Input issue body for create.")
	if !input.Scan() {
		log.Fatal(input.Err())
		os.Exit(1)
	}
	createRequest.Body = input.Text()
	resp, err := github.CreateIssues(username, token, owner, repo, &createRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", resp.URL)
}

func InputAuth(input *bufio.Scanner) (string, string, error) {
	print("Input your username.")
	if !input.Scan() {
		log.Fatal(input.Err())
		os.Exit(1)
	}
	username := input.Text()
	print("Input your access token.")
	if !input.Scan() {
		return "", "", input.Err()
	}
	token := input.Text()
	return username, token, nil
}
