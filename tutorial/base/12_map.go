package base

func Run_map() {
	// 定义map
	// var mapVar map[key_type]data_type
	// mapVar := make(map[key_type]data_type)

	var countrys map[string]string
	countrys = make(map[string]string)

	countrys["France"] = "Paris"
	countrys["China"] = "Peking"
	countrys["Germany"] = "Berlin"
	countrys["UK"] = "London"
	countrys["US"] = "NewYork"

	for country := range countrys {
		println("Capital of", country, "is", countrys[country])
	}

	capital, hasCountry := countrys["India"]
	println("what is", capital)
	println("if i know about the country", hasCountry)

	// 删除map键值对
	delete(countrys, "France")
	for country := range countrys {
		println("Capital of", country, "is", countrys[country])
	}
}
