package ticket

import "time"

func (m *Model) IsUsable(now int64) bool {
	if now == 0 {
		now = time.Now().Unix()
	}
	if m.Status != StatusUnused {
		return false
	}
	if m.ExpireAt > 0 && now > m.ExpireAt {
		return false
	}
	return true
}
