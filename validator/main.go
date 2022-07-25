package main

import (
    "encoding/json"
    "fmt"
    "github.com/go-playground/validator/v10"
)

type Users struct {
    Phone  string `form:"phone" json:"phone" validate:"required"`
    Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
    Code   string `form:"code" json:"code" validate:"required,len=6"`
    Num    interface{}   `form:"num" json:"num" validate:"number"`
}

type Users0 struct {
    Phone  string `form:"phone" json:"phone" validate:"required"`
    Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
    Code   string `form:"code" json:"code" validate:"required,len=6"`
    Num    *int   `form:"num" json:"num" validate:"number,gte=0"`
}

func main() {
    // num := 0
    userJson := "{\"phone\":\"13800138000\",\"passwd\":\"123456\",\"code\":\"123456\",\"num\":\"qweqw\" }"
    // userJson := "{\"phone\":\"13800138000\",\"passwd\":\"123456\",\"code\":\"123456\"}"
    fmt.Println(userJson)
    // users := &Users{
    //     Phone:  "1326654487",
    //     Passwd: "123456",
    //     Code:   "123456",
    //     // Num:    num,
    // }
    var users Users0
    json.Unmarshal([]byte(userJson), &users)
    fmt.Println(users)
    fmt.Println(*users.Num)
    // numN := users.Num.(float64)
    // fmt.Println(int(numN))
    validate := validator.New()
    err := validate.Struct(users)
    if err != nil {
        errors := err.(validator.ValidationErrors)
        fmt.Println(errors) //Key: 'Users.Passwd' Error:Field validation for 'Passwd' failed on the 'min' tag
        return
        //for _, err := range err.(validator.ValidationErrors) {
        //    fmt.Println(err) //Key: 'Users.Passwd' Error:Field validation for 'Passwd' failed on the 'min' tag
        //    return
        //}
    }
    return
}
