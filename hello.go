package main



import (
    "fmt"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "Hello,yo!")
    // // fmt.Fprint(w, "Hello,yo!")
    // fmt.Println("Hey")
    
    
}

func (h Hello) printHey() {
	fmt.Println("Hey")
}

func main() {
    var h Hello
    go h.printHey()
   	http.ListenAndServe("localhost:4000", h)
    
}
