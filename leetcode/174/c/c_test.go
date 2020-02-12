// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	exampleIns := [][]string{{`[1,2,3,4,5,6]`}, {`[1,null,2,3,4,null,null,5,6]`}, {`[2,3,9,10,7,8,6,5,4,11,1]`}, {`[1,1]`}}
	exampleOuts := [][]string{{`110`}, {`90`}, {`1025`}, {`1`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFuncWithCase(t, maxProduct, exampleIns, exampleOuts, 0); err != nil {
		t.Fatal(err)
	}
}
