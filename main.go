package main

import (
	"log"
	"os"
	"sendmessages/dieta"
	"sendmessages/ping_pong"

	"github.com/golift/imessage"
)

func main() {
	logger := log.New(os.Stdout, "[imessage-bot] ", log.LstdFlags)
	c := &imessage.Config{
		QueueSize: 10,
		Retries:   3,
		Timeout:   0,
		SQLPath:   getEnvWithFallback("IMESSAGE_DB_PATH", ""),
	}
	imessage, err := imessage.Init(c)
	if err != nil {
		log.Fatalln(err)
	}

	dietaRepository := dieta.NewRepository(getEnvWithFallback("DIET_FILE_PATH", "/tmp/dieta.md"))
	dietaRoutes := dieta.NewRoutes(&dietaRepository, logger, imessage)
	dietaRoutes.SetupRoutes()

	pingPongRoutes := ping_pong.NewRoutes(logger, imessage)
	pingPongRoutes.SetupRoutes()

	logger.Print("waiting for msgs")
	err = imessage.Start()
	if err != nil {
		log.Fatalln(err)
	}
	for {
	}
}

func getEnvWithFallback(env, fallback string) string {
	if actual, present := os.LookupEnv(env); present {
		return actual
	}
	return fallback
}
