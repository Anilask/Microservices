package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		//d,_:= ioutil.ReadAll(r.Body) //read the body and return body
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops"))
			return
		}
		fmt.Fprint(rw, "Hello %s\n", d)

	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodboy world")
	})

	http.ListenAndServe(":8080", nil)
}
/*
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// reqeusts to the path /goodbye with be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// any other request will be handled by this function
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		// read the body
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
*/