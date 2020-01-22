package singleton

type singleton struct{}

var s *singleton

func GetInstance() *singleton {
	if s == nil {
		s = &singleton{}
	}
	return s
}
