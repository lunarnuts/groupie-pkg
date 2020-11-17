package groupie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
}

func Artists() ([]Artist, int) {
	api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Status: 502,BG, Error connecting to API/artists")

		return nil, 502
	}
	if api.StatusCode != 200 {
		fmt.Println("Status: 500,IS, Error connecting to API/artists")
		return nil, 500
	} else {
		fmt.Println("Status: 200,OK, Connected to API/artists")
	}
	defer api.Body.Close()
	var artists []Artist
	body, err := ioutil.ReadAll(api.Body)
	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 500
	}
	//var relate []Relation
	relate, er := Relations()
	if er != 200 {
		return nil, er
	}
	for index := range artists {
		artists[index].Relations = MapToString(relate[index].DatesLocations)
	}
	return artists, 200
}

func MapToString(m map[string][]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		buf := ""
		for _, v := range value {
			buf += ",\n" + v
		}
		buf = buf[2:]
		fmt.Fprintf(b, "{%s: [%s]}\n", key, buf)
	}
	return b.String()
}
