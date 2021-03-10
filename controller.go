package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

func serve() {
	m := &Model{1000}
	http.HandleFunc("/api/v1/counter", m.counterHandeler)
	http.Handle("/", http.FileServer(getFileSystem()))

	server := &http.Server{Addr: ":8080"}
	fmt.Println("open:", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Model) counterHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%-4s Counter in: %22d", r.Method, p.Counter())
	defer func() { fmt.Printf(" out: %22d\n", p.Counter()) }()

	switch r.Method {
	case "GET":
		w.Write([]byte(fmt.Sprintf(`{"counter":%d}`, p.Counter())))

	case "POST":
		var m map[string]int64
		err := json.NewDecoder(r.Body).Decode(&m)
		// The Server will close the request body. The ServeHTTP
		// Handler does not need to.
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		v, ok := m["add"]
		if !ok {
			http.Error(w, "No 'add:value' insilde the JSON payload.", http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf(`{"counter":%d}`, p.Add(v))))

	case "PUT":
		var m map[string]int64
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		v, ok := m["counter"]
		if !ok {
			http.Error(w, "No 'counter:value' insilde the JSON payload.", http.StatusBadRequest)
			return
		}
		p.SetCounter(v)
		w.Write([]byte(fmt.Sprintf(`{"counter":%d}`, p.Counter())))

	default:
		http.Error(w, "Method NotImplemented", http.StatusNotImplemented)
	}
}

//go:embed view/*
var dist embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(dist, "view")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}
