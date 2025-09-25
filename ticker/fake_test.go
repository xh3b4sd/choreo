package ticker

import (
	"testing"
)

// Test_Ticker_Interface_Fake ensures that the fake ticker implementation
// complies with the ticker interface. This test already fails at compile time
// if Fake does not implement Interface.
func Test_Ticker_Interface_Fake(t *testing.T) {
	var _ Interface = Fake{}
}
