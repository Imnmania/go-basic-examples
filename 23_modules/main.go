package main

/*	Important Commands:
	-------------------
	go get -u github.com/gorilla/mux (gets the package)
	go build . (builds the app)
	go mod tidy (cleans up the packages)
	go mod verify (verify the added packages by hashcode)
	go list (lists modules)
	go list all (lists all dependent modules)
 	go list -m all (lists modules that are used)
	go list -m -versions github.com/gorilla/mux (lists all compatible versions of the package)
	go mod why github.com/gorilla/mux (lists why we use this package)
	go mod graph (which module uses which package)
	go mod edit -go 1.16 (changes go version to 1.16, or you can manually edit the mod file)
	go mod edit -module <new-name> (changes the module name)
	go mod vendor (brings all the dependencies from cache to project directory)
	go run -mod=vendor main.go (runs the app from vendor dependency files)
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on youtube</h1>"))
}
