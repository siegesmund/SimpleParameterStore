package simplestore

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

// Creates the aws session
func service(region string) (*ssm.SSM, error) {

	config := aws.Config{}

	if region != "" {
		config.Region = aws.String(region)
	}

	options := session.Options{
		Config:            config,
		SharedConfigState: session.SharedConfigEnable,
	}

	sess, err := session.NewSessionWithOptions(options)

	if err != nil {
		return nil, err
	}

	return ssm.New(sess, aws.NewConfig().WithRegion(region)), nil

}


// Fetches a set of parameters from AWS SSM Parameter Store
func getParameters(keynames []*string, region string, withDecryption bool) ([]*ssm.Parameter, error) {

	var err error


	ssmsvc, err := service(region)

	if err != nil {
		return nil, err
	}

	params, err := ssmsvc.GetParameters(&ssm.GetParametersInput{
		Names:           keynames,
		WithDecryption: &withDecryption,
	})

	return params.Parameters, err
}

// Puts an ssm parameter
func putParameter(name, value, parameterType, description, region string, overwrite bool) (*ssm.PutParameterOutput, error) {

	var err error

	ssmsvc, err := service(region)

	if err != nil {
		return nil, err
	}

	// Also: AllowedPattern
	// A regular expression used to validate the parameter value. For example, for
	// String types with values restricted to numbers, you can specify the following:
	// AllowedPattern=^\d+$

	input := ssm.PutParameterInput{
		Name: &name,
		Value: &value,
		Type: &parameterType,
		Description: &description,
		Overwrite: &overwrite,
	}

	return ssmsvc.PutParameter(&input)
}

func deleteParameter(name, region string) (*ssm.DeleteParameterOutput, error) {

	var err error

	ssmsvc, err := service(region)

	if err != nil {
		return nil, err
	}

	return ssmsvc.DeleteParameter(&ssm.DeleteParameterInput{Name: &name})

}

func deleteParameters(names []*string, region string) (*ssm.DeleteParametersOutput, error) {

	var err error

	ssmsvc, err := service(region)

	if err != nil {
		return nil, err
	}

	return ssmsvc.DeleteParameters(&ssm.DeleteParametersInput{Names: names})

}

