package main

import (
	"test-learning/learn1"
	"testing"
)

// 程序需要在一个名为 xxx_test.go 的文件中编写
// 测试函数的命名必须以单词 Test 开始
// 测试函数只接受一个参数 t *testing.T
func TestHello(t *testing.T) {
	// get := strings.ToLower(learn1.SayHelloWolrd())
	// want := "hello world"
	//  if get != want {
	// t.Errorf("get '%q' want '%q'", get, want)
	// }
	t.Run("saying hello to people ", func(t *testing.T) {
		got := learn1.GetString("hello")
		want := "hello"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
	t.Run("say something empty", func(t *testing.T) {
		got := learn1.GetString("")
		want := "hello"
		if got == want {
			t.Errorf("got '%q' want '%q' ", got, want)
		}
	})
}
