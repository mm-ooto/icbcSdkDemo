package lib

type ResponseBizContentM struct {
	ReturnCode int64  `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`
	MsgId      string `json:"msg_id"`
}

type NotifyResponseReturn struct {
	ResponseBizContent ResponseBizContentM `json:"response_biz_content"`
	SignType           string              `json:"sign_type"`
	Sign               string              `json:"sign,omitempty"`
}
