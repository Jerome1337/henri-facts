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

	return facts.Facts[randomInt(0, len(facts.Facts))]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
