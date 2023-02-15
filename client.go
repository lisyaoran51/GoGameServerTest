package main

type Client struct {
	IP        uint32 /**< \brief 客戶端當前 IP */
	UserName  string //不帶pid的帳號名稱
	SessionID uint32
}
