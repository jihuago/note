package common

import (
	"fmt"
	"regexp"
	"sync"
)

// 正则表达式
// 简单模式
//  ok, _ := regexp.Match(pat, []byte(searchIn))
// ok, _ := regexp.MatchString(pat, searchIn)
/*
		// 正则模式通过Compile返回一个regexp对象
		re, _ := regexp.Compile(pat)
		re.ReplaceAllString(SearchIn, replaceWith)
 */

func Test()  {
	searchIn := "this is a test 129" // 搜索的字符串
	pattern := "[0-9]+" // 正则

	ok, _ := regexp.Match(pattern, []byte(searchIn))
	fmt.Println(ok)

	info := new(Info)
	Update(info)

}

// sync.Mutex 互斥锁
type Info struct {
	mu sync.Mutex
	Str string
}

func Update(info *Info)  {
	info.mu.Lock()

	info.Str = "test"

	info.mu.Unlock()

	fmt.Println(info.Str)
}


