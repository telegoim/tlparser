package mtproto

import (
	"fmt"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
)

func (c *Conn) newEncryptedMessage(id int64, seq int32, payload bin.Encoder, b *bin.Buffer) error {
	if err := payload.Encode(b); err != nil {
		return err
	}
	{
		typeID, err := b.PeekID()
		if err == nil {
			c.log.With(
				zap.Int64("msg_id", id),
				zap.String("message_type", fmt.Sprintf("0x%x", typeID)),
				zap.String("message_type_str", c.types.Get(typeID)),
			).Debug("Request")
		}
	}
	d := crypto.EncryptedMessageData{
		SessionID:              atomic.LoadInt64(&c.sessionID),
		Salt:                   atomic.LoadInt64(&c.salt),
		MessageID:              id,
		SeqNo:                  seq,
		MessageDataLen:         int32(b.Len()),
		MessageDataWithPadding: b.Copy(),
	}
	if err := c.cipher.Encrypt(c.authKey, d, b); err != nil {
		return err
	}

	return nil
}
