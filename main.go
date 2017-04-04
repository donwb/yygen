package main

import "fmt"
import "math/rand"
import "time"

func main() {

	var id = generateID(100)

	var event = generateEvent()

	fmt.Println("ID: ", id)
	fmt.Println(event)
}

func generateID(seed int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(seed)
}

func generateEvent() string {
	event := "none"

	b := [...]string{"getMessages", "postMessage", "upvote", "downvote", "report", "delete", "hot"}
	weights := [...]int{5, 10, 14, 18, 22, 24, 25}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	idx := r.Intn(25)

	fmt.Println("index: ", idx)

	previousValue := 0
	for loopIndex, weightValue := range weights {
		if idx > previousValue && idx <= weightValue {
			event = b[loopIndex]
			break
		} else {
			previousValue = weightValue
		}
	}

	return event
}
