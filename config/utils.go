package config

import (
	"strings"
)

const (
	monadTestnetChainId = uint64(10143)
)

func IsMonad(chainId uint64) bool {
	return chainId == monadTestnetChainId
}

func ToMonadVersion(version *string) string {
	_version := GetVersion(version)

	if strings.HasSuffix(_version, "-monad") {
		return _version
	}

	return _version + "-monad"
}
