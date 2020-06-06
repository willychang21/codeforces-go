// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[1,4,6,7,8,20]`, `[2,7,15]`, 
			`11`,
		},
		{
			`[1,2,3,4,5,6,7,8,9,10,30,31]`, `[2,7,15]`, 
			`17`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, mincostTickets, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-121/problems/minimum-cost-for-tickets/
