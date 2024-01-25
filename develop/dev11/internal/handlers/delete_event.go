package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Nahbox/L2-wb/develop/dev11/internal/eventmanager"
)

// DeleteEvent ...
func DeleteEvent(em *eventmanager.EventManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendResponse(w, http.StatusMethodNotAllowed, Response{Error: "method error"})
			return
		}

		// Получаем параметр из тела запроса
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

		eventID, err := strconv.Atoi(r.FormValue("event_id"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		// Валидируем параметр
		if userID <= 0 || eventID <= 0 {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Удаляем в хранилище
		if err = em.Delete(userID, eventID); err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "delete in storage error"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: "event deleted successful!"})

		log.Println("event deleted successful!")
	}
}
