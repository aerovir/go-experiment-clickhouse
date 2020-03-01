package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"reflect"
	"strconv"
)

type Parser struct {
	DataReader *csv.Reader
	Header     []string
}

func Parse(path string, hasHeader bool) (*Parser, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	parser := &Parser{
		DataReader: csv.NewReader(file),
	}

	if hasHeader {
		parser.Header, err = parser.DataReader.Read()
		if err != nil {
			return nil, err
		}
	}

	return parser, nil
}

func (c *Parser) Next() (map[string]string, error) {
	row, err := c.DataReader.Read()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for col, value := range row {
		key := strconv.Itoa(col)
		if len(c.Header) > col {
			key = c.Header[col]
		}
		result[key] = value
	}

	return result, nil
}

func (c *Parser) FillStruct(data map[string]string, s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return errors.New("expect pointer")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("expect pointer to struct")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := ""

		if field.Tag != "" {
			if column, ok := field.Tag.Lookup("column"); ok {
				if mapValue, mapOk := data[column]; mapOk {
					value = mapValue
				}
			}
		}

		if err := c.castToField(v, field.Name, value); err != nil {
			return err
		}
	}

	return nil
}

func (c *Parser) castToField(s reflect.Value, field, data string) error {
	fieldVal := s.FieldByName(field)
	fieldType := fieldVal.Type()

	if !fieldVal.CanSet() {
		return errors.New("can not set any value to field [" + field + "]")
	}

	switch fieldType.Kind() {
	case reflect.String:
		fieldVal.SetString(data)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		fieldVal.SetInt(val)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(data, 10, 64)
		if err != nil {
			return err
		}
		fieldVal.SetUint(val)
		return nil
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return err
		}
		fieldVal.SetFloat(val)
		return nil
	case reflect.Bool:
		val := true
		if data != "" && data != "0" && data != "false" {
			val = false
		}
		fieldVal.SetBool(val)
		return nil
	}

	return errors.New("unsupported field type [" + fieldType.Kind().String() + "] in struct")
}
