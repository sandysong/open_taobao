// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/changkong/open_taobao"
)

/* 淘宝客商品转换。 */
type TbkItemsConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:click_url */
func (r *TbkItemsConvertRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 推广者的淘宝会员昵称.注：指的是淘宝的会员登录名 */
func (r *TbkItemsConvertRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 淘宝客商品数字id串.最大输入40个.格式如:"value1,value2,value3" 用" , "号分隔商品数字id */
func (r *TbkItemsConvertRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TbkItemsConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TbkItemsConvertRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 点击串跳转类型，1：单品，2：单品中间页（无线暂无） */
func (r *TbkItemsConvertRequest) SetReferType(value string) {
	r.SetValue("refer_type", value)
}

/* 商品track_iid串（带有追踪效果的商品id),最大输入40个,与num_iids必填其一 */
func (r *TbkItemsConvertRequest) SetTrackIids(value string) {
	r.SetValue("track_iids", value)
}

func (r *TbkItemsConvertRequest) GetResponse(accessToken string) (*TbkItemsConvertResponse, []byte, error) {
	var resp TbkItemsConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.items.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemsConvertResponse struct {
	TbkItems     []*TbkItem `json:"tbk_items"`
	TotalResults int        `json:"total_results"`
}

type TbkItemsConvertResponseResult struct {
	Response *TbkItemsConvertResponse `json:"tbk_items_convert_response"`
}

/* 查询淘宝客推广商品详细信息。 */
type TbkItemsDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:num_iid,seller_id,nick,title,price,volume,pic_url,item_url,shop_url
;字段之间用","分隔. */
func (r *TbkItemsDetailGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 淘宝客商品数字id串.最大输入40个.格式如:"value1,value2,value3" 用" , "号分隔商品数字id */
func (r *TbkItemsDetailGetRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 商品track_iid串（带有追踪效果的商品id),最大输入40个,与num_iids必填其一 */
func (r *TbkItemsDetailGetRequest) SetTrackIids(value string) {
	r.SetValue("track_iids", value)
}

func (r *TbkItemsDetailGetRequest) GetResponse(accessToken string) (*TbkItemsDetailGetResponse, []byte, error) {
	var resp TbkItemsDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.items.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemsDetailGetResponse struct {
	TbkItems     []*TbkItem `json:"tbk_items"`
	TotalResults int        `json:"total_results"`
}

type TbkItemsDetailGetResponseResult struct {
	Response *TbkItemsDetailGetResponse `json:"tbk_items_detail_get_response"`
}

/* 查询淘宝客推广商品。 */
type TbkItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 商品所在地 */
func (r *TbkItemsGetRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 是否自动发货 */
func (r *TbkItemsGetRequest) SetAutoSend(value string) {
	r.SetValue("auto_send", value)
}

/* 是否支持抵价券，设置为true表示该商品支持抵价券，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetCashCoupon(value string) {
	r.SetValue("cash_coupon", value)
}

/* 是否支持货到付款，设置为true表示该商品是支持货到付款，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetCashOndelivery(value string) {
	r.SetValue("cash_ondelivery", value)
}

/* 标准商品后台类目id。该ID可以通过taobao.itemcats.get接口获取到。<br /> 支持最大值为：2147483647 */
func (r *TbkItemsGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 30天累计推广量（与返回数据中的commission_num字段对应）上限. */
func (r *TbkItemsGetRequest) SetEndCommissionNum(value string) {
	r.SetValue("end_commissionNum", value)
}

/* 佣金比率上限，如：2345表示23.45%。注：start_commissionRate和end_commissionRate一起设置才有效。 */
func (r *TbkItemsGetRequest) SetEndCommissionRate(value string) {
	r.SetValue("end_commissionRate", value)
}

/* 可选值和start_credit一样.start_credit的值一定要小于或等于end_credit的值。注：end_credit与start_credit一起使用才生效 */
func (r *TbkItemsGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 最高价格 */
func (r *TbkItemsGetRequest) SetEndPrice(value string) {
	r.SetValue("end_price", value)
}

/* 商品总成交量（与返回字段volume对应）上限。 */
func (r *TbkItemsGetRequest) SetEndTotalnum(value string) {
	r.SetValue("end_totalnum", value)
}

/* 需返回的字段列表.可选值:num_iid,seller_id,nick,title,volume,pic_url,item_url,shop_url;字段之间用","分隔 */
func (r *TbkItemsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否查询消保卖家 */
func (r *TbkItemsGetRequest) SetGuarantee(value string) {
	r.SetValue("guarantee", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TbkItemsGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 商品标题中包含的关键字. 注意:查询时keyword,cid至少选择其中一个参数 */
func (r *TbkItemsGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 是否商城的商品，设置为true表示该商品是属于淘宝商城的商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetMallItem(value string) {
	r.SetValue("mall_item", value)
}

/* 是否30天维修，设置为true表示该商品是支持30天维修，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetOnemonthRepair(value string) {
	r.SetValue("onemonth_repair", value)
}

/* 是否海外商品，设置为true表示该商品是属于海外商品，默认为false */
func (r *TbkItemsGetRequest) SetOverseasItem(value string) {
	r.SetValue("overseas_item", value)
}

/* 结果页数.1~10<br /> 支持最大值为：10 */
func (r *TbkItemsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页返回结果数.最大每页40<br /> 支持最大值为：400 */
func (r *TbkItemsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 是否如实描述(即:先行赔付)商品，设置为true表示该商品是如实描述的商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetRealDescribe(value string) {
	r.SetValue("real_describe", value)
}

/* 是否支持7天退换，设置为true表示该商品支持7天退换，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetSevendaysReturn(value string) {
	r.SetValue("sevendays_return", value)
}

/* 默认排序:default

price_desc(价格从高到低)
price_asc(价格从低到高)
credit_desc(信用等级从高到低)
commissionRate_desc(佣金比率从高到低)
commissionRate_asc(佣金比率从低到高)
commissionNum_desc(成交量成高到低)
commissionNum_asc(成交量从低到高)
commissionVolume_desc(总支出佣金从高到低)
commissionVolume_asc(总支出佣金从低到高)
delistTime_desc(商品下架时间从高到低)
delistTime_asc(商品下架时间从低到高) */
func (r *TbkItemsGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 30天累计推广量（与返回数据中的commission_num字段对应）下限.注：该字段要与end_commissionNum一起使用才生效 */
func (r *TbkItemsGetRequest) SetStartCommissionNum(value string) {
	r.SetValue("start_commissionNum", value)
}

/* 佣金比率下限，如：1234表示12.34% */
func (r *TbkItemsGetRequest) SetStartCommissionRate(value string) {
	r.SetValue("start_commissionRate", value)
}

/* 卖家信用:

1heart(一心)
2heart (两心)
3heart(三心)
4heart(四心)
5heart(五心)
1diamond(一钻)
2diamond(两钻)
3diamond(三钻)
4diamond(四钻)
5diamond(五钻)
1crown(一冠)
2crown(两冠)
3crown(三冠)
4crown(四冠)
5crown(五冠)
1goldencrown(一黄冠)
2goldencrown(二黄冠)
3goldencrown(三黄冠)
4goldencrown(四黄冠)
5goldencrown(五黄冠) */
func (r *TbkItemsGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 起始价格.传入价格参数时,需注意起始价格和最高价格必须一起传入,并且 start_price <= end_price */
func (r *TbkItemsGetRequest) SetStartPrice(value string) {
	r.SetValue("start_price", value)
}

/* 商品总成交量（与返回字段volume对应）下限。 */
func (r *TbkItemsGetRequest) SetStartTotalnum(value string) {
	r.SetValue("start_totalnum", value)
}

/* 是否支持VIP卡，设置为true表示该商品支持VIP卡，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemsGetRequest) SetVipCard(value string) {
	r.SetValue("vip_card", value)
}

func (r *TbkItemsGetRequest) GetResponse(accessToken string) (*TbkItemsGetResponse, []byte, error) {
	var resp TbkItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemsGetResponse struct {
	TbkItems     []*TbkItem `json:"tbk_items"`
	TotalResults int        `json:"total_results"`
}

type TbkItemsGetResponseResult struct {
	Response *TbkItemsGetResponse `json:"tbk_items_get_response"`
}

/* 无线淘宝客商品转换。 */
type TbkMobileItemsConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:click_url. */
func (r *TbkMobileItemsConvertRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 淘宝客商品数字id串.最大输入40个.格式如:"value1,value2,value3" 用" , "号分隔商品数字id */
func (r *TbkMobileItemsConvertRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TbkMobileItemsConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

func (r *TbkMobileItemsConvertRequest) GetResponse(accessToken string) (*TbkMobileItemsConvertResponse, []byte, error) {
	var resp TbkMobileItemsConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.mobile.items.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkMobileItemsConvertResponse struct {
	TbkItems     []*TbkItem `json:"tbk_items"`
	TotalResults int        `json:"total_results"`
}

type TbkMobileItemsConvertResponseResult struct {
	Response *TbkMobileItemsConvertResponse `json:"tbk_mobile_items_convert_response"`
}

/* 无线淘宝客店铺转换。 */
type TbkMobileShopsConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:click_url. */
func (r *TbkMobileShopsConvertRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TbkMobileShopsConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 卖家昵称串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔。
注意：sids和seller_nicks两个参数任意必须输入一个，如果同时输入，则以seller_nicks为准 */
func (r *TbkMobileShopsConvertRequest) SetSellerNicks(value string) {
	r.SetValue("seller_nicks", value)
}

/* 店铺id串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔店铺id.
注意：sids和seller_nicks两个参数任意必须输入一个，如果同时输入，则以seller_nicks为准 */
func (r *TbkMobileShopsConvertRequest) SetSids(value string) {
	r.SetValue("sids", value)
}

func (r *TbkMobileShopsConvertRequest) GetResponse(accessToken string) (*TbkMobileShopsConvertResponse, []byte, error) {
	var resp TbkMobileShopsConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.mobile.shops.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkMobileShopsConvertResponse struct {
	TbkShops []*TbkShop `json:"tbk_shops"`
}

type TbkMobileShopsConvertResponseResult struct {
	Response *TbkMobileShopsConvertResponse `json:"tbk_mobile_shops_convert_response"`
}

/* 淘客店铺转换。 */
type TbkShopsConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:click_url. */
func (r *TbkShopsConvertRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 推广者的淘宝会员昵称.注：这里指的是淘宝的登录会员名 */
func (r *TbkShopsConvertRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TbkShopsConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TbkShopsConvertRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 卖家昵称串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔。
注意：sids和seller_nicks两个参数任意必须输入一个，如果同时输入，则以seller_nicks为准 */
func (r *TbkShopsConvertRequest) SetSellerNicks(value string) {
	r.SetValue("seller_nicks", value)
}

/* 店铺id串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔店铺id.
注意：sids和seller_nicks两个参数任意必须输入一个，如果同时输入，则以seller_nicks为准 */
func (r *TbkShopsConvertRequest) SetSids(value string) {
	r.SetValue("sids", value)
}

func (r *TbkShopsConvertRequest) GetResponse(accessToken string) (*TbkShopsConvertResponse, []byte, error) {
	var resp TbkShopsConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shops.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopsConvertResponse struct {
	TbkShops []*TbkShop `json:"tbk_shops"`
}

type TbkShopsConvertResponseResult struct {
	Response *TbkShopsConvertResponse `json:"tbk_shops_convert_response"`
}

/* 淘客店铺详情。 */
type TbkShopsDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:user_id,seller_nick,shop_title,pic_url,shop_url;字段之间用","分隔. */
func (r *TbkShopsDetailGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会生成无线店铺URL.如果不传值,则默认是false. */
func (r *TbkShopsDetailGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 卖家昵称串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔。
注意：sids和seller_nicks两个参数任意必须输入一个，如果同时输入，则以seller_nicks为准 */
func (r *TbkShopsDetailGetRequest) SetSellerNicks(value string) {
	r.SetValue("seller_nicks", value)
}

/* 店铺id串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔店铺id.
注意：sids和seller_nicks两个参数任意必须输入一个，如果同时输入，则以seller_nicks为准 */
func (r *TbkShopsDetailGetRequest) SetSids(value string) {
	r.SetValue("sids", value)
}

func (r *TbkShopsDetailGetRequest) GetResponse(accessToken string) (*TbkShopsDetailGetResponse, []byte, error) {
	var resp TbkShopsDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shops.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopsDetailGetResponse struct {
	TbkShops []*TbkShop `json:"tbk_shops"`
}

type TbkShopsDetailGetResponseResult struct {
	Response *TbkShopsDetailGetResponse `json:"tbk_shops_detail_get_response"`
}

/* 提供对参加了淘客推广的店铺的搜索。 */
type TbkShopsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 店铺前台展示类目id，可以通过taobao.shopcats.list.get获取。 */
func (r *TbkShopsGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 店铺商品数查询结束值。需要跟start_auctioncount同时设置才生效，只设置该值不生效。 */
func (r *TbkShopsGetRequest) SetEndAuctioncount(value string) {
	r.SetValue("end_auctioncount", value)
}

/* 店铺佣金比例查询结束值 */
func (r *TbkShopsGetRequest) SetEndCommissionrate(value string) {
	r.SetValue("end_commissionrate", value)
}

/* 店铺掌柜信用等级查询结束
店铺的信用等级总共为20级 1-5:1heart-5heart;6-10:1diamond-5diamond;11-15:1crown-5crown;16-20:1goldencrown-5goldencrown */
func (r *TbkShopsGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 店铺累计推广数查询结束值 */
func (r *TbkShopsGetRequest) SetEndTotalaction(value string) {
	r.SetValue("end_totalaction", value)
}

/* 需要返回的字段，目前包括如下字段user_id,seller_nick,shop_title,pic_url,shop_url */
func (r *TbkShopsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TbkShopsGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 店铺主题关键字查询 */
func (r *TbkShopsGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 是否只显示商城店铺 */
func (r *TbkShopsGetRequest) SetOnlyMall(value string) {
	r.SetValue("only_mall", value)
}

/* 页码.结果页1~99 */
func (r *TbkShopsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.最大每页40 */
func (r *TbkShopsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 排序字段。目前支持的排序字段有：
commission_rate，auction_count，total_auction。必须输入这三个任意值，否则排序无效 */
func (r *TbkShopsGetRequest) SetSortField(value string) {
	r.SetValue("sort_field", value)
}

/* 排序类型.必须输入desc,asc任一值，否则无效
desc-降序,asc-升序 */
func (r *TbkShopsGetRequest) SetSortType(value string) {
	r.SetValue("sort_type", value)
}

/* 店铺宝贝数查询开始值。需要跟end_auctioncount同时设置才生效，只设置该值不生效。 */
func (r *TbkShopsGetRequest) SetStartAuctioncount(value string) {
	r.SetValue("start_auctioncount", value)
}

/* 店铺佣金比例查询开始值，注意佣金比例是x10000的整数.50表示0.5% */
func (r *TbkShopsGetRequest) SetStartCommissionrate(value string) {
	r.SetValue("start_commissionrate", value)
}

/* 店铺掌柜信用等级起始
店铺的信用等级总共为20级 1-5:1heart-5heart;6-10:1diamond-5diamond;11-15:1crown-5crown;16-20:1goldencrown-5goldencrown */
func (r *TbkShopsGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 店铺累计推广量开始值 */
func (r *TbkShopsGetRequest) SetStartTotalaction(value string) {
	r.SetValue("start_totalaction", value)
}

func (r *TbkShopsGetRequest) GetResponse(accessToken string) (*TbkShopsGetResponse, []byte, error) {
	var resp TbkShopsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shops.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopsGetResponse struct {
	TbkShops     []*TbkShop `json:"tbk_shops"`
	TotalResults int        `json:"total_results"`
}

type TbkShopsGetResponseResult struct {
	Response *TbkShopsGetResponse `json:"tbk_shops_get_response"`
}
