package main

import (
	"fmt"
	"github.com/fuck/serviceTypes/services/service1"
	"github.com/fuck/serviceTypes/services/service2"
	"reflect"
	"time"
)

type ServiceInterface interface {
	Print0() // 获取每月平均币价
}

type Name struct {
	Name string `json:"name"`
}

func main() {
	var names []Name
	names = append(names, Name{Name: "service1"})
	names = append(names, Name{Name: "service2"})

	serviceMap := map[string]interface{}{
		"service1": service1.Service{},
		"service2": service2.Service{},
	}

	for _, name := range names {
		fmt.Println(name.Name)
		serviceType := reflect.TypeOf(serviceMap[name.Name])
		fmt.Println(serviceType)
		serviceInstance := reflect.New(serviceType)
		s := serviceInstance.Interface().(ServiceInterface)
		s.Print0()
	}
	time.Sleep(time.Second * 10)
}
