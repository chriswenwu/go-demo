package interview

import (
	"fmt"
	"testing"
)

/**
因为for遍历时，变量stu指针不变，每次遍历仅进行struct值拷贝，
故m[stu.Name]=&stu实际上是将所有值指向同一个指针
最终该指针的值为遍历的最后一个struct的值拷贝
*/
type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//for遍历时，变量stu指针不变
	//最终将所有值都指向同一个指针（最后一个）
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}

func TestForRange(t *testing.T) {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

/**
key=li,value=&{wang 22}
key=wang,value=&{wang 22}
key=zhou,value=&{wang 22}
*/
