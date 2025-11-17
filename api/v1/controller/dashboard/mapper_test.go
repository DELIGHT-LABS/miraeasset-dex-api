package dashboard

import (
	"testing"
	"time"

	ds "github.com/dezswap/dezswap-api/api/v1/service/dashboard"
	"github.com/stretchr/testify/assert"
)

func TestMapper_TxsToRes(t *testing.T) {
	m := &mapper{}

	txs := ds.Txs{
		{
			Action:     string(ds.TX_TYPE_SWAP),
			Hash:       "0xabc",
			Address:    "sender1",
			Sender:     "sender1",
			Timestamp:  time.Now(),
			TotalValue: "1000",

			Asset0:       "axpla",
			Asset0Amount: "-100.01",
			Asset0Symbol: "XPLA",
			Asset0Name:   "XPLA",

			Asset1:       "xpla1efgh",
			Asset1Amount: "20.01",
			Asset1Symbol: "WON",
			Asset1Name:   "원",
		},
		{
			Action:     string(ds.TX_TYPE_PROVIDE),
			Hash:       "0xdef",
			Address:    "sender2",
			Sender:     "sender2",
			Timestamp:  time.Now(),
			TotalValue: "500",

			Asset0:       "xpla1ijkl",
			Asset0Amount: "1.0",
			Asset0Symbol: "SAMSUNG",
			Asset0Name:   "삼성전자",

			Asset1:       "xpla1abcd",
			Asset1Amount: "45.0",
			Asset1Symbol: "ETH",
			Asset1Name:   "이더리움",
		},
	}

	res := m.txsToRes(txs)

	// --- CASE 1: SWAP transaction ---
	swap := res[0]
	assert.Equal(t, "xpla1efgh", swap.Asset0)
	assert.Equal(t, "axpla", swap.Asset1)
	assert.Equal(t, "100.01", swap.Asset1Amount)
	assert.Equal(t, "Swap WON for XPLA", swap.ActionDisplay)
	assert.Equal(t, "원화-XPLA 자산 전환", swap.ActionDisplayKo)

	// --- CASE 2: PROVIDE (Add liquidity) ---
	provide := res[1]
	assert.Equal(t, "add", provide.Action)
	assert.Equal(t, "xpla1ijkl", provide.Asset0)
	assert.Equal(t, "xpla1abcd", provide.Asset1)
	assert.Equal(t, "Add SAMSUNG and ETH", provide.ActionDisplay)
	assert.Equal(t, "삼성전자-이더리움 유동성 공급", provide.ActionDisplayKo)
}
