package invalid

func Add(a, b int) int { // want "function Add has unnamed returns"
	return a + b
}

func Divide(a, b int) (int, error) { // want "function Divide has unnamed returns"
	if b == 0 {
		return 0, nil
	}
	return a / b, nil
}

func GetUser() (string, int, error) { // want "function GetUser has unnamed returns"
	return "John", 30, nil
}
