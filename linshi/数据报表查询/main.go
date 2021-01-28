package main

import (
	"encoding/json"
	"fmt"
	"github.com/FirewoodBloody/Golang/linshi/数据报表查询/module"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

var (
	v = "0.3.1"
)

type V struct {
	V      string
	Remark string
}

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetShowMainForm(false)
	TForm := new(module.Windows)
	TForm.Init()
	TForm.Onclick()
	TForm.LoginTForm.Init()
	TForm.LoginOnclick()
	TForm.LoginTForm.WinLogin.Show()

	bady := url.Values{}
	bady.Add("Version", "Version")

	resp, err := http.PostForm(module.Url2, bady)
	if err != nil {
		vcl.ShowMessage(fmt.Sprintf("%v", err))
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	V := V{}
	json.Unmarshal(data, &V)

	//升级检测
	if v != V.V {
		Taskdlg := vcl.NewTaskDialog(TForm.LoginTForm.WinLogin)
		//defer w.Taskdlg.Free()
		Taskdlg.SetCaption("升级提示！")
		Taskdlg.SetTitle("检测到有新版本，是否升级？")
		Taskdlg.SetText(V.Remark)

		//w.Taskdlg.SetExpandButtonCaption("展开")
		//w.Taskdlg.SetExpandedText("展开的文本")
		Taskdlg.SetCommonButtons(0) //rtl.Include(0, 0))
		btn := vcl.AsTaskDialogButtonItem(Taskdlg.Buttons().Add())
		btn.SetCaption("确定")
		btn.SetModalResult(types.MrYes)

		btn = vcl.AsTaskDialogButtonItem(Taskdlg.Buttons().Add())
		btn.SetCaption("取消")
		btn.SetModalResult(types.MrNo)

		Taskdlg.SetFooterText("新版本号：" + V.V)

		if Taskdlg.Execute() {
			if Taskdlg.ModalResult() == types.MrYes {
				cmd := exec.Command("./update.exe")
				cmd.Start()
				Taskdlg.Free()
				TForm.Win.Free()
				TForm.LoginTForm.WinLogin.Free()
				os.Exit(0)
			} else if Taskdlg.ModalResult() == types.MrNo {
				Taskdlg.Free()
			}
		}

	}

	vcl.Application.Run()
}
