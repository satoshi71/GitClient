package main

import (
    "fmt"
	"os/exec"
    "io/ioutil"
	"os"
	"time"
)

func main() {
	initial();
}

func main_manu(){
	fmt.Println();
	fmt.Println("/_/_/_/_ Main Manu /_/_/_/_/")
	fmt.Println()
	fmt.Println("1. ファイル追加")
	fmt.Println("2. 追加したファイルを取り消し")
	fmt.Println("3. コミット")
	fmt.Println("4. リモートリポジトリURL設定")
	fmt.Println("5: リモートリポジトリにpush")
	fmt.Println("99: exit")
	fmt.Println()
	fmt.Printf("番号を選んでください。: ")
	var ans int
	fmt.Scan(&ans)

	switch ans{
	case 1:
		add_menu(0)
	case 2:
		add_menu(1)
	case 3:
		commit_menu()
	case 4:
		setURL_menu()
	case 5:
		push_menu()
	case 99:
		goodby()
	default:
		fmt.Println("入力した番号が間違っています。やり直してください。")
		time.Sleep(1 * time.Second)
		main_manu()
	}
}

func push_menu(){
	remoteURL, _ := exec.Command("git", "remote", "-v").Output()

	fmt.Println();
	fmt.Println("/_/_/_/_ Push Manu /_/_/_/_/")
	fmt.Println()
	fmt.Println("1. pushする")
	fmt.Println("99. Main Menu に戻る")
	fmt.Println()	
	fmt.Println("※現在、以下のリモートリポジトリのURLが設定されています。")
	fmt.Println(string(remoteURL))
	fmt.Println()
	fmt.Printf("番号を選んでください。: ")
	var ans int
	fmt.Scan(&ans)

	if ans==1{
		fmt.Printf("push中...")
		out, err := exec.Command("git", "push", "-u", "origin", "master").Output()
		if err != nil {
			fmt.Println("pushに失敗しました。 ")
			fmt.Println(string(out))
			time.Sleep(1 * time.Second)
			push_menu()
		}else{
			fmt.Println("pushしました。 ")
			fmt.Println(string(out))
			time.Sleep(1 * time.Second)
			main_manu()
		}
	}else if ans==99{
		main_manu()		
	}else{
		fmt.Println("入力した番号が間違っています。やり直してください。")
		time.Sleep(1 * time.Second)
		push_menu()
	}
}

func setURL_menu(){
	out, err := exec.Command("git", "remote", "-v").Output()
	if err == nil {
		fmt.Println();
		fmt.Println("/_/_/_/_ リモートリポジトリURL設定 /_/_/_/_/")
		fmt.Println();
		if string(out)=="" {
			fmt.Println("リモートリポジトリのURLが設定されていません。URLを入力してください。")
			fmt.Printf("URL(99:戻る): ")
			var ans string
			fmt.Scan(&ans)
			if ans!="99"{
				out2, err2 := exec.Command("git", "remote", "add", "origin", ans).Output()
				if err2!=nil{
					fmt.Println("設定失敗!")
					fmt.Println(string(out2))
				}else{
					fmt.Println("以下のURLに設定しました。")
					out3, _ := exec.Command("git", "remote", "-v").Output()
					fmt.Println(string(out3))
				}
				time.Sleep(1 * time.Second)
			}	
		}else{
			fmt.Println("リモートリポジトリのURLは以下の通り設定されています。変更しますか？")
			fmt.Println(string(out))
			fmt.Printf("URL(99:戻る): ")
			var ans string
			fmt.Scan(&ans)
			if ans!="99"{
				out2, err2 := exec.Command("git", "remote", "set-url", "origin", ans).Output()
				if err2!=nil{
					fmt.Println("設定失敗!")
					fmt.Println(string(out2))
				}else{
					fmt.Println("以下のURLに設定しました。")
					out3, _ := exec.Command("git", "remote", "-v").Output()
					fmt.Println(string(out3))
				}
				time.Sleep(1 * time.Second)	
			}
		}
	}
	main_manu()
}

