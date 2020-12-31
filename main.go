package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"sort"
	C "github.com/Hafez/Downloads/Music/ZinabCpl/Function+data"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
func CreatFile(nameOfFile string) {

}

func RemoveFile(nameOfFile string) {
	e := os.Remove(nameOfFile)
	if e != nil {
		log.Fatal(e)
	}
}

func OpenFileByOsToRead(nameOfFile string) *os.File {
	file2, err := os.Open(nameOfFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file2
}

func OpenFileByOsToWrite(nameOfFile string) *os.File {
	file2, err := os.OpenFile(nameOfFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file2
}

func GetRows(nameOfFile string) []string {

	file := OpenFileByOsToRead(nameOfFile)

	var rows []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	defer file.Close()

	return rows

}

func AddRow(nameOfFile string, row string) {
	file := OpenFileByOsToWrite(nameOfFile)
	_, err2 := file.WriteString(row)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("done")

	}
	defer file.Close()
}



func main() {
	
	m := mux.NewRouter()
	
	 m.HandleFunc("/reader", C.CreateReader).Methods("POST")
	 m.HandleFunc("/reader/{id}", C.DeleteReader).Methods("DELETE")
	 m.HandleFunc("/readers", getAllReaders).Methods("GET")
	 m.HandleFunc("/reader/{id}", C.SearchByIdForReader).Methods("GET")
	 m.HandleFunc("/readerName/{name}", C.SearchByName).Methods("GET")
	 m.HandleFunc("/book", C.CreateBook).Methods("POST")
	 m.HandleFunc("/book/{id}", C.SearchByIdForBook).Methods("GET")
	 m.HandleFunc("/books", C.GetBooks).Methods("GET")
	 m.HandleFunc("/bookEdit/{id}", C.EditBook).Methods("PUT")
	 m.HandleFunc("/book-sort-by-title", C.SortByTitle).Methods("GET")
	 m.HandleFunc("/book-sort-by-publication", C.SortByPublication).Methods("GET")
	log.Fatal(http.ListenAndServe(":8070", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}




