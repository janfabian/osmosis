package types

// GenesisState defines the arbot module's genesis state.
type GenesisState struct {
	// Add any fields your module needs for genesis.
	// Since your module currently doesn't have state, this can be empty.
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() *GenesisState {
	return &GenesisState{
		// Initialize fields if necessary
	}
}

// DefaultGenesis returns the default genesis state for the arbot module
func DefaultGenesis() *GenesisState {
	return NewGenesisState()
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	// Perform validation of genesis state if needed.
	// For example, check that parameters are within acceptable ranges.
	// Since there's nothing to validate here, return nil.
	return nil
}
