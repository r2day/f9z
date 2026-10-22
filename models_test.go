package f9z_test

import (
	"testing"

	"github.com/r2day/f9z"
	"github.com/r2day/f9z/address"
	"github.com/r2day/f9z/campaign"
	"github.com/r2day/f9z/categories"
	"github.com/r2day/f9z/coupon"
	"github.com/r2day/f9z/finance"
	"github.com/r2day/f9z/level"
	"github.com/r2day/f9z/member"
	"github.com/r2day/f9z/menu"
	"github.com/r2day/f9z/order"
	"github.com/r2day/f9z/pay"
	"github.com/r2day/f9z/product"
	"github.com/r2day/f9z/property"
	"github.com/r2day/f9z/record"
	"github.com/r2day/f9z/scm/category"
	"github.com/r2day/f9z/scm/inbound"
	"github.com/r2day/f9z/scm/inventory"
	"github.com/r2day/f9z/scm/material"
	"github.com/r2day/f9z/scm/outbound"
	"github.com/r2day/f9z/scm/process"
	"github.com/r2day/f9z/scm/purchase"
	"github.com/r2day/f9z/scm/recipe"
	"github.com/r2day/f9z/scm/stocktake"
	"github.com/r2day/f9z/scm/stocktxn"
	"github.com/r2day/f9z/scm/supplier"
	"github.com/r2day/f9z/shift"
	"github.com/r2day/f9z/store"
	"github.com/r2day/f9z/ticket"
	"github.com/r2day/f9z/topup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type named interface {
	ResourceName() string
	CollectionName() string
}

// 集合名是各服务共用的契约（同一张表会被多个服务读写），改名会让线上数据"消失"，这里逐个固定。
func TestResourceAndCollectionNames(t *testing.T) {
	cases := []struct {
		name               string
		m                  named
		resource, collName string
	}{
		{"address", &address.Model{}, "address", "member_address_config"},
		{"campaign", &campaign.Model{}, "campaign", "active_campaign_config"},
		{"categories", &categories.Model{}, "categories", "product_categories_config"},
		{"coupon", &coupon.Model{}, "coupon", "active_coupon_config"},
		{"finance", &finance.Model{}, "flow", "finance_flow_flow"},
		{"level", &level.Model{}, "level", "member_level_config"},
		{"member", &member.Model{}, "detail", "member_detail_data"},
		{"menu", &menu.Model{}, "menu", "product_menu_config"},
		{"order", &order.Model{}, "detail", "order_detail_flow"},
		{"pay", &pay.Model{}, "detail", "pay_detail_flow"},
		{"product", &product.Model{}, "detail", "product_detail_config"},
		{"property", &property.Model{}, "property", "product_property_config"},
		{"record", &record.Model{}, "record", "active_record_flow"},
		{"scm/category", &category.Model{}, "categories", "scm_categories_config"},
		{"scm/inbound", &inbound.Model{}, "inbound", "scm_inbound_flow"},
		{"scm/inventory", &inventory.Model{}, "inventory", "scm_inventory_stock"},
		{"scm/material", &material.Model{}, "material", "scm_material_config"},
		{"scm/outbound", &outbound.Model{}, "outbound", "scm_outbound_flow"},
		{"scm/process", &process.Model{}, "process", "scm_process_flow"},
		{"scm/purchase", &purchase.Model{}, "purchase", "scm_purchase_flow"},
		{"scm/recipe", &recipe.Model{}, "recipe", "scm_recipe_config"},
		{"scm/stocktake", &stocktake.Model{}, "stocktake", "scm_stocktake_flow"},
		{"scm/stocktxn", &stocktxn.Model{}, "stock_txn", "scm_stock_txn_flow"},
		{"scm/supplier", &supplier.Model{}, "supplier", "scm_supplier_config"},
		{"shift", &shift.Model{}, "shift", "finance_shift_log"},
		{"store", &store.Model{}, "store", "mini_store_manage"},
		{"ticket", &ticket.Model{}, "ticket", "active_ticket_flow"},
		{"topup", &topup.Model{}, "topup", "store_topup_config"},
	}
	seen := map[string]string{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.ResourceName(); got != tc.resource {
				t.Fatalf("ResourceName=%q want %q", got, tc.resource)
			}
			if got := tc.m.CollectionName(); got != tc.collName {
				t.Fatalf("CollectionName=%q want %q", got, tc.collName)
			}
		})
		if prev, dup := seen[tc.collName]; dup {
			t.Fatalf("集合名重复: %s 与 %s", tc.name, prev)
		}
		seen[tc.collName] = tc.name
	}
}

func TestConvertToObjectID(t *testing.T) {
	a, b := primitive.NewObjectID(), primitive.NewObjectID()
	cases := []struct {
		name    string
		in      []string
		want    []primitive.ObjectID
		wantErr bool
	}{
		{"nil 输入返回 nil", nil, nil, false},
		{"空切片返回 nil", []string{}, nil, false},
		{"单个", []string{a.Hex()}, []primitive.ObjectID{a}, false},
		{"保持顺序", []string{b.Hex(), a.Hex()}, []primitive.ObjectID{b, a}, false},
		{"任一非法即报错", []string{a.Hex(), "bad"}, nil, true},
		{"空字符串非法", []string{""}, nil, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := f9z.ConvertToObjectID(tc.in)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v", err)
			}
			if len(got) != len(tc.want) || (tc.want == nil) != (got == nil) {
				t.Fatalf("got=%v want=%v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("[%d] got=%v want=%v", i, got[i], tc.want[i])
				}
			}
		})
	}
}

// 写入数据库的枚举值，调整顺序会导致历史数据含义错位。
func TestEnumValues(t *testing.T) {
	cases := []struct {
		name      string
		got, want int
	}{
		{"order.HelpYourSelf", int(order.HelpYourSelf), 0},
		{"order.Takeout", int(order.Takeout), 1},
		{"order.DineIn", int(order.DineIn), 2},
		{"order.NoNeed", int(order.NoNeed), 4},
		{"order.PosClient", int(order.PosClient), 4},
		{"order.WebClient", int(order.WebClient), 5},
		{"order.CustomerApply", int(order.CustomerApply), 2},
		{"order.RefundReasonRefundOrder", int(order.RefundReasonRefundOrder), 2},
		{"order.RefundProcessed", int(order.RefundProcessed), 1},
		{"order.RefundDone", int(order.RefundDone), 2},
		{"member.PhoneUser", int(member.PhoneUser), 4},
		{"member.PosTerminal", int(member.PosTerminal), 5},
		{"member.Other", int(member.Other), 2},
		{"store.Franchise", int(store.Franchise), 1},
		{"store.JointVenture", int(store.JointVenture), 6},
		{"store.Large", int(store.Large), 2},
	}
	for _, tc := range cases {
		if tc.got != tc.want {
			t.Errorf("%s=%d want %d", tc.name, tc.got, tc.want)
		}
	}
}
