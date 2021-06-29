package panic

func iWillPanic() {
	panic("something")
}

func catchPanic() {
        defer func() {
            recover()
        }()
	iWillPanic()
}
