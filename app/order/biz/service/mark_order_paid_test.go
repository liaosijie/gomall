package service

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD
	"testing"

	// order "douyin-gomall/gomall/rpc_gen/kitex_gen/order"
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
=======
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
	"testing"
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	"testing"

	"github.com/PiaoAdmin/gomall/app/order/biz/dal"
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
)

func TestMarkOrderPaid_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewMarkOrderPaidService(ctx)
	// init req and assert value

	req := &order.MarkOrderPaidReq{
		OrderId: 1893995460704608256,
		UserId:  1892469484459921408,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
