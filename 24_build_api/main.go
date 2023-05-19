package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// -----------------------
// Model for Course - file
// -----------------------
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// -----------------------
// Model for Author - file
// -----------------------
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// -------
// Fake DB
// -------
var courses []Course

// -------------------------
// Middleware, Helper - file
// -------------------------
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// ------------------
// Controllers - file
// ------------------

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Accessing home page...")
	w.Write([]byte("<h1>Welcome to Golang API</h1>"))
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get course by course id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course...")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// - loop through the courses
	// - find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

// create course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course...")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what if: {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, item := range courses {
		if strings.EqualFold(item.CourseName, course.CourseName) {
			json.NewEncoder(w).Encode("Course already exists with this name")
			return
		}
	}

	// generate a unique id, convert to string
	// append course into courses
	rand.New(rand.NewSource(time.Now().Unix()))
	course.CourseId = strconv.Itoa(rand.Intn(10000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// update course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course...")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var courseFromReq Course
			_ = json.NewDecoder(r.Body).Decode(&courseFromReq)
			courseFromReq.CourseId = params["id"]
			courses = append(courses, courseFromReq)
			json.NewEncoder(w).Encode(courseFromReq)
		}
	}

	// TODO: send a response when id is not found
}

// delete course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course...")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove(index, index + 1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted course with given id")
			break
		}
	}
}

func main() {
	fmt.Println("Welcome to building API with golang...")
	router := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Flutter",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "Niloy Biswas",
			Website:  "niloybiswas.me",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Go",
		CoursePrice: 150,
		Author: &Author{
			Fullname: "Niloy Biswas",
			Website:  "niloybiswas.me",
		},
	})

	// routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", router))
}
