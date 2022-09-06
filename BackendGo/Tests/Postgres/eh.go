func indexHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}
Now start the app again and you see it runs without errors. Here’s the full code up to here for reference.
package main

import (
   "database/sql" // add this
   "fmt"
   "log"
   "os"

   _ "github.com/lib/pq" // add this

   "github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func main() {
   connStr := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
   // Connect to database
   db, err := sql.Open("postgres", connStr)
   if err != nil {
       log.Fatal(err)
   }
   app := fiber.New()

   app.Get("/", func(c *fiber.Ctx) error {
       return indexHandler(c, db)
   })

   app.Post("/", func(c *fiber.Ctx) error {
       return postHandler(c, db)
   })

   app.Put("/update", func(c *fiber.Ctx) error {
       return putHandler(c, db)
   })

   app.Delete("/delete", func(c *fiber.Ctx) error {
       return deleteHandler(c, db)
   })

   port := os.Getenv("PORT")
   if port == "" {
       port = "3000"
   }
   log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
