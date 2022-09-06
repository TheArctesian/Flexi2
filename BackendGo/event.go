/*
Flexi Data Struct
------------------
Name: string;
Teacher: string; # maybe a hotlink to the teachers page
Room: int;
Zoom: string;
Date: string; # unixTime
YearGroup: string;
Participants: string[];
*/

package event

import (
	"github.com/gofiber/fiber"
)

type event struct {
	name 			string
	teacher			string
	room 			int
	zoom 			int
	yearGroup 		string
	participants 	string[];
}

func getEvent(c *fiber.Ctx) {
	c.send("A fucking object man suck my balls")
}
func getEventName(c *fiber.Ctx) {
	c.Send("This is a big obj")
}

func getEventTeacher(c *fiber.Ctx) {
	c.Send("Edwin Lagos")
}
func getEventRoom(c *fiber.Ctx) {
	c.Send("513")
}
func getEventZoom(c *fiber.Ctx) {
	c.Send("https://zoom.us/j/5384991571#success")
}

func getEventYearGroup(c *fiber.Ctx) {
	c.Send("7")
}

func getEventParticipants(c *fiber.Ctx) {
	c.Send("Stephen Okita", "Micheal Kim")
}

func setEvent(c *fiber.Ctx) {
	c.send("A fucking object man suck my balls")
}
func setEventName(c *fiber.Ctx) {
	c.Send("This is a big obj")
}

func setEventTeacher(c *fiber.Ctx) {
	c.Send("Edwin Lagos")
}
func setEventRoom(c *fiber.Ctx) {
	c.Send("513")
}
func setEventZoom(c *fiber.Ctx) {
	c.Send("https://zoom.us/j/5384991571#success")
}

func setEventYearGroup(c *fiber.Ctx) {
	c.Send("7")
}

func setEventParticipants(c *fiber.Ctx) {
	c.Send("Stephen Okita", "Micheal Kim")
}

