package main

import (
	"fmt"
	"io/ioutil"
)

// ioutil 方便的IO操作函数集

// 列出某目录下的所有文件（包括子目录
func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info :=  range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}

			fmt.Println(info.Name(), "\\", curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
