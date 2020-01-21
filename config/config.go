package config

import (
	"testbot/log"
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
	"fmt"
)

type configsect struct {
	sect string
}

var cfg map[string]map[string]interface{}

func NewCfg(sect string) *configsect {
	//TODO: add a log message if the config section was not found
	return &configsect{sect}
}

//Returns either
func (sct *configsect) GetString(ind string, def string) string {
	if v, ok := cfg[sct.sect][ind]; ok {
		t := reflect.ValueOf(v)
		if (t.Kind() == reflect.String) {
			return t.Interface().(string)
		}
	}
	return def
}

func (sct *configsect) GetFloat(ind string, def float64) float64 {
	if v, ok := cfg[sct.sect][ind]; ok {
		t := reflect.ValueOf(v)
		if (t.Kind() == reflect.Float64) {
			return t.Interface().(float64)
		}
	}
	return def
}

func sanityCheck(m map[string]interface{}) (map[string](map[string]interface{}), bool){
	res := make(map[string](map[string]interface{}))
	//The config must be of form map[string](map[string]interface{}), no top level parameters must exist. We should return the corresponding map
	t := reflect.ValueOf(m)
	if (t.Kind() != reflect.Map) {
		log.Log(log.LOG_ERROR, "First level of the config is not good")
		return res, false
	}
	for k, m2 := range m {
		vof := reflect.ValueOf(m2)
		if (vof.Kind() != reflect.Map) {
			fmt.Println(m2)
			fmt.Println(reflect.TypeOf(m2).Elem().Kind())
			log.Log(log.LOG_ERROR, "Second level of the config is not good for key " + k)
			return res, false
		}
		res[k] = vof.Interface().(map[string]interface{})
	}
	fmt.Println("Here", res)
	return res, true
}



func readFile(path string) map[string]map[string]interface{} {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Log(log.LOG_ERROR, err.Error())
	}
	log.Log(log.LOG_INFO, "Successfully opened config file")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
    json.Unmarshal(byteValue, &res)
	//TODO:
	if  conf, ok := sanityCheck(res); ok {
		return conf
	} else {
		log.Log(log.LOG_ERROR, "Error with config format") //TODO: be more explicit
		return make(map[string]map[string]interface{})
	}
}


func loadCfg(path string) {
	cfg = readFile(path)
}