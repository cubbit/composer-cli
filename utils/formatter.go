package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

func PrintFormattedData(data interface{}, format string) {
	switch format {
	case "json":
		printJSON(data)
	case "csv":
		printCSV(data)
	default:
		printSemantic(data, 0)
	}
}

func printSemantic(data interface{}, indentLevel int) {
	val := reflect.ValueOf(data)

	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			fmt.Println(indent(indentLevel) + "null")
			return
		}
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		if val.Type().String() == "time.Time" {
			fmt.Println(indent(indentLevel) + val.Interface().(time.Time).String())
			return
		}
		if indentLevel > 0 {
			fmt.Println()
		}
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldValue := val.Field(i)

			if !fieldValue.CanInterface() {
				continue
			}

			fieldName := field.Tag.Get("json")
			if fieldName == "" {
				fieldName = field.Name
			}
			fmt.Printf("%s%s: ", indent(indentLevel), style(fieldName, boldStyle))

			printSemantic(fieldValue.Interface(), indentLevel+1)
		}
	case reflect.Map:
		fmt.Println()
		iter := val.MapRange()
		for iter.Next() {
			key := iter.Key().Interface()
			value := iter.Value().Interface()
			fmt.Printf("%s%s: ", indent(indentLevel), style(fmt.Sprintf("%v", key), boldStyle))
			printSemantic(value, indentLevel+1)
		}
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			fmt.Println(indent(indentLevel) + "[]")
			return
		}
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface()
			printSemantic(elem, indentLevel+1)
		}
	case reflect.Bool:
		fmt.Println(fmt.Sprintf("%v", val.Bool()))
	case reflect.String:
		fmt.Println(val.String())
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println(val.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Println(val.Float())
	default:
		fmt.Println(fmt.Sprintf("%v", val.Interface()))
	}
}

func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Print("error encoding json:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func printCSV(data interface{}) {

	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Struct:
		writer := csv.NewWriter(os.Stdout)
		defer writer.Flush()
		var header []string
		var row []string

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldValue := getString(val.Field(i))
			header = append(header, field.Tag.Get("json"))
			row = append(row, fieldValue)

		}
		writer.Write(header)
		writer.Write(row)

	case reflect.Map:
		writer := csv.NewWriter(os.Stdout)
		defer writer.Flush()

		writer.Write([]string{"Key", "Value"})

		iter := val.MapRange()
		for iter.Next() {
			key := iter.Key().Interface()
			value := iter.Value().Interface()
			writer.Write([]string{fmt.Sprintf("%v", key), fmt.Sprintf("%v", value)})
		}

	case reflect.Slice, reflect.Array:
		writer := csv.NewWriter(os.Stdout)
		defer writer.Flush()

		var header []string
		elemType := val.Type().Elem()
		for i := 0; i < elemType.NumField(); i++ {
			header = append(header, elemType.Field(i).Tag.Get("json"))
		}
		writer.Write(header)

		for i := 0; i < val.Len(); i++ {
			var row []string
			elem := val.Index(i)
			for j := 0; j < elem.NumField(); j++ {
				fieldValue := getString(elem.Field(j))
				row = append(row, fieldValue)
			}
			writer.Write(row)
		}

	case reflect.Ptr:
		if val.IsNil() {
			fmt.Println("null")
		} else {
			derefVal := val.Elem().Interface()
			printCSV(derefVal)
		}

	default:
		fmt.Println("Unsupported data type")
	}
}

func getString(field reflect.Value) string {
	switch field.Kind() {
	case reflect.String:
		return field.String()
	case reflect.Bool:
		return fmt.Sprintf("%t", field.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", field.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", field.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", field.Float())
	case reflect.Ptr:
		if field.IsNil() {
			return "null"
		} else {
			return getString(field.Elem())
		}
	case reflect.Struct:
		if field.Type() == reflect.TypeOf(time.Time{}) {
			return field.Interface().(time.Time).String()
		} else {
			return structToString(field.Interface())
		}
	case reflect.Map:
		return fmt.Sprintf("%v", field.Interface())
	default:
		return fmt.Sprintf("Unsupported type: %v", field.Kind())
	}
}

func structToString(data interface{}) string {
	val := reflect.ValueOf(data)
	var str strings.Builder

	str.Write([]byte("{"))
	for i := 0; i < val.NumField(); i++ {
		if i > 0 {
			str.WriteString(", ")
		}
		field := val.Type().Field(i)
		fieldValue := val.Field(i)
		str.WriteString(field.Tag.Get("json") + ": " + getString(fieldValue))
	}
	str.Write([]byte("}"))

	return str.String()
}

func indent(level int) string {
	return strings.Repeat("  ", level)
}
