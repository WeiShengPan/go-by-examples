package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// 行过滤器
func TestLineFilters(t *testing.T) {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
