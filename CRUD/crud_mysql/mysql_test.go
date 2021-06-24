package crud_mysql

import "testing"

func TestInitDB(t *testing.T) {
	ret := InitDB()
	t.Log(ret)

}
func TestInsert(t *testing.T) {
	name := "yang"
	passwd := "123456"
	Insert(name, passwd)
}
func TestDeleteByName(t *testing.T) {
	name := "zhang"
	DeleteByName(name)
}
func TestUpdate(t *testing.T) {
	Update()
}
func TestSelect(t *testing.T) {
	//SelectAll()
	name := "yang"
	SelectByName(name)
}
