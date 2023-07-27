package tests

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

func GenerateRequest() map[string]int {
	var req map[string]int
	req = map[string]int{}
	randomizer := rand.New(rand.NewSource(10))
	req["wind"]=randomizer.Intn(100)
	req["water"]=randomizer.Intn(100)
	return req
}

func GetStatus(jenis string,level float64) (status string) {
	if strings.ToLower(jenis) == "water" {
		if level < 5 {
			status = "aman"
		}else if level >=6 && level <=8 {
			status = "siaga"
		}else if level >8 {
			status = "bahaya"
		}
	}else if strings.ToLower(jenis) == "wind" {
		if level < 6 {
			status = "aman"
		}else if level >=7 && level <=15 {
			status = "siaga"
		}else if level > 15 {
			status = "bahaya"
		}
	}else{
		return ""
	}

	return "status "+jenis+" : "+status
}

func PostReq(){
	req:=GenerateRequest()
	reqString,_ := json.Marshal(req)
	client := resty.New()

	resp,_ := client.R().
	SetHeader("Content-Type", "application/json").
	SetBody([]byte(reqString)).
	Post("https://jsonplaceholder.typicode.com/posts")

	response := make(map[string]interface{})
	json.Unmarshal(resp.Body(),&response)
	fmt.Println(resp)
	for i, v := range response {
		fmt.Println(GetStatus(i,v.(float64)))
	}
}

func TestChallenge1(t *testing.T) {
	ticker := time.NewTicker(15 * time.Second)

	// Creating channel using make
	tickerChan := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChan:
				return
			// interval task
			case tm := <-ticker.C:
				fmt.Println("The Current time is: ", tm)
				PostReq()
			}
		}
	}()

	// Calling Sleep() method
	time.Sleep(60 * time.Second)

	// Calling Stop() method
	ticker.Stop()

	// Setting the value of channel
	tickerChan <- true

	// Printed when the ticker is turned off
	fmt.Println("Ticker is turned off!")
}