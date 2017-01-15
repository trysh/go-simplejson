package simplejson

import (
	"strconv"
	"log"
	"fmt"
	"encoding/json"
)

// NewJsonFromStr returns a pointer to a new `Json` object
func NewJsonFromStr(jsonStr string) (*Json) {
	a, _ := NewJson([]byte(jsonStr))
	return a
}

func NewEmpJson() (*Json) {
	return NewJsonFromStr(`{}`)
}

// ToJsonStr obj to jsonStr
func (j *Json) ToJsonStr() (string) {
	a, _ := j.MarshalJSON()
	return string(a)
}

func (j *Json) ToJsonStrPretty() (string) {
	a, _ := json.MarshalIndent(&j.data,"","\t")
	return string(a)
}

// M return Map
func (j *Json) M() (map[string]interface{}) {
	a, _ := j.Map()
	return a
}

func (j *Json) GetInt(key string, defaultV ...int) int {
	var def int
	
	switch len(defaultV) {
	case 0:
	case 1:
		def = defaultV[0]
	default:
		log.Panicf("too many arguments %d", len(defaultV))
	}
	
	gj:=j.Get(key)
	
	i, err := gj.Int()
	if err == nil {
		return i
	}
	
	s,err:=gj.String()
	if err==nil{
		i,err:=strconv.Atoi(s)
		if err==nil{
			return i
		}
		f,err:=strconv.ParseFloat(s,64)
		if err==nil{
			return int(f)
		}
	}
	
	f,err:=gj.Float64()
	if err==nil{
		return int(f)
	}
	
	return def
}

func (j *Json) GetString(key string, defaultV ...string) string {
	var def string
	
	switch len(defaultV) {
	case 0:
	case 1:
		def = defaultV[0]
	default:
		log.Panicf("too many arguments %d", len(defaultV))
	}
	
	gj:=j.Get(key)
	
	s, err := gj.String()
	if err == nil {
		return s
	}
	
	i,err:=gj.Int()
	if err==nil{
		s:=strconv.Itoa(i)
		return s
	}
	
	f,err:=gj.Float64()
	if err==nil{
		return fmt.Sprintf("%f",f)
	}
	
	return def
}

func (j *Json) GetFloat(key string, defaultV ...float64) float64 {
	var def float64
	
	switch len(defaultV) {
	case 0:
	case 1:
		def = defaultV[0]
	default:
		log.Panicf("too many arguments %d", len(defaultV))
	}
	
	gj:=j.Get(key)
	
	f,err:=gj.Float64()
	if err==nil{
		return f
	}
	
	s, err := gj.String()
	if err == nil {
		f,err:= strconv.ParseFloat(s,64)
		if err==nil {
			return f
		}
	}
	
	i,err:=gj.Int()
	if err==nil{
		return float64(i)
	}
	
	return def
}