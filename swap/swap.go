package swap

func Swap(in []int) (out int) {
	for ii := 0; ii < len(in); ii++ {
		for jj := ii; jj < len(in); jj++ {
			if in[ii] > in[jj] {
				in[ii], in[jj] = in[jj], in[ii]
				out++
			}
		}
	}
	return
}
