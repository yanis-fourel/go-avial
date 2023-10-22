package hgtp

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"math/big"
	"net"
	"os"
	"path"
	"strings"

	"github.com/silenceper/pool"
)

type channPool struct {
	pool pool.Pool
}

func newChannPool(conf *ClientConfig) (*channPool, error) {
	channConf, err := NewChannelConfig(conf)

	if err != nil {
		return nil, fmt.Errorf("Failed to create TLS config: %s", err.Error())
	}

	pool, err := newPool(conf, channConf)
	if err != nil {
		return nil, fmt.Errorf("Failed to create hgtp channel pool: %s", err.Error())
	}

	return &channPool{pool}, nil
}

// Closes all the connection properly
func (p *channPool) CloseAll() {
	p.pool.Release()
}

func (p *channPool) Borrow() (*hgtpChannel, error) {
	conn, err := p.pool.Get()
	if err != nil {
		return nil, err
	}

	return conn.(*hgtpChannel), nil
}

// Returns channel for future use
func (p *channPool) Return(channel *hgtpChannel) error {
	return p.pool.Put(channel)
}

// Deletes channel, in case there was an error and its state is uncertain
func (p *channPool) Close(channel *hgtpChannel) error {
	return p.pool.Close(channel)
}

func newPool(clientConf *ClientConfig, channConf *channelConfig) (pool.Pool, error) {

	chanFactory := func() (interface{}, error) {
		return NewHGTPChannel(channConf)
	}

	chanClose := func(conn interface{}) error {
		c := conn.(*hgtpChannel)

		var msg hgtpMessage
		msg.PackCommand(CommandBye)
		_ = c.Send(msg)
		return c.Close()
	}

	poolConfig := &pool.Config{
		InitialCap:  clientConf.InitialChannelCount,
		MaxIdle:     clientConf.MaxIdleChannelCount,
		MaxCap:      clientConf.MaxChannelCount,
		Factory:     chanFactory,
		Close:       chanClose,
		IdleTimeout: CHANNEL_IDLE_TIMEOUT,
	}

	return pool.NewChannelPool(poolConfig)
}

// Everything we need to create a new channel and connect it to the server.
type channelConfig struct {
	tlsCongfig *tls.Config
	tcpAddr    *net.TCPAddr
}

func NewChannelConfig(config *ClientConfig) (*channelConfig, error) {
	address := fmt.Sprintf("%s:%d", config.Address, config.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	numAddr := fmt.Sprint(ip4toInt(tcpAddr.IP))
	certificateRoots, err := loadCertificateRoots(config.CertificateDir, numAddr)
	if err != nil {
		return nil, err
	}

	tlsconfig := tls.Config{
		RootCAs:            certificateRoots,
		InsecureSkipVerify: true,
		KeyLogWriter:       config.KeyLogWriter,
	}

	return &channelConfig{
		tlsCongfig: &tlsconfig,
		tcpAddr:    tcpAddr,
	}, nil
}

func ip4toInt(IPv4Address net.IP) int64 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(IPv4Address.To4())
	return IPv4Int.Int64()
}

func loadCertificateRoots(dirpath, decimaladdr string) (*x509.CertPool, error) {
	paths := [...]string{
		path.Join(dirpath, decimaladdr+".pem"),
		path.Join(dirpath, "local.pem"),
		path.Join(dirpath, "avesterra.pem"),
	}
	var pempath string
	var data []byte
	var err error

	fileinfo, err := os.Stat(dirpath)
	if err != nil {
		return nil, err
	}

	if !fileinfo.IsDir() {
		return nil, fmt.Errorf("'%s' must be a directory, see ClientConfig doc", dirpath)
	}

	for i := range paths {
		pempath = paths[i]
		data, err = os.ReadFile(pempath)
		if !os.IsNotExist(err) {
			break
		}
	}

	if os.IsNotExist(err) {
		return nil, fmt.Errorf(
			"Failed to find .pem, looked up these %d paths: '%s'",
			len(paths),
			strings.Join(paths[:], "', '"),
		)
	}

	if err != nil {
		return nil, err
	}

	certificateRoots := x509.NewCertPool()
	ok := certificateRoots.AppendCertsFromPEM(data)
	if !ok {
		return nil, fmt.Errorf("Failed to decode .pem certificate '%s'", pempath)
	}

	return certificateRoots, nil
}
