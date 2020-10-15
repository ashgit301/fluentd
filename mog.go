package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name   string `json:"name" bson:"name"`
	Age    int    `json:"age" bson:"age"`
	IsMale bool   `json:"isMale" bson:"isMale"`
}

var session, _ = mgo.Dial("localhost:27017")
var c = session.DB("trydb").C("trycollection")

// func AddLogs() {
// 	fmt.Printf("Writing to a file in Go lang\n")
// 	file, err := os.Create("data/logs.txt")

// 	if err != nil {
// 		log.Fatalf("failed creating file: %s", err)
// 	}
// 	defer file.Close()
// 	log.SetOutput(file)
// 	log.Println("creating file")
// 	len, err := file.WriteString("This is a read/write file")
// 	if err != nil {
// 		log.Fatalf("failed writing to file: %s", err)
// 	}
// 	fmt.Printf("\nFile Name: %s", file.Name())
// 	fmt.Printf("\nLength: %d bytes", len)
// }

// func ReadLogs() {
// 	fmt.Printf("\n\nReading a file in Go lang\n")
// 	fileName := "/data/logs.txt"

// 	// The ioutil package contains inbuilt
// 	// methods like ReadFile that reads the
// 	// filename and returns the contents.
// 	data, err := ioutil.ReadFile("/data/logs.txt")
// 	if err != nil {
// 		log.Panicf("failed reading data from file: %s", err)
// 	}
// 	// file, _ := os.OpenFile("/data/logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	// log.SetOutput(data)
// 	log.Println("reading file")
// 	fmt.Printf("\nFile Name: %s", fileName)
// 	fmt.Printf("\nSize: %d bytes", len(data))
// 	fmt.Printf("\nData: %s", data)
// }
// func WriteData(w http.ResponseWriter, r *http.Request) {
// 	//w.Header().Set("Content-Type", "application/json")
// 	fmt.Println("@@")
// 	AddLogs()
// 	response := "writing logs to logs.txt file"
// 	// response.Put("message", "Logs Added Successfully")
// 	w.Write([]byte(response))
// }

// func ReadData(w http.ResponseWriter, r *http.Request) {
// 	//w.Header().Set("Content-Type", "application/json")
// 	ReadLogs()
// 	response := "reading logs from logs.txt file"
// 	//response.Put("message", "Logs Readed Successfully")
// 	w.Write([]byte(response))

// }

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.OpenFile("data/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Println("entering the getallData call")
	log.Println("@@@@@@@@@@@@@@@@@@@@@@@")
	result := []Person{}
	c.Find(bson.M{}).All(&result)
	json.NewEncoder(w).Encode(result)
	log.Println("fetched all data from db")
	//w.Write([]byte(`{"message":"fetching data from db"}`))
}
func getOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.OpenFile("data/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Println("@@@@@@@@@@@@@@@@@@@@@@@")
	log.Println("entering the getoneData call")
	params := mux.Vars(r)
	Name := params["name"]
	var data Person
	c.Find(bson.M{"name": Name}).One(&data)
	json.NewEncoder(w).Encode(data)
	log.Println("fetched single data from db")
	//w.Write([]byte(`{"message":"fetching one data from db"}`))

}
func insertData(w http.ResponseWriter, r *http.Request) {
	file, _ := os.OpenFile("data/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Println("entering the insertData call")
	log.Println("@@@@@@@@@@@@@@@@@@@@@@@")
	w.Header().Set("Content-Type", "application/json")
	var data Person
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
	c.Insert(&data)
	log.Println("inserting data into the db")
	//w.Write([]byte(`{"message":"inserting data from db"}`))
	//json.NewEncoder(w).Encode(data)
}

func updateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.OpenFile("data/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Println("entering the updateData call")
	log.Println("@@@@@@@@@@@@@@@@@@@@@@@")
	params := mux.Vars(r)
	Name := params["name"]
	var data Person
	json.NewDecoder(r.Body).Decode(&data)
	c.Update(bson.M{"name": Name}, bson.M{"$set": bson.M{"name": data.Name, "age": data.Age, "isMale": data.IsMale}})
	log.Println("updating data in the db")
	//w.Write([]byte(`{"message":"updating data into db"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	file, _ := os.OpenFile("data/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Println("entering the deleteData call")
	log.Println("@@@@@@@@@@@@@@@@@@@@@@@")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Name := params["name"]
	c.Remove(bson.M{"name": Name})
	log.Println("deleting data from db")
	log.SetOutput(file)
	//w.Write([]byte(`{"message":"deleting data from db"}`))
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getAll).Methods("GET")
	router.HandleFunc("/insert", insertData).Methods("POST")
	router.HandleFunc("/{name}", getOne).Methods("GET")
	router.HandleFunc("/update/{name}", updateData).Methods("PUT")
	router.HandleFunc("/delete/{name}", delete).Methods("DELETE")
	//router.HandleFunc("/write", WriteData).Methods("POST")
	//router.HandleFunc("/read", ReadData).Methods("GET")
	http.ListenAndServe(":8050", router)
}
