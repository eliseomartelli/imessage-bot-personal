package main

import (
	"log"
	"os"
	"sendmessages/calendar"
	"sendmessages/dieta"
	"sendmessages/hashset"
	"sendmessages/ping_pong"
	"strings"

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

	allowed, err := getAllowedList(logger, getEnvWithFallback("ALLOWED_FILE_PATH", "allowed"))
	if err != nil {
		log.Fatalln(err)
	}

	calRepository, err := calendar.NewRepository(getEnvWithFallback("UNI_CALENDAR", ""), getEnvWithFallback("WORK_CALENDAR", ""))
	if err != nil {
		log.Fatalln(err)
	}
	calRoutes := calendar.NewRoutes(&calRepository, allowed, logger, imessage)
	calRoutes.SetupRoutes()


	dietaRepository, err := dieta.NewRepository(getEnvWithFallback("DIET_FILE_PATH", "/tmp/dieta.md"))
	if err != nil {
		log.Fatalln(err)
	}
	dietaRoutes := dieta.NewRoutes(&dietaRepository, allowed, logger, imessage)
	dietaRoutes.SetupRoutes()

	pingPongRoutes := ping_pong.NewRoutes(logger, imessage)
	pingPongRoutes.SetupRoutes()

	logger.Print("waiting for msgs")
	err = imessage.Start()
	if err != nil {
		log.Fatalln(err)
	}
	defer imessage.Stop()
	for {
	}
}

func getEnvWithFallback(env, fallback string) string {
	if actual, present := os.LookupEnv(env); present {
		return actual
	}
	return fallback
}

func getAllowedList(logger *log.Logger, path string) (*hashset.Hashset, error) {
	allowed := hashset.New()

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	allowedContacts := strings.Split(string(file), "\n")

	for _, contact := range allowedContacts {
		if len(contact) > 1 {
			allowed.Add(contact)
			logger.Printf("Added %s to allowed contacts.", contact)
		}
	}
	return allowed, nil
}
