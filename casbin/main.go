package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"strings"

	"github.com/casbin/xorm-adapter/v2"
)

func main() {
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := xormadapter.NewAdapter("mysql", "root:1007@tcp(127.0.0.1:3306)/dxfufu", true) // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e, _ := casbin.NewEnforcer("./model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如 results[0] 是 {"alice", "data1", "read"} 的结果
	//added, err := e.AddPolicy("eve", "data3", "read")
	added, err := e.AddPolicy("eve2", "data*", "read")
	added, err = e.AddRoleForUser("user1", "admin")
	added, err = e.AddRoleForUser("admin", "business")
	added, err = e.AddRoleForUser("admin", "finance")
	added, err = e.AddRoleForUser("admin", "editor")
	added, err = e.AddRoleForUser("user2", "user")
	added, err = e.AddRoleForUser("business", "user")


	// added, err = e.AddPolicy("role2", "qqq", "read")
	added, err = e.AddPolicy("business", "qqq", "read")
	added, err = e.AddPolicy("finance", "ooo", "read")
	added, err = e.AddPolicy("editor", "iii", "read")
	added, err = e.AddPolicy("user", "sss", "read")
	added, err = e.AddPolicy("user", "ssss", "read")
	e.AddFunction("self_match", KeyMatchFunc)
	fmt.Println(added)
	fmt.Println(err)
	results, err := e.BatchEnforce([][]interface{}{{"user1", "qqq", "read"}})
	fmt.Println("result:", results)
	results, err = e.BatchEnforce([][]interface{}{{"user1", "ooo", "read"}})
	fmt.Println("result:", results)
	results, err = e.BatchEnforce([][]interface{}{{"user1", "iii", "read"}})
	fmt.Println("result:", results)
	results, err = e.BatchEnforce([][]interface{}{{"user1", "sss", "read"}})
	fmt.Println("result:", results)
	results0, err := e.Enforce("user1", "ooo", "read")
	fmt.Println("result0:", results0)
	results0, err = e.Enforce("user1", "sssss", "read")
	fmt.Println("result0:", results0)
	e.GetRolesForUser("role")
	// results, err := e.BatchEnforce([][]interface{}{{"eve2", "data3", "read"}, {"zhangSan", "data1", "read"}, {"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
	// fmt.Println("result:", results)

	roles, err := e.GetImplicitRolesForUser("role123")
	fmt.Println("roles:", roles)

	// 查找v0是alice的
	filteredPolicy := e.GetFilteredPolicy(0, "alice")
	fmt.Println(filteredPolicy)
	filteredPolicy1 := e.GetFilteredPolicy(1, "data1")
	fmt.Println(filteredPolicy1)

	// 增删改查 add remove update get

	//p就是policy g就是group 具体查文档的时候的关键词

	//keymatch 自定义函数
}

func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}
	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	if key2 == "*" {
		return true
	}
	return key1 == key2[:i]
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}

//func main() {
//	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
//	sub := "alice" // 想要访问资源的用户。
//    obj := "data1"     // 将被访问的资源。
//	act := "read"  // 用户对资源执行的操作。
//
//	ok, err := e.Enforce(sub, obj, act)
//
//	if err != nil {
//		// 处理err
//	}
//
//	if ok == true {
//		// 允许alice读取data1
//	} else {
//		// 拒绝请求，抛出异常
//	}
//
//	// 您可以使用BatchEnforce()来批量执行一些请求
//	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
//	// 例如 results[0] 是 {"alice", "data1", "read"} 的结果
//	added, err := e.AddPolicy("eve", "data3", "read")
//	fmt.Println(added)
//	results, err := e.BatchEnforce([][]interface{}{{"zhangSan", "data1", "read"},{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
//	fmt.Println(results)
//}
