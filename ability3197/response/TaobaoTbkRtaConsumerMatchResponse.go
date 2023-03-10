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
	   θΏεη»ζ
	*/
	Data domain.TaobaoTbkRtaConsumerMatchData `json:"data,omitempty" `
}
