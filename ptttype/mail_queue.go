package ptttype

import "github.com/PichuChen/go-bbs/types"

type MQType int

const (
	MQ_TEXT MQType = iota
	MQ_UUENCODE
	MQ_JUSTIFY
)

var MQTypeString = []string{
	"text",
	"uuencode",
	"mq-justify",
}

func (m MQType) String() string {
	if int(m) < len(MQTypeString) {
		return MQTypeString[m]
	}

	return "unknown"
}

type MailQueue struct {
	FilePath [FNLEN]byte
	Subject  [STRLEN]byte
	MailTime types.Time4
	Sender   [IDLEN + 1]byte
	Username [USERNAMESZ]byte
	RCPT     [RCPTSZ]byte
	Method   int32
	Niamod   []byte
}
