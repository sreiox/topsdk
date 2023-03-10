package response

import (
	"github.com/sreiox/topsdk/ability2474/domain"
)

type TaobaoTbkDgVegasSendReportResponse struct {

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
	Result domain.TaobaoTbkDgVegasSendReportResult `json:"result,omitempty" `
}
