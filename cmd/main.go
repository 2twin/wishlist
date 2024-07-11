package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"wishes/internal/server"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	srv := server.NewServer()

	go srv.ListenAndServe()

	oscall := <-stop
	log.Printf("System call: %+v\n", oscall)

	err := srv.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
