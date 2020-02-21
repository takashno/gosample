# gosample

## モジュール管理

Go Modulesを利用する．

## GOPATH

どこでもいい。  
pkgのキャッシュが散在するのを防ぐためある程度まとまった単位で作る方が良い。  

複数のGOPATHを使い分けると、環境変数への `SET` や `export` が面倒かも?  
なんか便利ツールあったりすんのかな。

## なんか作るときに始めること

1. GOPATH決める

```
export GOPATH=/path/to/your/development
```

2. GOPATH配下に作るモジュールのディレクトリを作る

```
mkdir gosample
```

3. 初期化

```
cd gosample
go mod init github.com/takashno/gosample
```

4. 利用するモジュールの追加

```
cd gosample
go get github.com/takashno/gosample2
```

`$GOPATH` にpkgというディレクトリが作成されてダウンロード等が行われる。  
go.mod に依存関係が追加される。  
go.sum にチェックサムみたいなのが追加される。

5. ビルド

```
cd gosample
go build
```

Windowsとかだと、exeファイルがビルドして作成される。

```
go build -o ./bin/gosample.exe
```

出力先とか指定もできる。

## Go Modules の replace ディレクティブについて

`Go Modules` には `replace` というディレクティブが用意されていて、
モジュールの解決パスを変更できるというもの。  
ローカル環境以外に使うことがあるのか良く分からない。

プロジェクト内でモジュール分割をしていて同時開発している場合などに、  
モジュールAから一時的に自分が改修をしているモジュールBを参照させたい場合などに有効かもしれない。

```
$GOPATH
  |
  |-- moduleA  ★並列にCloneしている状態★
  |     |-- go.mod
  |     |-- main.go ★ここで利用している★
  |     `-- *
  |
  |-- moduleB  ★並列にCloneしている状態★
  |     |-- go.mod
  |     |-- main.go
  |     |-- helper
  |     |     `-- utils.go ★このファイルを改修している★
```

このとき、ModuleAのgo.modには、ModuleBが `require` で定義されている。

```
module github.com/takashno/gosample

go 1.12

require (
	github.com/takashno/gosample2 v0.0.0-20200221022707-dbe481dd8466  ★これがModuleB★
	golang.org/x/text v0.3.2
)
```

これを一時的にローカル環境のものを利用するようにするには、

```
module github.com/takashno/gosample

go 1.12

replace github.com/takashno/gosample2 => ../gosample2  ★ replace ディレクティブを利用する★

require (
	github.com/takashno/gosample2 v0.0.0-20200221022707-dbe481dd8466
	golang.org/x/text v0.3.2
)
```

[公式ドキュメント](https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive) にはきちんと記載があった。
