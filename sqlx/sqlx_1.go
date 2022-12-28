package main

import (
	"fmt"
	"github.com/fuck/sqlx/mysql"
)

type GetOrderPaymentParams struct {
	Id      int    `json:"id" db:"id"`
	OrderId int    `json:"user_id" db:"order_id"`
	Code    string `json:"code" db:"code"`
	Status  int    `json:"status" db:"status"`
}

type CountStruct struct {
	Count int `json:"count" db:"count(*)"`
}

type GetOrderPaymentsParams struct {
	Id       int    `query:"id" json:"id" db:"id" validate:"omitempty,gt=0"`
	OrderId  int    `query:"order_id" json:"order_id" db:"order_id" validate:"omitempty,gt=0"`
	Code     string `query:"code" json:"code" db:"code" validate:"omitempty"`
	Status   int    `query:"status" json:"status" db:"status" validate:"omitempty,oneof=1 2 3 4"`
	PageSize int    `query:"page_size" db:"page_size" json:"page_size" validate:"omitempty,gt=0,lte=100"`
	Offset   int    `query:"offset" db:"offset" json:"offset" validate:"omitempty,gte=0"`
	SortBy   string `query:"sort_by" db:"sort_by" json:"sort_by" validate:"omitempty,alphanumunicode,oneof=id"`
	OrderBy  string `query:"order_by" db:"order_by" json:"order_by" validate:"omitempty,oneof=ASC DESC"`
}

type CompleteOrderPaymentStruct struct {
	Id int `json:"id" db:"id"`
	BasicOrderPaymentStruct
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type BasicOrderPaymentStruct struct {
	OrderId        int     `json:"order_id,omitempty" db:"order_id"`
	Code           string  `json:"code" db:"code"`
	PayTime        string  `json:"pay_time" db:"pay_time"`
	PayMethod      int     `json:"pay_method" db:"pay_method"`
	PayCurrency    string  `json:"pay_currency" db:"pay_currency"`
	Amount         float64 `json:"amount" db:"amount"`
	Status         int     `json:"status" db:"status"`
	Expire         int     `json:"expire" db:"expire"`
	HashrateFee    float64 `json:"hashrate_fee" db:"hashrate_fee"`
	ElectricityFee float64 `json:"electricity_fee" db:"electricity_fee"`
	ServiceFee     float64 `json:"service_fee" db:"service_fee"`
	OriginPrice    float64 `json:"origin_fee" db:"origin_price"`
	Price          float64 `json:"price" db:"price"`
	Discount       float64 `json:"discount" db:"discount"`
}

func main() {
	//var params GetOrderPaymentParams
	//params.OrderId = 1
	//err, payments := GetOrderPayment(params)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(payments)

	var params GetOrderPaymentsParams
	params.OrderId = 1
	params.OrderBy = "DESC"
	params.SortBy = "id"
	params.Offset = 1
	err, payments := GetOrderPayments(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(payments)
	err, count := GetOrderPaymentsCount(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

}

func GetOrderPayment(params GetOrderPaymentParams) (err error, order CompleteOrderPaymentStruct) {
	db := mysql.GetInstance()
	s := `SELECT * FROM order_payment where id > 0 `
	if params.Id != 0 {
		s = fmt.Sprintf(`%s and id=:id`, s)
	}
	if params.Code != "" {
		s = fmt.Sprintf(`%s and code=:code`, s)
	}
	if params.OrderId != 0 {
		s = fmt.Sprintf(`%s and order_id=:order_id`, s)
	}
	if params.Status != 0 {
		s = fmt.Sprintf(`%s and status=:status`, s)
	}

	s = fmt.Sprintf(`%s limit 1`, s)
	rows, err := db.NamedQuery(s, params)
	if err != nil {
		println(err.Error())
		return
	}

	for rows.Next() {
		err = rows.StructScan(&order)
		if err != nil {
			println("scan failed, err :", err.Error())
			return
		}
	}
	return
}

func GetOrderPayments(params GetOrderPaymentsParams) (err error, orders []CompleteOrderPaymentStruct) {
	db := mysql.GetInstance()
	s := fmt.Sprintf(`select * from order_payment where id > 0`)
	if params.OrderId != 0 {
		s = fmt.Sprintf(`%s and order_id=:order_id`, s)
	}
	if params.Code != "" {
		s = fmt.Sprintf(`%s and code=:code`, s)
	}
	if params.Status != 0 {
		s = fmt.Sprintf(`%s and status=:status`, s)
	}
	if params.Id != 0 {
		s = fmt.Sprintf(`%s and id=:id`, s)
	}
	if params.SortBy != "" && params.OrderBy != "" {
		s = fmt.Sprintf(`%s order by %s %s`, s, params.SortBy, params.OrderBy)
	}
	if params.PageSize == 0 {
		params.PageSize = 20
	}

	s = fmt.Sprintf(`%s limit :offset, :page_size`, s)
	rows, err := db.NamedQuery(s, params)
	if err != nil {
		println(err.Error())
		return
	}
	for rows.Next() {
		var order CompleteOrderPaymentStruct
		if err := rows.StructScan(&order); err != nil {
			fmt.Printf("struct sacn failed, err:%v\n", err)
			continue
		}
		fmt.Printf("%#v\n", order)
		orders = append(orders, order)
	}
	return
}

func GetOrderPaymentsCount(params GetOrderPaymentsParams) (err error, count int) {
	db := mysql.GetInstance()
	s := fmt.Sprintf(`select count(*) from order_payment where id > 0`)
	if params.OrderId != 0 {
		s = fmt.Sprintf(`%s and order_id=:order_id`, s)
	}
	if params.Status != 0 {
		s = fmt.Sprintf(`%s and status=:status`, s)
	}

	rows, err := db.NamedQuery(s, params)

	for rows.Next() {
		var c CountStruct
		if err := rows.StructScan(&c); err != nil {
			fmt.Printf("struct sacn failed, err:%v\n", err)
			continue
		}
		fmt.Printf("%#v\n", c)
		count = c.Count
	}

	if err != nil {
		println(err.Error())
		return
	}
	return
}
