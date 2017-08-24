/*
함수 정의 및 등록해서 입력에 따른 Msg에 따라 실행함수 분기
*/
package test_execute

import (
	"fmt"
)

var Executors = make(map[string]Executor)

type Executor interface {
	Exec()
}

func init() {
	var hello HelloExecutor
	Register("hello", hello)

	var test TestExecutor
	Register("test", test)
	
	var def DefaultExecutor
	Register("default", def)
}

func Register(execName string, executor Executor) {
	if _, exists := Executors[execName]; exists {
		fmt.Println(execName, "실행함수가 이미 등록되었습니다.")
	}

	fmt.Println("등록 완료:", execName, " 실행함수")
	Executors[execName] = executor
}

type HelloExecutor struct{}

func (hello HelloExecutor) Exec() {
	fmt.Println("Call Hello")
}

type TestExecutor struct{}

func (test TestExecutor) Exec() {
	fmt.Println("Call Test")
}

type DefaultExecutor struct{}

func (def DefaultExecutor) Exec(){
	fmt.Println("Call Default")
}
