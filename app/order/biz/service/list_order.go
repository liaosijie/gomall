<<<<<<< HEAD
/*
 * @Author: liaosijie
 * @Date: 2025-02-18 22:28:50
 * @Last Modified by: liaosijie
 * @Last Modified time: 2025-02-18 23:30:55
 */

=======
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
package service

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD

	"github.com/PiaoAdmin/gomall/app/cart"

	// "douyin-gomall/gomall/app/order/biz/dal/mysql"
	// "douyin-gomall/gomall/app/order/biz/model"
	"github.com/PiaoAdmin/gomall/app/order/biz/dal/mysql"
	"github.com/PiaoAdmin/gomall/app/order/biz/model"

	// order "douyin-gomall/gomall/rpc_gen/kitex_gen/order"
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
=======
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======

	"github.com/PiaoAdmin/gomall/app/order/biz/dal/mysql"
	"github.com/PiaoAdmin/gomall/app/order/biz/model"
	"github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
)

type ListOrderService struct {
	ctx context.Context
<<<<<<< HEAD
}

// NewListOrderService new ListOrderService
=======
} // NewListOrderService new ListOrderService
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
<<<<<<< HEAD
<<<<<<< HEAD
	model.ListOrder(s.ctx,mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001,err.Error())
	}

	var orders []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item:&cart.CartItem{
					ProductId:oi.ProductId,
					Quantity:oi.Quantity,
				},
				Cost:oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:		v.OrderId,
			UserId:			v.UserId,
			UserCurrency:	v.UserCurrency,
			Email:			v.Consignee.Email,
			Address:		&order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				Country:	   v.Consignee.Country,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				ZipCode:       v.Consignee.ZipCode,
			},
			Items: items,
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
=======

>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	orders, err := model.ListOrder(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	var list []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, oi := range v.OrderItem {
			items = append(items, &order.OrderItem{
				Cost: oi.Cost,
				Item: &cart.CartItem{
					ProductId: oi.ProductID,
					Quantity:  oi.Quantity,
				},
			})
		}
		o := &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Country:       v.Consignee.Country,
				City:          v.Consignee.City,
				StreetAddress: v.Consignee.StreetAddress,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		}
		list = append(list, o)
	}
	resp = &order.ListOrderResp{
		Orders: list,
	}
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
	return
}
