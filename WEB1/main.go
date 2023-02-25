package main

import (
	// "encoding/json"
	// "fmt"

	"net/http"
)

// type User struct{
// 	FirstName string
// 	LastName string
// 	Email string
// 	CreatedAt time.Time
// }

// type fooHandler struct{}

// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	user := new(User)
// 	//json 형태로 파싱 //제이슨으로 디코딩하는데 리턴으로 error를 뱉음
// 	err := json.NewDecoder(r.Body).Decode(user)
// 	if err := nil{
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, err)
// 		return
// 	}

// user.CreatedAt = time.Now()

// data,_=json.Marshal(user)
// w.Header().Add("content-type","application/json")
// w.WriteHeader(http.StatusOK)
// fmt.Fprint(w, string(data))

// }

// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Foo")
// }

func main() { //핸들러 등록 , 어떤 도메인 올떄 어떻게 할것인지
	//경로를 나눠주는에를 동적으로 등록하려고 Mux
	// mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World")
	// })

	// mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello bar")
	// })

	// mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type User struct {
// 	FirstName string
// 	LastName  string
// 	Email     string
// 	CreatedAt time.Time
// }

// type fooHandler struct{}

// func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	user := new(User)
// 	err := json.NewDecoder(r.Body).Decode(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "Bad Request: ", err)
// 		return
// 	}
// }

// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	if name == "" {
// 		name = "World"
// 	}
// 	fmt.Fprintf(w, "Hello %s!", name)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello World")
// 	})

// 	mux.HandleFunc("/bar", barHandler)

// 	mux.Handle("/foo", &fooHandler{})

// 	http.ListenAndServe(":3000", mux)
// }
