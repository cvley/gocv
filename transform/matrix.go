func (m *Matrix) Sub(startH, endH, startW, endW int) (*Matrix, error) {
	if startH > endH || startW > endW || startH < 0 || startW < 0 ||
		endH > m.height || endW > m.width || startH >= m.height || startW >= m.width {
		return nil, errors.New("sub matrix parameters out of range")
	}

	sub, err := New(endH-startH, endW-startW)
	if err != nil {
		return nil, err
	}
	for h := startH; h < endH; h++ {
		for w := startW; w < endW; w++ {
			sub.Set(h-startH, w-startW, m.At(h, w))
		}
	}

	return sub, nil
}
