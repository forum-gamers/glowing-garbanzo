package base

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateInsertQueryAndValue(table DBNAME, data any) (string, []any) {
	val := reflect.ValueOf(data)
	typ := val.Type()

	var columns, placeholders []string
	var values []any
	var placeholderCount int

	for i := range val.NumField() {
		if typ.Field(i).Name != "Id" {
			columns = append(columns, typ.Field(i).Name)
			placeholderCount++
			placeholders = append(placeholders, fmt.Sprintf("$%d", placeholderCount))
			values = append(values, val.Field(i).Interface())
		}
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, strings.Join(columns, ","), strings.Join(placeholders, ",")), values
}
