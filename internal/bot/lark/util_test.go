package lark

import (
	"fmt"
	"log/slog"
	"os"
	"xiaoxiaojiqiren/config"
	"xiaoxiaojiqiren/internal/pkg/slogor"
)

func getClient() *Client {
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		TimeFormat: "2006-01-02 15:04:05.000",
		ShowSource: true,
		Level:      slog.LevelDebug,
	})))

	config := config.NewConfig()

	fmt.Printf("config: %+v\n", config)

	return NewClient(config)
}
