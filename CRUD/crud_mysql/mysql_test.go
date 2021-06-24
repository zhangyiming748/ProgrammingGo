package crud_mysql

import "testing"

func TestInitDB(t *testing.T) {
	ret := InitDB()
	t.Log(ret)

}
func TestInsert(t *testing.T) {
	Insert()
}
func TestDeleteByID(t *testing.T) {
	DeleteByID()
}
func TestUpdate(t *testing.T) {
	Update()
}
func TestSelect(t *testing.T) {
	Select()
}
