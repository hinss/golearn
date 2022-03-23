package singleton

/**
	饿汉模式
 */
type singleton struct {

}

var s *singleton = &singleton{}


func GetSingleton() *singleton {
	return s
}
