package base

func Run_if() {

	var a int = 10
	if a < 100 {
		println("a < 100")
	} else {
		println("a >= 100")
	}

	// switch need break no more
	// but can also use break to break case
	var grade int = 80

	switch grade {
	case 90:
		println("Your Grade is 90")
	case 80:
		println("Your Grade is 80")
	case 70, 60:
		println("Your Grade is 70/60") // 多个条件
	default:
		println("Your Grade in not in [90, 80, 70/60]")
	}

	switch {
	case grade >= 90:
		println("Grade is A")
	case grade >= 80:
		println("Grade is B")
	case grade >= 70:
		println("Grade is C")
	case grade >= 60:
		println("Grade is D")
	default:
		println("Grade is E")
	}

	// switch type
	var x interface{}
	switch i := x.(type) {
	case int:
		println("Typeof grade is int", i)
	case float32:
		println("Typeof grade is float32")
	default:
		println("Other Types")
	}

	// TODO: select ? commucation ? channel ?
}
