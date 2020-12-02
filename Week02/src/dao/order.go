package dao

import (
	"github.com/pkg/errors"
	"jauhwan.com/Go-000/Week02/src/init"
	"jauhwan.com/Go-000/Week02/src/models"
)

type OrderDao struct {
}

func (o *OrderDao) OrderList() ([]*models.OrderListItem, error) {
	sql := `SELECT order_id FROM order_info`
	rows, err := init.DB.Query(sql)
	if err != nil {
		return nil, errors.Wrap(err, "[mysql].[query].[order_info]")
	}
	orderList := make([]*models.OrderListItem, 0)
	for rows.Next() {
		orderItem := &models.OrderListItem{}
		if err := rows.Scan(&orderItem.OrderId); err != nil {
			return nil, errors.Wrap(err, "[mysql].[scan].[order_info]")
		}
		orderList = append(orderList, orderItem)
	}
	return orderList, nil
}
