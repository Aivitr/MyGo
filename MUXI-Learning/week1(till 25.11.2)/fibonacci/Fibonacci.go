package main

func fibonacci(n int)int{
	if n <= 2 {
		return 1
	}

	x, y, t := 1, 1, 0 

	for i := 0; i < n-2; i ++ {
		t = x + y
		x = y
		y = t
	}

	return t
}
