package maputils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

// StructToMap converts a struct to a map.
// It returns the map and error.
func StructToMap(v interface{}) (map[string]interface{}, error) {
	str, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(str, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// MapToStruct converts a map to a struct.
// It returns error.
func MapToStruct(omap map[string]interface{}, s interface{}) error {
	j, err := json.Marshal(omap)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(j, s); err != nil {
		return err
	} else {
		return nil
	}
}

// UpdateMapValues updates map with new values with matching keys.
// It returns the updated map.
func UpdateMapValues(omap, vmap map[string]interface{}) map[string]interface{} {
	for k, v := range vmap {
		if _, ok := omap[k]; ok {
			omap[k] = v
		}
	}

	return omap
}

// UpdateStructFromMap updates values in struct from map.
// It returns error.
func UpdateStructFromMap(s interface{}, vmap map[string]interface{}) error {
	omap, err := StructToMap(s)
	if err != nil {
		return err
	}

	omap = UpdateMapValues(omap, vmap)

	return MapToStruct(omap, s)
}

// YAMLToMap parses yaml file to convert it into a map.
// It returns the map and error.
func YAMLToMap(filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var iMap map[interface{}]interface{}
	if err = yaml.Unmarshal(data, &iMap); err != nil {
		return nil, err
	}

	sMap, _ := ConvMapKeysToString(iMap).(map[string]interface{})

	return sMap, nil
}

// ConvMapKeysToString converts map keys to string.
// It returns the updated map.
func ConvMapKeysToString(v interface{}) interface{} {
	if v != nil && reflect.TypeOf(v).Kind() == reflect.Map {
		var mStr = make(map[string]interface{})
		switch v := v.(type) {
		case map[string]interface{}:
			for key, val := range v {
				keyStr := fmt.Sprintf("%v", key)
				mStr[keyStr] = ConvMapKeysToString(val)
			}
		default:
			m, _ := v.(map[interface{}]interface{})
			for key, val := range m {
				keyStr := fmt.Sprintf("%v", key)
				mStr[keyStr] = ConvMapKeysToString(val)
			}
		}
		return mStr
	} else {
		return v
	}
}

// DeepMerge merges first map with second map.
// It returns the merged map.
func DeepMerge(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			if v != nil && reflect.TypeOf(v).Kind() == reflect.Map {
				m1[k], _ = StructToMap(v)
			} else {
				m1[k] = v
			}
		} else if m1[k] != nil && reflect.TypeOf(m1[k]).Kind() == reflect.Map {
			m1[k] = DeepMerge(m1[k].(map[string]interface{}), v.(map[string]interface{}))
		}
	}
	m3, _ := StructToMap(m1)
	return m3
}

// MapInterfaceToObject maps any interface to any object.
// It returns error.
func MapInterfaceToObject(v interface{}, obj interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}

// Keys fetches all string keys of given map.
// It returns all the keys.
func Keys(m map[string]interface{}, exclude ...string) []string {
	var keys = make([]string, 0)
	for k := range m {
		var isExclude = false
		for _, e := range exclude {
			if e == k {
				isExclude = true
				break
			}
		}
		if !isExclude {
			keys = append(keys, k)
		}
	}
	return keys
}

func init() {
	fmt.Println("map init util _____ maputils")
}