func commit_menu(){
	comment := getCommitMessage()
	fmt.Println();
	fmt.Println("/_/_/_/_ Commit Manu /_/_/_/_/")
	fmt.Println()
	fmt.Println("1. commit (comment:" + comment +")")
	fmt.Println("2. commit 取り消し")
	fmt.Println("----------------")
	fmt.Println("99: Main Manu に戻る")
	fmt.Println()	
	fmt.Printf("番号を選んでください。: ")
	var ans int
	fmt.Scan(&ans)

	if ans==99{
		main_manu()
	}else if ans==1{
		err := exec.Command("git", "commit", "-m", comment).Run()
		if err != nil {
			fmt.Println("commitに失敗しました。 ")
			fmt.Println(err)
		}else{
			fmt.Println("commitしました。 ")
		}
	}else if ans==2{
		err := exec.Command("git", "reset").Run()
		if err != nil {
			fmt.Println("commitの取り消しに失敗しました。 ")
			fmt.Println(err)
			}else{
			fmt.Println("commitを取り消しました。 ")
		}
	}else{
		fmt.Println("入力した番号が間違っています。やり直してください。")
		time.Sleep(1 * time.Second)
		commit_menu()
	}
	time.Sleep(1 * time.Second)
	main_manu()
}

func add_menu(code int){
    dir, _ := os.Getwd() //カレントディレクトリ
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

	fmt.Println();
	fmt.Println("/_/_/_/_ Add File /_/_/_/_/")
	fmt.Println()
	if code==0 {
		fmt.Println("追加するファイルNo.を選んでください。")		
	}else{
		fmt.Println("取り消しするファイルNo.を選んでください。")	
	}
	fmt.Println()
	
	var paths []string
	var cnt int = 0
    for _, file := range files {
        if !file.IsDir() {
			paths = append(paths, file.Name())
			out, err := exec.Command("git", "status", file.Name(), "-s").Output() //ファイルごとのstatus
			if(err==nil){
				var filestatus string
				mess := "--"
				if(len(string(out)) < 2){
					mess = "commit済み"
				}else{
					//filestatus = strings.TrimSpace(string(out)[0:2])
					filestatus = string(out)[0:2]			
				}
				if filestatus==" M" || filestatus=="AM" || filestatus=="MM" { mess = "内容変更あり" }
				if filestatus==" A" || filestatus=="A " { mess = "追加済み" }
				if filestatus=="M " { mess = "追加済み" }
				if filestatus==" R" { mess = "ファイル名変更" }
				if filestatus=="??" { filestatus="--" }
				//fmt.Println(string(out))
				fmt.Printf("%d: %s\t [%s]:%s\n", cnt, file.Name(), filestatus, mess)
				cnt++
			}
        }
	}
	fmt.Println("----------------")
	fmt.Println("99: Main Manu に戻る")
	fmt.Println()
	fmt.Printf("番号を選んでください。: ")
	var ans int
	fmt.Scan(&ans)

	if ans==99{
		main_manu()
	}else{
		if ans<len(paths){
			if code==0 {
				exec.Command("git", "add", paths[ans]).Run()
			}else{
				exec.Command("git", "reset", paths[ans]).Run()
			}
			add_menu(code)
		}else{
			fmt.Println("入力した番号が間違っています。やり直してください。")
			time.Sleep(1 * time.Second)
			add_menu(code)
		}
	}
}

func initial(){
	_, err := exec.Command("git", "status").Output()
	if(err!=nil){
		fmt.Println()
		fmt.Printf("リポジトリが作成されていません。git initしますか？ (y/n) :")
		fmt.Println()
		var ans string
		fmt.Scan(&ans)
		if ans=="y" {
			exec.Command("git", "init").Run()
			main_manu()
		}else{
			fmt.Println("終了します。")
		}
	}else{
		main_manu()
	}
}

func getCommitMessage() string{
	return "Commit " + time.Now().Format("2006-01-02 15:04:05")
}

func goodby(){
	fmt.Println("終了します。")
	time.Sleep(1 * time.Second)
	os.Exit(0)
}