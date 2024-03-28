package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func runinitial() {
	router := httprouter.New()
	//request := httptest.NewRequest("GET", "https://localhost:3000/", nil)
	//recorder := httptest.NewRecorder()

	//router.ServeHTTP(recorder, request)
	//response := recorder.Result()

	router.GET("/tesrouting", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, err := fmt.Fprint(writer, "hello get")
		if err != nil {
			return
		}
	})
	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
func methodprint(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	userid := params.ByName("userid")
	_, err := fmt.Fprintf(writer, "tesmethod %s %s", id, userid)
	if err != nil {
		return
	}

}
func teshttp() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "gak ketemu ")
		if err != nil {
			return
		}
	})
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		_, err := fmt.Fprint(writer, "panic", i)
		if err != nil {
			return
		}
	}
	router.GET("/book/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "product" + params.ByName("id") + "tes"
		_, err := fmt.Fprintf(writer, text)
		if err != nil {
			return
		}
		panic("errrorrrrrrrrrrrrrrrrrrrrrrrrrrrr")
	})
	router.GET("/buku2/*id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "product" + params.ByName("id") + "tes"
		_, err := fmt.Fprintf(writer, text)
		if err != nil {
			return
		}
	})
	router.GET("/books/:id/user/:userid", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		methodprint(writer, request, params)
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
func main() {
	teshttp()
}
