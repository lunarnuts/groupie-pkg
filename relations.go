package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func Relations() ([]Relation, int) {
	api, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Status: 502,BG, Error connecting to API/relations")
		return nil, 502
	}
	if api.StatusCode != 200 {
		fmt.Println("Status: 502,BG, Error connecting to API/relations")
		return nil, 502
	} else {
		fmt.Println("Status: 200,OK, Connected to API/relations")
	}
	defer api.Body.Close()
	var relations []Relation
	body, err := ioutil.ReadAll(api.Body)
	err = json.Unmarshal(body[9:len(body)-2], &relations)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 500
	}
	return relations, 200
}
