package store

import "testing"

func TestResolvedChannels(t *testing.T) {
	cases := []struct {
		name string
		in   PayChannelFlags
		want PayChannelFlags
	}{
		{
			name: "all off is cash",
			want: PayChannelFlags{Cash: true, Default: "cash"},
		},
		{
			name: "cash only ignores dirty default",
			in:   PayChannelFlags{Cash: true, Default: "wechat"},
			want: PayChannelFlags{Cash: true, Default: "cash"},
		},
		{
			name: "lakala only",
			in:   PayChannelFlags{Lakala: true},
			want: PayChannelFlags{Lakala: true, Default: "lakala"},
		},
		{
			name: "two channels keep valid default",
			in:   PayChannelFlags{Cash: true, Lakala: true, Default: "lakala"},
			want: PayChannelFlags{Cash: true, Lakala: true, Default: "lakala"},
		},
		{
			name: "dirty default falls back cash then wechat then lakala",
			in:   PayChannelFlags{WeChat: true, Lakala: true, Default: "cash"},
			want: PayChannelFlags{WeChat: true, Lakala: true, Default: "wechat"},
		},
		{
			name: "empty default among cash and balance picks cash",
			in:   PayChannelFlags{Cash: true, Balance: true},
			want: PayChannelFlags{Cash: true, Balance: true, Default: "cash"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := (FinanceConfig{Channels: tc.in}).ResolvedChannels()
			if got != tc.want {
				t.Fatalf("got %+v want %+v", got, tc.want)
			}
		})
	}
}
