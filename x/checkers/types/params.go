package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxTurnDuration = []byte("MaxTurnDuration")
	// TODO: Determine the default value
	DefaultMaxTurnDuration string = "max_turn_duration"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxTurnDuration string,
) Params {
	return Params{
		MaxTurnDuration: maxTurnDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxTurnDuration,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxTurnDuration, &p.MaxTurnDuration, validateMaxTurnDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxTurnDuration(p.MaxTurnDuration); err != nil {
		return err
	}

	return nil
}

// validateMaxTurnDuration validates the MaxTurnDuration param
func validateMaxTurnDuration(v interface{}) error {
	maxTurnDuration, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxTurnDuration

	return nil
}
