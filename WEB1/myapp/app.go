package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct{
	FirstName string
	LastName string
	Email string
	CreatedAt time.Time
}


type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	//json 형태로 파싱 //제이슨으로 디코딩하는데 리턴으로 error를 뱉음
	err := json.NewDecoder(r.Body).Decode(user) 
	if err := nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

user.CreatedAt = time.Now()

data,_=json.Marshal(user)
w.Header().Add("content-type","application/json")
w.WriteHeader(http.StatusOK)
fmt.Fprint(w, string(data))

}

func NewHttpHandler() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello bar")
	})

	mux.Handle("/foo", &fooHandler{})

	return mux

}
