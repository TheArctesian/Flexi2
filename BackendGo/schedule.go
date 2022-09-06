/*
In this situation OOP is a good idea

Schedule Data Struct
------------------
Date: string;
Week: string;
Year7Event: flexi[];
Year8Event: flexi[];
Year9Event: flexi[];
Year10Event: flexi[];
Year11Event: flexi[];
Year12Event: flexi[];
Year13Event: flexi[];
*/
package schedule

import (
	"github.com/gofiber/fiber"
)

type schedule struct {
	ID primitive.ObjectID
	Week string
	Date string
	Year7Events flexi[]
	Year8Events flexi[]
	Year9Events flexi[]
	Year10Events flexi[]
	Year11Events flexi[]
	Year12Events flexi[]
	Year13Events flexi[]
}


func getSchedule(c *fiber.Ctx) {
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

func createSchedule(c *fiber.Ctx){
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var Schedule Schedule
	json.Unmarshal([]byte(c.Body()), &Schedule)

	res, err := collection.InsertOne(context.Background(), Schedule)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func updateSchedule(c *fiber.Ctx){
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var Schedule Schedule
	json.Unmarshal([]byte(c.Body()), &Schedule)

	update := bson.M{
		"$set": Schedule,
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

func deleteSchedule(c *fiber.Ctx) {
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


func getFlexiWeek(c *fiber.Ctx) {

	c.Send("This is a big obj")
}

func getFlexiYr(c *fiber.Ctx) {

	c.Send("This is a big obj")
}
func postFlexiYr(c *fiber.Ctx) {

	c.Send("Need to set up db")
}
func deleteFlexiYr(c *fiber.Ctx) {

	c.Send("All Books")
}
