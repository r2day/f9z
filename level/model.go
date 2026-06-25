package level

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "member_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "level"
)

// Model 会员层级配置模型
type Model struct {
	// 模型继承（包含 meta、tenant_id 等通用字段）
	model.Model `json:"_" bson:"_"`

	// 主键
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// ==================== 基本信息 ====================
	Name string `json:"name" bson:"name"` // 层级名称（如 黄金会员）
	Code string `json:"code" bson:"code"` // 层级编码（唯一，如 gold、vip1）

	Level     int  `json:"level" bson:"level"`           // 层级序号（数值越大等级越高）
	SortOrder int  `json:"sort_order" bson:"sort_order"` // 显示排序
	IsActive  bool `json:"is_active" bson:"is_active"`   // 是否启用
	IsDefault bool `json:"is_default" bson:"is_default"` // 是否为默认层级（新用户默认归属）

	// ==================== 升级条件 ====================
	// 推荐使用这个字段替代或补充原来的 Min/MaxPoints
	Condition Condition `json:"condition" bson:"condition"`

	// ==================== 核心权益 ====================
	DiscountRate     float64 `json:"discount_rate" bson:"discount_rate"`         // 折扣率（0.95 = 95折）
	CashbackRate     float64 `json:"cashback_rate" bson:"cashback_rate"`         // 返现比例
	PointsMultiplier float64 `json:"points_multiplier" bson:"points_multiplier"` // 积分倍率（默认 1.0）

	// ==================== 附加权益 ====================
	FreeShipping    bool `json:"free_shipping" bson:"free_shipping"`       // 免运费
	PriorityService bool `json:"priority_service" bson:"priority_service"` // 优先服务
	BirthdayGift    bool `json:"birthday_gift" bson:"birthday_gift"`       // 生日礼遇

	// ==================== 其他 ====================
	Description string `json:"description" bson:"description"` // 层级描述与权益说明（支持富文本）
	UserCount   int    `json:"user_count" bson:"user_count"`   // 涉及到的用户，这个一般只会在返回的时候统计算出来
}

// Condition 会员升级条件（支持多种类型，未来扩展友好）
type Condition struct {
	Type      ConditionType `json:"type" bson:"type"`           // 条件类型：points, amount, orders, custom 等
	MinValue  int64         `json:"min_value" bson:"min_value"` // 最小值
	MaxValue  int64         `json:"max_value" bson:"max_value"` // 最大值（0 表示无上限）
	Unlimited bool          `json:"unlimited" bson:"unlimited"` // 是否无上限

	// 扩展字段（未来用）
	// Unit      string `json:"unit,omitempty" bson:"unit,omitempty"` // 如 "元", "单", "天"
	// Expression string `json:"expression,omitempty" bson:"expression,omitempty"` // 复杂表达式（可选）
}

// ConditionType 条件类型枚举
type ConditionType string

const (
	ConditionPoints ConditionType = "points" // 积分
	ConditionAmount ConditionType = "amount" // 消费金额
	ConditionOrders ConditionType = "orders" // 订单数量
	ConditionCustom ConditionType = "custom" // 自定义（预留）
)

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
