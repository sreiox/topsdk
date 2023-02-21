package ability3197

import (
	"log"
	"errors"
	"github.com/sreiox/topsdk"
	"github.com/sreiox/topsdk/ability3197/request"
	"github.com/sreiox/topsdk/ability3197/response"
	"github.com/sreiox/topsdk/util"
)

type Ability3197 struct {
	Client *topsdk.TopClient
}

func NewAbility3197(client *topsdk.TopClient) *Ability3197 {
	return &Ability3197{client}
}

/*
   淘宝客-推广者-定向活动目标发布
*/
func (ability *Ability3197) TaobaoTbkRtaConsumerMatch(req *request.TaobaoTbkRtaConsumerMatchRequest) (*response.TaobaoTbkRtaConsumerMatchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability3197 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.rta.consumer.match", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkRtaConsumerMatchResponse{}
	if err != nil {
		log.Println("taobaoTbkRtaConsumerMatch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
