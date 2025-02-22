package converter

import (
	"log"
	"strconv"
)

func StoI(s string) int {
	str, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Некорректный тип данных для конвертации: %v ", err)
	}
	return str

}

//func MtoS(model interface{}) string {
//	return fmt
//}
