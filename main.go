package SimpleSSM

import (
	"reflect"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)


type structData struct {
	plain struct {
		fields   []*string       // list of parameter store names
		position map[string]int // list of names -> struct field position
	}

	encrypted struct {
		fields   []*string       // list of parameter store names
		position map[string]int // list of names -> struct field position
	}
}

// fields ssm_name, ssm_type
func GetParameters(obj interface{}, region string) (interface{}, error) {

	var err error
	data := structData{}
	data.plain.position = make(map[string]int)
	data.encrypted.position = make(map[string]int)

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

		encrypted := strings.ToUpper(tag.Get("ssm_type")) == "SECURESTRING"

		if encrypted {
			data.encrypted.fields = append(data.encrypted.fields, &storeParamName)
			data.encrypted.position[storeParamName] = fieldPosition
		} else {
			data.plain.fields = append(data.plain.fields, &storeParamName)
			data.plain.position[storeParamName] = fieldPosition
		}
	}

	var params []*ssm.Parameter

	if len(data.plain.fields) > 0 {
		if plainFields, err := getSSMParameters(data.plain.fields, region, false); err == nil {
			params = append(params, plainFields...)
		}
	}

	if len(data.encrypted.fields) > 0 {
		if secureFields, err := getSSMParameters(data.encrypted.fields, region, true); err == nil {
			params = append(params, secureFields...)
		}
	}

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

	return obj, err

}
