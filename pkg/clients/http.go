package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/yggdrion/shelly-cloudwatch/pkg/structs"
)

func GetWatt(baseUrl string) (float64, error) {

	// get current usage of the shelly

	shellyUrl := fmt.Sprintf("%s/status", baseUrl)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", shellyUrl, nil)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}

	responseObject := structs.ShellyStatusStruct{}
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return 0.0, err
	}

	return responseObject.Meters[0].Power, nil

}

func GetName(baseUrl string) (string, error) {
	// get the name of the shelly

	shellyUrl := fmt.Sprintf("%s/settings", baseUrl)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", shellyUrl, nil)
	if err != nil {
		log.Println(err)
		return "unknown", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "unknown", err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "unknown", err
	}

	responseObject := structs.ShellySettingsStruct{}
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return "unknown", err
	}

	return responseObject.Name, nil

}

func ShellyCheck(url string) error {

	// Returns nil if its no shelly with valid json output
	// TODO search json for shelly specific string

	c := &http.Client{
		Timeout: 3 * time.Second,
	}

	webPage := fmt.Sprintf("%s/status", url)
	resp, err := c.Get(webPage)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	bString := string(b)

	if isJSON(bString) {
		return nil
	}

	return errors.New("error")

}
