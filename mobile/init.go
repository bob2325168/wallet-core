package mobile

import (
	"strings"

	"github.com/lomocoin/HDWallet-Core/core"
	"github.com/lomocoin/HDWallet-Core/core/bch"
	"github.com/lomocoin/HDWallet-Core/core/btc"
	"github.com/lomocoin/HDWallet-Core/core/eth"
	"github.com/lomocoin/HDWallet-Core/core/lmc"
	"github.com/lomocoin/HDWallet-Core/core/mgd"
	"github.com/lomocoin/HDWallet-Core/core/nxt"
	"github.com/lomocoin/HDWallet-Core/core/omni"
	"github.com/lomocoin/HDWallet-Core/core/trx"
	"github.com/lomocoin/HDWallet-Core/core/wcg"
	"github.com/lomocoin/HDWallet-Core/core/xrp"
	"github.com/pkg/errors"
)

func (c Wallet) initCoin(symbol string) (coin core.Coin, err error) {
	if len(c.seed) == 0 {
		err = errors.New("seed is empty")
		return
	}
	switch symbol {
	case "BTC":
		coin, err = btc.New(c.seed, c.testNet)
	case "USDT(Omni)":
		coin, err = omni.New(c.seed, c.testNet)
	case "BCH":
		coin, err = bch.New(c.seed, c.testNet)
	case "MGD":
		coin, err = mgd.New(c.seed)
	case "LMC":
		coin, err = lmc.New(c.seed)
	case "ETH", "XT", "THM", "ALI", "RED", "USO", "BTK", "EGT", "HOTC(HOTCOIN)":
		coin, err = eth.New(c.seed)
	case "XRP":
		coin, err = xrp.New(c.seed)
	case "TRX", "BTT":
		coin, err = trx.New(c.seed)
	case "WCG", "USDTK", "MTR", "DRT", "MAT", "WOS", "EQT", "ENX", "NRT", "CTM":
		coin, err = wcg.New(c.seed)
	case "NXT", "RMB":
		coin, err = nxt.New(c.seed)
	default:
		err = errors.Errorf("no entry for coin (%s) was found.", symbol)
	}
	return
}

// GetAvailableCoinList 获取支持币种列表 (case sensitive)
// return : "BTC LMC ETH WCG"
func GetAvailableCoinList() string {
	availableCoin := []string{
		// BTC series
		"BTC",
		"BCH",
		"MGD",
		"LMC",

		// OMNI series
		"USDT(Omni)",

		// ETH series
		"ETH",
		// ERC20 series
		"XT",
		"THM",
		"ALI",
		"RED",
		"USO",
		"BTK",
		"EGT",
		"HOTC(HOTCOIN)",

		// ripple
		"XRP",

		// tron
		"TRX",
		// TRC10
		"BTT",

		// WCG series
		"WCG",
		"USDTK",
		"MTR",
		"DRT",
		"MAT",
		"WOS",
		"EQT",
		"ENX",
		"NRT",
		"CTM",

		// NXT series
		"NXT",
		"RMB",
	}
	return strings.Join(availableCoin, " ")
}