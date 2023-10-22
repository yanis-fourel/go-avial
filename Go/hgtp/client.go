package hgtp

import (
	"errors"
	"time"
)

type Client struct {
	channPool *channPool
}

func NewClient(conf *ClientConfig) (*Client, error) {
	err := validateAndAddDefaultFields(conf)
	if err != nil {
		return nil, err
	}

	channPool, err := newChannPool(conf)
	return &Client{channPool}, err

}

// Cleanly shutdowns
// Also issues an HGTP 'BYE' command to the server.
// Client should not be used after calling this function
func (c *Client) Close() {
	c.channPool.CloseAll()
}

func validateAndAddDefaultFields(c *ClientConfig) error {
	if c == nil {
		return errors.New("Config cannot be nil")
	}

	if c.Address == "" {
		return errors.New("Address cannot be empty")
	}
	if c.Port == 0 {
		c.Port = DEFAULT_PORT
	}
	if c.CertificateDir == "" {
		c.CertificateDir = DEFAULT_CERTIFICATE_DIR
	}
	if c.InitialChannelCount == 0 {
		c.InitialChannelCount = 1
	}
	if c.MaxIdleChannelCount == 0 {
		c.MaxIdleChannelCount = 8
	}
	if c.MaxChannelCount == 0 {
		c.MaxChannelCount = 16
	}

	return nil
}

func (c *Client) transact(msg hgtpMessage, maxDuration time.Duration) (hgtpMessage, error) {
	chann, err := c.channPool.Borrow()
	if err != nil {
		return msg, err
	}

	msg, err = chann.SendRecv(msg, maxDuration)
	if err != nil {
		c.channPool.Close(chann)
		return msg, err
	}

	c.channPool.Return(chann)

	return msg, msg.Validate()
}
