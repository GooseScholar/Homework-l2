package models

type Result struct {
	Result Answer `json:"result"`
}

//конструктор результата
func NewResult(date string, id string, procedure string) *Result {
	return &Result{
		Result: Answer{
			Date:      date,
			Id:        id,
			Procedure: procedure,
		},
	}
}

type Answer struct {
	Id        string `json:"id"`
	Date      string `json:"date"`
	Procedure string `json:"procedure"`
}

type Post struct {
	Id      string `json:"user_id"`
	Date    string `json:"date"`
	NewDate string `json:"new_date"`
}

type Err struct {
	Error     string `json:"error"`
	Procedure string `json:"procedure"`
}

//конструктор ошибки
func NewErr(err string, procedure string) *Err {
	return &Err{
		Error:     err,
		Procedure: procedure,
	}
}
