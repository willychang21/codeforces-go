// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, divideArray, "b.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-376/problems/divide-array-into-arrays-with-max-difference/
// https://leetcode.cn/problems/divide-array-into-arrays-with-max-difference/