package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"t11/internal/cache"
	"t11/internal/eventerrors"
	"t11/internal/models"
)

//DeleteEvent удаление события
func DeleteEvent(w http.ResponseWriter, r *http.Request, cache *cache.Cache) {
	if r.Method == "POST" {
		//чтение BODY
		b, errBody := ioutil.ReadAll(r.Body) // do error checking!
		if errBody != nil {
			log.Println(errBody)
			return
		}
		defer r.Body.Close()

		//копирование BODY согласно модели models.Post{}
		jsn := models.Post{}
		errUnm := json.Unmarshal(b, &jsn)
		if errUnm != nil {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutError(errUnm, "copy body"))))
			return
		}
		log.Println(jsn)

		//валидация даты
		date := jsn.Date
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
		id := jsn.ID
		mID, _ := regexp.MatchString(`^\d+$`, id)
		if mID != true {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("invalid id format", "id validation"))))
			return
		}

		//удаление события
		cache.DeleteEvent(date, id)

		//ответ
		out, err := json.Marshal(models.NewResult(date, id, "event removed"))
		if err != nil {
			io.WriteString(w, fmt.Sprintf(string(eventerrors.OutError(err, "response formation"))))
			return
		}

		io.WriteString(w, fmt.Sprintf(string(out)))

	} else {
		io.WriteString(w, fmt.Sprintf(string(eventerrors.OutNewError("wrong method, want method POST", "method check"))))
	}
}
