package util

import (
	"bytes"
	"errors"
)

func calculateColLen(headers []string, datas [][]string) ([]int, error) {
	headerSize := len(headers)
	res := make([]int, headerSize)
	for i, header := range headers {
		res[i] = len(header)
	}

	for _, cols := range datas {
		if len(cols) != headerSize {
			return nil, errors.New("data and header size not match")
		}
		for idx, col := range cols {
			res[idx] = max(res[idx], len(col))
		}
	}
	return res, nil
}

func renderLine(colLengths []int) []byte {
	content := []byte{}
	for _, v := range colLengths {
		content = append(content, '+')
		content = append(content, bytes.Repeat([]byte{'-'}, v+2)...)
	}
	content = append(content, []byte("+\n")...)
	return content
}

func renderData(data []string, colLengths []int) []byte {
	content := []byte{}
	for i, v := range data {
		content = append(content, []byte("| ")...)
		content = append(content, []byte(v)...)
		content = append(content, bytes.Repeat([]byte{' '}, colLengths[i]-len(v)+1)...)
	}
	content = append(content, []byte("|\n")...)
	return content
}

func RenderTable(headers []string, datas [][]string) (string, error) {
	colLengths, err := calculateColLen(headers, datas)
	if err != nil {
		return "", err
	}

	table := []byte{}
	// render upper bodder
	table = append(table, renderLine(colLengths)...)
	// render header
	table = append(table, renderData(headers, colLengths)...)
	// render divider
	table = append(table, renderLine(colLengths)...)
	// render data
	for _, data := range datas {
		table = append(table, renderData(data, colLengths)...)
	}
	// render bottom bodder
	table = append(table, renderLine(colLengths)...)
	return string(table), nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
