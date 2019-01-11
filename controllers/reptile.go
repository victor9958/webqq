package controllers

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
	"strconv"
	"regexp"
	"time"
	"github.com/astaxie/beego"
)

type ReptileController struct {
	BaseController
}

type Spider struct {
	url string
	header map[string]string
}

//爬虫
func(c *ReptileController)Get(){
	t1 := time.Now()
	parse()
	elapsed :=time.Since(t1)
	c.Ctx.WriteString("爬虫结束,总共耗时" + string(elapsed))
}


func(keyword Spider)get_html_header() string{
	client :=&http.Client{}
	req,err :=http.NewRequest("GET",keyword.url,nil)
	if err !=nil {

	}

	for k,v :=range keyword.header{
		req.Header.Add(k,v)
	}

	resp,err :=client.Do(req)

	if err!=nil {

	}

	defer resp.Body.Close()


	body ,err := ioutil.ReadAll(resp.Body)

	if err!=nil {

	}
	return string(body)


}
func parse(){
	header :=map[string]string{
		"Host":"movie.douban.com",
		"Connectin":"keep-alive",
		"Cache-Control":"max-age=0",
		"Upgrade-Insecure-Requests":"1",
		"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer": "https://movie.douban.com/top250",
	}

	//创建execl文件
	f,err :=os.Create("D:/ceshi.xlsx")
	if err!=nil{
		panic(err)
	}
	//写入标题
	f.WriteString("电影名称"+"\t"+"评分"+"\t"+"评价人数"+"\t"+"\r\n")

	//循环每页解析并把结果写入execl
	for i:=0;i<10 ;i++  {
		fmt.Println("正在抓取第"+strconv.Itoa(i)+"页.........")

		url:="https://movie.douban.com/top250?start="+strconv.Itoa(i*25)+"&filter="
		spider :=&Spider{url,header}

		html:= spider.get_html_header()

		//评价人数
		pattern2 :=`<span>(.*?)评价</span>`

		rp2 :=regexp.MustCompile(pattern2)
		find_txt2 := rp2.FindAllStringSubmatch(html ,-1)


		//评分
		pattern3 :=`property="v:average">(.*?)</span>`
		rp3 :=regexp.MustCompile(pattern3)
		find_txt3 := rp3.FindAllStringSubmatch(html,-1)

		//电影名称
		pattern4 := `img width="(.*?)" alt="(.*?)" src=`
		rp4 := regexp.MustCompile(pattern4)
		find_txt4 := rp4.FindAllStringSubmatch(html,-1)

		beego.Info(len(find_txt4))

		//写入UTF-8 BOM
		f.WriteString("\xEF\xBB\xBF")
		//打印全部数据和写入execl文件
		for i:=0;i<len(find_txt2) ;i++  {
			fmt.Sprintf("%s %s %s\n",find_txt4[i][1],find_txt3[i][1],find_txt2[i][1])
			f.WriteString(find_txt4[i][2]+"\t"+find_txt3[i][1]+"\t"+find_txt2[i][1]+"\t"+"\r\n")
		}
	}

}

