package src

import "time"

// 先頭文字を大文字にするとexported(=別packageからも参照可能)となる
// ※メンバやメソッドも同様
type User struct {
	Name string
	Birth time.Time
}

// 構造体の値を変更する場合、レシーバーに構造体のポインタを指定
func (u *User) AddAgeByRef() {
	u.Birth = u.Birth.AddDate(1,0,0)
}

// 値渡しの場合は値が変更されない
func (u User) AddAgeByVal() {
	u.Birth = u.Birth.AddDate(1,0,0)
}
