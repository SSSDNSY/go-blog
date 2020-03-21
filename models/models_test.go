package models

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestAddCateGory(t *testing.T) {
	AddCateGory("name")
}

func TestCookieSet(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ck := &http.Cookie{
			Name:       "myCookie",
			Value:      "myasdlkfjaslk;dfasjlk;dfj",
			Path:       "/",
			Domain:     "localhost",
			Expires:    time.Time{},
			RawExpires: "120",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(writer, ck)
		ck2, err := request.Cookie("myCookie")
		if nil != err {
			io.WriteString(writer, err.Error())
			return
		}
		io.WriteString(writer, ck2.Value)
	})
	http.ListenAndServe(":8080", nil)
}

func TestCookieSet2(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ck := &http.Cookie{
			Name:       "myCookie",
			Value:      "hello ,TestCookieSet2",
			Path:       "/",
			Domain:     "localhost",
			Expires:    time.Time{},
			RawExpires: "120",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		writer.Header().Set("Set-Cookie", strings.Replace(ck.String(), " ", "%20", -1))
		ck2, err := request.Cookie("myCookie")
		if nil != err {
			io.WriteString(writer, err.Error())
			return
		}
		io.WriteString(writer, ck2.Value)
	})
	http.ListenAndServe(":8080", nil)
}

func TestTimeNow(t *testing.T) {
	tq := time.Now().Local().Local()
	t1 := tq.Format("2006-01-02 15:04:05")
	t2, _ := time.Parse("2006-01-02 15:04:05", t1)
	fmt.Println(tq)
	fmt.Println(t1)
	fmt.Println(t2)
}
