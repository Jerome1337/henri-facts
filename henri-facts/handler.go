package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type facts struct {
	Facts []string `json:"facts"`
}

type response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func Handle(req []byte) string {
	rand.Seed(time.Now().UTC().UnixNano())

	jsonFile, err := os.Open("facts.json")

	if err != nil {
		fmt.Println("Failed to open JSON file:", err)
	}

	defer jsonFile.Close()

	byteArray, _ := ioutil.ReadAll(jsonFile)

	var facts facts

	if err := json.Unmarshal(byteArray, &facts); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
	}

	marshalJSON, _ := json.Marshal(
		response{
			"in_channel",
			facts.Facts[randomInt(0, len(facts.Facts))],
		},
	)

	return string(marshalJSON)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
