package main

import (
	"fmt"
	"net/http"
)

func defult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func baidu(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	par := r.URL.RawQuery
	url := "http://www.baidu.com"
	if par != "" {
		url += "?" + par
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.HandleFunc("/", defult)
	http.HandleFunc("/baidu", baidu)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
