package services

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Summary(c *gin.Context) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	counts := GetBeef(string(body))
	c.JSON(http.StatusOK, gin.H{"beef": counts})
}

func GetBeef(rawWords string) map[string]int {
	wordsToCount := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}
	counts := make(map[string]int)
	lowerRawWords := strings.ToLower(rawWords)

	for _, word := range wordsToCount {
		counts[word] = strings.Count(lowerRawWords, word)
	}
	return counts
}
