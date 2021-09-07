package main

import (
	"log"
	"os"
	"sendmessages/dieta"

	"github.com/golift/imessage"
)

func main() {
	c := &imessage.Config{
		SQLPath:   getEnvWithFallback("IMESSAGE_DB_PATH", ""),
		QueueSize: 10,
		Retries:   3,
	}
	s, err := imessage.Init(c)
	if err != nil {
		log.Fatalln(err)
	}

	dietaRepository := dieta.NewDietaRepository(getEnvWithFallback("DIET_FILE_PATH", "/tmp/dieta.md"))
	dietaRoutes := dieta.New(&dietaRepository, s)
	dietaRoutes.SetupRoutes()

	log.Print("waiting for msgs")
	err = s.Start()
	if err != nil {
		log.Fatalln(err)
	}
	for {
	}
}

func getEnvWithFallback(env, fallback string) string {
	actual, present := os.LookupEnv(env)
	if present == false {
		return fallback
	}
	return actual
}
