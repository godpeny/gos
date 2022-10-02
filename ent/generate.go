package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
// for GRPC protobuf
//go:generate go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema
