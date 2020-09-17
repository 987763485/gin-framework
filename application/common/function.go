/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  function
 * @Time: 2020/9/17 10:08 上午
 */

package common

import (
	"reflect"
)

func StructureToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func ParamsToModel(params, model interface{}) map[string]interface{} {
	p := StructureToMap(params)
	m := StructureToMap(model)
	for k1, v1 := range p {
		for k2, v2 := range m {
			a := reflect.TypeOf(v1)
			b := reflect.TypeOf(v2)
			if k1 == k2 && a == b {
				m[k2] = v1
			}
		}
	}
	return m
}
