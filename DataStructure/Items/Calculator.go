package Items

import (
	"DataStructure/StackArray"
	"errors"
)

// 中缀转后缀
func InToPost(exp string) string {
	Stack := StackArray.NewStack()
	result := ""
	for _, s := range exp{
		s := string(s)
		if s >= "0" && s <= "9" {
			result += s
		} else {
			result += " "
			if s == "*" || s == "/" || s == "(" {
				Stack.Push(s)
			}
			if s == "+" || s == "-" {
				for Stack.Top() != "("{
					if Stack.IsEmpty() {
						break
					}
					result += Stack.Pop().(string)
					result += " "
				}
				Stack.Push(s)
			}
			if s == ")" {
				for Stack.Top() != "(" {
					result += Stack.Pop().(string)
					result += " "
				}
				Stack.Pop()
			}
		}
	}

	for !Stack.IsEmpty(){
		result += " "
		result += Stack.Pop().(string)
		result += " "
	}

	//strings.Replace(result, "  ", " ", -1)
	return result
}

func Compute(x float64, y float64, op rune) (float64, error){
	switch op {
	case '+': return x + y, nil
	case '-': return x - y, nil
	case '*': return x * y, nil
	case '/': {
		if y == 0 {
			return 0, errors.New("除数不能为零")
		}
		return x / y, nil
	}
	default:
		return 0, errors.New("err")
	}
}

func GetResult(exp string) float64 {
	Stack := StackArray.NewStack()
	var (
		x, y, z, tp float64
		err error
	)

	for i := 0; i < len(exp); i++ {
		if exp[i] == ' ' {
			continue
		}
		if exp[i] >= '0' && exp[i] <= '9' {
			tp = float64(exp[i] - '0')
			for exp[i+1] >= '0' && exp[i+1] <= '9' {
				i += 1
				tp = tp*10 + float64(exp[i] - '0')
			}
			Stack.Push(tp)
		}else {
			x = Stack.Pop().(float64)
			y = Stack.Pop().(float64)
			z, err = Compute(y, x, rune(exp[i]))
			if err != nil {
				break
			}
			Stack.Push(z)
		}
	}
	return z
}
