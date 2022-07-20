package eventerrors

import (
	"encoding/json"
	"fmt"
	"log"
	"t11/internal/models"
)

//OutNewError log новой ошибки, Marshal новой ошибки
func OutNewError(err string, procedure string) string {
	log.Println(err)
	out, _ := json.Marshal(models.NewErr(err, procedure))
	return string(out)
}

//OutError log ошибки, Marshal ошибки
func OutError(err error, procedure string) string {
	log.Println(err)
	out, _ := json.Marshal(models.NewErr(fmt.Sprint(err), procedure))
	return string(out)
}
