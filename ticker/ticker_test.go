package ticker

import (
	"testing"
)

// Test_Ticker_Interface_Ticker ensures that the real ticker implementation
// complies with the ticker interface. This test already fails at compile time
// if Ticker does not implement Interface.
func Test_Ticker_Interface_Ticker(t *testing.T) {
	var _ Interface = &Ticker{}
}
