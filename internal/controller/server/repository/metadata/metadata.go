package metadata

import (
	"context"
	"fmt"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"octavius/internal/pkg/util"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/proto"
)

//Repository interface for functions related to metadata repository
type Repository interface {
	Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error)
	GetValue(ctx context.Context, jobName string) (*clientCPproto.Metadata, error)
	GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error)
	GetAvailableJobList(ctx context.Context) (*clientCPproto.JobList, error)
}

type metadataRepository struct {
	etcdClient etcd.Client
}

//NewMetadataRepository initializes metadataRepository with the given etcdClient
func NewMetadataRepository(client etcd.Client) Repository {
	return &metadataRepository{
		etcdClient: client,
	}
}

//Save marshals metadata and saves the value in etcd database with the given key
func (c *metadataRepository) Save(ctx context.Context, key string, metadata *clientCPproto.Metadata) (*clientCPproto.MetadataName, error) {

	val, err := proto.Marshal(metadata)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	dbKey := constant.MetadataPrefix + key

	gr, err := c.etcdClient.GetValue(ctx, dbKey)
	if gr != "" {
		return nil, status.Error(codes.AlreadyExists, constant.Etcd+constant.KeyAlreadyPresent)
	}

	if err != nil {
		if err.Error() != constant.NoValueFound {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	log.Info(fmt.Sprintf("Request ID: %v, saving metadata to etcd with value %s", ctx.Value(util.ContextKeyUUID), metadata.String()))
	err = c.etcdClient.PutValue(ctx, dbKey, string(val))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &clientCPproto.MetadataName{Name: key}
	return res, nil
}

//GetAll returns array of metadata
func (c *metadataRepository) GetAll(ctx context.Context) (*clientCPproto.MetadataArray, error) {
	res, err := c.etcdClient.GetAllValues(ctx, constant.MetadataPrefix)
	if err != nil {
		var arr []*clientCPproto.Metadata
		res := &clientCPproto.MetadataArray{Values: arr}
		return res, status.Error(codes.Internal, err.Error())
	}

	var resArr []*clientCPproto.Metadata
	for _, val := range res {
		metadata := &clientCPproto.Metadata{}
		err := proto.Unmarshal([]byte(val), metadata)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		resArr = append(resArr, metadata)
	}
	resp := &clientCPproto.MetadataArray{Values: resArr}
	return resp, nil
}

// GetAvailableJobList returns list of available jobs
func (c *metadataRepository) GetAvailableJobList(ctx context.Context) (*clientCPproto.JobList, error) {
	keys, _, err := c.etcdClient.GetAllKeyAndValues(ctx, constant.MetadataPrefix)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())

	}
	var jobList []string

	for index := range keys {
		jobList = append(jobList, strings.Split(keys[index], "/")[1])
	}
	return &clientCPproto.JobList{Jobs: jobList}, nil
}

func (c *metadataRepository) GetValue(ctx context.Context, jobName string) (*clientCPproto.Metadata, error) {
	dbKey := constant.MetadataPrefix + jobName
	gr, err := c.etcdClient.GetValue(ctx, dbKey)
	if err != nil {
		return &clientCPproto.Metadata{}, err
	}
	metadata := &clientCPproto.Metadata{}
	err = proto.Unmarshal([]byte(gr), metadata)
	return metadata, err

}
