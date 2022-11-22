package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbMzYsNTcsODAsNTY0XSwiY2FydGFvLXBvc3RhZ2VtIjp7Im51bWVybyI6IjAwNzM0MjM3NzciLCJjb250cmF0byI6Ijk5OTIxNTc5NzAiLCJkciI6MzYsImFwaSI6WzU3LDgzXX0sImlwIjoiMTkxLjE2Mi4xMjQuNywxOTEuMTYyLjEyNC43IiwiaWF0IjoxNjY4Nzc3NjQwLCJpc3MiOiJ0b2tlbi1zZXJ2aWNlIiwiZXhwIjoxNjY4ODY0MDQwLCJqdGkiOiJmNzZiZWI5Zi1kZjA0LTQxZDYtYTlmMy02ODU1M2UzZDM0MmQifQ.3q4lXeMbdBnoSH32JG5_m8DemtSyyAmGk_zwv3EY8_J67rD_Gdifx0KA4TtV17ECIbbD0GUl9vQGSeHx-t4d5mIyzY3nExtmn8kEXMUmhndVNEj8sucK2PCeKrO6GfgYEQRTm8rtGGQCpWv8i7WwW6msXs2SotNrKJSWAov3enGXYZimzbJuUPh5YcfkoAwNvC7eui1PEI5WcvFCCj2DYloYMVJOKJeTqjpjAqb8rUjUUTlVjCTT_SMeDs9L8LMCC6-QVQphfTrTyugoc07jgdoxXtJZDYDquF94-isGLExHhkAF99J43PWBvtj2dQ_lvUXF6-uOgb0-3Nl_EpakRg",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
