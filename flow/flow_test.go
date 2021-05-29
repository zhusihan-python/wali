package flow

import (
	"fmt"
	"testing"

	"github.com/wali/test_util"
)

func TestFullRun(t *testing.T) {
	var stepResult = true
	flow := NewFlow("test", func() bool {
		return stepResult
	})

	flow.ErrorCallback(func() {
		fmt.Printf("find error on step %v \n", flow.GetCurrentStepName())
	})

	var counter int
	flow.Do("add_1", func() {
		counter++
	}).
		AndThen("add_2", func() {
		counter += 2
	}).
		AndThen("minus_100", func() {
			counter -= 100
	}).Run()
	fmt.Printf("the final result is %+v \n", counter)
	test_util.AssertEqual(t, counter, -97)
}

func TestHalfStop(t *testing.T) {
	var stepResult = true
	flow := NewFlow("test", func() bool {
		return stepResult
	})

	flow.ErrorCallback(func() {
		fmt.Printf("find error on step %v \n", flow.GetCurrentStepName())
	})
	var counter int
	flow.Do("add_1", func() {
		counter++
	}).
		AndThen("add_2", func() {
			counter += 2
		    stepResult = false
		}).
		AndThen("minus_100", func() {
			counter -= 100
		}).Run()
	fmt.Printf("the final result is %+v \n", counter)
	test_util.AssertEqual(t, counter, 3)
}