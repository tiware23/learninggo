package main

import (
	"fmt"
	"encoding/json"
)


type vehicle interface {
	start() string
}


type Car struct {
	CarName string `json:"car"`
	CarYear int `json:"year"`
}


func (c *Car) CarAccelerate() {
	c.CarName = "Fusion"
	//fmt.Printf("The Car %s", c.CarName)
}


func(c Car) start() string {
	return "startou"
}


func exemplo(car vehicle) {
	
}

func main() {

	car2 := Car{
		CarName: "Ford",
		CarYear: 2021,
	} 
	
	exemplo(car2)
	

	j := []byte(`{"car": "BMW", "year": 2020}`)
	
	var car Car
	
	json.Unmarshal(j, &car)
	
	fmt.Println(car)
	
	car1 := Car{
		CarName: "Ferrari",
		CarYear: 2021,
	}
	
	result, _ := json.Marshal(car1)
	
	fmt.Println(string(result))
}
