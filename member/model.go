package member

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
	collectionNameSuffix = "_data"
	// 这个需要用户根据具体业务完成设定
	modelName = "detail"
)

// Model 订单信息
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 用户来源
	From UserFrom `json:"source" bson:"source"`
	// 客户信息
	Identity IdentityInfo `json:"identity" bson:"identity"`
	// 登陆审计信息
	Login LoginInfo `json:"login" bson:"login"`
	// 个人资产
	Assets AssetsInfo `json:"assets" bson:"assets"`
	// 当前邀请人情况
	Invitee InviterInfo `json:"invitee" bson:"invitee"`
	// 权限管理
	Permission PermissionInfo `json:"permission" bson:"permission"`
}

// IdentityInfo 标识信息
type IdentityInfo struct {
	// 备注信息 （后台管理员操作）
	AccountID string `json:"account_id" bson:"account_id"`
	// 昵称
	Nickname string `json:"nickname" bson:"nickname"`
	// 用户名
	UserName string `json:"user_name" bson:"user_name"`
	// 生日
	Birthday string `json:"birthday" bson:"birthday"`
	// 加入时间(注册时间）
	JoinAt string `json:"join_at" bson:"join_at"`
	// 性别
	Gender UserGender `json:"gender" bson:"gender"`
	// 微信登陆信息
	WxLoginID string `json:"wx_login_id" bson:"wx_login_id"`
	// 手机号
	Phone string `json:"phone" bson:"phone"`
	// 注册ip
	RegisterIP string `json:"register_ip" bson:"register_ip"`
	// 注册地区
	RegisterArea string `json:"register_area" bson:"register_area"`
	// Avatar 头像
	Avatar string `json:"avatar" bson:"avatar"`
	// Level 等级
	Level string `json:"level" bson:"level"`
	// Badge 标签
	Badge string `json:"badge" bson:"badge"`
}

// LoginInfo 登陆信息
type LoginInfo struct {
	// 登陆时长
	Times uint64 `json:"times" bson:"times"`
	// 登陆次数
	Counter uint64 `json:"counter" bson:"counter"`
	// 上次登陆时间
	LastLoginAt int64 `json:"last_login_at" bson:"last_login_at"`
	// 当前登陆时间
	LoginAt int64 `json:"login_at" bson:"login_at"`
	// 当前登陆IP
	LoginIp string `json:"login_ip" bson:"login_ip"`
	// 上次登陆IP
	LastLoginIp string `json:"last_login_ip" bson:"last_login_ip"`
}

// PermissionInfo 特权信息
// 例如：设置会员是否属于内部员工等
type PermissionInfo struct {
	// 权限等级
	Access uint64 `json:"access" bson:"access"`
	// 登陆次数1,2,4,8, 不存储（在创建的时候用于参数接收与渲染展示
	AccessBits []uint64 `json:"access_bits" bson:"-"`
	// 上次登陆时间
	IsInternalStaff bool `json:"is_internal_staff" bson:"is_internal_staff"`
}

// AssetsInfo 资产
type AssetsInfo struct {
	// 点数
	PointNum int `json:"point_num" bson:"point_num"`
	// 优惠券
	CouponNum int `json:"coupon_num" bson:"coupon_num"`
	// 余额
	Balance float64 `json:"balance" bson:"balance"`
	// 礼物
	GiftBalance int `json:"gift_balance" bson:"gift_balance"`
	// 积分
	Integrals int `json:"integrals" bson:"integrals"`
}

// InviterInfo 邀请人
type InviterInfo struct {
	// 当前阶段，当前已经完成度 （完成三次就可以得到一次奖励
	CurrentValue int `json:"current_value" bson:"current_value"`
	// 当前阶段，还需要完成值
	NeedValue int `json:"need_value" bson:"need_value"`
	// 邀请人数要求完成次数/每轮
	Turn int `json:"turn" bson:"turn"`
	// 总共完成
	Total int `json:"total" bson:"total"`
	// 被邀请人id列表
	Invitee []string `json:"invitee" bson:"invitee"`
	// 邀请人id
	Inviter string `json:"inviter" bson:"inviter"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
