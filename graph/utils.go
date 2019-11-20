package graph

import (
	"math"
	"reflect"
	"strings"
)

var priceRipple = 0.27


func DropToXrp(drop float64) (xrp float64){
	xrp = drop/math.Pow(10, 6)
	return
}


func DropToPriceInUSD(drop int) (usd float64){
	return DropToXrp(float64(drop))*priceRipple
}



func getFieldName(tag, key string, s interface{}) (fieldname string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(key), ",")[0] // use split to ignore tag "options"
		if v == tag {
			return f.Name
		}
	}
	return ""
}
