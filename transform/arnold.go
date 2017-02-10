package transform

func Arnold(data *Matrix) {
	N := data.Width()
	result, _ := New(N, N)
	for h := 0; h < N; h++ {
		for w := 0; w < N; w++ {
			x := (h + w) % N
			y := (2*h + w) % N
			result.Set(y, x, data.At(h, w))
		}
	}

	for h := 0; h < N; h++ {
		for w := 0; w < N; w++ {
			data.Set(h, w, result.At(h, w))
		}
	}
}
