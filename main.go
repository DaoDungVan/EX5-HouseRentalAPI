package main

import (
	"fmt"
	"net/http"

	configs "HouseRental/Configs"
	routers "HouseRental/Routers"
)

func main() {
	if err := configs.ConnectMySQL(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	routers.BillRouter()

	fmt.Println("Server is running at http://localhost:9999")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
