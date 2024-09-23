package main

import "strconv"

type Human struct {
	Liver  ///inhetance olarak kullanılabilie ama tam karşılığı değil
	Name   string
	Age    int
	wiht   int
	height int
}

func (h Human) calculate(de string) string { //class metdoları olarak kullanıla bilir ama tam karşılığı değil
	cal := h.height * h.wiht
	return de + strconv.Itoa(cal)
}
