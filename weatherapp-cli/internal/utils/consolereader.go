package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Function reads the input from a console and returns the value
func ReadConsoleString() string {
	data, _ := r.ReadString('\n')
	data = strings.Replace(data, "\r\n", "", -1) // Note: Windows fix
	data = strings.Replace(data, "\n", "", -1)   // Note: Windows fix
	return data
}

// Function reads the input from a console and returns the value
func ReadConsoleInt64() int64 {
	data, _ := r.ReadString('\n')
	data = strings.Replace(data, "\r\n", "", -1) // Note: Windows fix
	data = strings.Replace(data, "\n", "", -1)   // Note: Windows fix
	dataInt, _ := strconv.ParseInt(data, 10, 64)
	return dataInt
}

func ReadConsoleFloat64() float64 {
  data, _ := r.ReadString('\n')
	data = strings.Replace(data, "\r\n", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	dataFloat, _ := strconv.ParseFloat(data, 64)
	return dataFloat
}

var r *bufio.Reader

func init() {
	r = bufio.NewReader(os.Stdin)
}
