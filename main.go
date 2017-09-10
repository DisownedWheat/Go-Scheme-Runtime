package RT

type Env struct {
	Variables map[string]SchemeInterface
	Functions map[string]SchemeFunc
	Parent    *Env
	type_     string
}

func (e *Env) Init(p *Env) Env {
	e.Variables = map[string]SchemeInterface{}
	e.Functions = map[string]SchemeFunc{}
	e.Parent = p
	e.type_ = "Env"
	return *e
}

func (e *Env) GetType() string {
	return e.type_
}

func Call(e *Env, originalEnv *Env, funcName string, args []SchemeInterface) SchemeInterface {
	funcError := "NO FUNCTION BY The Name " + funcName
	if i, ok := e.Variables[funcName]; ok {

		f, ok := i.(SchemeFunc)
		if !ok {
			if e.Parent == nil {
				panic("This is not a function")
			}
			return Call(e.Parent, originalEnv, funcName, args)
		}
		return f.Value(originalEnv, args)
	}

	if e.Parent == nil {
		panic(funcError)
	}
	return Call(e.Parent, originalEnv, funcName, args)
}

func Get(e *Env, originalEnv *Env, varName string) SchemeInterface {
	varError := "Undefined variable: " + varName
	if i, ok := e.Variables[varName]; ok {
		v, ok := i.(SchemeInterface)
		if !ok {
			if e.Parent == nil {
				panic(varError)
			}
			return Get(e.Parent, e, varName)
		}
		return v
	}

	if e.Parent == nil {
		panic(varError)
	}
	return Get(e.Parent, e, varName)
}

type SchemeString struct {
	Value string
}

func (s SchemeString) GetType() string       { return "String" }
func (s SchemeString) GetValue() interface{} { return s.Value }

type SchemeNumber struct {
	Value int
}

func (s SchemeNumber) GetType() string       { return "Number" }
func (s SchemeNumber) GetValue() interface{} { return s.Value }

type SchemeFunc struct {
	Value func(e *Env, args []SchemeInterface) SchemeInterface
	Env   *Env
}

func (sf SchemeFunc) GetType() string {
	return "Function"
}
func (s SchemeFunc) GetValue() interface{} { return nil }

type SchemeSymbol struct {
	Value string
}

func (s SchemeSymbol) GetType() string       { return "Symbol" }
func (s SchemeSymbol) GetValue() interface{} { return s.Value }

type SchemeInterface interface {
	GetType() string
	GetValue() interface{} // This is for debugging purposes
}

func main() {}

func MakeRootEnv() *Env {
	var root Env
	root.Init(nil)

	root.Variables["def"] = SchemeFunc{
		Value: func(e *Env, args []SchemeInterface) SchemeInterface {
			name := args[0].(SchemeSymbol)
			value := args[1].(SchemeInterface)
			e.Variables[name.Value] = value
			return value
		},
	}

	root.Variables["+"] = SchemeFunc{
		Value: func(e *Env, args []SchemeInterface) SchemeInterface {
			total := 0
			for _, arg := range args {
				switch v := arg.(type) {
				case SchemeNumber:
					total += v.Value
				case SchemeSymbol:
					x := Get(e, e, v.Value)
					val, ok := x.(SchemeNumber)
					if !ok {
						panic("Cannot add this type")
					}
					total += val.Value
				default:
					panic("Cannot add this type")
				}
			}
			return SchemeNumber{Value: total}
		},
	}

	root.Variables["print"] = SchemeFunc{
		Value: func(e *Env, args []SchemeInterface) SchemeInterface {
			for _, arg := range args {
				switch v := arg.(type) {
				case SchemeNumber:
					print(v.Value, "")
				case SchemeString:
					print(v.Value, "")
				case SchemeFunc:
					print("Function", "")
				default:
					print(v.GetValue(), " ")
				}
			}
			return SchemeNumber{Value: 0}
		},
	}

	return &root
}
