package response

import (
	"github.com/sreiox/topsdk/ability3197/domain"
)

type TaobaoTbkRtaConsumerMatchResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Data domain.TaobaoTbkRtaConsumerMatchData `json:"data,omitempty" `
}
