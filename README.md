# serr is the simplest way I could think of to return errors with a human readable stacktrace.
### I'm not saying it's fast or that you should use it everywhere, but if you just want a stacktrace with your error, `serr.New` is a drop-in replacement for `fmt.Errorf` and `errors.New`.

###  Usage:

	err := serr.New("yeah, it failed here with %v", badStuff)
	// <meanwhile, back in main()...>
	if err := doStuff(); err != nil {
		if se, ok := err.(serr.Serr); ok {
			fmt.Println(err, "because", se.Stack())
		} else {
			fmt.Println(err)
		}
	}

which prints something like:

	yeah, it failed here with 123 because goroutine 1 [running]:
	runtime/debug.Stack(0x10235a, 0x1f, 0x1042bf28, 0x1)
		/usr/local/go/src/runtime/debug/stack.go:24 +0xc0
	serr.New(0x10235a, 0x1f, 0x1042bf28, 0x1, 0x1, 0x0, 0xca8c0, 0x4a29)
		/tmp/sandbox404242805/serr.go:29 +0x80
	main.doRiskyStuff(0x3c, 0x0, 0x0, 0x9)
		/tmp/sandbox404242805/main.go:26 +0xc0
	main.doStuff(0x6, 0x2, 0x2, 0x10410020)
		/tmp/sandbox404242805/main.go:32 +0x40
	main.main()
		/tmp/sandbox404242805/main.go:42 +0xc0
