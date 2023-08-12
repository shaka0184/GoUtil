package log

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func Info(ctx context.Context, projectID, logName string, m string) {
	logPrintln(ctx, projectID, logName, m, logging.Info)
}

func Error(ctx context.Context, projectID, logName string, err error) {
	var convertedErr interface{ StackTrace() errors.StackTrace }

	if errors.As(err, &convertedErr) {
		// スタックトレースがある場合
		logPrintln(ctx, projectID, logName, convertedErr, logging.Error)
	} else {
		logPrintln(ctx, projectID, logName, err, logging.Error)
	}
}

func logPrintln(ctx context.Context, projectID, logName string, v interface{}, logLevel logging.Severity) {
	if len(projectID) == 0 {
		// 通常のログ出力
		log.Printf("%+v\n", v)
	} else {
		client, err2 := logging.NewClient(ctx, projectID)
		if err2 != nil {
			// ログクライアントのインスタンス作成時にエラーが発生した場合
			// インスタンス作成エラーを出力
			log.Println(err2)

			// 通常のログ出力
			log.Printf("%+v\n", v)
		}
		defer client.Close()

		logger := client.Logger(logName).StandardLogger(logLevel)

		// GCPのフォーマットでログ出力
		logger.Println(fmt.Sprintf("%+v\n", v))
	}
}
