package test

import (
	"errors"
	"fmt"
	"github.com/dbinggo/gerr"
	"testing"
)

// 测试ERR返回是否正常
func TestErr1(t *testing.T) {
	// 这里模拟一个正常的函数调用
	test1()
}

func TestErr2(t *testing.T) {
	// 这里模拟是否可以进行error.Is 判断原来类型

	err := test2()
	if err != nil {
		//这里查看包装之后的error是否可以正常被error.Is比较

		ok := errors.Is(err, FlagError)

		fmt.Printf("%v", ok)
		if !ok {
			t.Fail()
		}
		// OK = true 这里说明我们包装后的error 可以使用error.Is进行解包比较
		// 原因：实现了unwarp方法

	}
}

func test1() gerr.Error {
	//我们不再返回 原声error 而是强制所有函数返回包装版本的error
	// 目的 请在第一时间获得error之后将error加上堆栈信息和友好日志
	err := returnErr()
	// 这里发现了有err 所以我们要进行包装和打印，不包装返回就会报错
	//return err
	// 这样直接报错，也是一个强制要求你要在源头处理error

	if err != nil {
		err = gerr.WrapSysErrf(err, "服务内部异常") // 这里的信息就是包装给前端的信息
		fmt.Println("--------------------------------")
		fmt.Println(err) // 这里的信息就是包装给前端的信息
		fmt.Println("--------------------------------")
		fmt.Println(err.Error()) // 这里的信息就是包装给前端的信息
		fmt.Println("--------------------------------")
		fmt.Printf("%+v", err) // 这里是我们自己的堆栈信息
		fmt.Println("--------------------------------")
		return err.(gerr.Error) // 这里是强制返回包装err进行断言
	}
	return nil
}

func test2() gerr.Error {
	// 这个测试主要测试解包是否正确
	err := returnFlagErr()
	// 进行包装
	err = gerr.WrapSysErrf(err, "服务内部异常")
	return err.(gerr.Error)
}

var FlagError = errors.New("FlagError")

func returnErr() error {
	return errors.New("这里返回了一个Err")
}

func returnFlagErr() error {
	return FlagError
}
