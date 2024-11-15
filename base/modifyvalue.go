package base

type person struct {
	name string
}

func modifyValue(ptr *int) {
	*ptr = 20
}

func modifyValue2(ptr int) {
	ptr = 20
}

func modifyStruct(ptr *person) {
	(*ptr).name = "b"
}

func modifyStruct2(ptr person) {
	ptr.name = "c"
}

func ModifyMap(ptr *map[string]string) {
	(*ptr)["a"] = "b"
}

func ModifyMap2(ptr map[string]string) {
	ptr["a"] = "c"
}

func Mv() {
	/*var num int = 10
	modifyValue(&num)
	println(num)
	num = 30
	modifyValue2(num)
	println(num)*/

	/*p := person{name: "a"}
	modifyStruct(&p)
	println(p.name)
	p.name = "d"
	modifyStruct2(p)
	println(p.name)*/

	m := make(map[string]string)
	m["a"] = "a"
	ModifyMap(&m)
	println(m["a"])
	m["a"] = "d"
	ModifyMap2(m)
	println(m["a"])
}
