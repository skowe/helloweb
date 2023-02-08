package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Funkcija samo treba da popuni ResponseWriter objekat
func SayHelloHandleFunction(w http.ResponseWriter, r *http.Request) {
	dataToWrite := []byte("Hello World!")

	_, err := w.Write(dataToWrite)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func main() {
	// Ucitaj vrednost porta iz APPPORT promenljive
	port, ok := os.LookupEnv("APP_PORT")

	if !ok {
		port = "8080"
	}
	// Ispisi "Hello tty!" na standardnom izlazu
	fmt.Println("Hello TTY!")
	
	// Dodaj putanju u defoltnom ServeMux objektu funkciji SayHelloHandleFunction
	http.HandleFunc("/", SayHelloHandleFunction)

	// Formiraj adresu na kojoj aplikacija slusa za zahteve
	listenOn := fmt.Sprintf(":%s", port)

	// Pokreni multiplekser na adresi definisanoj u listenOn
	log.Println(http.ListenAndServe(listenOn, http.DefaultServeMux))
}
