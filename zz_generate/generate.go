package zz_generate

//go:generate rsync -avz --exclude generate.go  ../ent/proto/delivery/ ../delivery/grpc/
