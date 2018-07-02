package SimpleSSM

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

// Fetches a parameter from AWS SSM Parameter Store
func getSSMParameters(keynames []*string, region string, withDecryption bool) ([]*ssm.Parameter, error) {

	var err error

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return nil, err
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))

	params, err := ssmsvc.GetParameters(&ssm.GetParametersInput{
		Names:           keynames,
		WithDecryption: &withDecryption,
	})

	return params.Parameters, err
}

