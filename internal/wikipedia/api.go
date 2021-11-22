package wikipedia

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type PageInfo struct {
	Title string
	Description string
	URL string
}

type infoResponse struct {
	Query struct {
		Pages map[string]struct{
			Title string `json:"title"`
			FullURL string `json:"fullurl"`
			Terms struct {
				Description []string `json:"description"`
			} `json:"terms"`
		} `json:"pages"`
	} `json:"query"`
}

type randomReponse struct {
	Query struct {
		Random []struct {
			ID int64 `json:"id"`
		} `json:"random"`
	} `json:"query"`
}

const baseURL = "https://de.wikipedia.org"

var noDescription = errors.New("no description found")
var noPageInfo = errors.New("no page info found")

func request(params map[string]string) ([]byte, error) {
	builder := strings.Builder{}
	builder.WriteString(baseURL)
	builder.WriteString("/w/api.php?")

	for key, value := range params {
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(value)
		builder.WriteString("&")
	}

	builder.WriteString("format=json")

	response, err := http.Get(builder.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func responseToPageInfo(response infoResponse) (PageInfo, error) {
	for _, page := range response.Query.Pages {
		if len(page.Terms.Description) < 1 {
			return PageInfo{}, noDescription
		}
		return PageInfo{
			Title:       page.Title,
			Description: page.Terms.Description[0],
			URL:         page.FullURL,
		}, nil
	}

	return PageInfo{}, noPageInfo
}

func queryInfo(id int64) (PageInfo, error) {
	params := map[string]string {
		"action": "query",
		"pageids": strconv.FormatInt(id, 10),
		"prop": "info|pageterms",
		"inprop": "url",
	}

	content, err := request(params)
	if err != nil {
		return PageInfo{}, err
	}

	var response infoResponse
	err = json.Unmarshal(content, &response)
	if err != nil {
		return PageInfo{}, err
	}

	return responseToPageInfo(response)
}

func queryRandom() (int64, error) {
	params := map[string]string {
		"action": "query",
		"list": "random",
		"rnnamespace": "0",
	}

	content, err := request(params)
	if err != nil {
		return 0, err
	}

	var response randomReponse
	err = json.Unmarshal(content, &response)
	if err != nil {
		return 0, err
	}

	if len(response.Query.Random) < 1 {
		return 0, errors.New("could not get random result")
	}

	id := response.Query.Random[0].ID

	if id == 0 {
		return 0, errors.New("no page id in result")
	}

	return id, nil
}
