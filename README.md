**Simple Parameter Store**

Simple Parameter Store is a wrapper around the AWS Go SDK SSM Parameter Store API that aims to reduce boilerplate 
for simple use cases. It facilitates creation, retrieval and deletion of parameters via a simplified API using
annotated Go structs. 

```go
import "simplestore"

// Create an instance of Store
store := Store{Region: "us-east-1"}

// Create a struct that will hold your parameters
type Parameters struct {
	Postgres string `ssm_key:"PG_LOGIN_STRING" ssm_type:"SecureString"`
}

// Create an instance of the Parameters struct
params := Parameters{}

// Fetch the parameters from the SSM Parameter Store
err := store.Get(&params) // Fetch parameter

if err != nil {
	panic(err)
}

fmt.Println(params.Postgres) // Print Postgres login string

```