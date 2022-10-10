package zz_generate

// copy GRPC ent-go generated protobuf codes to 'delivery' layer.
//go:generate rsync -avz --exclude generate.go  ../ent/proto/delivery/ ../delivery/grpc/
// generate grpc-gateway code in 'delivery' layer.
//go:generate protoc -I .. --grpc-gateway_out .. --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ../delivery/grpc/delivery.proto
