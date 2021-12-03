package server

import (
	"context"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

//下面的函数是服务器的原始API。（实现TinyKvServer）。
//在当前目录的sever.go中可以找到一些帮助器方法
//RawGet根据RawGetRequest的CF和Key字段返回相应的Get响应
func (server *Server) RawGet(_ context.Context, req *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error) {
	
	// Your Code Here (1).
	return nil, nil
}

//RawPut将目标数据放入存储器并返回相应的响应
func (server *Server) RawPut(_ context.Context, req *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error) {
	// Your Code Here (1).
	//提示：考虑使用存储。修改以存储要修改的数据。
	return nil, nil
}

//RawDelete从存储器中删除目标数据并返回相应的响应
func (server *Server) RawDelete(_ context.Context, req *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error) {
	// Your Code Here (1).
	//提示：考虑使用存储。修改以存储要删除的数据。
	return nil, nil
}

//RawScan从开始键开始扫描数据，直到限制。并返回相应的结果
func (server *Server) RawScan(_ context.Context, req *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error) {
	//提示：考虑使用Realer-TyrCF。
	return nil, nil
}
