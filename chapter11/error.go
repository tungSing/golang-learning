/**
错误
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
*/
package error

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Demo() error {
	//1. 一般处理方式-如果该方法或函数已经是最顶层，通过记录错误日志方式供排查错误
	NormalProcessingLog()
	//2. 一般处理方式-如果该方法或函数是提供API，可以把错误返回，让上层处理
	err := NormalProcessingReturn()
	if err != nil {
		//错误处理
		//return err
	}
	//继续正常的业务逻辑

	//2的变种写法，代码更简洁
	if ss := NormalProcessingReturn(); ss != nil {
		//错误处理
		//return ss
	}

	//定义错误
	f, err := DefineError(-1)
	if err != nil {
		//错误处理
		log.Println(err)
	}
	fmt.Println(f)

	//自定义错误结构
	if bizErr := fileUpload(""); bizErr != nil {
		log.Println(bizErr)
	}

	return nil
}

func NormalProcessingLog() {
	content, err := ioutil.ReadFile("C:\\Users\\tungS\\Downloads\\1.txt")
	if err != nil {
		//记录日志
		log.Println(err)
		return
	} else {
		fmt.Println(string(content))
	}
}

func NormalProcessingReturn() error {
	f, err := os.Open("C:\\Users\\tungS\\Downloads\\2.txt")
	if err != nil {
		return err
	}
	fmt.Println(f.Name())
	return nil
}

func DefineError(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	//具体函数实现
	return f, nil
}

//模拟业务错误
type BizError struct {
	//业务编码
	code string
	//错误信息
	message string
}

func (e *BizError) Error() string {
	return e.code + ":" + e.message
}

//模拟文件上传
func fileUpload(fileName string) error {
	if fileName == "" {
		return &BizError{"22222", "文件名称不能为空"}
	}
	//真正的业务处理
	return nil
}
