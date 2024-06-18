// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/gym/105139/problem/L
// https://codeforces.com/gym/105139/status/L
func Test_cfL(t *testing.T) {
	testCases := [][2]string{
		{
			`3
3 4
10 15
2 4`,
			`10
25
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cfL)
}