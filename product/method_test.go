package product

import "testing"

// 没有属性时直接返回，不访问数据库（这里未初始化 Handler）。
func TestGetPropertyEmpty(t *testing.T) {
	got, err := (&Model{}).GetProperty()
	if got != nil || err != nil {
		t.Fatalf("got=%v err=%v", got, err)
	}
}
