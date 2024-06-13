package GoClass

// 类的特性

// 封装、继承、多态
type Animal struct {
	Name string
}

func (a Animal) Call() string {
	return "动物的叫声..."
}
func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}
func (a *Animal) GetName() string {
	return a.Name
}

type Pet struct {
	Name string
}

func (p *Pet) GetName() string {
	return p.Name
}

type Dog struct {
	animal *Animal
	pet    *Pet
}

func (d Dog) FavorFood() string {
	return "骨头"
}
func (d Dog) Call() string {
	return "汪汪汪"
}
