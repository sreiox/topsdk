package domain


type TaobaoTbkRtaConsumerMatchData struct {
    /*
        返回结果列表     */
    ResultList  *[]TaobaoTbkRtaConsumerMatchResultlist `json:"result_list,omitempty" `

}

func (s *TaobaoTbkRtaConsumerMatchData) SetResultList(v []TaobaoTbkRtaConsumerMatchResultlist) *TaobaoTbkRtaConsumerMatchData {
    s.ResultList = &v
    return s
}
