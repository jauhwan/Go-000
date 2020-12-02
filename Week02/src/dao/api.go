package dao

import "jauhwan.com/Go-000/Week02/src/models"

type DaoInterface interface {
	OrderList() ([]*models.OrderListItem, error)
}
