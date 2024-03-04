package main

import "sync"

type singleton struct{}

var ins *singleton //
var mu sync.Mutex  // 1. 需要加鎖

func GetIns() *singleton {
	/* 存在並發安全問題, 需要加鎖*/
	/* 第一層非nil直接回傳,不用每次進來都lock*/
	/* 確定nil後第二層lock解決並發問題*/
	if ins == nil {
		mu.Lock()         // 1.
		defer mu.Unlock() // 1.
		if ins == nil {   // 1.
			ins = &singleton{}
		} // 1.
	}

	return ins
}
