package stream

import (
	"fmt"
	"github.com/zhusihan-python/wali/flow"
)

// - Q1: 输入 employees，返回 年龄 >22岁 的所有员工，年龄总和
func Question1Sub1(employees []*Employee) int64 {
	var stepResult = true
	var adultEmployees []*Employee
	var sumAge int64
	flowQ1 := flow.NewFlow("sum_age_gt_22", func() bool {
		return stepResult
	})
	flowQ1.ErrorCallback(func() {
		fmt.Printf("find error on step %v \n", flowQ1.GetCurrentStepName())
	})
	flowQ1.Do("filter_age", func() {
        for _, e := range employees {
        	if *e.Age > 22 {
				adultEmployees = append(adultEmployees, e)
			}
		}
	}).
		AndThen("sum_age", func() {
			for _, ae := range adultEmployees {
				sumAge += int64(*ae.Age)
			}
	})
	return sumAge
}

// - Q2: - 输入 employees，返回 id 最小的十个员工，按 id 升序排序
func Question1Sub2(employees []*Employee) []*Employee {
	return nil
}

// - Q3: - 输入 employees，对于没有手机号为0的数据，随机填写一个
func Question1Sub3(employees []*Employee) []*Employee {
	return nil
}

// - Q4: - 输入 employees ，返回一个map[int][]int，其中 key 为 员工年龄 Age，value 为该年龄段员工ID
func Question1Sub4(employees []*Employee) map[int][]int64 {
	return nil
}