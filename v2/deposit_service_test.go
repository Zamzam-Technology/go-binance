package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type depositServiceTestSuite struct {
	baseTestSuite
}

func TestDepositService(t *testing.T) {
	suite.Run(t, new(depositServiceTestSuite))
}

func (s *depositServiceTestSuite) TestListDeposits() {
	data := []byte(`
	[
    {
        "amount":"0.00999800",
        "coin":"PAXG",
        "network":"ETH",
        "status":1,
        "address":"0x788cabe9236ce061e5a892e1a59395a81fc8d62c",
        "addressTag":"",
        "txId":"0xaad4654a3234aa6118af9b4b335f5ae81c360b2394721c019b5d1e75328b09f3",
        "insertTime":1599621997000,
        "transferType":0,
        "confirmTimes":"12/12"
    },
    {
        "amount":"0.50000000",
        "coin":"IOTA",
        "network":"IOTA",
        "status":1,
        "address":"SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
        "addressTag":"342341222",
        "txId":"ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
        "insertTime":1599620082000,
        "transferType":0,
        "confirmTimes":"1/1"
    }
]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":      "BTC",
			"status":    1,
			"startTime": 1508198532000,
			"endTime":   1508198532001,
		})
		s.assertRequestEqual(e, r)
	})

	deposits, err := s.client.NewListDepositsService().
		Coin("BTC").
		Status(1).
		StartTime(1508198532000).
		EndTime(1508198532001).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	r.Len(deposits, 2)
	s.assertDepositEqual(&Deposit{
		InsertTime: 1599621997000,
		Amount:     "0.00999800",
		Coin:       "PAXG",
		Address:    "0x788cabe9236ce061e5a892e1a59395a81fc8d62c",
		AddressTag: "",
		TxID:       "0xaad4654a3234aa6118af9b4b335f5ae81c360b2394721c019b5d1e75328b09f3",
		Status:     1,
	}, deposits[0])
	s.assertDepositEqual(&Deposit{
		InsertTime: 1599620082000,
		Amount:     "0.50000000",
		Coin:       "IOTA",
		Address:    "SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
		AddressTag: "342341222",
		TxID:       "ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
		Status:     1,
	}, deposits[1])
}

func (s *depositServiceTestSuite) assertDepositEqual(e, a *Deposit) {
	r := s.r()
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TxID, a.TxID, "TxID")
}

func (s *depositServiceTestSuite) TestGetDepositAddress() {
	data := []byte(`
	{
		"address": "0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000",
		"success": true,
		"tag": "1231212",
		"coin": "ETH",
		"url": "https://etherscan.io/address/0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "ETH"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin": asset,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetDepositAddressService().
		Coin(asset).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000", res.Address)
	r.Equal("1231212", res.Tag)
	r.Equal("ETH", res.Coin)
	r.Equal("https://etherscan.io/address/0xbf1f86b3c8ff4f8cbfc195e9713b6f0000000000", res.URL)
}
