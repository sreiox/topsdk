package domain


type TaobaoTbkRtaConsumerMatchOfferList struct {
    /*
        活动id     */
    OfferId  *string `json:"offer_id,omitempty" `

    /*
        淘礼金领取URL，不支持使用短链接     */
    TljUrl  *string `json:"tlj_url,omitempty" `

    /*
        商品id     */
    ItemId  *string `json:"item_id,omitempty" `

}

func (s *TaobaoTbkRtaConsumerMatchOfferList) SetOfferId(v string) *TaobaoTbkRtaConsumerMatchOfferList {
    s.OfferId = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchOfferList) SetTljUrl(v string) *TaobaoTbkRtaConsumerMatchOfferList {
    s.TljUrl = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchOfferList) SetItemId(v string) *TaobaoTbkRtaConsumerMatchOfferList {
    s.ItemId = &v
    return s
}
