// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

const VersionNo = "20140607"

/* 淘宝客商品 */
type TbkItem struct {
	ClickUrl string  `json:"click_url"`
	ItemUrl  string  `json:"item_url"`
	Nick     string  `json:"nick"`
	NumIid   int     `json:"num_iid"`
	PicUrl   string  `json:"pic_url"`
	Price    float64 `json:"price"`
	SellerId int     `json:"seller_id"`
	ShopUrl  string  `json:"shop_url"`
	Title    string  `json:"title"`
	Volume   int     `json:"volume"`
}

/* 淘宝客店铺 */
type TbkShop struct {
	ClickUrl   string `json:"click_url"`
	PicUrl     string `json:"pic_url"`
	SellerNick string `json:"seller_nick"`
	ShopTitle  string `json:"shop_title"`
	ShopUrl    string `json:"shop_url"`
	UserId     int    `json:"user_id"`
}
