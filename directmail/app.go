package directmail

import (
	"time"
)

type DirectMailApp struct {
	Version          string
	AccessKeyId      string
	SignatureMethod  string
	SignatureVersion string
	RegionId         string
	Timeout          time.Duration
}

func NewDirectMailApp(version, accessKeyId, signatureMethod, signatureVersion, regionId string, timeout time.Duration) *YunzhixunApp {
	return &DirectMailApp{
		Version:          version,
		AccessKeyId:      accessKeyId,
		SignatureMethod:  signatureMethod,
		SignatureVersion: signatureVersion,
		RegionId:         regionId,
		Timeout:          timeout,
	}
}

func (dm *DirectMailApp) SetHttpClientTimeout() {
	if dm.Timeout != 0 {
		SetClientTimeout(dm.Timeout)
	}
}
