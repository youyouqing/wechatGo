package api

import (
	"fmt"
	"net/http"
)

func ApiHandle(w http.ResponseWriter,r *http.Request)  {

	// TODO  register route here
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Fprintf(w,"HelloWorld!")

}
