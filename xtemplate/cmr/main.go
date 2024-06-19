package cmr

// customized marshal json
import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"
)

type JSONOrderedMap map[string]any

func (j JSONOrderedMap) MarshalJSON() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := genJSONMap(buf, j); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func genDefaultJSON(buf *bytes.Buffer, i any) error {
	switch v := i.(type) {
	case string:
		buf.WriteString("\"")
		buf.WriteString(v)
		buf.WriteString("\"")
	case float64:
		buf.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		buf.WriteString(strconv.FormatBool(v))
	case nil:
		buf.WriteString("null")
	default:
		bts, err := json.Marshal(v)
		if err != nil {
			return err
		}

		buf.Write(bts)
	}
	return nil
}

func genJSONArr(buf *bytes.Buffer, vals []any) error {
	buf.WriteString("[")
	for i, val := range vals {
		if i > 0 {
			buf.WriteString(",")
		}
		switch v := val.(type) {
		case []any:
			err := genJSONArr(buf, v)
			if err != nil {
				return err
			}
		case map[string]any:
			err := genJSONMap(buf, v)
			if err != nil {
				return err
			}
		default:
			err := genDefaultJSON(buf, v)
			if err != nil {
				return err
			}
		}

	}
	buf.WriteString("]")
	return nil
}

func genJSONMap(buf *bytes.Buffer, j JSONOrderedMap) error {
	keys := make([]string, 0)
	for k := range j {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	buf.WriteString("{")

	for i, k := range keys {
		if i > 0 {
			buf.WriteString(",")
		}

		buf.WriteString("\"")
		buf.WriteString(k)
		buf.WriteString("\": ")

		switch v := j[k].(type) {
		case []any:
			err := genJSONArr(buf, v)
			if err != nil {
				return err
			}
		case map[string]any:
			err := genJSONMap(buf, v)
			if err != nil {
				return err
			}
		default:
			err := genDefaultJSON(buf, v)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString("}")

	return nil
}
