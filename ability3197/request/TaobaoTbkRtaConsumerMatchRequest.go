package request

import (
	"github.com/sreiox/topsdk/ability3197/domain"
	"github.com/sreiox/topsdk/util"
)

type TaobaoTbkRtaConsumerMatchRequest struct {
	/*
	   mm_xxx_xxx_xxx的第3段数字     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   活动列表     */
	OfferList *[]domain.TaobaoTbkRtaConsumerMatchOfferList `json:"offer_list" required:"true" `
	/*
	   消费者对应的会员ID（会员ID或设备信息同时填时，优先使用会员ID）     */
	SpecialId *string `json:"special_id,omitempty" required:"false" `
	/*
	   设备信息，加密后的值(IMEI,IDFA,OAID,MOBILE需要加密)，需用MD5加密，32位小写     */
	DeviceValue *string `json:"device_value,omitempty" required:"false" `
	/*
	   设备信息，入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)：IMEI, 或者IDFA, 或者OAID, 或者MOBILE, 或者ALIPAY_ID     */
	DeviceType *string `json:"device_type,omitempty" required:"false" `
}

func (s *TaobaoTbkRtaConsumerMatchRequest) SetAdzoneId(v int64) *TaobaoTbkRtaConsumerMatchRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkRtaConsumerMatchRequest) SetOfferList(v []domain.TaobaoTbkRtaConsumerMatchOfferList) *TaobaoTbkRtaConsumerMatchRequest {
	s.OfferList = &v
	return s
}
func (s *TaobaoTbkRtaConsumerMatchRequest) SetSpecialId(v string) *TaobaoTbkRtaConsumerMatchRequest {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkRtaConsumerMatchRequest) SetDeviceValue(v string) *TaobaoTbkRtaConsumerMatchRequest {
	s.DeviceValue = &v
	return s
}
func (s *TaobaoTbkRtaConsumerMatchRequest) SetDeviceType(v string) *TaobaoTbkRtaConsumerMatchRequest {
	s.DeviceType = &v
	return s
}

func (req *TaobaoTbkRtaConsumerMatchRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.OfferList != nil {
		paramMap["offer_list"] = util.ConvertStructList(*req.OfferList)
	}
	if req.SpecialId != nil {
		paramMap["special_id"] = *req.SpecialId
	}
	if req.DeviceValue != nil {
		paramMap["device_value"] = *req.DeviceValue
	}
	if req.DeviceType != nil {
		paramMap["device_type"] = *req.DeviceType
	}
	return paramMap
}

func (req *TaobaoTbkRtaConsumerMatchRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
    return fileMap
}