package es

import (
	"reflect"

	// "github.com/k0kubun/pp"
	"github.com/ikeikeikeike/gopkg/convert"
	elastigo "github.com/mattbaird/elastigo/lib"

	"github.com/astaxie/beego"
)

func Index(model interface{}) (elastigo.BaseResponse, error) {
	t := reflect.TypeOf(model).Elem()
	pv := reflect.ValueOf(model).Elem()

	index := beego.AppConfig.String("appname")
	itype := t.Name()
	id := convert.ToStr(pv.FieldByName("Id").Int())
	// pp.Println(index, itype, id, nil, model)

	c := elastigo.NewConn()
	return c.Index(index, itype, id, nil, model)
}
