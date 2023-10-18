package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

type Course struct {
	Course_id    string  `json:"course_id"`
	Course_Name  string  `json:"coursename"`
	Course_price string  `json:"price"`
	Author       *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// Middle ware
func (c *Course) isEmpty() bool {
	// return c.Course_id == "" && c.Course_Name == ""
	return c.Course_Name == ""

}

func main() {

	fmt.Println("hello this is api shit ")
	r := mux.NewRouter()
	courses = append(courses, Course{Course_id: "1", Course_Name: "GOLANF", Course_price: "1234", Author: &Author{Fullname: "Lauda Prasad", Website: "www.archit.com"}})
	courses = append(courses, Course{Course_id: "2", Course_Name: "GOLANF", Course_price: "1234", Author: &Author{Fullname: "Lauda Prasad", Website: "www.archit.com"}})

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/courses", getcourses).Methods("GET")
	r.HandleFunc("/getone/{id}", getOne).Methods("GET")
	r.HandleFunc("/createcourse", createcoutse).Methods("POST")
	r.HandleFunc("/createcourse/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/createcourse/{id}", DeleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}

// controllers
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>>THis is the home route</h1>"))

}
func getcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GEt all courses")
	w.Header().Set("Content=-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GEtting 1 course")
	w.Header().Set("Content=-Type", "application/json")
	// grab id form request
	params := mux.Vars(r)
	// loop through courses ,find matchin gid an return the respose
	for _, course := range courses {
		if course.Course_id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found wirh given ID")
	return

}

func createcoutse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create 1 course")
	w.Header().Set("Content=-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some Data")

	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("NO  Data")

		return

	}
	// check if the course with the same name exist 
	

	// generate a unique id,string
	id := uuid.New()

	course.Course_id = (id.String())
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
	// course=append(courses,course)
}
func updateCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create 1 course")
	w.Header().Set("Content=-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Course_id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.Course_id = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}

	}

	// send a response when id is not found

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting course")
	w.Header().Set("Content=-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {

		if course.Course_id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			break

		}
	}

}
