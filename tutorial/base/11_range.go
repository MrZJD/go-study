package base

// range 用于迭代 array slice channel map (index-value)(key-value)

func Run_range() {
	nums := []int{1, 2, 3, 4}
	for ind, num := range nums { // range array
		println("in nums (ind, val) ->", ind, num)
	}

	var sum int = 0
	for _, num := range nums { // with _
		sum += num
	}
	println("Total nums is -> ", sum)

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for key, val := range kvs { // range map
		println("in maps (key, val) ->", key, val)
	}

	for i, v := range "abcdefg" {
		println("in str (ind, charcode) ->", i, v)
	}
}
