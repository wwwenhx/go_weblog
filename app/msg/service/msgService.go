package service

import (
	"context"
	"gin_gomicro/idl/pb"
	"sync"
)

type MsgSrv struct {
	pb.UnimplementedMsgServer
}

var MsgIns *MsgSrv
var MsgOne sync.Once

func GetMsgSrv() *MsgSrv {
	MsgOne.Do(func() {
		MsgIns = &MsgSrv{}
	})
	return MsgIns
}

func (m *MsgSrv) GetMsg(ctx context.Context, req *pb.MsgReq) (*pb.MsgRes, error) {
	res := &pb.MsgRes{}
	return res, nil
}

func (m *MsgSrv) SendMsg(ctx context.Context, req *pb.MsgReq) (*pb.MsgRes, error) {
	res := &pb.MsgRes{}
	return res, nil
}
