.PHONY: all go

all: go

go: go/ServiceCmdb.pb.go go/ServiceEvents.pb.go  go/ServiceIdm.pb.go go/ServiceMetrics.pb.go

cmdb:  go/ServiceCmdb.pb.go

events: go/ServiceEvents.pb.go

idm: go/ServiceIdm.pb.go

metrics: go/ServiceMetrics.pb.go

clean:
	rm -f cmdb/*.pb.go
	rmdir cmdb* 2>/dev/null || true
	rm -f events/*.pb.go
	rmdir events* 2>/dev/null || true
	rm -f metrics/*.pb.go
	rmdir metrics* 2>/dev/null || true
	rm -f idm/*.pb.go
	rmdir idm* 2>/dev/null || true

clean_cmdb:
	rm -f cmdb/*.pb.go
	rmdir cmdb* 2>/dev/null || true

clean_events:
	rm -f events/*.pb.go
	rmdir events* 2>/dev/null || true

clean_metrics:
	rm -f metrics/*.pb.go
	rmdir metrics* 2>/dev/null || true

clean_idm:
	rm -f idm/*.pb.go
	rmdir idm* 2>/dev/null || true


go/ServiceCmdb.pb.go: cmdb.proto
	mkdir -p ./cmdb # make directory for go package
	protoc $$PROTO_PATH --go_opt=paths=source_relative   --go-grpc_out=. --go_out=./cmdb cmdb.proto


go/ServiceEvents.pb.go: events.proto
	mkdir -p ./events # make directory for go package
	protoc $$PROTO_PATH --go_opt=paths=source_relative   --go-grpc_out=. --go_out=./events events.proto

go/ServiceIdm.pb.go: idm.proto
	mkdir -p ../idm/proto/idm # make directory for go package
	protoc $$PROTO_PATH --go_opt=paths=source_relative   --go-grpc_out=../idm/proto/ --go_out=../idm/proto/idm idm.proto

go/ServiceMetrics.pb.go: metrics.proto
	mkdir -p ./metrics # make directory for go package
	protoc $$PROTO_PATH --go_opt=paths=source_relative   --go-grpc_out=. --go_out=./idm metrics.proto

