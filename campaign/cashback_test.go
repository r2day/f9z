package campaign

import "testing"

func TestCalcCashback(t *testing.T) {
	if got := (Rule{CashbackAmount: 3}).CalcCashback(100); got != 3 {
		t.Fatalf("fixed got %v", got)
	}
	if got := (Rule{CashbackPercent: 0.1, CashbackMax: 5}).CalcCashback(100); got != 5 {
		t.Fatalf("percent cap got %v", got)
	}
	if got := (Rule{}).CalcCashback(100); got != 5 {
		t.Fatalf("default 5%% got %v", got)
	}
	if got := (Rule{}).CalcCashback(0); got != 0 {
		t.Fatalf("zero order got %v", got)
	}
	if (Rule{ConsumeThreshold: 20}).ConsumeMinAmount() != 20 {
		t.Fatal("consume threshold")
	}
	if (Rule{Threshold: 10}).ConsumeMinAmount() != 10 {
		t.Fatal("threshold fallback")
	}
}
