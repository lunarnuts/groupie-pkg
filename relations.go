package groupie

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func Relations() ([]Relation, error) {
	api, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Status: 502,BG, Error connecting to API/relations")
		return nil, errors.New("Bad Gateway")
	}
	if api.StatusCode != 200 {
		fmt.Println("Status: 502,BG, Error connecting to API/relations")
		return nil, errors.New("Bad Gateway")
	} else {
		fmt.Println("Status: 200,OK, Connected to API/relations")
	}
	defer api.Body.Close()
	var relations []Relation
	body, err := ioutil.ReadAll(api.Body)
	err = json.Unmarshal(body[9:len(body)-2], &relations)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return relations, nil
}
