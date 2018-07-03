package simplestore

import (
	"reflect"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)


type Store struct {
	Region string
}

func (s *Store) Get(obj interface{}) error {

	var err error

	data := newStuctData(obj)

	var params []*ssm.Parameter

	if len(data.plain.fields) > 0 {

		if plainFields, err := getParameters(data.plain.fields, s.Region, false); err == nil {
			params = append(params, plainFields...)
		} else {
			return err
		}
	}

	if len(data.encrypted.fields) > 0 {

		if secureFields, err := getParameters(data.encrypted.fields, s.Region, true); err == nil {
			params = append(params, secureFields...)
		} else {
			return err
		}
	}

	val := reflect.ValueOf(obj).Elem()

	for _, field := range params {
		if *field.Type == "SecureString" {
			i := data.encrypted.position[*field.Name]
			val.Field(i).SetString(*field.Value)
		}

		if *field.Type == "String" {
			i := data.plain.position[*field.Name]
			val.Field(i).SetString(*field.Value)
		}
	}

	return err
}

// Fetch a single parameter without using a struct; Returns a string if successful.
func (s *Store) GetParameter(name string, withEncryption bool) (string, error) {

	result, err := getParameter(name, s.Region, withEncryption)

	if err != nil {
		return "", err
	}

	return *result.Value, nil
}

// Put a single parameter without using a struct; Returns the version number if successful
func (s *Store) PutParameter(name, value, parameterType, description string, overwrite bool) (int, error) {

	result, err := putParameter(name, value, parameterType, description, s.Region, overwrite)

	if err != nil {
		return 0, err
	}

	return int(*result.Version), err
}

func (s *Store) DeleteParameter(name string) error {
	_, err := deleteParameter(name, s.Region)
	return err
}



type structData struct {
	plain struct {
		fields   []*string      // list of parameter store names
		position map[string]int // list of names -> struct field position
		values map[string]string
	}

	encrypted struct {
		fields   []*string       // list of parameter store names
		position map[string]int  // list of names -> struct field position
		values map[string]string
	}
}





func newStuctData(obj interface{}) structData {

	data := structData{}
	data.plain.position = make(map[string]int)
	data.plain.values = make(map[string]string)
	data.encrypted.position = make(map[string]int)
	data.encrypted.values = make(map[string]string)

	// Use reflection to loop over the structs fields,
	// evaluating and recording the tags
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {

		// valueField := val.Field(i) // The struct field name
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		// Store parameter name
		storeParamName := tag.Get("ssm_key")
		fieldPosition := i

		// Ignore fields that don't have an ssm_key tag
		if storeParamName != "" {

			encrypted := strings.ToUpper(tag.Get("ssm_type")) == "SECURESTRING"

			if encrypted {
				data.encrypted.fields = append(data.encrypted.fields, &storeParamName)
				data.encrypted.position[storeParamName] = fieldPosition
			} else {
				data.plain.fields = append(data.plain.fields, &storeParamName)
				data.plain.position[storeParamName] = fieldPosition
			}
		}
	}

	return data
}