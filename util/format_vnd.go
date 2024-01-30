package util

import "strconv"

func FormatVND(number int) string {
    output := strconv.Itoa(number)
    startOffset := 3
    if number < 0 {
        startOffset++
    }
    for outputIndex := len(output); outputIndex > startOffset; {
        outputIndex -= 3
        output = output[:outputIndex] + "," + output[outputIndex:]
    }
    return "VND " + output
}
