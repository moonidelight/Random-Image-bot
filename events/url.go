package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lab4/repository"
	"lab4/user"
	"log"
	"net/http"
	"os"
	"time"
)

type Photo struct {
	Urls struct {
		Image string `json:"regular"`
	} `json:"urls"`
}

var d *user.Data
var repo *repository.DataRepository

func URL() (string, error) {
	key := os.Getenv("UNSPLASH_ACCESS_KEY")
	UnsplashUrl := fmt.Sprintf("https://api.unsplash.com/photos/random/?client_id=%s", key)
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}
	req, err := http.NewRequest(http.MethodGet, UnsplashUrl, nil)
	if err != nil {
		//log.Fatal(err)
		return "", err
	}
	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		//log.Fatal(getErr)
		return "", getErr
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		//return "", readErr
	}
	var image Photo
	jsonErr := json.Unmarshal(body, &image)
	if jsonErr != nil {
		return "", jsonErr
	}
	out, err := json.Marshal(image)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(string(out)), &data)
	if err != nil {
		return "", err
	}
	url := data["urls"].(map[string]interface{})["regular"].(string)

	return url, nil
}
