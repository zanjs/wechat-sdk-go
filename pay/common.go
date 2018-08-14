package pay

import (
	"encoding/xml"
	"errors"

	"github.com/zanjs/wechat-sdk-go/common"
	"github.com/zanjs/wechat-sdk-go/utils"
)

type (
	// UnifiedOrderResp 统一下单返回
	UnifiedOrderResp struct {
		ReturnCode string `xml:"return_code"`
		ReturnMsg  string `xml:"return_msg"`
		AppID      string `xml:"appid"`
		MchID      string `xml:"mch_id"`
		DeviceInfo string `xml:"device_info"`
		NonceStr   string `xml:"nonce_str"`
		Sign       string `xml:"sign"`
		ResultCode string `xml:"result_code"`
		ErrCode    string `xml:"err_code"`
		ErrCodeDes string `xml:"err_code_des"`
		TradeType  string `xml:"trade_type"`
		PrepayID   string `xml:"prepay_id"`
	}

	// UnifiedOrder 统一下单公共参数
	UnifiedOrder struct {
		XMLName        xml.Name `xml:"xml"`                                                //xml标签
		AppID          string   `xml:"appid" json:"appid"`                                 //Appid
		MchID          string   `xml:"mch_id" json:"mch_id"`                               //微信支付分配的商户号，必须
		DeviceInfo     string   `xml:"device_info" json:"device_info"`                     //微信支付填"WEB"，必须
		NonceStr       string   `xml:"nonce_str" json:"nonce_str"`                         //随机字符串，必须
		Sign           string   `xml:"sign" json:"sign"`                                   //签名，必须
		SignType       string   `xml:"sign_type" json:"sign_type"`                         //"HMAC-SHA256"或者"MD5"，非必须，默认MD5
		Body           string   `xml:"body" json:"body"`                                   //商品简单描述，必须
		Detail         string   `xml:"detail,omitempty" json:"detail,omitempty"`           //商品详细列表，使用json格式
		Attach         string   `xml:"attach" json:"attach"`                               //附加数据，如"贵阳分店"，非必须
		OutTradeNo     string   `xml:"out_trade_no" json:"out_trade_no"`                   //订单号，必须
		FeeType        string   `xml:"fee_type,omitempty" json:"fee_type,omitempty"`       //默认人民币：CNY，非必须
		TotalFee       int      `xml:"total_fee" json:"total_fee"`                         //订单金额，单位分，必须
		SpBillCreateIP string   `xml:"spbill_create_ip" json:"spbill_create_ip"`           //支付提交客户端IP，如“123.123.123.123”，必须
		TimeStart      string   `xml:"time_start,omitempty" json:"time_start,omitempty"`   //订单生成时间，格式为yyyyMMddHHmmss，如20170324094700，非必须
		TimeExpire     string   `xml:"time_expire,omitempty" json:"time_expire,omitempty"` //订单结束时间，格式同上，非必须
		GoodsTag       string   `xml:"goods_tag,omitempty" json:"goods_tag,omitempty"`     //商品标记，代金券或立减优惠功能的参数，非必须
		NotifyURL      string   `xml:"notify_url" json:"notify_url"`                       //接收微信支付异步通知回调地址，不能携带参数，必须
		TradeType      string   `xml:"trade_type" json:"trade_type"`                       //交易类型，小程序写"JSAPI"，APP 写 APP
		LimitPay       string   `xml:"limit_pay,omitempty" json:"limit_pay,omitempty"`     //限制某种支付方式，非必须
	}

	// AppUnifiedOrder APP统一下单
	AppUnifiedOrder struct {
		UnifiedOrder
		SceneInfo string `json:"scene_info"` //场景信息
	}

	// WxaUnifiedOrder 微信小程序统一下单
	WxaUnifiedOrder struct {
		UnifiedOrder
		OpenID string `xml:"openid" json:"openid"` //微信用户唯一标识，必须
	}
)

// NewUnifiedOrder 统一下单
func NewUnifiedOrder(unifiedOrder interface{}) (unifiedOrderResp UnifiedOrderResp, err error) {

	data, err := xml.Marshal(unifiedOrder)
	if err != nil {
		return unifiedOrderResp, err
	}

	body, err := utils.NewRequest("POST", common.UnifiedOrderURL, data)
	if err != nil {
		return unifiedOrderResp, err
	}

	err = xml.Unmarshal(body, &unifiedOrderResp)
	if err != nil {
		return unifiedOrderResp, err
	}

	if unifiedOrderResp.ReturnCode != "SUCCESS" {
		return unifiedOrderResp, errors.New(unifiedOrderResp.ReturnMsg)
	}

	return unifiedOrderResp, err
}
