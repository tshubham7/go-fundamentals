package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var jsonBytes = []byte(`{
    "code": 200,
	"data": {
		"message": "long unpredicted data",
		"date": "15/08/2020",
		"creation_date": "08 Sep 2020 15:09:47 +0530",
		"from": "13 JUL 2020",
		"to": "08 SEP 2020",
		"modification_date": "08 Sep 2020 15:09:47 +0530"
	}}`)

func main() {
	var data interface{}
	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		panic(err)
	}

	// these keys contains date in string format
	keys := List{"date", "creation_date", "to", "from", "modification_date"}
	finalData := iterateWithOperation(data, keys, convertDateToEpoch)
	fmt.Println(finalData)
}

func iterateWithOperation(data interface{}, keys List, fn func(interface{}) (interface{}, error)) interface{} {
	// in this case my operator function is converting string date into epoche and modifying the value
	switch t := data.(type) {
	case map[string]interface{}:
		for k, v := range t {
			if keys.contains(k) {
				cnv, err := convertDateToEpoch(v)
				if err == nil {
					t[k] = cnv
				} else {
					print(err)
				}
			} else {
				t[k] = iterateWithOperation(v, keys, fn)
			}
		}
		return t
	case []interface{}:
		for i, v := range t {
			t[i] = iterateWithOperation(v, keys, fn)
		}
		return t
	}
	return data
}

type List []string

// contains checks if a string is present in a slice
func (l *List) contains(str string) bool {
	for _, v := range *l {
		if v == str {
			return true
		}
	}
	return false
}

func convertDateToEpoch(date interface{}) (interface{}, error) {
	dateStr := date.(string)
	dateStrs := strings.Split(dateStr, " ")

	layout := "02/01/2006 00:00:00"
	if len(dateStrs) > 1 {
		dateStr = strings.Join(dateStrs[:3], "/")
		layout = "02/Jan/2006 00:00:00"
	}

	dateStr += " 00:00:00"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return t.Unix(), nil
}
