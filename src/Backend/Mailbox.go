package Backend

import (
	"io"
	"time"

	ORM "antegral.net/mailmix/src/Database/Sqlc"
	"antegral.net/mailmix/src/Log"
	"github.com/DusanKasan/parsemail"
	"github.com/emersion/go-imap"
	"github.com/google/uuid"
	"github.com/zeebo/blake3"
)

type MailboxDesc struct {
	Id    uuid.UUID
	Owner uuid.UUID
	Name  string
}

type Mailbox struct {
	Desc    MailboxDesc
	Queries *ORM.Queries
}

// Name returns this mailbox name.
func (b Mailbox) Name() string {
	return b.Desc.Name
}

// Info returns this mailbox info.
func (b Mailbox) Info() (*imap.MailboxInfo, error) {
	// TODO: Getting Mailbox attributes
	// Mailbox Attributes (https://www.iana.org/assignments/imap-mailbox-name-attributes/imap-mailbox-name-attributes.xhtml)
	MailboxInfo := imap.MailboxInfo{
		Attributes: []string{""},
		Delimiter:  "/",
		Name:       b.Name(),
	}

	return &MailboxInfo, nil
}

// Status returns this mailbox status. The fields Name, Flags, PermanentFlags
// and UnseenSeqNum in the returned MailboxStatus must be always populated.
// This function does not affect the state of any messages in the mailbox. See
// RFC 3501 section 6.3.10 for a list of items that can be requested.
func (b Mailbox) Status(items []imap.StatusItem) (*imap.MailboxStatus, error) {
	// TODO: Implementing StatusItem

	Result := imap.MailboxStatus{
		Name: b.Name(),
	}
	return &Result, nil
}

// SetSubscribed adds or removes the mailbox to the server's set of "active"
// or "subscribed" mailboxes.
func (b Mailbox) SetSubscribed(subscribed bool) error {
	return nil
}

// Check requests a checkpoint of the currently selected mailbox. A checkpoint
// refers to any implementation-dependent housekeeping associated with the
// mailbox (e.g., resolving the server's in-memory state of the mailbox with
// the state on its disk). A checkpoint MAY take a non-instantaneous amount of
// real time to complete. If a server implementation has no such housekeeping
// considerations, CHECK is equivalent to NOOP.
func (b Mailbox) Check() error {
	// TODO: housekeeping considerations
	return nil
}

// ListMessages returns a list of messages. seqset must be interpreted as UIDs
// if uid is set to true and as message sequence numbers otherwise. See RFC
// 3501 section 6.4.5 for a list of items that can be requested.
//
// Messages must be sent to ch. When the function returns, ch must be closed.
func (b Mailbox) ListMessages(uid bool, seqset *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
	// section 6.4.5, RFC 3501: https://www.rfc-editor.org/rfc/rfc3501#section-6.4.5
	return nil
}

func (b Mailbox) SearchMessages(uid bool, criteria *imap.SearchCriteria) ([]uint32, error) {
	// uid가 true인 경우 반환값에 uid로, false인 경우 메시지 시퀸스 번호로 반환
	return nil, nil
}

func (b Mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	// TODO: 본 메일박스에 메일 넣기 (\Recent Flag 추가로 붙여서 넣어야 함)

	IsFlagged := false
	for _, flag := range flags {
		if flag == "\\Recent" {
			IsFlagged = true
			break
		} else {
			continue
		}
	}

	if !IsFlagged {
		Log.Verbose.Printf("Add \\Recent to flags")
		flags = append(flags, "\\Recent")
	}

	Log.Verbose.Printf("flags: %v\n", flags)

	ParsedMail, err := parsemail.Parse(body)
	if err != nil {
		return err
	}

	h := blake3.New()
	if _, err := io.Copy(h, body); err != nil {
		return err
	}
	Hash := h.Sum(nil)

	b.Queries.CreateMail(ctx, ORM.CreateMailParams{
		Uuid:     uuid.New(),
		Boxuuid:  uuid.New(), // TODO: User의 기본 Mailbox UUID 가져와서 넣기
		Header:   ParsedMail.Header.,
		Sentfrom: ParsedMail.Sender.String(),
		Sentto:   ParsedMail.To[0],
		Sentat:   ParsedMail.Date,
		hash:     Hash,
		Flags:    flags,
		Size:     2,
	})

	return nil
}

func (b Mailbox) UpdateMessagesFlags(uid bool, seqset *imap.SeqSet, operation imap.FlagsOp, flags []string) error {
	// TODO: 본 메일박스에 있는 특정 메일의 Flag를 재설정 (seqset에 특정 메일의 고유번호 포함)
	// uid가 true인 경우 seqset을 uid로 해석하고, false인 경우 메시지 시퀸스 번호로 해석
	return nil
}

func (b Mailbox) CopyMessages(uid bool, seqset *imap.SeqSet, dest string) error {
	// TODO: dest 메일박스로 본 메일박스에 있는 특정 메일을 복사 (seqset에 특정 메일의 고유번호 포함)
	// uid가 true인 경우 seqset을 uid로 해석하고, false인 경우 메시지 시퀸스 번호로 해석
	return nil
}

func (b Mailbox) Expunge() error {
	// TODO: 본 메일박스에 포함되고 \Deleted 플래그가 설정된 모든 메일을 제거
	return nil
}
