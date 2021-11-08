package logfile

import (
	"bufio"
	"fmt"
	"github.com/Jay-zjy/common_use/tools/util"
	"os"
	"time"
)
func CheckFile(filePath string){
	is,err1 := util.PathExists(filePath)
	if !is || err1 != nil{
		util.CreateDir(filePath)
	}
}
func SaveErrLog(contentStr string,path string)  {
	if path == ""{
		path = "./logs/" +time.Now().Format("2006-01-02")+"_err_log"+ ".txt"
	}
	//os.O_WRONLY | os.O_CREATE
	file, err1 := os.OpenFile(path, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err1!=nil{
		fmt.Println("file open error:",err1)
		//return
	}
	defer file.Close()

	str := "请求时间: "+time.Now().Format("2006-01-02 15:04:05")+" "+" 内容："+contentStr+"\n"
	//使用缓存方式写入
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(str)
	//将缓存中数据刷新到文本中
	writer.Flush()
}