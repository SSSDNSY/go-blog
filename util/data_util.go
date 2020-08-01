package util

import (
	"strings"
	"time"
)

func PaserDateTime(d time.Time) string {
	return strings.Split(d.String(), ".")[0]
}

//func SetPrivatePage(data * map[interface{}]interface{} ,image, title, subtitle string ) {
//	data
//	data["image"] = image
//	data["title"] = title
//	data["subtitle"] = subtitle
//}
