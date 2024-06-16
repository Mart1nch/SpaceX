package test

import (
	"encoding/json"
	"testing"
	"xtemplate/cmr"
)

func TestJSONOrderedMapStruct(t *testing.T) {
	jsonString := `{
		"yolo": "covfefe",
		"stuff": {
			"b": "12",
			"d": "654",
			"a": "1"
		},
		"yay": 5,
		"foo": ["b","c","a"]
	}`

	var ordered cmr.JSONOrderedMap
	if err := json.Unmarshal([]byte(jsonString), &ordered); err != nil {
		t.Error(err)
		return
	}

	jsonStr, err := json.Marshal(ordered)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("JSON: %s", jsonStr)
	//fmt.Println(jsonStr == `{"foo":["b","c","a"],"stuff":{"a":"1","b":"12","d":"654"},"yay":5,"yolo":"covfefe"}`)
}
