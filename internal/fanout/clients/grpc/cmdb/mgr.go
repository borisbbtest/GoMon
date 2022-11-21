// Package пакет по работе с сервисом CMDB
package integrationcmdb

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// pbCiMapCi - функция размапа обертки
func pbCiMapCi(v *cmdb.Ci) *Ci {
	return &Ci{
		Name:        v.Name,
		Description: v.Description,
		Update:      v.Update.AsTime(),
		Created:     v.Created.AsTime(),
		CreatedBy:   v.CreatedBy,
		Type:        v.Type,
	}
}

// GetBachListItem - получает список КЕ по имени
func (hook *ServiceWrapperCmdb) GetBachListItem(ctx context.Context, items *[]string) (resp *[]*Ci, err error) {

	req := cmdb.GetBatchObjectsRequest{
		Item: *items,
	}
	cis_, err := hook.Ci.GetBatchObject(ctx, &req)
	if err != nil {
		return nil, err
	}
	order := []*Ci{}
	for _, v := range cis_.Item {
		buff := pbCiMapCi(v)
		order = append(order, buff)
	}
	resp = &order
	return resp, nil
}

// GetItem - получает одно ке
func (hook *ServiceWrapperCmdb) GetItem(ctx context.Context, items *string) (resp *Ci, err error) {

	req := cmdb.GetObjectRequest{
		Name: *items,
	}
	cis_, err := hook.Ci.GetObject(ctx, &req)
	if err != nil {
		return nil, err
	}
	resp = pbCiMapCi(cis_.Item)
	return resp, nil
}
