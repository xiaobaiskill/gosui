package transactions

import (
	"github.com/W3Tools/gosui/client"
	"github.com/W3Tools/gosui/types"
)

func getNormalizedMoveFunctionFromCache(suiClient *client.SuiClient, pkg, mod, fn string) (*types.SuiMoveNormalizedFunction, error) {
	entry := cache.GetMoveFunctionDefinition(pkg, mod, fn)
	if entry != nil {
		return entry.Normalized, nil
	}

	result, err := suiClient.GetNormalizedMoveFunction(types.GetNormalizedMoveFunctionParams{Package: pkg, Module: mod, Function: fn})
	if err != nil {
		return nil, err
	}

	cache.AddMoveFunctionDefinition(&MoveFunctionCacheEntry{Package: pkg, Module: mod, Function: fn, Normalized: result})
	return result, nil
}
