package main

import (
	"io"
	"os"
)

func arguments(args []string) {
	println(args)
	f, err := os.Open(args[0])
	if err != nil {
		println("whoopsie")
	} else {
		io.Copy(os.Stdout, f)
		// Lines below give you control over buffer size
		/*
			buf := make([]byte, 1024)
			n, err := f.Read(buf)
			if err != nil {
				println("whoopsie")
			}
			fmt.Printf("%d bytes: %s\n", n, string(buf[:n]))
		*/
	}
}
