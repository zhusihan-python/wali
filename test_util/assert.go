package test_util

import "testing"

// 建议使用
func AssertD(t *testing.T, condition bool, des string) {
	if !condition {
		t.Fail()
		t.Error("assert failed: " + des)
	} else {
		t.Log("assert success: " + des)
	}
}

// 不建议使用
func Assert(t *testing.T, condition bool) {
	if !condition {
		t.Fail()
	}
}

func AssertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fail()
	}
}

func AssertEqualWithDes(t *testing.T, a, b interface{}, des string) {
	if a != b {
		t.Error("assert failed: " + des)
	} else {
		t.Log("assert success: " + des)
	}
}
