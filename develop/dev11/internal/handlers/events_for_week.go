package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Nahbox/L2-wb/develop/dev11/internal/eventmanager"
)

// EventsForWeek ...
func EventsForWeek(em *eventmanager.EventManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			SendResponse(w, http.StatusMethodNotAllowed, Response{Error: "method error"})
			return
		}

		// Получаем параметры из тела запроса
		err := r.ParseForm()
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "parse error"})
			return
		}

		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		eventWeek, err := strconv.Atoi(r.FormValue("event_week"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		// Валидируем параметры
		if userID <= 0 || eventWeek <= 0 {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Ищем данные в хранилище
		dates, err := em.ForWeek(userID, eventWeek)
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "event for week error"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: dates})

		log.Println("event for week successful!")

	}
}
