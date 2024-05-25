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
	// 性别
	Gender UserGender `json:"gender" bson:"gender"`
	// 微信登陆信息
	WxLoginID string `json:"wx_login_id" bson:"wx_login_id"`
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

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
