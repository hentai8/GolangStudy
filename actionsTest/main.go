package main

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type Config struct {
	Name     string   `toml:"name"`
	Age      int      `toml:"age"`
	Height   float64  `toml:"height"`
	IsActive bool     `toml:"is_active"`
	Hobbies  []string `toml:"hobbies"`
	Address  struct {
		Street  string `toml:"street"`
		City    string `toml:"city"`
		ZipCode string `toml:"zip_code"`
	} `toml:"address"`
}

func main() {
	var cfg Config
	err := configor.Load(&cfg, "config.toml")
	if err != nil {
		fmt.Println("read config err=", err)
	} else {
		fmt.Println(cfg)
	}
}
