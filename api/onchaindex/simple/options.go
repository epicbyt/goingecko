package simple

import (
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// priceOption is an interface that extends internal.Option with a marker method
// to ensure type safety for price-specific options
type priceOption interface {
	internal.Option
	isPriceOption()
}

// WithIncludeMarketCapOption returns a priceOption that configures whether to include
// market cap data in the price response
func WithIncludeMarketCapOption(includeMarketCap bool) priceOption {
	return includeMarketCapOption{includeMarketCap}
}

// WithIncludeDayVolumeOption returns a priceOption that configures whether to include
// 24-hour volume data in the price response
func WithIncludeDayVolumeOption(includeDayVolume bool) priceOption {
	return includeDayVolumeOption{includeDayVolume}
}

// WithIncludeDayChangeOption returns a priceOption that configures whether to include
// 24-hour price change data in the price response
func WithIncludeDayChangeOption(includeDayChange bool) priceOption {
	return includeDayPriceChangeOption{includeDayChange}
}

type (
	includeMarketCapOption       struct{ include bool }
	includeDayVolumeOption       struct{ include bool }
	includeDayPriceChangeOption  struct{ include bool }
	includeLastTimeUpdatedOption struct{ include bool }
	precisionOption              struct{ precision string }
)

func (o includeMarketCapOption) Apply(v *url.Values) {
	v.Set("include_market_cap", strconv.FormatBool(o.include))
}

func (o includeDayVolumeOption) Apply(v *url.Values) {
	v.Set("include_24hr_vol", strconv.FormatBool(o.include))
}

func (o includeDayPriceChangeOption) Apply(v *url.Values) {
	v.Set("include_24hr_price_change", strconv.FormatBool(o.include))
}

func (o includeMarketCapOption) isPriceOption()       {}
func (o includeDayVolumeOption) isPriceOption()       {}
func (o includeDayPriceChangeOption) isPriceOption()  {}
func (o includeLastTimeUpdatedOption) isPriceOption() {}
func (o precisionOption) isPriceOption()              {}
