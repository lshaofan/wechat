package auth

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	checkEncryptedDataURL = "https://api.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=%s"

	getPhoneNumber = "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s"
)
