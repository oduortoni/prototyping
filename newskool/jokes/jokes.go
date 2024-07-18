package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Id        int    `json:"id"`
}

func Jokes(response http.ResponseWriter, request *http.Request) {

	joke := Joke{
		Type:      "general",
		Setup:     "Why is peter pan always flying?",
		Punchline: "Because he neverlands",
		Id:        10932,
	}

	// convert a golang struct to json
	jokeson, _ := json.Marshal(joke)

	// inform tthe browser of the json data type to expect
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprint(response, string(jokeson))
}
