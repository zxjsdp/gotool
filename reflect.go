package gotool

import "reflect"

func GetNumberOfField(i interface{}) int {
	total := 0
	if i == nil {
		return total
	}
	v := reflect.ValueOf(i)
	return v.NumField()
}
