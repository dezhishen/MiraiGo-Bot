package entity

// Rule 规则
type Rule struct {
	ID           string
	Name         string
	Type         string
	RespTemplate string
	Command      string
	Min          int
	Max          int
}
