package models

import (
	"context"
	"fmt"

	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// PushCi - метод, отправляющий КЕ в cmdb с конвертацией Ci в protobuf Ci
func (cw *ConfigWrapper) PushCi(ctx context.Context, ci *Ci) error {
	var req pb.PushObjectRequest
	req.Item = ci.ToPB()
	_, err := cw.Conns.Cmdb.PushObject(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

// PushCi - метод, отправляющий КЕ пакетно в cmdb с конвертацией Ci в protobuf Ci
func (cw *ConfigWrapper) PushBatchCis(ctx context.Context, cis []Ci) error {
	var req pb.PushBatchObjectsRequest
	var pbcis []*pb.Ci
	for _, v := range cis {
		pbci := v.ToPB()
		pbcis = append(pbcis, pbci)
	}
	req.Item = pbcis
	resp, err := cw.Conns.Cmdb.PushBatchObject(ctx, &req)
	if err != nil {
		return err
	}
	// Так как gRPC возвращает либо response либо err, проверяется дополнительно внутренний код ответа приложения при пакетной загрузке
	if resp.Code != "OK" {
		return fmt.Errorf("cmdb return code: %s", resp.Code)
	}
	return nil
}

// PushCi - метод, удаляющий КЕ в cmdb
func (cw *ConfigWrapper) DeleteCi(ctx context.Context, ciname string) error {
	var req pb.DeleteObjectRequest
	req.Name = ciname
	_, err := cw.Conns.Cmdb.DeleteObject(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

// PushCi - метод, удаляющий КЕ в cmdb пакетно
func (cw *ConfigWrapper) DeleteBatchCis(ctx context.Context, cis []string) error {
	var req pb.DeleteBatchObjectsRequest
	req.Item = cis
	resp, err := cw.Conns.Cmdb.DeleteBatchObject(ctx, &req)
	if err != nil {
		return err
	}
	// Так как gRPC возвращает либо response либо err, проверяется дополнительно внутренний код ответа приложения при пакетной загрузке
	if resp.Code != "OK" {
		return fmt.Errorf("cmdb return code: %s", resp.Code)
	}
	return nil
}
