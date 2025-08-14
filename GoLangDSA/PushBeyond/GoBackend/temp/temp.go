package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// This file used for notes

// Equivalent of JSON.stringify() in GO
func main() {
	data := map[string] interface{}{
		"name": "Saleh Abdulla",
		"age":  30,
	}

	//          Marshal means stringify converts json -> string
	jsonString, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonString))

	// Equivalent of JSON.parse(), converts string -> json
	dataStringThisTime := `{"name":"John Doe", "age":30}`

	var data2 map[string] interface{}
	err2 := json.Unmarshal([]byte(dataStringThisTime), &data2)
	if err2 != nil {fmt.Println(err); return}

	fmt.Println(data)

	// Equivalent to nodemon \\ air -- there are other tool like fresh as well
	// go install github.com/cosmterk/air@latest
	
	// Equivalent to dotenv in Go
	// go get github.com/joho/godotenv
	
	err3 := godotenv.Load(".env")
	if err3 != nil {log.Fatal("Error loading .env file")}

	// Assign it to variable
	MONOGODB_URL := os.Getenv("MONGOFB_")
	
	
	// Equivalent of Go Express.js in GO Fiber
	// other frameworks called Gin and Echo
	// go get github.con/gofiber/fiber/v2

	// Equivalent of Express.js Middleware in Go

	app := fiber.New() // Initilise fiber instance
	app.Use(middleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// app.Listen(":3000")

	// Route Handling in Go
	helloHandler := func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
	
	app.Get("/", helloHandler)

	

}

func middleware(next http.Handler) http.Handler {
	return http.HandleFunc(func http.ResponseWriter, r *http.Requesnt {
		// Middleware logic

		// .
		// .
		// .
		// .
		next.ServeHTTP(w, r)
	})
}
