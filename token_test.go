package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbMzYsNTcsODAsNTY0XSwiY2FydGFvLXBvc3RhZ2VtIjp7Im51bWVybyI6IjAwNzM0MjM3NzciLCJjb250cmF0byI6Ijk5OTIxNTc5NzAiLCJkciI6MzYsImFwaSI6WzU3LDgwLDgzXX0sImlwIjoiMTg5Ljc3LjkzLjI4LDE4OS43Ny45My4yOCIsImlhdCI6MTY2OTgyOTU2MSwiaXNzIjoidG9rZW4tc2VydmljZSIsImV4cCI6MTY2OTkxNTk2MSwianRpIjoiOWNmMTZlNDYtZjQ5OC00OGJlLWFmNWUtZjJiZTMxZjE2YWE1In0.PW6W88kQQx9uAlSueTe-eoiY9T2zFW32OIXiO9uONo8RZTOsmH5RQusv3vR6eIuLRKnIGlcD3iz095Ia1kwAPYBN2dyKYTWzfoTTYFD7CbFuxfxtYZdCTTto0fVysxxmlj13W7JaAFcv466B9GLdGUYw9dKJatOJ5zkc-r_esBSob4V9yZrt1FDRMqf56X164O7KmYQqiJaOO4Ybfr7aNveSmFS8hQJ27R_vfZPIBUtxr_lj3Ywbsb-t5ugJQl0HEL5zUia2CUykV5LpuCg8HMlcRUW0Xlk49TgjDBYMP_C-OENwmRDbjEll5KxnFDJF96Ib9bJO7HOYuH73jl-ggQ",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "JSIV2hUBvbdkFQykkxfnqSGBREwmBJhdarZXjTLW", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
