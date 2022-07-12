package service1

import (
	"fmt"
	"time"
)

type Service struct {

}

func (service *Service) Print0()  {
	go func() {
		for {
			fmt.Println("service1")
			time.Sleep(time.Second * 1)
		}
	}()
}
