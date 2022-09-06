/*
Teacher Data Struct
------------------
Name: string;
Room: int;
Email: string;
Picture: string; This wll just pont to and endpoint on the server where the img i stored
Zoom: string;
Description: string;
*/

package teacher

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type student struct {
	name    string
	room    int
	email   string
	zoom    string
	picture string
	picture string
}

func getStudent(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}

func createTeacher(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var teacher teacher
	json.Unmarshal([]byte(c.Body()), &teacher)

	res, err := collection.InsertOne(context.Background(), teacher)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func updateTeacher(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var teacher teacher
	json.Unmarshal([]byte(c.Body()), &teacher)

	update := bson.M{
		"$set": teacher,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func deleteStudent(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
}
func getTeacher(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getTeacherName(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getTeacherRoom(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getTeacherEmail(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getTeacherPicture(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getTeacherZoom(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getTeacherDescription(c *fiber.Ctx) {
	c.Send("The Arctesian")
}

func postTeacher(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func postTeacherName(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func postTeacherRoom(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func postTeacherEmail(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func postTeacherPicture(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func postTeacherZoom(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func postTeacherDescription(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func deleteTeacher(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
