package matrix

func Conv3(input [][][]int64, kernel [][][]int64, origin [3]int) [][][]int64 {
	output := make([][][]int64, len(input))
	for ri := range input {
		output[ri] = make([][]int64, len(input[ri]))
		for ci := range input[ri] {
			output[ri][ci] = make([]int64, len(input[ri][ci]))
			for di := range input[ri][ci] {
				var acc int64
				for rk := range kernel {
					for ck := range kernel[rk] {
						for dk, k := range kernel[rk][ck] {
							x := ri - origin[0] + rk
							y := ci - origin[1] + ck
							z := di - origin[2] + dk
							if k != 0 && x >= 0 && y >= 0 && z >= 0 && x < len(input) && y < len(input) && z < len(input) {
								acc += input[x][y][z]
							}
						}
					}
				}
				output[ri][ci][di] = acc
			}
		}
	}
	return output
}
