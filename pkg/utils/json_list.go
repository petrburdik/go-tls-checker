package utils

import (
	"log"

	jsoniter "github.com/json-iterator/go"
)

// JSONListTagged - get tagged objects list
func JSONListTagged(obj interface{}, tag string) []byte {
	if len(tag) == 0 {
		tag = "json"
	}
	// custom marshal with tag
	var json = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		//		MaxDepth:               -1, // encoding/json has no max depth (stack overflow at 2581101)	}.Froze()
		TagKey: tag,
	}.Froze()

	var jsonData []byte

	jsonData, err := json.Marshal(obj)

	if err != nil {
		log.Println(err)
	}
	return jsonData
}
