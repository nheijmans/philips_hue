package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Device is a struct containing the data
type Device struct {
	Name  string
	State Status
}

// Status is a struct containing the info of devices
type Status struct {
	On          bool
	Reachable   bool
	Temperature int
	Lightlevel  int
}

func main() {
	hueIP := flag.String("ip", "192.168.1.2", "Your Hue IP address")
	apiKey := flag.String("key", "yourapikey", "Provide your Philips Hue API key")
	lightID := flag.String("lid", "", "Provide device ID to read out")
	sensorID := flag.String("sid", "", "Provide device ID to read out")
	rawData := flag.Bool("raw", false, "Set this if you want the raw data")

	flag.Parse()

	var data Device
	switch {
	case *lightID != "":
		url := fmt.Sprintf("http://%s/api/%s/lights/%s", *hueIP, *apiKey, *lightID)
		content := contentFromServer(url)
		data = dataFromJSON(content)
		fmt.Printf("Name: %v - On: %v - Reachable: %v\n", data.Name, data.State.On, data.State.Reachable)
	case *sensorID != "":
		url := fmt.Sprintf("http://%s/api/%s/sensors/%s", *hueIP, *apiKey, *sensorID)
		content := contentFromServer(url)
		data = dataFromJSON(content)
		fmt.Printf("Name: %v - Temperature: %v - Brightness: %v\n", data.Name, data.State.Temperature, data.State.Lightlevel)
	case *rawData == true:
		fmt.Println("No options given.")
		url := fmt.Sprintf("http://%s/api/%s", *hueIP, *apiKey)
		content := contentFromServer(url)
		fmt.Println(content)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func contentFromServer(url string) string {
	// retrieve content from a website
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	// extract the content of the body
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func dataFromJSON(content string) Device {
	content = "[" + content + "]"

	// decode the JSON content from the web request
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	// create an object of the Tour structure
	var device Device
	for decoder.More() {
		//parse the json content and it only gets the values in the struct
		err := decoder.Decode(&device)
		checkError(err)
	}

	return device
}
