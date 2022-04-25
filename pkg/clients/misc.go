package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}
*/

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func GenerateIpList() map[string]string {

	ipPrefix := os.Getenv("IP_PREFIX")
	if ipPrefix == "" {
		log.Printf("IP_PREFIX env not set correctly")
		os.Exit(1)
	}
	ipStartStr := os.Getenv("IP_START")
	if ipStartStr == "" {
		log.Printf("IP_START env not set correctly")
		os.Exit(1)
	}
	ipStart, _ := strconv.Atoi(ipStartStr)

	ipEndStr := os.Getenv("IP_END")
	if ipEndStr == "" {
		log.Printf("IP_END env not set correctly")
		os.Exit(1)
	}
	ipEnd, _ := strconv.Atoi(ipEndStr)

	var ipListMap = make(map[string]string)

	var baseUrl string

	// Generate a list of base urls with alive shellys
	log.Printf("Bulding IP List from %d to %d\n", ipStart, ipEnd)
	for octectFour := ipStart; octectFour <= ipEnd; octectFour++ {

		baseUrl = fmt.Sprintf("http://%s%d", ipPrefix, octectFour)
		err := ShellyCheck(baseUrl)

		if err == nil {
			log.Printf("Checking %d/%d shelly\n", octectFour, ipEnd)

			name, err := GetName(baseUrl)
			if err != nil {
				continue
			}

			ipListMap[name] = baseUrl

		} else {
			log.Printf("Checking %d/%d unknown\n", octectFour, ipEnd)
		}

	}
	return ipListMap
}
