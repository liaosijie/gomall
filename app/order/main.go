<<<<<<< HEAD
/*
 * @Author: liaosijie
 * @Date: 2025-02-18 16:47:56
 * @Last Modified by: liaosijie
 * @Last Modified time: 2025-02-18 16:48:56
 */

=======
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
package main

import (
	"net"
	"time"

<<<<<<< HEAD
<<<<<<< HEAD
	//"douyin-gomall/gomall/app/order/biz/dal"
	//"douyin-gomall/gomall/app/order/conf"
	//"douyin-gomall/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/PiaoAdmin/gomall/app/order/biz/dal"
	"github.com/PiaoAdmin/gomall/app/order/conf"
	"github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
=======
=======
	"github.com/PiaoAdmin/gomall/app/order/biz/dal"
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
	"github.com/PiaoAdmin/gomall/app/order/conf"
	"github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
<<<<<<< HEAD
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	consul "github.com/kitex-contrib/registry-consul"
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	_ = godotenv.Load()
	dal.Init()
=======
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	_ = godotenv.Load()
	dal.Init()
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
	opts := kitexInit()

	svr := orderservice.NewServer(new(OrderServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithRegistry(r))
=======
	opts = append(opts, server.WithServiceAddr(addr))
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	//consul注册
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
