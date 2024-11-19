package config

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	// ChainConfigUrl = "https://raw.githubusercontent.com/FastLane-Labs/atlas-config/refs/heads/main/configs/chain-config.json"
	ChainConfigUrl = "https://raw.githubusercontent.com/FastLane-Labs/atlas-config/refs/heads/atlas-multi-version/configs/chain-config.json"
)

func downloadChainConfig() (map[uint64]map[string]*ChainConfig, error) {
	resp, err := http.Get(ChainConfigUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	remoteConfig := make(map[uint64]map[string]*ChainConfig)

	err = json.Unmarshal(body, &remoteConfig)
	if err != nil {
		return nil, err
	}

	return remoteConfig, nil
}
