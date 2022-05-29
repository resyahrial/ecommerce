package transformers

import "encoding/json"

func StructToMap(structs ...interface{}) (result map[string]interface{}, err error) {
	for _, v := range structs {
		if structJson, errMarshal := json.Marshal(v); errMarshal != nil {
			err = errMarshal
		} else {
			if errUnmarshal := json.Unmarshal(structJson, &result); errUnmarshal != nil {
				err = errUnmarshal
			}
		}
	}

	return
}
