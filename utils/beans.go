package utils

import (
	"errors"
	"fmt"
	"k8s.io/klog/v2"
	"reflect"
)

// 存储结构类实例的容器
var beans = make(map[string]interface{})

// Register 将结构类实例指针注册到容器中
func Register(bean interface{}) error {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr {
		return errors.New("register value is not pointer type")
	}

	// 存入容器中
	elem := t.Elem()
	beans[fmt.Sprintf("%v/%v", elem.PkgPath(), elem.Name())] = bean
	klog.Infof("register %v -> %v", fmt.Sprintf("%v/%v", elem.PkgPath(), elem.Name()), bean)
	return nil
}

// Obtain 获取结构体指针实例
func Obtain(bean interface{}) interface{} {
	t := reflect.TypeOf(bean)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	klog.Infof("obtain key: %v", fmt.Sprintf("%v/%v", t.PkgPath(), t.Name()))
	return beans[fmt.Sprintf("%v/%v", t.PkgPath(), t.Name())]
}
