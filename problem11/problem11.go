// https://leetcode.com/problems/encode-and-decode-tinyurl/

package main

import (
	"fmt"
)

type Codec struct {
	url         string
	encoded_url string
}

func Constructor() Codec {
	var c Codec
	return c
}

// global list of encodes and urls
var url_list []Codec

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	retUrl := "http://tinyurl.com/"

	a_count := 0
	e_count := 0
	i_count := 0
	o_count := 0
	m_count := 0

	for _, char := range longUrl {
		if char == 'a' {
			a_count++
		}

		if char == 'e' {
			e_count++
		}

		if char == 'i' {
			i_count++
		}

		if char == 'o' {
			o_count++
		}

		if char == 'm' {
			m_count++
		}
	}
	retUrl += fmt.Sprintf("%da", a_count)
	retUrl += fmt.Sprintf("%de", e_count)
	retUrl += fmt.Sprintf("%di", i_count)
	retUrl += fmt.Sprintf("%do", o_count)
	retUrl += fmt.Sprintf("%dm", m_count)
	retUrl += fmt.Sprintf("L%d", len(longUrl))
	retUrl += fmt.Sprintf("S%s", longUrl[:1])
	retUrl += fmt.Sprintf("E%s", longUrl[1:])

	var c Codec
	c.url = longUrl
	c.encoded_url = retUrl

	url_list = append(url_list, c)

	return retUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	for i := 0; i < len(url_list); i++ {
		if shortUrl == url_list[i].encoded_url {
			return url_list[i].url
		}
	}
	return "UNKNOWN"
}

func main() {
	fmt.Print("Use 'go test' to test this function\n")
}
