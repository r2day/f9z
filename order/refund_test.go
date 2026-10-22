package order

import "testing"

func TestPendingApply(t *testing.T) {
	a := &RefundApply{ApplyId: "a", Status: RefundApplying}
	b := &RefundApply{ApplyId: "b", Status: RefundProcessed}
	c := &RefundApply{ApplyId: "c", Status: RefundProcessed}
	d := &RefundApply{ApplyId: "d", Status: RefundDone}
	cases := []struct {
		name   string
		refund []*RefundApply
		want   *RefundApply
	}{
		{"无申请", nil, nil},
		{"只有待处理（未审批）不返回", []*RefundApply{a}, nil},
		{"只有已完成不返回", []*RefundApply{d}, nil},
		{"返回第一条已审批", []*RefundApply{a, b, c, d}, b},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Model{Refund: tc.refund}
			if got := m.PendingApply(); got != tc.want {
				t.Fatalf("got=%v want=%v", got, tc.want)
			}
		})
	}
}

func TestUpdateApply(t *testing.T) {
	a := &RefundApply{ApplyId: "a", Status: RefundApplying}
	b := &RefundApply{ApplyId: "b", Status: RefundProcessed}
	newB := &RefundApply{ApplyId: "b", Status: RefundDone}
	cases := []struct {
		name   string
		refund []*RefundApply
		target *RefundApply
		want   []*RefundApply
	}{
		{"空列表返回空（非 nil）", nil, newB, []*RefundApply{}},
		{"替换匹配项并保持顺序", []*RefundApply{a, b}, newB, []*RefundApply{a, newB}},
		{"无匹配时原样返回，不追加", []*RefundApply{a}, newB, []*RefundApply{a}},
		{"重复 ID 全部替换", []*RefundApply{b, a, b}, newB, []*RefundApply{newB, a, newB}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Model{Refund: tc.refund}
			got := m.UpdateApply(tc.target)
			if got == nil || len(got) != len(tc.want) {
				t.Fatalf("got=%v want=%v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("[%d] got=%+v want=%+v", i, got[i], tc.want[i])
				}
			}
		})
	}
}
