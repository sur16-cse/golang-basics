package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//controllers file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Api building"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id fromr request
	params := mux.Vars(r)

	//loop through courses and find matching id and return the course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//What if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside json")
		return
	}

	//check only if title is duplicate
	//loop title matches with course.coursename, JSON
	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicate data found")
			return
		}
	}

	//generate unique id,string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id fromr request
	params := mux.Vars(r)

	//loop id, remove it, add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	 json.NewEncoder(w).Encode("No Course found with given id") 
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete All Courses")
	courses = []Course{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//loop,id,remove (index,index+1)

	for index,course:=range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted Successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id") 
}

func main() {
	fmt.Println("Api building")
	r:=mux.NewRouter()

	//seeding
	courses=append(courses, Course{CourseId: "2",CourseName: "Reactjs",CoursePrice: 500,Author: &Author{Fullname: "Max",Website: "lco.dev"}})
	courses=append(courses, Course{CourseId: "4",CourseName: "Mern",CoursePrice: 600,Author: &Author{Fullname: "Hari",Website: "go.dev"}})

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("Post")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses",deleteAllCourses).Methods("DELETE")
	
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))
}
