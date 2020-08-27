package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/controller/server/execution"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
)

type customCTXKey string

type clientCPServicesServer struct {
	procExec execution.Execution
}

// NewProcServiceServer used to create a new execution context
func NewProcServiceServer(exec execution.Execution) clientCPproto.ClientCPServicesServer {
	return &clientCPServicesServer{
		procExec: exec,
	}
}

func (s *clientCPServicesServer) PostMetadata(ctx context.Context, request *clientCPproto.RequestToPostMetadata) (*clientCPproto.MetadataName, error) {
	uuid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request ID: %v, PostMetadata request received", uuid))

	name, err := s.procExec.SaveMetadata(ctx, request.Metadata)
	if err != nil {
		log.Error(err, fmt.Sprintf("request ID: %v, error in saving to etcd", uuid))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return name, nil
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *clientCPproto.RequestToGetAllMetadata) (*clientCPproto.MetadataArray, error) {
	uuid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request ID: %v, GetAllMetadata request received", uuid))

	dataList, err := s.procExec.ReadAllMetadata(ctx)
	if err != nil {
		log.Error(err, fmt.Sprintf("request ID: %v, error in getting all metadata from etcd", uuid))
	}
	return dataList, status.Error(codes.Internal, err.Error())
}

func (s *clientCPServicesServer) GetStreamLogs(request *clientCPproto.RequestForStreamLog, stream clientCPproto.ClientCPServices_GetStreamLogsServer) error {
	uuid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "error while assigning is to the request")
		return status.Error(codes.Internal, err.Error())
	}

	// TODO: relay stream logs from executor
	logString := &clientCPproto.Log{RequestId: uuid, Log: "lorem ipsum logger logger logger dumb"}
	log.Info(fmt.Sprintf("request ID: %v, getstream request received", uuid))

	err = stream.Send(logString)
	if err != nil {
		log.Error(errors.New(fmt.Sprintf("request id: %v, error in streaming logs, error details: %v", uuid, err)), "")
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *clientCPServicesServer) ExecuteJob(ctx context.Context, execute *clientCPproto.RequestForExecute) (*clientCPproto.Response, error) {
	//will be utilized after implementation
	//uid, err := id_generator.NextID()
	return nil, errors.New("not implemented yet")
}