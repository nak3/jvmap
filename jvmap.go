package jvmap

import (
	"encoding/json"
	"log"
)

// jsonValueSearch returns valueList found under searchRoot
func jsonValueSearch(decoded map[string]interface{}, searchRoot string, valueList []map[string]interface{}) []map[string]interface{} {
	for k, v := range decoded {
		if k != searchRoot {
			switch v.(type) {
			case map[string]interface{}:
				valueList = jsonValueSearch(v.(map[string]interface{}), searchRoot, valueList)
			default:
				continue
			}
		} else {
			valueMap := make(map[string]interface{})
			valueMap[k] = v
			valueList = append(valueList, valueMap)
			switch v.(type) {
			case map[string]interface{}:
				valueList = jsonValueSearch(v.(map[string]interface{}), searchRoot, valueList)
			default:
				continue
			}
		}
	}
	return valueList
}

// JsonValueMap makes map(s) list from jsondata by searchKey under the values of rootKey
func JsonValueMap(jsondata []byte, keys ...string) [][]map[string]interface{} {
	searchKey := keys[0]
	rootKey := searchKey
	if keys[1] != "" {
		rootKey = keys[1]
	}

	decoded := map[string]interface{}{}
	err := json.Unmarshal(jsondata, &decoded)
	if err != nil {
		log.Fatal(err)
	}
	var scopedData []map[string]interface{}

	// NOTE: Should not skip rootKey == searchKey, since rootKey is not only one in JSON data.
	scopedData = jsonValueSearch(decoded, rootKey, scopedData)

	var jsonValueMapList [][]map[string]interface{}
	for _, val := range scopedData {
		var foundValues []map[string]interface{}
		foundValues = jsonValueSearch(val, searchKey, foundValues)
		if foundValues != nil {
			jsonValueMapList = append(jsonValueMapList, foundValues)
		}
	}
	return jsonValueMapList
}
