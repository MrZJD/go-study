package base

// error Interface
// type error interface {
// 	Error() string
// }

// errors.New("error msg: -> such as ...")

// 自定义error
type DivideError struct {
	dividee int
	divider int
}

func (de DivideError) Error() string {
	return "You can't inter a zero divider"
}

func divide(dividee int, divider int) (int, error) {
	if divider == 0 {
		d := DivideError{
			dividee: dividee,
			divider: 0,
		}
		return 0, d
	} else {
		return dividee / divider, nil
	}
}

func Run_error() {
	if res, err := divide(100, 0); err == nil {
		println("Your divide result is", res)
	} else {
		println(err.Error())
	}
}
