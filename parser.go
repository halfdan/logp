package logp

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	COMMON_LOG_FORMAT        = "^(?P<RemoteAddress>\\S+) \\S+ (?P<RemoteUser>\\S+) \\[(?P<date>.*?) (?P<timezone>.*?)\\] \"\\S+ (?P<path>.*?) \\S+\" (?P<Status>\\S+) (?P<BodyBytesSent>\\S+)$"
	NCSA_EXTENDED_LOG_FORMAT = COMMON_LOG_FORMAT + "(?P<referrer>.*?)\" \"(?P<user_agent>.*?)"
	S3_LOG_FORMAT            = "\\S+ (?P<host>\\S+) \\[(?P<date>.*?) (?P<timezone>.*?)\\] (?P<RemoteAddress>\\S+) \\S+ \\S+ \\S+ \\S+ \"\\S+ (?P<path>.*?) \\S+\" (?P<Status>\\S+) \\S+ (?P<BodyBytesSent>\\S+) \\S+ \\S+ \\S+ \"(?P<referrer>.*?)\" \"(?P<user_agent>.*?)\""
	ICECAST2_LOG_FORMAT      = NCSA_EXTENDED_LOG_FORMAT + " (?P<session_time>\\S+)"
	COMBINED_LOG_FORMAT      = "^(?P<RemoteAddress>(?:\\d{1,3}\\.){3}\\d{1,3}) \\- (?P<RemoteUser>\\S+) \\[(?P<LocalTime>.+?)\\] \"(?P<Request>.+?)\" (?P<Status>\\d{3}) (?P<BodyBytesSent>\\d+) \"(?P<HttpReferer>.+?)\" \"(?P<HttpUserAgent>.+?)\"$"
)

type Line struct {
	RemoteAddress string
	RemoteUser    string
	LocalTime     time.Time
	Request       string
	Status        int
	BodyBytesSent int
	HttpReferer   string
	HttpUserAgent string
}

type Parser struct {
	lineParser *regexp.Regexp
}

func (li *Line) String() string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s",
		li.RemoteAddress,
		li.RemoteUser,
		li.LocalTime,
		li.Request,
		li.Status,
		li.BodyBytesSent,
		li.HttpReferer,
		li.HttpUserAgent,
	)
}

func NewParser(format string) (*Parser, error) {
	regex, err := regexp.Compile(format)
	if err != nil {
		return nil, errors.New("Cannot compile regular expression")
	}

	parser := &Parser{lineParser: regex}
	return parser, nil
}

func (p *Parser) Parse(line string) (*Line, error) {
	var err error
	fmt.Printf("%q\n", p.lineParser.SubexpNames())
	return nil, err
}
