// Site project: i-zemlya.ru
// License: MIT
package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

func main() {
	//var g *int32
	//*g = 637
	fmt.Printf("%s\n", "https://i-zemlya.ru")
	//fmt.Printf("%+v\n", &g)

	os.Remove("cli.sock")
	l, err := net.Listen("unix", "cli.sock")
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
		return
	}
	r := http.NewServeMux()
	r.HandleFunc("/", index)
	r.HandleFunc("/favicon.ico", favicon)
	fcgi.Serve(l, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "stok")
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/ico")
	fmt.Fprint(w, "")
}