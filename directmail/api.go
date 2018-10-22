package directmail

import (
	"encoding/json"
)

type DirectMailReq struct {
	Format           string
	Version          string
	AccessKeyId      string
	Signature        string
	SignatureMethod  string
	Timestamp        string
	SignatureVersion string
	SignatureNonce   string
	RegionId         string
}

func (r *DirectMailReq) setAppInfo(version, accessKeyId, signatureMethod, signatureVersion, regionId string) {
	r.Version = version
	r.AccessKeyId = accessKeyId
	r.SignatureMethod = signatureMethod
	r.SignatureVersion = signatureVersion
	r.RequestId = regionId
}

type DirectMailResp struct {
	RequestId string
}

type EmailSendSingleReq struct {
	DirectMailReq
	Action         string
	ReplyToAddress string
	AddressType    string
	ToAddress      string
	FromAlias      string
	Subject        string
	HtmlBody       string
	TextBody       string
	ClickTrace     string
}

type EmailSendSingleResp struct {
	DirectMailResp
}

type EmailSendGroupReq struct {
	DirectMailReq
	Action        string
	AccountName   string
	AddressType   string
	TemplateName  string
	ReceiversName string
	TagName       string
	ClickTrace    string
}

type EmailSendGroupResp struct {
	DirectMailResp
}

func (dm *DirectMailApp) SendSingleVerifyCode(req DirectSendSingleReq) *DirectMailSendSingleResp {
	req.setAppInfo(dm.Version, dm.AccessKeyId, dm.SignatureMethod, dm.SignatureVersion, dm.RequestId)

	resp, err := httpReqWithParams("POST", DIRECT_MAIL_API_HOST, req)
	if err != nil {
		return &EmailSendSingleResp{
			DirectMailResp: DirectMailRespResp{},
		}
	}

	var result EmailSendSingleResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return &EmailSendSingleResp{
			DirectMailResp: DirectMailResp{},
		}
	}

	return &result
}
