package utils

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	syaml "gopkg.in/yaml.v3"
)

func PrintFormattedData(data interface{}, format string) {
	switch format {
	case "json":
		printJSON(data)
	case "yaml":
		printYAML(data)
	case "xml":
		printXML(data)
	case "csv":
		printCSV(data)
	default:
		PrintVerbose(data, false)
	}
}

func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Print("error encoding json:", err)
		return
	}

	if string(jsonData) == "null" {
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

		if val.Len() == 0 {
			fmt.Println("Empty slice/array")
			return
		}

		firstElem := val.Index(0)
		if firstElem.Kind() == reflect.Ptr {
			if firstElem.IsNil() {
				firstElem = reflect.New(firstElem.Type().Elem()).Elem()
			} else {
				firstElem = firstElem.Elem()
			}
		}

		var header []string
		elemType := firstElem.Type()
		for i := 0; i < elemType.NumField(); i++ {
			header = append(header, elemType.Field(i).Tag.Get("json"))
		}
		writer.Write(header)

		for i := 0; i < val.Len(); i++ {
			var row []string
			elem := val.Index(i)

			if elem.Kind() == reflect.Ptr {
				if elem.IsNil() {
					for j := 0; j < elemType.NumField(); j++ {
						row = append(row, "")
					}
				} else {
					elem = elem.Elem()
					for j := 0; j < elem.NumField(); j++ {
						fieldValue := getString(elem.Field(j))
						row = append(row, fieldValue)
					}
				}
			} else {
				for j := 0; j < elem.NumField(); j++ {
					fieldValue := getString(elem.Field(j))
					row = append(row, fieldValue)
				}
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
	case reflect.Slice, reflect.Array:
		if field.IsNil() {
			return "[]"
		}
		var elements []string
		for i := 0; i < field.Len(); i++ {
			elem := field.Index(i)
			elements = append(elements, getString(elem))
		}
		return "[" + strings.Join(elements, ",") + "]"
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

func printYAML(data interface{}) {
	yamlData, err := syaml.Marshal(data)
	if err != nil {
		fmt.Print("error encoding yaml:", err)
		return
	}
	fmt.Println(string(yamlData))
}

func printXML(data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		xmlStr, err := mapToXML(v, "root")
		if err != nil {
			fmt.Print("error encoding xml:", err)
			return
		}
		fmt.Println(xmlStr)
	case []map[string]interface{}:
		fmt.Println("<root>")
		for _, item := range v {
			xmlStr, err := mapToXML(item, "item")
			if err != nil {
				fmt.Print("error encoding xml:", err)
				return
			}
			fmt.Println(xmlStr)
		}
		fmt.Println("</root>")
	default:
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			fmt.Println("<root>")
			for i := 0; i < val.Len(); i++ {
				elem := val.Index(i).Interface()
				b, err := json.Marshal(elem)
				if err == nil {
					var m map[string]interface{}
					if err := json.Unmarshal(b, &m); err == nil {
						xmlStr, err := mapToXML(m, "item")
						if err != nil {
							fmt.Print("error encoding xml:", err)
							return
						}
						fmt.Println(xmlStr)
						continue
					}
				}

				xmlData, err := xml.MarshalIndent(elem, "", "  ")
				if err != nil {
					fmt.Print("error encoding xml:", err)
					return
				}
				fmt.Println(string(xmlData))
			}
			fmt.Println("</root>")
			return
		}

		b, err := json.Marshal(data)
		if err == nil {
			var m map[string]interface{}
			if err := json.Unmarshal(b, &m); err == nil {
				xmlStr, err := mapToXML(m, "root")
				if err != nil {
					fmt.Print("error encoding xml:", err)
					return
				}
				fmt.Println(xmlStr)
				return
			}
		}

		xmlData, err := xml.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Print("error encoding xml:", err)
			return
		}
		fmt.Println(string(xmlData))
	}
}

func mapToXML(data map[string]interface{}, rootName string) (string, error) {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("<%s>", rootName))
	for k, v := range data {
		buf.WriteString(fmt.Sprintf("<%s>", k))
		switch val := v.(type) {
		case map[string]interface{}:
			inner, _ := mapToXML(val, k)
			buf.WriteString(inner)
		case []interface{}:
			for _, item := range val {
				if m, ok := item.(map[string]interface{}); ok {
					inner, _ := mapToXML(m, k)
					buf.WriteString(inner)
				} else {
					buf.WriteString(fmt.Sprintf("%v", item))
				}
			}
		default:
			buf.WriteString(fmt.Sprintf("%v", val))
		}
		buf.WriteString(fmt.Sprintf("</%s>", k))
	}
	buf.WriteString(fmt.Sprintf("</%s>", rootName))
	return buf.String(), nil
}
