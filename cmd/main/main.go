package main

import (
	_ "Api/config"
	"Api/pkg/router"
	"fmt"
	"os"
)

func main() {
	fmt.Println("http://localhost" + os.Getenv("SERVER_PORT") + "/swagger/index.html")
	e := router.New()

	e.Start(os.Getenv("SERVER_PORT"))
}
