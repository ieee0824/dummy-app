package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "<h1>It works!</h1>"
	str += fmt.Sprintf("%v<br>\n", runtime.GOOS)
	str += fmt.Sprintf("%v<br>\n", runtime.GOARCH)
	str += fmt.Sprintf("%v<br>\n", runtime.Version())

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalln(err)
		}
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				str += fmt.Sprintf("%v<br>\n", v.IP)
			case *net.IPAddr:
				str += fmt.Sprintf("%v<br>\n", v.IP)
			}
		}
	}

	fmt.Fprintf(w, str)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
