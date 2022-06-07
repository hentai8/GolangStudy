package main

import "fmt"

func main() {
	s := fmt.Sprintf(`select * from miner_product_order_electric_item_detail where id > 0`)
	s = fmt.Sprintf(`%s and electric_item_id=?`, s)
	fmt.Println(s)
}
