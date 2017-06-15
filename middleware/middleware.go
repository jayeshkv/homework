package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
)

const (
	crlf       = "\r\n"
	colonspace = ": "
)

func ChecksumMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// your code goes here ...
		//https://groups.google.com/forum/#!topic/golang-nuts/Jk785WB7F8I
		recorder := httptest.NewRecorder()
		h.ServeHTTP(recorder, r)

		cod := recorder.Code
		junkString := strconv.Itoa(cod) //str conversion
		junkString += crlf

		keys := []string{}
		headerJunk := ""

		for k := range recorder.Header() {
			keys = append(keys, k) //extracting headers
		} //endfor

		sort.Strings(keys)
		for _, vals := range keys {
			junkString += vals + ": " + recorder.Header().Get(vals) + crlf
			if vals != keys[len(keys)-1] { //compares with last element
				headerJunk += vals + ";"
			} else {
				headerJunk += vals + crlf + crlf
			}
		} //end for

		finString := []byte(junkString + "X-Checksum-Headers" + colonspace + headerJunk + recorder.Body.String()) //final canonical string
		// converting finString to Sha1
		sha := sha1.New()
		sha.Write(finString)
		bs := sha.Sum(nil)
		//converting to hex
		hexEncoded := hex.EncodeToString([]byte(bs))
		w.Header().Set("X-Checksum", hexEncoded)
		h.ServeHTTP(w, r)

	})
}

// Do not change this function.
func main() {
	var listenAddr = flag.String("http", "localhost:8080", "address to listen on for HTTP")
	flag.Parse()

	http.Handle("/", ChecksumMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Foo", "bar")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Date", "Sun, 08 May 2016 14:04:53 GMT")
		msg := "Curiosity is insubordination in its purest form.\n"
		w.Header().Set("Content-Length", strconv.Itoa(len(msg)))
		fmt.Fprintf(w, msg)
	})))

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
