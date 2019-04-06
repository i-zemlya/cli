// Site project: i-zemlya.ru
// License: MIT
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags)

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

	/*for {
		b, err := net.Accept()
		if err != nil {
			continue
		}

	}*/

	server := NewServer()
	log.Println("Listening server...")
	go server.Listen()

	ss := &http.Server{
		Addr:    ":12190",
		Handler: server.Handler(),
	}
	ss.ListenAndServe()

	s := new(FCgiServer)
	fcgi.Serve(l, s)
	//fcgi.Serve(l, r)
}

type (
	FCgiServer struct{}
)

func (s *FCgiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`500 Error`))
			//log.Println(err.Error())
		}
	}()

	fmt.Printf("%s\n", r.URL.String())
	fmt.Printf("%+v\n\n\n", r)
	if r.URL.String() == "/favicon.png" {
		w.Write([]byte(``))
	} else {
		//w.Write([]byte(`<title>ok</title><link rel="icon" href="/favicon.png" type="image/x-icon"><p>stok</p>`))
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "stok")
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/ico")
	fmt.Fprint(w, "")
}
