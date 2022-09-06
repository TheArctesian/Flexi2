package main

import (
	"github.com/gofiber/fiber/v2"
)

// config := fiber.Config{
// ServerHeader: "Flexi Server", // add custom server header
//}

func setupRoutes(app *fiber.App) {
	// mainCalls

	// Event
	app.Get("api/event/:id", getEvent)
	app.Post("apu/event/", createTeacher)
	app.Put("api/event/:id", updateEvent)
	app.Delete("api/event/:id", deleteEvent)

	// Teacher
	app.Get("api/teacher/:id", getTeacher)
	app.Post("apu/teacher/", createTeacher)
	app.Put("api/teacher/:id", updateTeacher)
	app.Delete("api/teacher/:id", deleteTeacher)

	// student
	app.Get("api/student/:id", getStudent)
	app.Post("apu/student/", createStudent)
	app.Put("api/student/:id", updateStudent)
	app.Delete("api/student/:id", deleteStudent)

	// schedule
	app.Get("api/schedule/:id", getSchedule)
	app.Post("apu/schedule/", createSchedule)
	app.Put("api/schedule/:id", updateSchedule)
	app.Delete("api/schedule/:id", deleteSchedule)

	// Get Schedule
	app.Get("/api/schedule/day/:id", getFlexiDate)
	app.Get("/api/schedule/week/:id", getFlexiWeek)
	app.Get("/api/schedule/yr/:id", getFlexiYr)
	// app.Get("/api/schedule/Year7")
	// appapp.Get("/api/schedule/Year8")
	// appapp.Get("/api/schedule/Year9")
	// appapp.Get("/api/schedule/Year10")
	// appapp.Get("/api/schedule/Year11")
	// appapp.Get("/api/schedule/Year12")
	// appapp.Get("/api/schedule/Year13")

	// Set Schedule
	app.Post("/api/schedule/day", setFlexiDate)
	app.Post("/api/schedule/week", setFlexiWeek)
	app.Post("/api/schedule/yr", setFlexiYr)
	// app.Post("/api/schedule/Year7")
	// app.Post("/api/schedule/Year8")
	// app.Post("/api/schedule/Year9")
	// app.Post("/api/schedule/Year10")
	// app.Post("/api/schedule/Year11")
	// app.Post("/api/schedule/Year12")
	// app.Post("/api/schedule/Year13")

	// Get Events
	app.Get("api/event/name", getEventName)
	app.Get("api/event/teacher", getEventTeacher)
	app.Get("api/event/room", getEventRoom)
	app.Get("api/event/zoom", getEventZoom)
	app.Get("api/event/yearGroup", getEventYearGroup)
	app.Get("api/event/participants", getEventParticipants)

	// Set Events
	app.Post("api/event/", setEvent)
	app.Post("api/event/name", setEventName)
	app.Post("api/event/teacher", setEventTeacher)
	app.Post("api/event/room", setEventRoom)
	app.Post("api/event/zoom", setEventZoom)
	app.Post("api/event/yearGroup", setEventYearGroup)
	app.Post("api/event/participants", setEventParticipants)

	// Get Teacher
	app.Get("api/teacher", getTeacher)
	app.Get("api/teacher/description", getTeacherDescription)
	app.Get("api/teacher/email", getTeacherEmail)
	app.Get("api/teacher/name", getTeacherName)
	app.Get("api/teacher/picture", getTeacherPicture)
	app.Get("api/teacher/room", getTeacherRoom)
	app.Get("api/teacher/zoom", getTeacherZoom)

	// Set Teacher
	app.Post("api/teacher", setTeacher)
	app.Post("api/teacher/description", setTeacherDescription)
	app.Post("api/teacher/email", setTeacherEmail)
	app.Post("api/teacher/name", setTeacherName)
	app.Post("api/teacher/picture", setTeacherPicture)
	app.Post("api/teacher/room", setTeacherRoom)
	app.Post("api/teacher/zoom", setTeacherZoom)

	//Get Student
	app.Get("api/student", getStudent)
	app.Get("api/student/advisory", getStudentAdvisory)
	app.Get("api/student/name", getStudentName)
	app.Get("api/student/picture", getStudentPicture)
	app.Get("api/student/teachers", getStudentTeachers)

	// Set Student
	app.Post("api/student", setStudent)
	app.Post("api/student/advisory", setStudentAdvisory)
	app.Post("api/student/name", setStudentName)
	app.Post("api/student/picture", setStudentPicture)
	app.Post("api/student/teachers", setStudentTeachers)

}
func main() {
	app := fiber.New(config)
	setupRoutes(app)
	app.Listen(":6000")
}
