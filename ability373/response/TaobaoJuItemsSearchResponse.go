package response

import (
	"github.com/sreiox/topsdk/ability373/domain"
)

type TaobaoJuItemsSearchResponse struct {

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
	Result domain.TaobaoJuItemsSearchPaginationResult `json:"result,omitempty" `
}
