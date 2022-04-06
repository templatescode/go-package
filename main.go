package main

import (
	"fmt"
	"go-package/configapp"
	"go-package/external/database"
	"log"
	"time"
)

func init() {

	if err := configapp.Start(); err != nil {
		log.Panic(err)
	}

	err := database.NewDatabase().Open()
	if err != nil {
		log.Panic(err)
	}

}

func main() {

	mode := configapp.GetMode()
	fmt.Println("mode:", mode)

	data := configapp.GetData()
	fmt.Println("data:", data)

	db := database.NewDatabase()
	defer db.Close()

	conn := db.GetConnect()
	fmt.Println("conn db:", conn)

	time.Sleep(5 * time.Second)

}
