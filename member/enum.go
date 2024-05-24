package member

// UserFrom 用户来源
type UserFrom int

const (
	// WxUser 微信小程序用户
	WxUser UserFrom = iota
	// DouYinUser 抖音小程序
	DouYinUser
	// AlipayUser 支付宝小程序
	AlipayUser
	// MeiTuanUser 美团小程序
	MeiTuanUser
	// PhoneUser 手机号注册用户
	PhoneUser
)

// UserGender 用户性别
type UserGender int

const (
	// Male 男
	Male UserGender = iota
	// Female 女
	Female
	// Other 其他
	Other
)
