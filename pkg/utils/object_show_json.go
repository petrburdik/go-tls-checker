package utils

import (
	"bytes"
	"encoding/json"
)

// JSONPrettyPrint - udela hezky JSON
func JSONPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// ObjectPrettyPrint - tady udelej hezke zobrazeni objektu skrz JSON
func ObjectPrettyPrint(obj interface{}) string {
	jsonTmp, _ := json.Marshal(obj)
	return JSONPrettyPrint(string(jsonTmp))
}
