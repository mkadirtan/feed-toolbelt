package common_paths

import (
	"bufio"
	"bytes"
	_ "embed"
)

//go:embed common_paths.txt
var commonPathsRaw []byte

// CommonPaths requires go build flag
// GOEXPERIMENT=rangefunc https://go.dev/wiki/RangefuncExperiment
func CommonPaths(yield func(string) bool) {
	scanner := bufio.NewScanner(bytes.NewReader(commonPathsRaw))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if !yield(scanner.Text()) {
			return
		}
	}
}
