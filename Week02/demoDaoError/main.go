package main

import (
	"database/sql"
	"fmt"
	perror "github.com/pkg/errors"
)

func main(){
	v, err := GetValue(10)
	// 先判断错误，再读取返回值
	if err != nil {
		processError(err)
	} else {
		processValue(v)
	}
}

func processError(err error){
	// 显示完整错误
	fmt.Println("Error:", err)
	// 显示初始错误
	fmt.Println("Root error:", perror.Cause(err))
	// 显示错误调用栈
	fmt.Printf("Error stack: %+v", perror.WithStack(err))
	// 按照错误类型进行专门处理，如错误码转换
	// 此处可添加多种错误类型判断处理
	if perror.Is(err, sql.ErrNoRows){
		fmt.Println("processing sql.ErrNoRows")
	}
}

func processValue(v string) {
	fmt.Println("processing value ", v)
}