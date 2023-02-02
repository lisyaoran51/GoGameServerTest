package main

type Client struct {
	IP           uint32 /**< \brief 客戶端當前 IP */
	LoginName    string //帶pid的username
	UserName     string //不帶pid的帳號名稱
	Token        string
	AgCode       string
	ProductID    string
	SessionID    string
	AgType       int    //0真錢 1試玩
	CurrentTable string //目前桌台，拿來判定是否在桌台上
	LastTable    string //殘留桌台，如果還有注單沒結算，離桌時這個值不會被清除
}
