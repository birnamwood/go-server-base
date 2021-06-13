package main

import (
	"flag"

	"github.com/birnamwood/go-server-base/initialize/alert"
	"github.com/birnamwood/go-server-base/initialize/config"
	"github.com/birnamwood/go-server-base/initialize/database"
	"github.com/birnamwood/go-server-base/initialize/logger"
	"github.com/birnamwood/go-server-base/initialize/route"

	"go.uber.org/zap"
)

func main() {
	// 環境設定取得 flag.String(<パラメータ名>, <デフォルト値>, <パラメータの説明>)
	env := flag.String("e", "local", "動作環境名")
	//変数宣言のあとに、flag.Parseを実行することでコマンドラインのパラメータがパースされ、各変数に値が格納されます
	flag.Parse()
	//パラメータを渡してconfigの初期化を行う
	config.Init(*env, "environment")

	//Loggerの初期化
	logger := logger.Init("app_log")
	zap.ReplaceGlobals(logger)
	logger.Info("Logger Initialize")
	alert.Init()

	//Database初期化
	database.Init()
	logger.Info("DB Initialize")

	route.Init()
	database.Close()
}
