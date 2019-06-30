package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jlaffaye/ftp"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
	"time"
)

func Mail(adder, dody string) {
	//创建mail对象
	m := gomail.NewMessage()

	//添加设置Mail 的收件人 发件人，主题，和内容
	m.SetHeader("From", "1048549775@qq.com")
	m.SetHeader("To", "1048549775@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", adder)
	m.SetBody("text/html", fmt.Sprintf("%02d%02d%02d<b>ERROR:</b><i>%s</i>!", time.Now().Hour(), int(time.Now().Minute()), time.Now().Second(), dody))
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "1048549775@qq.com", "srwpzaglmmrpbdhh")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func FtpStor(name string) error {
	str := strings.Split(name, "\\")                                                 //处理字符串
	c, err := ftp.Dial("2shoucang.f3322.net:21", ftp.DialWithTimeout(5*time.Second)) //创建连接
	if err != nil {
		return err
	}
	err = c.Login("BOLONG", "131420") //登陆
	if err != nil {
		return err
	}
	c.li
	time.Sleep(time.Second * 5)
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	defer file.Close()
	//传输文件，指定传输的文件路径和文件名（针对于接收文件的服务器的根目录之下的），以及需要传输的文件IO
	err = c.Stor(fmt.Sprintf("%d%02d%02d/%s", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), str[len(str)-1]), file)
	if err != nil {
		return err
	}
	return nil
}

//判断文件是否存在，如不存在创建
func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) { //os.IsNotExist  判断 ERR 这个错误是否是文件不存在
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

func main() {
	//自定义监控文件绝对路径
	Path := fmt.Sprintf("D:/FTPROOT/%d%02d%02d", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	err := PathExists(Path) //判断监控文件是否存在，如不存在创建
	if err != nil {
		Mail("西安分公司FTP回传程序故障", fmt.Sprintf("%v", err))
		return
	}
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		Mail("西安分公司FTP回传程序故障", fmt.Sprintf("%v", err))
		return
	}
	defer watch.Close()
	//添加要监控的对象，文件或文件夹
	err = watch.Add(Path)
	if err != nil {
		Mail("西安分公司FTP回传程序故障", fmt.Sprintf("%v", err))
		return
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					//if ev.Op&fsnotify.Create == fsnotify.Create {
					//	log.Println("创建文件 : ", ev.Name)
					//}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						//处理写入完成的监控文件，使用FTP进行回传
						err := FtpStor(ev.Name)
						if err != nil {
							Mail("西安分公司FTP回传程序故障", fmt.Sprintf("%v：%v", err, ev.Name))
						}
					}
					//if ev.Op&fsnotify.Remove == fsnotify.Remove {
					//	log.Println("删除文件 : ", ev.Name)
					//}
					//if ev.Op&fsnotify.Rename == fsnotify.Rename {
					//	log.Println("重命名文件 : ", ev.Name)
					//}
					//if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					//	log.Println("修改权限 : ", ev.Name)
					//}
				}
			case err := <-watch.Errors:
				{
					Mail("西安分公司FTP回传程序故障", fmt.Sprintf("%v", err))
					return
				}

			}
		}
	}()
	//循环
	select {}

}
