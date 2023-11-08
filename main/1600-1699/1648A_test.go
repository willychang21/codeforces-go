// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1648/A
// https://codeforces.com/problemset/status/1648/problem/A
func TestCF1648A(t *testing.T) {
	testCases := [][2]string{
		{
			`2 3
1 2 3
3 2 1`,
			`7`,
		},
		{
			`3 4
1 1 2 2
2 1 1 2
2 2 1 1`,
			`76`,
		},
		{
			`4 4
1 1 2 3
2 1 1 2
3 1 2 1
1 1 2 1`,
			`129`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF1648A)
}
