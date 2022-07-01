// https://leetcode.com/problems/encode-and-decode-tinyurl/

package main

import "fmt"

type Codec struct {
}

func Constructor() Codec {
	var c Codec
	return c
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	// sample input: https://leetcode.com/problems/design-tinyurl
	// sample return: http://tinyurl.com/4e9iAk
	
	retUrl := "http://tinyurl.com/"
	return retUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return shortUrl
}

func main() {
	fmt.Print("Use 'go test' to test this function")
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
