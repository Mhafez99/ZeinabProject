package main

import (
	
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"sort"
	O "github.com/Hafez/data"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
type reader struct {
	Id         int    `json:"id"`
	Name      string `json:"name"`
	Birthday   string `json:"birthday"`
	Height    int    `json:"height"`
	Weight     int    `json:"weight"`
	Employment string `json:"employment"`
}
type book struct {
	Id        int    `json:"id"`
	Title      string `json:"title"`
	Publication string `json:"publication"`
	Date       string `json:"date"`
	Author     string `json:"author"`
	Genre    string `json:"genre"`
	Publisher  string `json:"publisher"`
	Language   string `json:"language"`
}

func ConvertReaderToRow(oneReader reader) string {
	row := strconv.Itoa(oneReader.Id) + " "
	row += oneReader.Name + " "
	row += oneReader.Birthday + " "
	row += strconv.Itoa(oneReader.Height) + " "
	row += strconv.Itoa(oneReader.Weight) + " "
	row += oneReader.Employment
	row += "\n"
	return row
}
func CreateReader(w http.ResponseWriter, r *http.Request) {
	// take data from body of request
	reqBody, _ := ioutil.ReadAll(m.Body)

	// store data into oneReader
	var oneReader reader
	json.Unmarshal(reqBody, &oneReader)

	// convert data to string (row)
	row := ConvertReaderToRow(oneReader)
	// write in file
	O.AddRow("data/readers.txt", row)

	json.NewEncoder(w).Encode(oneReader)

}

func GetArrayOfReaders() []reader {
	// open file and return rows
	rows := O.GetRows("data/readers.txt")

	// array of reader'sStruct to sotre reader'sdata
	readers := []reader{}

	for _, row := range rows {
		// convert row to array of words
		words := strings.Split(row, " ")

		// cast to int
		id, _ := strconv.Atoi(words[0])
		height, _ := strconv.Atoi(words[3])
		weight, _ := strconv.Atoi(words[4])

		// create one reader then store in array of reader'sStruct
		newReader := reader{Id: id,
			Name:       words[1],
			Birthday:   words[2],
			Height:     height,
			Weight:     weight,
			Employment: words[5],
		}

		readers = append(readers, newReader)

	}
	return readers
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func DeleteReader(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// get all readers then search for a specific reader
	// delete the file then write again without the deleted reader
	readers := GetArrayOfReaders()
	O.RemoveFile("data/readers.txt")
	for _, reader := range readers {
		if reader.Id != id {
			row := ConvertReaderToRow(reader)
			O.AddRow("data/readers.txt", row)
		}
	}

}


func getAllReaders(w http.ResponseWriter, r *http.Request) {
	// conect with file
	readers := GetArrayOfReaders()
	// enableCors(&w)
	//(w).Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(readers)
	// get data
	// convert data from string to struct
	// convert struct to json
	// return json
	// fmt.Print("getAllReaders")
}





func SearchByIdForReader(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	readers := GetArrayOfReaders()
	for _, reader := range readers {
		if reader.Id == id {
			json.NewEncoder(w).Encode(reader)
		}
	}

}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	name := vars["name"]

	readers := GetArrayOfReaders()
	for _, reader := range readers {
		if reader.Name == name {
			json.NewEncoder(w).Encode(reader)
		}
	}
}
func ConvertBookToRow(oneBook book) string {
	row := strconv.Itoa(oneBook.Id) + " "
	row += oneBook.Title + " "
	row += oneBook.Publicaton + " "
	row += oneBook.Date + " "
	row += oneBook.Author + " "
	row += oneBook.Genre + " "
	row += oneBook.Publisher + " "
	row += oneBook.Language + " "
	row += "\n"
	return row
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// take data from body of request
	reqBody, _ := ioutil.ReadAll(m.Body)

	// store data into oneBook
	var oneBook book
	json.Unmarshal(reqBody, &oneBook)

	// convert data to string (row)
	row := ConvertBookToRow(oneBook)

	// write in file
	O.AddRow("data/books.txt", row)

	json.NewEncoder(w).Encode(oneBook)

}

func GetArrayOfBooks() []book {
	// open file and return rows
	rows := O.GetRows("data/books.txt")

	// array of reader'sStruct to sotre reader'sdata
	books := []book{}

	for _, row := range rows {
		// convert row to array of words
		words := strings.Split(row, " ")

		// cast to int
		id, _ := strconv.Atoi(words[0])

		// create one reader then store in array of reader'sStruct
		newBook := book{Id: id,
			Title:      words[1],
			Publicaton: words[2],
			Date:       words[3],
			Author:     words[4],
			Genre:      words[5],
			Publisher:  words[6],
			Language:   words[7],
		}

		books = append(books, newBook)

	}
	return books
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := GetArrayOfBooks()
	json.NewEncoder(w).Encode(books)

}

func SearchByIdForBook(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	books := GetArrayOfBooks()
	for _, book := range books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
		}
	}

}

func SearchByNameForBook(w http.ResponseWriter, r *http.Request) {
	// take the path variable
	// convert it to int
	vars := mux.Vars(r)
	title := vars["title"]

	books := GetArrayOfBooks()
	for _, book := range books {
		if book.Title == title {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	// Get Data From Request'sbody
	reqBody, _ := ioutil.ReadAll(m.Body)

	// Convert Json To Book
	var updatedBook book
	json.Unmarshal(reqBody, &updatedBook)

	// Get Array Of Books
	books := GetArrayOfBooks()

	// Remove The File
	O.RemoveFile("data/books.txt")

	// Write In The File
	// Don't Forget To add The UpdatedBook By Condition
	for _, book := range books {
		if book.Id == updatedBook.Id {
			row := ConvertBookToRow(updatedBook)
			O.AddRow("data/books.txt", row)
			json.NewEncoder(w).Encode(updatedBook)
		} else {
			row := ConvertBookToRow(book)
			O.AddRow("data/books.txt", row)
		}
	}

}

func SortByTitle(w http.ResponseWriter, r *http.Request) {
	// Get Books
	books := GetArrayOfBooks()

	// Sort Books By Title
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})
	json.NewEncoder(w).Encode(books)
}

func SortByPublication(w http.ResponseWriter, r *http.Request) {
	// Get Books
	books := GetArrayOfBooks()

	// Sort Books By Publication
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].Publicaton < books[j].Publicaton
	})
	json.NewEncoder(w).Encode(books)
}