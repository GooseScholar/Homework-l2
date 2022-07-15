package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"t11/internal/cache"
	"t11/internal/eventerrors"
	"t11/internal/models"
)

//Проверить отмеченные события на текущем месяце для пользователя
func EventsForMonth(w http.ResponseWriter, r *http.Request, cache *cache.Cache) {
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

		//проверка всех дней ткущей условно недели
		answer := make([]string, 0, 31)
		for i := 0; i < 7; i++ {
			var startStr string
			if i < 10 {
				startStr = "0" + strconv.Itoa(i)
			} else {
				startStr = strconv.Itoa(i)
			}
			_, errGE := cache.GetEvents(date[:8]+startStr, id)
			if errGE != false {
				answer = append(answer, date[:8]+startStr)
			}
		}

		//вывод результата в зависимости от кол-ва обнаруженных событий
		switch len(answer) {
		case 0:
			out, err := json.Marshal(models.NewResult(date, id, "not found an event in the current month"))
			if err != nil {
				log.Println(err)
				io.WriteString(w, fmt.Sprintf("Why? %v", err))
				return
			}

			io.WriteString(w, fmt.Sprintf(string(out)))
		case 1:
			out, err := json.Marshal(models.NewResult(strings.Join(answer, " "), id, "found an event in the current month"))
			if err != nil {
				log.Println(err)
				io.WriteString(w, fmt.Sprintf(string(eventerrors.OutError(err, "response formation"))))
				return
			}

			io.WriteString(w, fmt.Sprintf(string(out)))
		default:
			out, err := json.Marshal(models.NewResult(strings.Join(answer, " "), id, "found an events in the current month"))
			if err != nil {
				log.Println(err)
				io.WriteString(w, fmt.Sprintf(string(eventerrors.OutError(err, "response formation"))))
				return
			}

			io.WriteString(w, fmt.Sprintf(string(out)))
		}

	} else {
		io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("wrong method, want method GET", "method check"))))
	}
}
