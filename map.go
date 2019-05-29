package main

import "fmt"

/*
map函数
*/
func main() {

	//声明一个map
	var fmap map[string]string
	print(fmap)

	//可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

	fmap = make(map[string]string)

	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roman"
	countryCapitalMap["Japn"] = "Tokyo"
	countryCapitalMap["India"] = "Madrid"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是:", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"]
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("存在,首都是：", countryCapitalMap[capital])
	} else {

		fmt.Println("不存在")
	}

	//	delete函数
	fmt.Println("删除前", countryCapitalMap)
	delete(countryCapitalMap, "India")
	fmt.Println("删除后", countryCapitalMap)

}
