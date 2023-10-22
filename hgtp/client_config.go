package hgtp

import (
	"io"
	"time"
)

const DEFAULT_PORT = 20057
const DEFAULT_CERTIFICATE_DIR = "/AvesTerra/Certificates/"

// Stolen from the python client
const CHANNEL_IDLE_TIMEOUT = 360 * time.Second

// Configuration for the AvesTerra client.
// Any optional field left unspecified will be set to its default value.
type ClientConfig struct {
	// Required: Address specifies the IP address or hostname of the AvesTerra
	// server to
	Address string

	// Optional: Port to use for the connection. Defaults to 20057, which is
	// the IANA approved AvesTerra port.
	Port int

	// Optional: CertificateDir specifies the path to the directory containing
	// .pem certificate file to use for the connection.
	// First tries to load "{address}.pem" where {address} is the decimal
	// representation of the 16bit integer representing target IPv4 address
	// (you can search "IPv4 to decimal" online to find a converter).
	// Defaults to "/AvesTerra/Certificates/".
	CertificateDir string

	// Optional: The number of channels to initially open. Defaults to 1.
	// If more channels are needed, they will lazily open as needed (up to
	// `MaxChannelCount`) at the cost of an initial performance hit
	// corresponding to the initial TCP+TLS handshake.
	InitialChannelCount int

	// Optional: The maximum number of idle channels to keep open. Defaults to
	// 8.
	// Channels are considered idle after `CHANNEL_IDLE_TIMEOUT`
	MaxIdleChannelCount int

	// Optional: The maximum number of channels to keep open. Defaults to 16.
	// Once reached, calls will block until a channel is available.
	MaxChannelCount int

	// Optional: KeyLogWriter optionally specifies a destination for TLS master
	// secrets in NSS key log format that can be used to allow external
	// programs such as Wireshark to decrypt TLS connections.
	// See https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format.
	// Use of KeyLogWriter compromises security and should only be
	// used for debugging.
	KeyLogWriter io.Writer
}
