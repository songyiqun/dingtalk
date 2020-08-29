package request

import (
	"fmt"
	"os"
	"reflect"
)

const ERR_CODE_PARAM_MISSING = 40
const ERR_CODE_PARAM_INVALID = 41
const ERR_MSG_PARAM_MISSING = "client-error:Missing required arguments:%s"
const ERR_MSG_PARAM_INVALID = "client-error:Invalid arguments:%s"

func ValidateRequired(name string, value interface{}) {
	if value == nil {
		fmt.Println(ERR_CODE_PARAM_MISSING, fmt.Sprintf(ERR_MSG_PARAM_MISSING, name))
		os.Exit(ERR_CODE_PARAM_MISSING)
	}
	if reflect.TypeOf(value).Kind() == reflect.String {
		strValue := value.(string)
		if strValue == "" {
			fmt.Println(ERR_CODE_PARAM_MISSING, fmt.Sprintf(ERR_MSG_PARAM_MISSING, name))
			os.Exit(ERR_CODE_PARAM_MISSING)
		}
	}
	//fmt.Println(reflect.TypeOf(value).Kind())
}
