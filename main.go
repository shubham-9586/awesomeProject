package main

import (
	"awesomeProject/appLogic"
	"awesomeProject/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	RollNo int    `json:"rollNo"`
	Class  string `json:"class"`
}

var students = []Student{
	{Name: "A", Age: 12, RollNo: 1, Class: "6th"},
	{Name: "B", Age: 13, RollNo: 2, Class: "7th"},
	{Name: "C", Age: 14, RollNo: 3, Class: "8th"},
	{Name: "D", Age: 15, RollNo: 4, Class: "9th"},
}

func main() {
	router := gin.Default()
	router.GET("/students", process)
	router.GET("/students/:id", processId)
	router.POST("/students", processAdd)
	router.Run(":8080")
	fmt.Println("server started successfully")
}

func process(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}

func processAdd(context *gin.Context) {
	var student Student
	if err := context.BindJSON(&student); err != nil {
		return
	}
	students = append(students, student)
	context.IndentedJSON(http.StatusOK, students)
}
func processId(context *gin.Context) {
	var student Student
	id, _ := strconv.Atoi(context.Param("id"))

	for _, val := range students {
		if val.RollNo == id {
			student = val
			break
		}
	}
	context.IndentedJSON(http.StatusOK, student)
}

func processInput(query []string) {
	switch query[0] {
	case constants.ADD_USER:
		{
			age, _ := strconv.Atoi(query[4])
			user := appLogic.GetNewUser(query[1], query[2], query[3], age, query[5], query[6])
			appLogic.AddUserToDB(user)
		}
	case constants.ADD_VACCINATION_CENTER:
		{
			var arr [10]int
			vaccinationCentre := appLogic.GetNewVaccinationCentre(query[3], query[1], query[2], arr)
			appLogic.AddNewVaccinationCentre(vaccinationCentre)
		}
	case constants.ADD_CAPACITY:
		{
			capacity, _ := strconv.Atoi(query[3])
			day, _ := strconv.Atoi(query[2])
			appLogic.AddCapacity(query[1], capacity, day)
		}
	case constants.BOOK_VACCINATION:
		{
			day, _ := strconv.Atoi(query[2])
			booking := appLogic.GetNewBooking(query[1], day, query[3], constants.TRUE)
			appLogic.AddNewBooking(booking)
		}
	case constants.CANCEL_BOOKING:
		{
			appLogic.CancelBooking(query[2])
		}
	case constants.LIST_ALL_BOOKINGS:
		{
			appLogic.ListAllBookings(query[1])
		}
	case constants.LIST_VACCINATION_CENTERS:
		{
			appLogic.ListAllVaccinationCentres(query[1])
		}

	}
}
