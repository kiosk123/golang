package channel

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

// example
func Main() {
	urls := []string{
		"https://www.naver.com",
		"https://www.daum.net/",
		"https://www.nate.com/",
		"https://www.saramin.co.kr/zf_user/",
		"https://github.com/",
	}

	resultMap := make(map[string]string)
	c := make(chan result)

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		resultMap[result.url] = result.status
	}

	for url, status := range resultMap {
		fmt.Println(url, status)
	}

}

// channel에 전송만 가능하고 받기는 불가능 : send only
func hitURL(url string, c chan<- result) {
	// fmt.Println(<-result)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Fail"
	}
	c <- result{url: url, status: status}
}
