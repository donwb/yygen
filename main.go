package main

import "fmt"
import "math/rand"
import "time"

func main() {

	id := getARand(100)
	yakarma := getARand(20000)

	event := generateEvent()
	ua := generateUserAgent()

	fmt.Println("ID: ", id)
	fmt.Println("Yakarma: ", yakarma)
	fmt.Println(event)
	fmt.Println(ua)
}

func generateEvent() string {
	event := "none"

	b := []string{"getMessages", "postMessage", "upvote", "downvote", "report", "delete", "hot"}
	weights := []int{5, 10, 14, 18, 22, 24, 25}

	event = getARandomThing(b, weights)

	return event
}

func generateUserAgent() string {

	agents := []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 6P Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586"}

	weights := []int{5, 10, 12, 16, 20}

	userAgent := getARandomThing(agents, weights)

	return userAgent
}

func getARandomThing(data []string, weights []int) string {
	found := "none"

	idx := getARand(weights[len(weights)-1])

	previousValue := -1
	for loopIndex, weightValue := range weights {
		if idx > previousValue && idx <= weightValue {
			found = data[loopIndex]
			break
		} else {
			previousValue = weightValue
		}
	}

	return found
}

func getARand(seed int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(seed)
}
