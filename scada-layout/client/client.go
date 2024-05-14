package client

import (
	"context"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/scada-layout/pb"
)

type ScadaLayoutClient interface {
	GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error)
	GetFieldsTagsReader(ctx context.Context, fields []string) (FieldTagsReader, error)
}

type clientImpl struct {
	address string
}

func New(address string) ScadaLayoutClient {
	return &clientImpl{
		address: address,
	}
}

func (impl *clientImpl) GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	return clt.GetFieldsTags(ctx, &pb.FieldList{
		Fields: fields,
	})
}

func (impl *clientImpl) GetFieldsTagsReader(ctx context.Context, fields []string) (FieldTagsReader, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetFieldsTags(ctx, &pb.FieldList{
		Fields: fields,
	})
	if err != nil {
		return nil, err
	}
	return NewFieldTagsReader(resp), nil
}

type FieldTagsReader interface {
	GetFieldTag(f string) *pb.GetFieldsTagsResponse_FieldTag
}

func NewFieldTagsReader(resp *pb.GetFieldsTagsResponse) FieldTagsReader {
	data := make(map[string]*pb.GetFieldsTagsResponse_FieldTag)
	for _, ft := range resp.FieldTags {
		data[ft.Field] = ft
	}
	return &fieldTagsReaderImpl{
		data: data,
	}
}

type fieldTagsReaderImpl struct {
	data map[string]*pb.GetFieldsTagsResponse_FieldTag
}

func (impl *fieldTagsReaderImpl) GetFieldTag(f string) *pb.GetFieldsTagsResponse_FieldTag {
	return impl.data[f]
}
