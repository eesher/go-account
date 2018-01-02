package main

import (
	"encoding/json"
	"go-account/util"
	"io"
	"net/http"
	//"reflect"
	//"regexp"
	//"strconv"
)

/*
type Login struct {
	Platform  string `json:"platform"`
	Device_id string `json:"device_id"`
}

var type_registry = make(map[string]reflect.Type)
*/

/*
func RegisterType(name string, elem interface{}) {
	type_registry[name] = reflect.TypeOf(elem).Elem()
}

func NewData(name string) (interface{}, bool) {
	elem, ok := type_registry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

func RoutesInit() {
	//type_registry["/login"] = Login
	//RegisterType("/login", (*Login)(nil))
		param_check["/auth"] = map[string]bool{
			"platform":     true,
			"channel":      true,
			"user":         true,
			"device_id":    true,
			"access_token": true,
		}
}

var (
	user_re   = "[^a-zA-Z0-9._@-]+|([aA][dD][mM][iI][nN])|([rR][oO][bB][oO][tT])|([gG][uU][eE][sS][tT])"
	passwd_re = "[^a-z0-9]+"
)
*/

type handler func(w http.ResponseWriter, r *http.Request)
type Method func(msg map[string]string, json_data *[]byte)

func Handler(pass Method) handler {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var json_data []byte
		var data = map[string]string{}
		//msg := make(map[string]string)
		if r.Header.Get("Content-Type") == "application/json" {
			/*
				data, ok := NewData(r.URL.Path)
				if !ok {
					json_data, _ = json.Marshal(GetResult(31001))
					io.WriteString(w, string(json_data))
					return
				}
			*/
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				json_data, _ = json.Marshal(util.GetResult(util.ERRCODE.INVA_PARAM))
				io.WriteString(w, string(json_data))
				return
			} /* else {
				for k, v := range data {
					switch t := v.(type) {
					case bool:
						msg[k] = strconv.FormatBool(t)
					case int:
						msg[k] = strconv.Itoa(t)
					case string:
						msg[k] = t
					}
				}
			}
			*/
		} else if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			r.ParseForm()
			for param := range r.Form {
				data[param] = r.Form.Get(param)
			}
		}

		pass(data, &json_data)
		io.WriteString(w, string(json_data))
		/*
			params := param_check[r.URL.Path]
			for param, is_necessary := range params {
				if is_necessary {
					tmp := msg[param]
					if tmp == "" {
						json_data, _ = json.Marshal(GetResult(31001))
						io.WriteString(w, string(json_data))
						return
					}
				}
			}

			if val, ok := msg["user"]; ok {
				matcon_data, _ = json.Marshal(GetResult(31001))
					io.WriteString(w, string(json_data))
					, _ := regexp.MatchString(user_re, val);
				if match || len(val) < 6 {
					json_data, _ = json.Marshal(GetResult(31002))
					io.WriteString(w, string(json_data))
					return
				}
			}
			if val, ok := msg["passwd"]; ok {
				match, _ := regexp.MatchString(passwd_re, val);
				if match || len(val) < 6 {
					json_data, _ = json.Marshal(GetResult(31003))
					io.WriteString(w, string(json_data))
					return
				}
			}

			if val, ok := msg["device_id"]; ok {
				if len(val) < 6 {
					json_data, _ = json.Marshal(GetResult(31004))
					io.WriteString(w, string(json_data))
					return
				}
			}
			pass(msg, &json_data)
			io.WriteString(w, string(json_data))
		*/
	}
}
