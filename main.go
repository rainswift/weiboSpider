package main

import (
	"fmt"
	"strconv"
	"time"
	"weiboSpider/common"
	"weiboSpider/structs"
)

func forFn(arr structs.Response,class int)  {
	//fmt.Printf("a[%d]:----------\n", arr.MaxId )
	var maxid = 0

	for i, v := range arr.Data {


		fmt.Printf("a[%d]-[%d]: %s:-%s\n", i,class, v.User.Name,v.Textraw)
		//fmt.Println(arr.MaxId)
		//if maxid != 0 && class == 2{
		//	forFn(*arr2,2)
		//}

		if arr.MaxId != 0 && class != -1{
			maxid = arr.MaxId

		}
		if v.TotalNumber != 0 {
			//fmt.Println("https://weibo.com/ajax/statuses/buildComments?flow=1&is_reload=1&id="+ v.Rootid +"&is_show_bulletin=2&is_mix=1&fetch_level=1&max_id="+ strconv.Itoa(maxid) +"&count=20&uid=6166129198")
			//fmt.Printf("zi--id=[%s]-maxid==[%d]\n", v.Rootid ,maxid )
			arr2,_ := common.WeiboH("https://weibo.com/ajax/statuses/buildComments?flow=1&is_reload=1&id="+ v.Rootid +"&is_show_bulletin=2&is_mix=1&fetch_level=1&max_id="+ strconv.Itoa(maxid) +"&count=20&uid=6166129198")
			//
			//forFn(*arr2,2)
			zi_1(*arr2)
		}

	}


}
func zi_1(arr structs.Response)  {
	t := time.NewTimer(1 * time.Second)
	<-t.C
	var maxId =  arr.MaxId
	for  i, v := range arr.Data {
		fmt.Printf("zi-[%d]: %s:-%s\n", i, v.User.Name,v.Textraw)
	}
	if maxId != 0{
		fmt.Println( arr)
		var http = "https://weibo.com/ajax/statuses/buildComments?flow=1&is_reload=1&id="+arr.Data[0].Rootid+"&is_show_bulletin=2&is_mix=1&fetch_level=1&max_id="+ strconv.Itoa(maxId) +"&count=20&uid=6166129198"
		fmt.Println( http)
		arr2,_ := common.WeiboH(http)
		fmt.Println( arr2)
		for  i, v := range arr2.Data {
			fmt.Printf("zi-[%d]: %s:-%s\n", i, v.User.Name,v.Textraw)
		}
		//forFn(*arr2,2)
		zi_1(*arr2)
	}


}

func main() {
	arr,_ := common.WeiboH("https://weibo.com/ajax/statuses/buildComments?is_reload=1&id=4855030950135054&is_show_bulletin=2&is_mix=0&count=10&uid=3277405011")
	forFn(*arr,-1)

	//for i, v := range arr.Data {
	//	arr2,_ := common.WeiboH("https://weibo.com/ajax/statuses/buildComments?is_reload=1&id=4851030888033780&is_show_bulletin=2&is_mix=0&count=10&uid=6166129198")
	//	for i, v := range arr2.Data {
	//
	//	}
	//	fmt.Printf("a[%d]: %s\n", i, v)
	//}


}
