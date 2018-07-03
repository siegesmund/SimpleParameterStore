**Simple Parameter Store**

Simple Parameter Store is a wrapper around the AWS Go SDK SSM Parameter Store API that aims to reduce boilerplate 
for simple use cases. It facilitates creation, retrieval and deletion of parameters via a simplified API using
annotated Go structs. 

```go
type Parameters struct {
	Postgres string `ssm_key:"PG_LOGIN_STRING" ssm_type:"SecureString"`
}

params := Parameters{}

err := GetParameters(&params, "us-east-1")

```