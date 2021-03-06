package db

import (
	"strings"

	"github.com/dennis-tra/nebula-crawler/pkg/models"
	swarm "github.com/libp2p/go-libp2p-swarm"
)

// KnownErrors contains a list of known errors. Property key + string to match for
var KnownErrors = map[string]string{
	models.DialErrorIoTimeout:                                  "i/o timeout",
	models.DialErrorConnectionRefused:                          "connection refused",
	models.DialErrorProtocolNotSupported:                       "protocol not supported",
	models.DialErrorPeerIDMismatch:                             "peer id mismatch",
	models.DialErrorNoRouteToHost:                              "no route to host",
	models.DialErrorNetworkUnreachable:                         "network is unreachable",
	models.DialErrorNoGoodAddresses:                            "no good addresses",
	models.DialErrorContextDeadlineExceeded:                    "context deadline exceeded",
	models.DialErrorNoPublicIP:                                 "no public IP address",
	models.DialErrorMaxDialAttemptsExceeded:                    "max dial attempts exceeded",
	models.DialErrorHostIsDown:                                 "host is down",
	models.DialErrorStreamReset:                                "stream reset",
	models.DialErrorNegotiateSecurityProtocolNoTrailingNewLine: "failed to negotiate security protocol: message did not have trailing newline",
}

// DialError extracts the appropriate error type from the given error.
func DialError(err error) string {
	if dialErr, ok := err.(*swarm.DialError); ok && dialErr.Cause != nil {
		return DialError(dialErr.Cause)
	}

	for key, errStr := range KnownErrors {
		if strings.Contains(err.Error(), errStr) {
			return key
		}
	}
	return models.DialErrorUnknown
}
