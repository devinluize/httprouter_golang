package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	//request := httptest.NewRequest("GET", "https://localhost:3000/", nil)
	//recorder := httptest.NewRecorder()

	//router.ServeHTTP(recorder, request)
	//response := recorder.Result()

	router.GET("/tesrouting", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "hello get")
	})
	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
