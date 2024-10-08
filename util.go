package ebct

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type State struct {
	CdState string
	NmState string
}

var States = [...]State{
	{CdState: "AC", NmState: "Acre"},
	{CdState: "AL", NmState: "Alagoas"},
	{CdState: "AM", NmState: "Amazonas"},
	{CdState: "AP", NmState: "Amapá"},
	{CdState: "BA", NmState: "Bahia"},
	{CdState: "CE", NmState: "Ceará"},
	{CdState: "DF", NmState: "Distrito Federal"},
	{CdState: "ES", NmState: "Espírito Santo"},
	{CdState: "GO", NmState: "Goiás"},
	{CdState: "MA", NmState: "Maranhão"},
	{CdState: "MG", NmState: "Minas Gerais"},
	{CdState: "MS", NmState: "Mato Grosso do Sul"},
	{CdState: "MT", NmState: "Mato Grosso"},
	{CdState: "PA", NmState: "Pará"},
	{CdState: "PB", NmState: "Paraíba"},
	{CdState: "PE", NmState: "Pernambuco"},
	{CdState: "PI", NmState: "Piauí"},
	{CdState: "PR", NmState: "Paraná"},
	{CdState: "RJ", NmState: "Rio de Janeiro"},
	{CdState: "RN", NmState: "Rio Grande do Norte"},
	{CdState: "RO", NmState: "Rondônia"},
	{CdState: "RR", NmState: "Roraima"},
	{CdState: "RS", NmState: "Rio Grande do Sul"},
	{CdState: "SC", NmState: "Santa Catarina"},
	{CdState: "SE", NmState: "Sergipe"},
	{CdState: "SP", NmState: "São Paulo"},
	{CdState: "TO", NmState: "Tocantin"},
}

func In(value string, list ...string) bool {

	for _, element := range list {
		if value == element {
			return true
		}
	}

	return false
}

func NotIn(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return false
		}
	}

	return true

}

func StringPtr(value string) *string {
	return &value
}

func IntPtr(value int) *int {
	return &value
}

func FloatPtr(value float64) *float64 {
	return &value
}

func BoolPtr(value bool) *bool {
	return &value
}

func IsUUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}

func TimePtr(value time.Time) *time.Time {
	return &value
}

func Serialize(obj interface{}) string {

	// Convert back to a pretty JSON string
	prettyJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return string(prettyJSON)
}

func AnyToString(value any) (str string) {

	switch v := value.(type) {

	case int:
		str = strconv.Itoa(v)

	case int32:
		str = strconv.FormatInt(int64(v), 10)

	case int8, int16, int64:
		str = strconv.FormatInt(v.(int64), 10)

	case float32, float64:
		str = strconv.FormatFloat(v.(float64), byte('f'), -1, 64)

	case bool:
		str = strconv.FormatBool(v)

	case []string:
		str = strings.Join(v, ",")

	case string, []byte:
		str = v.(string)

	case sql.NullString:
		str = v.String

	case time.Time:
		str = v.Format(time.RFC3339)

	case nil:
		str = ""

	default:
		str, _ = ToJson(v)

	}

	return

}

func ToJson(data interface{}) (jsn string, err error) {

	var bytes []byte

	if bytes, err = json.Marshal(data); err == nil {
		jsn = string(bytes)
	}

	return

}
