package GoClass

// 类定义
type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

// 初始化方法(全量定义)
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

// 初始化方法（部分定义）
func NewStudentPart(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}

// 成员方法
// 值方法
func (s Student) GetName() string {
	return s.name
}

// 指针方法
func (s *Student) SetName(name string) {
	s.name = name
}
