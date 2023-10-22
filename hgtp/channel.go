package hgtp

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type hgtpChannel tls.Conn

func NewHGTPChannel(config *channelConfig) (*hgtpChannel, error) {
	tcp_connection, err := net.DialTCP("tcp", nil, config.tcpAddr)
	if err != nil {
		return nil, err
	}

	tcp_connection.SetNoDelay(true)
	tls_connection := tls.Client(tcp_connection, config.tlsCongfig)

	err = tls_connection.Handshake()
	if err != nil {
		return nil, fmt.Errorf("TLS Handshake error: %s", err)
	}

	return (*hgtpChannel)(tls_connection), nil
}

func (_c *hgtpChannel) Close() error {
	c := (*tls.Conn)(_c)

	return (*tls.Conn)(c).Close()
}

func (_c *hgtpChannel) SetDeadline(time time.Time) {
	(*tls.Conn)(_c).SetDeadline(time)
}

func (c *hgtpChannel) SendRecv(msg hgtpMessage, maxDuration time.Duration) (hgtpMessage, error) {
	c.SetDeadline(time.Now().Add(maxDuration))

	err := c.Send(msg)
	if err != nil {
		return msg, err
	}
	err = c.Recv(&msg)
	return msg, err
}

func (_c *hgtpChannel) Send(msg hgtpMessage) error {
	c := (*tls.Conn)(_c)

	_, err := c.Write(msg.Frame[:])
	if err != nil {
		return err
	}

	extension := int(msg.UnpackExtension())
	// TODO: Why do we write this in chunks of 1024 instead of one big/bigger write?
	for i := 0; i < extension; i += 1024 {
		from := i
		to := i + 1024
		if to > len(msg.Unbounded) {
			to = len(msg.Unbounded)
		}

		_, err = c.Write(msg.Unbounded[from:to])
		if err != nil {
			return err
		}
	}
	return nil
}

// If an error is returned, the buffer is left in an undefined state
func (_c *hgtpChannel) Recv(buff *hgtpMessage) error {
	c := (*tls.Conn)(_c)

	_, err := c.Read(buff.Frame[:])
	if err != nil {
		return err
	}

	extension := int(buff.UnpackExtension())

	if extension > 0 {
		buff.Unbounded = make([]byte, extension)

		// TODO: Why do we read this in chunks of 1024 instead of one big/bigger write?
		for i := 0; i < extension; i += 1024 {
			from := i
			to := i + 1024
			if to > extension {
				to = extension
			}

			_, err = c.Read(buff.Unbounded[from:to])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
