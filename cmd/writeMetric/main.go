package main

import (
	"log"
	"os"
	"time"

	"github.com/yggdrion/shelly-cloudwatch/pkg/clients"
)

var (
	startTime   time.Time
	nowTime     time.Time
	diffTime    time.Duration
	diffSeconds float64
)

func main() {

	startTime = time.Now()

	// TODO
	// make ipMin, ipMax and cidr a ENV

	// Check if env is set, else exit
	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	if awsAccessKeyId == "" {
		log.Printf("AWS_ACCESS_KEY_ID env not set correctly")
		os.Exit(1)
	}
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if awsSecretAccessKey == "" {
		log.Printf("AWS_SECRET_ACCESS_KEY env not set correctly")
		os.Exit(1)
	}

	ipList := clients.GenerateIpList()

	// endless loop
	for {

		log.Printf("Collecting metrics\n")

		// iterate all alive shellys
		for name, address := range ipList {
			watt, err := clients.GetWatt(address)
			if err != nil {
				continue
			}

			log.Printf("%s\t%s\t%f\n", address, name, watt)

			go clients.PutCloudwatchMetric(name, watt)
		}

		// check if we need to generate new list
		nowTime = time.Now()

		diffTime = nowTime.Sub(startTime)
		diffSeconds = diffTime.Seconds()

		if diffSeconds >= 1800 {
			log.Printf("Need to rebuild list")
			ipList = clients.GenerateIpList()
			startTime = time.Now()
		} else {
			log.Printf("Difftime is: %f", diffSeconds)
		}

		time.Sleep(60 * time.Second)
	}

}
