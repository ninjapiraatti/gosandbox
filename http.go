package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct {
}

func httpInterface() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := LogWriter{}

	/*
		bs := make([]byte, 99999)
		res.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	io.Copy(lw, res.Body)
}

// Because Logwriter is now implementing Write, it can be used in io.Copy
func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
