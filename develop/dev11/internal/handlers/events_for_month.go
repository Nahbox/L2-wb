package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Nahbox/L2-wb/develop/dev11/internal/eventmanager"
)

// EventsForMonth ...
func EventsForMonth(em *eventmanager.EventManager) http.HandlerFunc {
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

		eventMonth, err := strconv.Atoi(r.FormValue("event_month"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		// Валидируем параметры
		if userID <= 0 || eventMonth <= 0 {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Ищем данные в хранилище
		dates, err := em.ForMonth(userID, eventMonth)
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "event for month error"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: dates})

		log.Println("event for month successful!")

	}
}
