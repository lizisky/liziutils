package utils

import (
	"encoding/json"
	"io/ioutil"
)

// ToJSONIndent is a helper method that converts the object to human readable format
func ToJSONIndent(v interface{}) string {
	ret, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(ret)
	} else {
		// logger.Info("MarshalIndent object to json failed", v)
		return ""
	}
}

func ToJSON(v interface{}) string {

	ret, err := json.Marshal(v)
	if err != nil {
		// logger.Info("MarshalIndent object to json failed", v)
		return ""
	}
	return string(ret)
}

func FromJson(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}

// ReadJSONFile reads a JSON format file into v
func ReadJSONFile(filename string, v interface{}) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//logger.Info("read json ",string(data))
	err = json.Unmarshal(data, v)
	return err
}
