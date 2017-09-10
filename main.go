package goScheme

type Env struct {
	Variables map[string]SchemeInterface
	CallStack []func(args ...SchemeInterface) SchemeInterface
	Closures  map[string]*Env
	RootEnv   *Env
	Parent    *Env
	Name      string
}

func (e *Env) Call(funcName string) {
	if nextEnv, ok := e.Closures[funcName]; ok {
		e.Closures[nextEnv.Name] = nextEnv
	} else {
		if e.Parent == nil {
			panic("NO FUNCTION BY THAT NAME")
		}
		e.Parent.Call(funcName)
	}
}

type SchemeType struct {
	type_ string
}

func (s *SchemeType) GetType() string {
	return s.type_
}

type SchemeInterface interface {
	GetType() string
}

func main() {}

func MakeRootEnv() {
	root := Env{}

	rootEnv.closures["define"] = func(e *Env, name SchemeType, value SchemeInterface) {
		type_ = value.GetType()
		if type_ == "CallExpression" {
			e.closures[name.(string)] = value
		}
	}
}
