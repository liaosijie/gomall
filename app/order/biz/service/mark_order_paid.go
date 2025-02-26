package service

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD
	// order "douyin-gomall/gomall/rpc_gen/kitex_gen/order"
=======
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	"fmt"

	"github.com/PiaoAdmin/gomall/app/order/biz/dal/mysql"
	"github.com/PiaoAdmin/gomall/app/order/biz/model"
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
<<<<<<< HEAD
}

// NewMarkOrderPaidService new MarkOrderPaidService
=======
} // NewMarkOrderPaidService new MarkOrderPaidService
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	if req.UserId == 0 || req.OrderId == 0 {
		err = fmt.Errorf("user_id or order_id can not be empty")
		return
	}
	_, err = model.GetOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStatePaid)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	resp = &order.MarkOrderPaidResp{}
	return
}
