package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func LoadMatrix(data []byte) ([][]int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var matrix [][]int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)

			if err != nil {
				return nil, fmt.Errorf("invalid value %q: %v", field, err)
			}

			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading matrix: %v", err)
	}

	return matrix, nil
}
