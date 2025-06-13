package Convert

type Convert struct {
	value any
}

var ConvertValue = new(Convert)

//// New creates and returns a new Var with given `value`.
//// The optional parameter `safe` specifies whether Var is used in concurrent-safety,
//// which is false in default.
//func New(value any) *Var {
//	return &Var{
//		value: value,
//	}
//}
