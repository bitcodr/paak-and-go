package main

import (
	"context"
	"fmt"
)

func main(){
	ctx := context.Background()

	//logger := initLogger(ctx)
	//_ = level.Info(logger).Log(
	//	"msg", fmt.Sprintf("run %s with commit %s & commit date %s by version %s",
	//		serviceName, versionCommit, versionDate, versionTag),
	//	"func", "main", "when", "Bootstaping project")
	//
	//
	//cfg, err := config.Load(ctx, serviceName)
	//if err != nil {
	//	_ = level.Error(logger).Log(
	//		"err", err,
	//		"msg", "Error occurred while creating configs.",
	//		"func", "config.Load()", "when", "LoadConfig")
	//
	//	panic(err)
	//}
}
