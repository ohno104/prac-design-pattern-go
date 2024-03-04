package main

type singleton struct{}

var ins *singleton = &singleton{} //package導入時初始化, 會導致加載時間比較長

func GetIns() *singleton {
	return ins
}
