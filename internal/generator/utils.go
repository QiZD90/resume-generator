package generator

import "math"

func DivideInColumns(items []string, nColumns int) [][]string {
	result := make([][]string, nColumns)
	itemsPerColumn := int(math.Ceil(float64(len(items)) / float64(nColumns)))

	for i, v := range items {
		column := i / itemsPerColumn
		result[column] = append(result[column], v)
	}

	return result
}
