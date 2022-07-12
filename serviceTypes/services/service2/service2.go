package service2

import (
	"fmt"
	"time"
)

type Service struct {

}

func (service *Service) Print0()  {
	go func() {
		for {
			fmt.Println("service2")
			time.Sleep(time.Second * 2)
		}
	}()
}
