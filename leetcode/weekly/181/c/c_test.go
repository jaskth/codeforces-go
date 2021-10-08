// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	exampleIns := [][]string{{`[[2,4,3],[6,5,2]]`}, {`[[1,2,1],[1,2,1]]`}, {`[[1,1,2]]`}, {`[[1,1,1,1,1,1,3]]`}, {`[[2],[2],[2],[2],[2],[2],[6]]`}}
	exampleOuts := [][]string{{`true`}, {`false`}, {`false`}, {`true`}, {`true`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	exampleIns = append(exampleIns, []string{`[[4,1],[6,1]]`})
	exampleOuts = append(exampleOuts, []string{`true`})
	targetCaseNum := 6
	if err := testutil.RunLeetCodeFuncWithCase(t, hasValidPath, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-181/problems/check-if-there-is-a-valid-path-in-a-grid/