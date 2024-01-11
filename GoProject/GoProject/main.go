package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	endpoints = "https://api.open-meteo.com/v1/forecast" //?latitude=52.52&longitude=13.41&hourly=temperature_2m"
)

type Sender interface {
	Send(*WeatherData) error
}

type SMSSender struct {
	number string
}

func NewSMSSender(number string) *SMSSender {
	return &SMSSender{
		number: number,
	}
}

func (s *SMSSender) Send(data *WeatherData) error {
	fmt.Println("sending weather data to sms: ", s.number)
	return nil
}

type EmailSender struct {
	email string
}

func NewEmailSender(email string) *EmailSender {
	return &EmailSender{
		email: email,
	}
}

func (e *EmailSender) Send(data *WeatherData) error {
	fmt.Println("sending weather to email: ", e.email)
	return nil
}

type WeatherData struct { //this structure defines the response data datatypes
	Temprature float64        //`json:"temperature"`
	Hourly     map[string]any //`json:"hourly"`
}

var pollinterval = time.Second * 3

type MyPoller struct {
	closechain chan struct{}
	senders    []Sender
} //defines a new type for MyPoller

func NewPoller(senders ...Sender) *MyPoller { //it is function for structure MyPoller which takes the address of MyPollr and return its data on that address. We will always call other methods with dattype MyPoller via this methods only. Its a good practice!!
	return &MyPoller{
		closechain: make(chan struct{}),
		senders:    senders,
	}
}

func (MP *MyPoller) close() {
	close(MP.closechain)
}

func (MP *MyPoller) start() { //defines a method names start which has MyPoller data type. It is function for the structure which returns start data
	fmt.Println("Starting the poller")
	timeWatch := time.NewTicker(pollinterval)
free:
	for {
		select {
		case <-timeWatch.C: //this for function creates infinte loop that listens from event timewatch.
			data, err := getWeatherDetails(52.4233, 13.2221)
			//we are requesting the temprature and time based on the values of latitude and longitude
			if err != nil {
				log.Fatal(err)
			}
			if err := MP.handledata(data); err != nil {
				log.Fatal(err)
			}
		case <-MP.closechain:
			break free
		}
	}
	fmt.Println("Poller is Stopped")
}

func (mp *MyPoller) handledata(Wdata *WeatherData) error {
	for _, s := range mp.senders {
		if err := s.Send(Wdata); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("")
	fmt.Println("------------")
	fmt.Println("The Weather JSON data : ")
	fmt.Println("------------")
	fmt.Println(Wdata)
	return nil
}

func getWeatherDetails(latitude float64, longitude float64) (*WeatherData, error /* We are returning data and error as these values are stored in varibales passed in main function */) {
	uri := fmt.Sprintf("%s?latitude=%.2f&longitude=%.2f&hourly=temperature_2m", endpoints, latitude, longitude) //here we are making the endpoints url via code by letting user pass the latitude and longitude value
	fmt.Println("------------")
	fmt.Println(uri)
	fmt.Println("------------")

	response, err := http.Get(uri) //it will get the json fied for the temprature and time on basis of specified longitude and latittude url
	if err != nil {                //if error is not equal to null --> if eror is there !!
		return nil, err
	}
	defer response.Body.Close()
	//closing the response function after it is completed

	var data WeatherData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}
	//coverts the data into the json format

	return &data, nil
}

func main() {
	smsSender := NewSMSSender("......94030")
	emailSender := NewEmailSender("akshitmittal20@gmail.com")
	ThePoller := NewPoller(smsSender, emailSender)
	ThePoller.start()
}
