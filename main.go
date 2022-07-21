//catコマンドの再現
package main

import (
	"bufio" //スキャナー関連で使う
	"flag"  //コマンドライン関連で使う
	"fmt"
	"os"            //ファイルの読み書き関連で使う
	"path/filepath" //ファイルのパス関連で使う
)

func main() {
	//　cat -n の「n」の部分を再現
	//  -n はオプションなのでON/OFFのようなスイッチを提供するflag.Boolを使用(
	var n = flag.Bool("n", false, "通し番号を付与する")
	//　catコマンドの引数を再現(ファイル名の部分を再現)
	flag.Parse()
	var (
		//ファルが複数あることを想定し、右辺の構文を使用
		files = flag.Args()
		//上記のファイルがあるパスを再現
		path, err = os.Executable()
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
	}

	//実行ファイルのディレクトリ名を取得
	path = filepath.Dir(path)
	//通し番号の用意
	i := 1

	//あるディレクトリの中にcatコマンドで出力したいファイルがある、といった想定
	//変数xはfilesのインデックス番号
	for x := 0; x < len(files); x++ {
		//上記パスにあるファイルの中身を読み込んで左辺に代入
		sf, err := os.Open(filepath.Join(path, files[x]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		} else {
			//上記で読み込まれた内容をターミナル出力するための準備
			scanner := bufio.NewScanner(sf)
			for ; scanner.Scan(); i++ {
				if *n {
					//オプションがある場合
					//cat -n の「n」を出力する役割
					fmt.Printf("%v: ", i)
				}
				//catコマンドのデフォルトの役割（ファイルの中身を表示する役割）
				fmt.Println(scanner.Text())
			}
		}
	}
}
