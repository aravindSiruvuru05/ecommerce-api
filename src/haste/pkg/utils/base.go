package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	mathrand "math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type HType map[string]string

// Get fetches key value from HType.
// It returns a string value.
func (ht HType) Get(key string) string {
	for k, v := range ht {
		if strings.ToLower(k) == key {
			return v
		}
	}
	return ""
}

var seededRand *mathrand.Rand = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

// NewUUID generates a random UUID according to RFC 4122.
// It returns uuid and error.
func NewUUID() (string, error) {
	uuid := make([]byte, 16) //nolint:makezero
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1.
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3.
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// IsValidUUID checks whether uuid is valid or not.
// It returns a boolean value.
func IsValidUUID(uuid string) bool {
	r, _ := regexp.Compile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// GetRandomString generates a random alphanumeric string of given length.
// It returns the random string.
func GetRandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	}

	b := make([]byte, length) //nolint:makezero
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// ToStringSlice converts any slice of interface to a slice of string.
// It returns the string slice.
func ToStringSlice(iSlice []interface{}) []string {
	var sSlice = make([]string, 0)
	for _, val := range iSlice {
		sSlice = append(sSlice, fmt.Sprintf("%v", val))
	}
	return sSlice
}

// BoolToInt equates false as 0 and true as 1.
// It returns the int equivalent.
func BoolToInt(bVal bool) int {
	if !bVal {
		return 0
	} else {
		return 1
	}
}

// GetAppConfigMap fetches app config as map for given key and with env vars replaced.
// It returns a map.
func GetAppConfigMap(key string, fn func(string) (interface{}, error), expand ...bool) map[string]interface{} {
	value, _ := fn(key)

	var mValue map[string]interface{}
	if len(expand) > 0 && expand[0] {
		j, _ := json.Marshal(value)
		strVal := os.ExpandEnv(string(j))
		_ = json.Unmarshal([]byte(strVal), &mValue)
	} else {
		mValue, _ = value.(map[string]interface{})
	}

	return mValue
}

// GetHeaderValue fetches the value from the header.
// It returns the value and error.
func GetHeaderValue(headers string, key string) (string, error) {
	var httpHeaders http.Header
	var headerMap HType
	var headerValue string
	if err := json.Unmarshal([]byte(headers), &httpHeaders); err == nil {
		headerValue = httpHeaders.Get(key)
	} else if err = json.Unmarshal([]byte(headers), &headerMap); err == nil {
		headerValue = headerMap.Get(key)
	} else {
		return "", err
	}
	return strings.ToLower(headerValue), nil
}

func init() {
	fmt.Println("utils init______ util")
}
