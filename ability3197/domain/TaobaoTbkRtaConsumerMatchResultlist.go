package domain


type TaobaoTbkRtaConsumerMatchResultlist struct {
    /*
        活动id     */
    OfferId  *string `json:"offer_id,omitempty" `

    /*
        活动状态：1-符合活动要求，2-淘客无活动权限，3-用户不匹配活动，4-系统异常，5-活动不存在     */
    Status  *string `json:"status,omitempty" `

    /*
        推广链接     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        微信小程序路径     */
    WxMiniprogramPath  *string `json:"wx_miniprogram_path,omitempty" `

    /*
        微信小程序码     */
    WxQrcodeUrl  *string `json:"wx_qrcode_url,omitempty" `

}

func (s *TaobaoTbkRtaConsumerMatchResultlist) SetOfferId(v string) *TaobaoTbkRtaConsumerMatchResultlist {
    s.OfferId = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchResultlist) SetStatus(v string) *TaobaoTbkRtaConsumerMatchResultlist {
    s.Status = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchResultlist) SetClickUrl(v string) *TaobaoTbkRtaConsumerMatchResultlist {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchResultlist) SetWxMiniprogramPath(v string) *TaobaoTbkRtaConsumerMatchResultlist {
    s.WxMiniprogramPath = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchResultlist) SetWxQrcodeUrl(v string) *TaobaoTbkRtaConsumerMatchResultlist {
    s.WxQrcodeUrl = &v
    return s
}
