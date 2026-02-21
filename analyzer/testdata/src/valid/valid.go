package valid

func Add(a, b int) (sum int) {
	return a + b
}

func Divide(a, b int) (result int, err error) {
	if b == 0 {
		return 0, nil
	}
	return a / b, nil
}

func GetUser() (name string, age int, err error) {
	return "John", 30, nil
}
