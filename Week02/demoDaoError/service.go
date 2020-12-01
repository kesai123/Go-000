package main

// service 层不处理error
func GetValue(id int) (string, error){
	return DaoGet(id)
}
