package form2json

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strings"
)

type f2j struct {
	input []input
}
type input struct {
	keys  []string
	value string
}
type der map[string]interface{}

//Unmarshal form转json
func Unmarshal(s string, str interface{}) (string, error) {
	var (
		err error
		vs  url.Values
	)
	F := new(f2j)
	vs, err = url.ParseQuery(s)
	if err != nil {
		return "", err
	}
	for k, r := range vs {
		k = strings.Replace(k, "]", "", -1)
		F.input = append(F.input, input{
			keys:  strings.Split(k, "["),
			value: r[0],
		})
	}
	data := F.unmarshal([]string{}, 0)
	d, _ := json.Marshal(data)
	jsonStr := regexp.MustCompile(`:\"([1-9]+(\.?)+[0-9]+|true|false)\"`).ReplaceAllString(string(d), ":$1")
	if str != nil {
		err = json.Unmarshal([]byte(jsonStr), str)
	}
	return jsonStr, err
}

func (F *f2j) unmarshal(prefix []string, skep int) der {
	data := make(der)
	nextSkep := skep + 1
	for _, r := range F.input {
		if len(r.keys) >= nextSkep {
			if strings.Join(r.keys[0:skep], ".") == strings.Join(prefix, ".") {
				if len(r.keys) == nextSkep {
					data[r.keys[skep]] = r.value
				} else {
					data[r.keys[skep]] = F.unmarshal(r.keys[0:nextSkep], nextSkep)
				}
			}
		}
	}
	return data
}
