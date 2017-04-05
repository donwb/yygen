package main

import "fmt"

func main() {
	/*
		id := getARand(100)
		yakarma := getARand(20000)

		event := generateEvent()
		ua := generateUserAgent()

		fmt.Println("ID: ", id)
		fmt.Println("Yakarma: ", yakarma)
		fmt.Println(event)
		fmt.Println(ua)
	*/
	evt := NewEvent(100, 10000)

	fmt.Println("ID: ", evt.Id)
	fmt.Println("Yakarma: ", evt.YakkerID)
	fmt.Println(evt.UserAgent)
	fmt.Println(evt.EventType)

	fmt.Println(evt.CreatedUTC)

}
