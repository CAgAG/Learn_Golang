package SetArray

type Set struct {
	buffer []interface{}
	num int // 数量
	hash map[interface{}] bool
}

func NewSet() *Set {
	return &Set{make([]interface{}, 0),
		0,
		make(map[interface{}] bool)}
}

func (this *Set)IsExist(value interface{}) bool {
	return this.hash[value]
}

func (this *Set)Add(value interface{}) bool {
	if this.IsExist(value) {
		return false
	} else {
		this.buffer = append(this.buffer, value)
		this.hash[value] = true
		this.num++
		return true
	}
}

func (this *Set)Strings() []interface{} {
	return this.buffer
}
