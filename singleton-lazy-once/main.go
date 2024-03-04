package main

import "sync"

type singleton struct{}

var ins *singleton
var once sync.Once //

func GetIns() *singleton {
	// 確保ins實例全局只會被創建一次
	once.Do(func() {
		ins = &singleton{}
	})

	return ins
}
