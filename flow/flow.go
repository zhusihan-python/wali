package flow

import (
	"fmt"
	"sync"
	"time"
)

var DebugState = false

type FunAndNamePair struct {
	name     string
	realFunc func()
}


// flower 针对每一步会做一些日志、metrics 相关的通用工作
// 如果提供了 logMessage 那会在错误的时候打日志
// 如果提供了 errorCallBack 那在异常时候会调用 errorCallBack
// 通过 Do 和 AndThan 方法来构造无分支的调用链，最后调用Run 方法来调用
// 必须要设置 continueChecker 用于确定每一步时候应该继续
type Flow struct {
	startTime       time.Time
	name            string
	timerName       string
	errorCallBack   func()
	endCallBack     func(success bool)
	logMessage      string
	logArgs         []interface{}
	continueChecker func() bool
	pipeline        []FunAndNamePair
	debugMode       bool
	//单步变化的参数
	stepName          string
	ignoreStepMetrics bool
	runningInStep     bool
	timerMap          map[string]time.Duration
}

// 不需要打 step metrics，只是你用了 pipeline 功能
func NewFlow(name string, continueChecker func() bool) *Flow {
	return &Flow{
		name:            name,
		continueChecker: continueChecker,
		debugMode:       DebugState,
	}
}

func (flow *Flow) ResetDebugMode(enable bool) {
	flow.debugMode = enable
}

func (flow *Flow) WithLogOnError(message string, args ...interface{}) {
	flow.logMessage = message
	flow.logArgs = args
}

// doFunc 返回 false 停止下一步调用
//todo p1 回调函数 传入step name
func (flow *Flow) ErrorCallback(do func()) {
	flow.errorCallBack = do
}

func (flow *Flow) FlowEndCallBack(do func(success bool)) {
	flow.endCallBack = do
}
func (flow *Flow) GetCurrentStepName() string {
	return flow.stepName
}

// doFunc 返回 false 停止下一步调用
func (flow *Flow) AndThen(stepName string, doFunc func()) *Flow {
	return flow.Do(stepName, doFunc)
}

func (flow *Flow) Do(stepName string, doFunc func()) *Flow {
	flow.pipeline = append(flow.pipeline, FunAndNamePair{
		name:     stepName,
		realFunc: doFunc,
	})
	return flow
}

func (flow *Flow) Run() bool {
	if flow.debugMode {
		fmt.Printf("start run flow %v with step count %d", flow.name, len(flow.pipeline))
	}
	flow.startTime = time.Now()
	flow.runningInStep = true
	var result bool
	var do FunAndNamePair
	for _, do = range flow.pipeline {
		result = flow.doRunAndCheck(do.name, do.realFunc)
		if !result {
			break
		}
	}
	if flow.endCallBack != nil {
		flow.endCallBack(result)
	}
	// reset
	flow.runningInStep = false
	flow.pipeline = nil
	return result
}

// 立即调用 do 方法，主要是用于单步执行
// do 返回一个boolean 标量，表示结果是否正确，true 表示正确
func (flow *Flow) ImmediatelyRun(step string, do func()) {
	flow.doRunAndCheck(step,  do)
}

func (flow *Flow) doRunAndCheck(step string, do func()) bool {
	flow.stepName = step
	if flow.debugMode {
		fmt.Printf("run_flow_%s on step %s start", flow.name, step)
	}
	do()
	result := flow.continueChecker()
	if flow.debugMode {
		fmt.Printf("run_flow_%s on step %s end with result %v", flow.name, step, result)
	}
	if !result {
		if flow.logMessage != "" {
			fmt.Printf(flow.logMessage, flow.logArgs)
		}
		if flow.errorCallBack != nil {
			flow.errorCallBack()
		}
	}
	return result
}

func (flow *Flow) AndParallel(stepName string, funs ...func()) *Flow {
	return flow.Do(stepName, func() {
		wg := sync.WaitGroup{}
		wg.Add(len(funs))
		for _, fun := range funs {
			fun := fun
			go func() {
				fun()
				wg.Done()
			}()
		}
		wg.Wait()
	})
}


func (flow *Flow) GetTimerMap() map[string]time.Duration {
	return flow.timerMap
}
