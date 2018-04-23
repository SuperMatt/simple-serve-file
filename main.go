package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	listen := flag.String("listen", ":8081", "Listen address")
	flag.Parse()
	a := flag.Args()
	if len(a) == 0 {
		fmt.Println("Please specify a file")
		os.Exit(1)
	}
	file := a[0]
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(f))
	})

	log.Fatal(http.ListenAndServe(*listen, nil))
}
