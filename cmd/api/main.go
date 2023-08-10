package main

import (
	"fmt"
	"log"
	"sweetweather.test/internal/api"
	"sweetweather.test/internal/usecase/calculate"
)

func main() {
	var err chan error

	calcUC := calculate.NewCalculate()

	srv := api.NewServer(calcUC)
	go func() {
		port := 3000
		fmt.Println(fmt.Sprintf("Listening on %d", port))
		err <- srv.Start(port)
	}()

	if srvErr := <-err; srvErr != nil {
		log.Fatal(srvErr)
	}
}
