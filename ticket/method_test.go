package ticket

import (
	"testing"
	"time"
)

func TestIsUsable(t *testing.T) {
	cases := []struct {
		name   string
		status string
		expire int64
		now    int64
		want   bool
	}{
		{"未使用且不过期", StatusUnused, 0, 100, true},
		{"未使用且未到期", StatusUnused, 200, 100, true},
		{"到期时刻仍可用", StatusUnused, 100, 100, true},
		{"已过期", StatusUnused, 50, 100, false},
		{"已使用", StatusUsed, 0, 100, false},
		{"已作废", StatusVoid, 0, 100, false},
		{"状态为 expired", StatusExpired, 0, 100, false},
		{"now=0 使用当前时间（已过期）", StatusUnused, time.Now().Unix() - 60, 0, false},
		{"now=0 使用当前时间（未过期）", StatusUnused, time.Now().Unix() + 3600, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Model{Status: tc.status, ExpireAt: tc.expire}
			if got := m.IsUsable(tc.now); got != tc.want {
				t.Fatalf("got %v", got)
			}
		})
	}
}
