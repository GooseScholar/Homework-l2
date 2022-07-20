package models

//Result модель результата
type Result struct {
	Result Answer `json:"result"`
}

//NewResult конструктор результата
func NewResult(date string, id string, procedure string) *Result {
	return &Result{
		Result: Answer{
			Date:      date,
			ID:        id,
			Procedure: procedure,
		},
	}
}

//Answer модель ответа
type Answer struct {
	ID        string `json:"user_id"`
	Date      string `json:"date"`
	Procedure string `json:"procedure"`
}

//Post модель POST запроса
type Post struct {
	ID      string `json:"user_id"`
	Date    string `json:"date"`
	NewDate string `json:"new_date"`
}

//Err модель ошибки
type Err struct {
	Error     string `json:"error"`
	Procedure string `json:"procedure"`
}

//NewErr конструктор ошибки
func NewErr(err string, procedure string) *Err {
	return &Err{
		Error:     err,
		Procedure: procedure,
	}
}
