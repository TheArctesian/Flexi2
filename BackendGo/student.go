/*
Student Data Struct
------------------
Name: string;
Advisory: int;
Email: string;
Picture: string; // This wll just pont to and endpoint on the server where the img i stored
FavoriteTeacher: string[];

*/

package student

import (
	"github.com/gofiber/fiber"
	
)

type student struct {
	name      string
	yearGroup int
	advisory  string
	email     string
	picture   string
	favoriteTeacher string[]
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

func createStudent(c *fiber.Ctx){
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var student student
	json.Unmarshal([]byte(c.Body()), &student)

	res, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func updatestudent(c *fiber.Ctx){
	collection, err := getMongoDbCollection(flexiDB, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var student student
	json.Unmarshal([]byte(c.Body()), &student)

	update := bson.M{
		"$set": student,
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

func getStudentName(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getStudentAdvisory(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getStudentPicture(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func getStudentTeachers(c *fiber.Ctx) {
	c.Send("The Arctesian")
}


func updateStudentName(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func updateStudentAdvisory(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func updateStudentPicture(c *fiber.Ctx) {
	c.Send("The Arctesian")
}
func updateStudentTeachers(c *fiber.Ctx) {
	c.Send("The Arctesian")
}


