package store

// Qrcode 微信二维码
type Qrcode struct {
	// 微信二维码的场景id，用于绑定门店id和tableId
	SceneID string `json:"scene_id" bson:"scene_id"`
	//// 门店id
	//StoreID string `json:"store_id" `
	// 桌子id 餐桌号
	TableID string `json:"table_id"  bson:"table_id"`
	// 二维码图片地址
	Url string `json:"url" bson:"url"`
}
