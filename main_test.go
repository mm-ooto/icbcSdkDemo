package icbc_api_sdk

import (
	"github.com/mm-ooto/icbcSdkDemo/example"
	"testing"
)

func TestT1(t *testing.T) {
	example.CardbusinessAggregatepayB2cOnlineConsumepurchase()
	example.CardbusinessZfbh5UiH5consumption()
	example.APICardbusinessQrcodeQrgenerate()
}

func TestN(t *testing.T) {
	data := "api=%2Fapi%2Fcardbusiness%2Faggregatepay%2Fb2c%2Fonline%2Fconsumepurchase%2FV1&app_id=10000011100000199068&biz_content=%7B%22access_type%22%3A%224%22%2C%22attach%22%3A%22%22%2C%22bank_disc_amt%22%3A%220%22%2C%22card_flag%22%3A%22%22%2C%22card_kind%22%3A%22%22%2C%22card_no%22%3A%22%22%2C%22coupon_amt%22%3A%220%22%2C%22cust_id%22%3A%22%22%2C%22decr_flag%22%3A%22%22%2C%22ecoupon_amt%22%3A%220%22%2C%22mer_disc_amt%22%3A%220%22%2C%22mer_id%22%3A%22101159931205%22%2C%22msg_id%22%3A%22050206851050818093451747020%22%2C%22open_id%22%3A%22%22%2C%22order_id%22%3A%22100159931205000512102190006259%22%2C%22out_trade_no%22%3A%22JM4689038966101079010%22%2C%22pay_time%22%3A%2220210228093451%22%2C%22pay_type%22%3A%2210%22%2C%22payment_amt%22%3A%221%22%2C%22point_amt%22%3A%220%22%2C%22return_code%22%3A%221%22%2C%22return_msg%22%3A%22%E4%BA%A4%E6%98%93%E5%A4%B1%E8%B4%A5%22%2C%22third_party_coupon_amt%22%3A%220%22%2C%22third_party_discount_amt%22%3A%220%22%2C%22third_trade_no%22%3A%22%22%2C%22total_amt%22%3A%221%22%2C%22total_disc_amt%22%3A%220%22%7D&charset=UTF-8&format=json&from=icbc-api&sign=g21oTFuB%2FOU04Xz%2FfsawlAS93QJf7ralHOf2QwKNvgZQiWm%2BhS5HHZlLdGfTV3WRi1MYyZrOzH3JOEtnyx5%2Fhka9dGnz7%2Bv%2BImtYPBoUE70IB7DEJvr02Lkr3paTat2PzLhFzbVgsRWIGi2L%2FB%2FriwbhvS7RB%2F6MDHA1UvlzRD0%3D&sign_type=RSA&timestamp=2021-02-19+09%3A45%3A27"
	example.Notify(data)
}
