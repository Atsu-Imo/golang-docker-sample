# golang-docker-sample

GolangでシンプルなREST APIを作成したサンプルです。  
`docker-compose up`するだけでローカルの環境が作成できることが目標。

## 起動方法

```bash
$ docker-compose up
```


## 使用ライブラリ、フレームワークなど

### echo

[Echo - High performance, minimalist Go web framework](https://echo.labstack.com/)

シンプルに見えたので採用しました。
他のフレームワーク(`grilla`, `Gin`などなど)については使用したことがないので比較はできませんが、現時点では問題ないです。