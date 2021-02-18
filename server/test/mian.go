package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// if 语句执行赋值操作
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// switch语句的fallthrough语法
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
// goto语句简化for循环的终结
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
//测试break
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
func httpDemo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "astaxie!") //这个写入到w的是输出到客户端的
}

//一般函数，匿名函数，方法，及指针
func funcDemo1(a ,b *int) int{
	*a++
	*b++
	return *a+*b
}

type InterfaceDemo interface {
	doPost()
	doGet()
}
type myRequest struct {
	id int
	locationName string
}

func (mr myRequest) doPost(){
	println("this is the post tag")
	mr.id = 1
}
func (mr myRequest)doGet()  {
	println("this is the get tag")
	mr.locationName = "you test"
}
func PostDemo(){
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		_,fileHeader,err:=c.Request.FormFile("upload")
		if err!=nil{
			c.String(http.StatusBadRequest,"请求失败")
			return
		}

		fileName := fileHeader.Filename
		fmt.Println("文件名",fileName)
		if err:=c.SaveUploadedFile(fileHeader,fileName);err!=nil{
			c.String(http.StatusBadRequest,"保存失败 Error:%s",err.Error())
			return
		}
		c.String(http.StatusOK,"上传成功")
		
	})
	router.Run()
}
func GetDemo(){
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		

	})
}
//1.VUE的界面,组件的形式，路由跳转
//2.server后端的API的编写
//3.后端API请求前端的界面然后获取主界面，在主界面里面实现面包屑的路由跳转，展示最近本的信息
//需求分析：
//前端请求上传Excel，后端获取之后
//func tcpDemo() {
//	listener,err:=net.Listen("tcp","127.0.0.1:12345")
//	if err!=nil{
//		fmt.Println("失败，err:",err)
//		return
//	}
//
//	conn,err:=listener.Accept()
//	if err!=nil{
//		fmt.Println("失败，err:",err)
//		return
//	}
//	var tmp  [128]byte
//
//}
func main() {
	var r1 myRequest
	ifDemo1()
	ifDemo2()
	switchDemo5()
	breakDemo1()
	var a1 int  =1
	var b1 int  =3
	var  a,b *int
	a = &a1
	b = &b1
	funcDemo1(a,b)
	println("after%d,%d,count%d",a1,b1)
	println("-------------test interface-------------")
	r1.doGet()
	r1.doPost()
	println("------------test emptyInterface--------")
	var i1 interface{}
	i1 = "777"
	v,ok :=  i1.(string)
	if ok {
		println("is string",v)
	} else {
		println("is not string")
	}
	PostDemo()
	//println("-------------test  httpDemo-------------")
	//http.HandleFunc("/",httpDemo)
	//err:=http.ListenAndServe(":9090",nil)
	//if err != nil{
	//	log.Fatal("ListenAndServe:", err)
	//}
}