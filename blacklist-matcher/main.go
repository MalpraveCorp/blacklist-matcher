package main

import (
	"context"
	"fmt"
	// "path"
	// "bufio"
	// "os"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/search"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/cloudflare/ahocorasick"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(ctx context.Context, request Request) (Response, error) {

	blacklist, err := loadWordBlacklist()
	if err != nil {
		return Response{
			StatusCode: 500,
			Body:       fmt.Sprintf("Cannot load blacklist database"),
		}, nil
	}

	word, found := findBlacklistWord(blacklist, request.Body)

	// return on first match
	if found == true {
		return Response{
			StatusCode: 400,
			Body:       fmt.Sprintf("Submission is invalid. Found black listed keyword: %s", word),
		}, nil
	}

	return Response{
		StatusCode: 200,
		Body:       "Submission is valid",
	}, nil

}

// https://github.com/cloudflare/ahocorasick
// supposedly better at matching large number of patterns at once
// but performance seem to be equivelent to standard library for
// tests with 100s of keywords and a body of text under 1000 words.
func findBlacklistWord(words []string, src string) (word string, found bool) {
	m := search.New(language.Thai, search.IgnoreCase)
	for _, v := range words {
		i, _ := m.IndexString(src, v)
		if i > 0 {
			return v, true
		}
	}
	return "", false
}

// func findBlackListWordAhocorasick(words []string, src string) (word string, found bool) {
// 	m := ahocorasick.NewStringMatcher(words)
// 	matches := m.Match([]byte(src))
// 	fmt.Println("", found)
// 	if(len(matches) > 0) {
// 		return matches, true
// 	}
// 	return "", false
// }

func loadWordBlacklist() (words []string, err error) {
	data, err := Asset("blacklist.txt")
	if err != nil {
		fmt.Println("Blacklist database could not be loaded")
	}
	// fmt.Printf(data)
	// scanner := bufio.Buffer(data)
	// scanner.Split(data)
	// for scanner.Scan() {
	// 	words = append(words, scanner.Text())
	// }
	// file.Close()

	return strings.Split(string(data), "\n"), nil
}

func main() {
	lambda.Start(Handler)
}
