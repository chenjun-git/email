package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chenjun-git/umbrella-email/common"
	"github.com/chenjun-git/umbrella-email/directmail"
	"github.com/chenjun-git/umbrella-email/pb"
)

func (s *Server) SendSignupVerifyCode(ctx context.Context, req *pb.EmailSendSingleReq) (*pb.EmailSendSingleResp, error) {

	sendReq := directmail.SmsSendSingleReq{
		TemplateId: common.Config.Yunzhixun.SignupTemplateId,
		Param:      req.VerifyCode,
		Mobile:     req.Mobile,
		Uid:        "0",
	}
	resp := s.DireactMail.SendSingleVerifyCode(sendReq)
	code, err := strconv.Atoi(resp.Code)
	if err != nil {
		return nil, common.NewRPCError(common.EmailServiceInternalErr, fmt.Sprintf("code: %v, msg: %v", resp.Code, resp.Msg))
	}

	if code != yunzhixun.YUNZHIXUN_RESULT_CODE_OK {
		return nil, common.NewRPCError(common.EmailServiceInternalErr, fmt.Sprintf("code: %v, msg: %v", resp.Code, resp.Msg))
	}

	return nil, nil
}
