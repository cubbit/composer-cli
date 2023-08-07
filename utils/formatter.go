package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func PrintFormattedData(data interface{}, format string) {
	switch format {
	case "json":
		printJSON(data)
	case "csv":
		printCSV(data)
	default:
		printSemantic(data)
	}
}

func printSemantic(data interface{}) {
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldValue := val.Field(i).Interface()
			fmt.Printf("%s: %v\n", field.Name, fieldValue)
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			key := iter.Key().Interface()
			value := iter.Value().Interface()
			fmt.Printf("%v: %v\n", key, value)
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface()
			fmt.Printf("%v\n", elem)
		}
	default:
		fmt.Println("Unsupported data type")
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
			header = append(header, field.Name)
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
			header = append(header, elemType.Field(i).Name)
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

	default:
		fmt.Println("Unsupported data type")
	}
}

func getString(field reflect.Value) string {
	switch field.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		jsonData, err := json.Marshal(field.Interface())
		if err != nil {
			return ""
		}
		return string(jsonData)
	default:
		return fmt.Sprintf("%v", field.Interface())
	}
}
