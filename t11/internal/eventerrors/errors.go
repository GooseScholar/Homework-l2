package eventerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"t11/internal/models"
)

func OutNewError(err string, procedure string) string {
	log.Println(err)
	out, _ := json.Marshal(models.NewErr(err, procedure))
	log.Println(errors.New(err))
	return string(out)
}

func OutError(err error, procedure string) string {
	log.Println(err)
	out, _ := json.Marshal(models.NewErr(fmt.Sprint(err), procedure))
	return string(out)
}
