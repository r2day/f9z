package store

import "testing"

// mongo.Connect 只解析 URI、不建立连接，所以这里不会访问任何数据库。
func TestNew(t *testing.T) {
	prev := MongoDatabase
	t.Cleanup(func() { MongoDatabase = prev })
	cases := []struct{ name, db string }{{"默认", "f9z"}, {"自定义", "f9z_other"}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			New("mongodb://127.0.0.1:1", tc.db)
			if MongoDatabase == nil || MongoDatabase.Name() != tc.db {
				t.Fatalf("MongoDatabase=%v", MongoDatabase)
			}
		})
	}
	defer func() {
		if recover() == nil {
			t.Fatal("非法 URI 应 panic")
		}
	}()
	New("not-a-uri", "x")
}

func TestEnabledChannelKeysOrder(t *testing.T) {
	cases := []struct {
		name string
		in   PayChannelFlags
		want []string
	}{
		{"全关", PayChannelFlags{}, nil},
		{"固定顺序 cash,wechat,lakala,balance", PayChannelFlags{Balance: true, Lakala: true, WeChat: true, Cash: true}, []string{"cash", "wechat", "lakala", "balance"}},
		{"只有余额", PayChannelFlags{Balance: true}, []string{"balance"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := enabledChannelKeys(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("got %v", got)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("got %v", got)
				}
			}
		})
	}
}
