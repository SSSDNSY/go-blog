package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestIdeaCode(t *testing.T) {
	resp, _ := http.Get("http://idea.medeming.com/jet/images/jihuoma.txt")
	fmt.Println(resp)

}

func Test4Code(t *testing.T) {
	url := "http://idea.medeming.com/jet/images/jihuoma.txt"
	var reader io.Reader
	r, _ := http.NewRequest(http.MethodGet, url, reader)
	fmt.Println(r)
	fmt.Println(reader)
}
