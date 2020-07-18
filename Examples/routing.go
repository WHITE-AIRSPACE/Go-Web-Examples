// +build ignore

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 路由使用 gorilla/mux 包

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8080", r)
}

//将请求处理程序限制为特定的HTTP方法。

// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// 将请求处理程序限制为特定的主机名或子域。

// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

//将请求处理程序限制为http / https。

// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

//将请求处理程序限制为特定的路径前缀。

// bookrouter := r.PathPrefix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBooks)
// bookrouter.HandleFunc("/{title}", GetBook)
