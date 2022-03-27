package goroutine

import (
	"fmt"
	"net/http"
	"time"
)

func DoAsynchronousTask() {
	urls := []string{
		"https://twitter.com/home/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.wikipedia.org/",
		"https://github.com/",
	}
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(urls); i++ {
		go checkUrl(urls[i])
		go printNumbers(numbers[i])
	}
	time.Sleep(5 * time.Second)

}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down!")
		return
	}
	fmt.Println(url, "is up and running.")
}

func printNumbers(number int) {
	fmt.Println(number)
}
