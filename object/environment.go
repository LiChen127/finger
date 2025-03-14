package object

/*
	Environment结构体定义了finger语言的执行环境。
	它包含一个存储变量和值的map，以及一个指向外部环境的指针。
*/
type Environment struct {
	store map[string]Object
	outer *Environment // 外层环境 用于闭包
}

/*
	创建一个新的环境
*/
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

/*
	获取变量
*/
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

/*
	设置变量
*/
func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}

/*
	创建一个新的封闭环境
	实现闭包
*/
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
