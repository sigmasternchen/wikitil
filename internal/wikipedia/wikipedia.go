package wikipedia

import (
	"errors"
	"fmt"
	"strings"
	. "wikitil/internal/config"
)

const retries = 5

var illegalDescriptionParts = []string{
	"disambiguation page",
	"Wikimedia list article",
	"Begriffsklärungsseite",
	"Wikimedia-Liste",
}

func Get(config Config) (PageInfo, error){
	retryLoop:
	for i := 0; i < retries; i++ {
		id, err := queryRandom(config.BaseUrl)
		if err != nil {
			return PageInfo{}, err
		}

		info, err := queryInfo(config.BaseUrl, id)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Retrying...")
			fmt.Println()
			continue
		}

		for _, part := range illegalDescriptionParts {
			if strings.Contains(info.Description, part) {
				fmt.Println("illegal description: " + info.Description)
				i-- // illegal descriptions don't count towards retry limit
				continue retryLoop
			}
		}

		return info, err
	}
	return PageInfo{}, errors.New("retries exceeded")
}

func Format(info PageInfo) string {
	var builder strings.Builder
	builder.WriteString(info.Title)
	builder.WriteString(":\n")
	builder.WriteString(info.Description)
	builder.WriteString("\n\n")
	builder.WriteString(info.URL)
	return builder.String()
}