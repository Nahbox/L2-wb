package main

import (
	"log"
	"net/http"

	"github.com/Nahbox/L2-wb/develop/dev11/internal/eventmanager"

	"github.com/Nahbox/L2-wb/develop/dev11/internal/config"
	"github.com/Nahbox/L2-wb/develop/dev11/internal/handlers"
	"github.com/Nahbox/L2-wb/develop/dev11/internal/logging"
)

func main() {
	// init config
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init storage
	store := initStorage()

	// init handlers
	initHandlers(store)

	//init server
	initServer(cfg)
}

func initConfig() (*config.Config, error) {
	cfg, err := config.FromEnv()
	return cfg, err
}

func initStorage() *eventmanager.EventManager {
	store := eventmanager.NewEventManager()
	log.Println("Storage init successful!")
	return store
}

func initHandlers(store *eventmanager.EventManager) {
	// POST
	http.HandleFunc("/create_event", logging.LogRequest(handlers.CreateEvent(store)))
	http.HandleFunc("/update_event", logging.LogRequest(handlers.UpdateEvent(store)))
	http.HandleFunc("/delete_event", logging.LogRequest(handlers.DeleteEvent(store)))

	// GET
	http.HandleFunc("/events_for_day", logging.LogRequest(handlers.EventsForDay(store)))
	http.HandleFunc("/events_for_week", logging.LogRequest(handlers.EventsForWeek(store)))
	http.HandleFunc("/events_for_month", logging.LogRequest(handlers.EventsForMonth(store)))

	log.Println("Handlers init successful!")
}

func initServer(cfg *config.Config) {
	log.Println("Server listening")
	http.ListenAndServe(":"+cfg.AppPort, nil)
}
