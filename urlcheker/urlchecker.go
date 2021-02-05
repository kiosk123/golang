package urlcheker

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New(("Request failed"))

// hit url
func HitURL(url string) error {
	fmt.Println("checking url :", url)
	resq, err := http.Get(url)
	if err != nil || resq.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
