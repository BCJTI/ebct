package ebct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSroRastroObjeto(t *testing.T) {

	sendObjects := []string{"QC963533014BR", "OV852367667BR", "SY530313828BR"}
	requestReturn, err := client.GetSroRastroObjetos(sendObjects)
	fmt.Println(Serialize(requestReturn))
	assert.NoError(t, err)
}
