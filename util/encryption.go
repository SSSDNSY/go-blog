package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

const (
	ApiAppKey     string = "A1MLDO616VDVBQHVQOKFBT6OQDJEJF280W5S"
	ApiAppSecrect string = "dJoBLS9I9kmWvE4PGiUiQK5NUNlWUmYw3hAnaxkcFo0s0xUglrcX5Fi5oeHnolcgEXAuJSILCgCZDkdAgufCMRlF"

	//ApiAppKeyBAKKKKKK     string = "13BECC0B90027D99D86E6E6E80EA8E3E"
	//ApiAppSecrectBAKKKKKK string = "rDJR1F9vEcvcdZh650JHZl8ngVNnBNyojlIrnigw3WqP23OoNRdyHhUqVdf3rZMGBmCHPEAL3"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomStr(length int) string {
	return StringWithCharset(length, charset)
}
