package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"t11/internal/cache"
	"t11/internal/eventerrors"
	"t11/internal/models"
)

//Проверить отмеченно ли событие в текущую дату
func EventsForDay(w http.ResponseWriter, r *http.Request, cache *cache.Cache) {
	if r.Method == "GET" {

		r.ParseForm()

		//валидация даты
		date := r.FormValue("date")
		mDate, _ := regexp.MatchString(`^20\d{2}-\d{2}-\d{2}`, date)
		if mDate != true {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("invalid date format", "date validation"))))
			return
		}

		//валидация месяца
		month, _ := strconv.Atoi(date[5:7])
		if month < 1 || month > 12 {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("invalid month format", "month validation"))))
			return
		}

		//валидация дня
		day, _ := strconv.Atoi(date[8:])
		if day < 1 || day > 31 {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("invalid day format", "day validation"))))
			return
		}

		//валидация id пользователя
		id := r.FormValue("user_id")
		mId, _ := regexp.MatchString(`^\d+$`, id)
		if mId != true {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("invalid id format", "id validation"))))
			return
		}

		//поиск события
		_, errGE := cache.GetEvents(date, id)
		if errGE == false {
			out, err := json.Marshal(models.NewResult(date, id, "no event found on this day"))
			if err != nil {
				io.WriteString(w, fmt.Sprintf(string(eventerrors.OutError(err, "response formation"))))
				return
			}

			io.WriteString(w, fmt.Sprintf(string(out)))
		} else {
			out, err := json.Marshal(models.NewResult(date, id, "event found"))
			if err != nil {
				io.WriteString(w, fmt.Sprintf(string(eventerrors.OutError(err, "response formation"))))
				return
			}

			io.WriteString(w, fmt.Sprintf(string(out)))
		}
	} else {
		io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("wrong method, want method GET", "method check"))))
	}
}
