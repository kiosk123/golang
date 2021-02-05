package main

import (
	"fmt"
	"time"

	"github.com/kiosk123/golang/banking"
	"github.com/kiosk123/golang/dict"
	"github.com/kiosk123/golang/urlcheker"
)

func main() {
	// -- loop --
	var numbers = []int{}
	for i := 1; i < 11; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	for index, number := range numbers {
		fmt.Println(index, ":", number)
	}

	// index not used
	for _, number := range numbers {
		fmt.Println(number)
	}

	// -- loop --

	// -- struct --
	account := banking.NewAccount("nico")
	fmt.Println(account)

	account.Deposit(100)
	fmt.Println(account)

	fmt.Println(account.Balance())
	err := account.Withdraw(200)

	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
	fmt.Println(account)

	// -- struct --
	// -- Dictionary --

	dictionary := dict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(definition)
	} else {
		fmt.Println(err)
	}

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(hello)
	}

	err = dictionary.Update("rank", "third")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Delete("hello")
	if err != nil {
		fmt.Println(err)
	}
	// -- Dictionary --

	// -- URL checker --
	urls := []string{
		"https://www.naver.com",
		"https://www.daum.net/",
		"https://www.nate.com/",
		"https://www.saramin.co.kr/zf_user/",
		"https://github.com/",
	}

	var results = make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err = urlcheker.HitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println("Checking :", url, "Result :", result)
	}

	// -- URL checker --

	// -- go routine test --
	// go 루틴은 메인함수가 실행될 돌안만 유효 - 메인 스레드는 무조건 하나 살아있어야함
	go sexyCount("nico")
	sexyCount("flynn")

	fmt.Println("--------------------------------")

	go sexyCount("cat")
	go sexyCount("dog")
	time.Sleep(time.Second * 10) // main이 먼저 종료되는 문제 예방
	// -- go routine test --
}

// -- go routine --
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep((time.Second))
	}
}
