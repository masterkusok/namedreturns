package anonymous

func Process() {
	fn := func() (int, error) { // want "anonymous function has unnamed returns"
		return 42, nil
	}
	fn()
}

func Handler() {
	callback := func(x int) error { // want "anonymous function has unnamed returns"
		return nil
	}
	callback(10)
}
