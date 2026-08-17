package campaign

import "testing"

func TestCalcBenefit(t *testing.T) {
	full := Model{Type: TypeFullReduce, Rule: Rule{Threshold: 50, Reduce: 10}}
	if full.CalcBenefit(49) != 0 || full.CalcBenefit(50) != 10 {
		t.Fatalf("full reduce")
	}
	off := Model{Type: TypeDiscount, Rule: Rule{Discount: 0.8, MaxReduce: 12}}
	if off.CalcBenefit(100) != 12 {
		t.Fatalf("discount cap got %v", off.CalcBenefit(100))
	}
}
