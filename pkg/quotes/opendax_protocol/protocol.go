// Events/metods names for open-finance.org protocol implementation client and server.
package opendax_protocol

import "fmt"

const (
	EventBalanceUpdate      = "bu"
	EventWithdrawalUpdate   = "wu"
	EventDepositUpdate      = "du"
	EventOrderCreate        = "on"
	EventOrderCancel        = "oc"
	EventOrderUpdate        = "ou"
	EventOrderReject        = "or"
	EventPrivateTrade       = "tr"
	EventTrade              = "trade"
	EventTickers            = "tickers"
	EventMarkets            = "markets"
	EventKLine              = "kline-"
	EventOrderBookIncrement = "obi"
	EventOrderBookSnapshot  = "obs"
	EventRawBookIncrement   = "rbi"
	EventRawBookSnapshot    = "rbs"
	EventSystem             = "sys"
	EventConfig             = "config"

	MethodAuth               = "authenticate"
	MethodSubscribe          = "subscribe"
	MethodListOrders         = "list_orders"
	MethodListOrdersByUUIDS  = "list_orders_by_uuids"
	MethodListOrderTrades    = "list_order_trades"
	MethodListChannelTrades  = "list_channel_trades"
	MethodCreateOrder        = "create_order"
	MethodCreateOrderBulk    = "create_bulk"
	MethodCancelOrder        = "cancel_order"
	MethodCancelOrdersAll    = "cancel_all"
	MethodGetKlines          = "klines"
	MethodGetMarkets         = "get_markets"
	MethodGetSymbols         = "get_symbols"
	MethodGetTokens          = "get_tokens"
	MethodGetTokenBySymbol   = "get_token_by_symbol_id"
	MethodGetTokenById       = "get_token_by_id"
	MethodGetNetworks        = "get_networks"
	MethodGetDeposits        = "get_deposits"
	MethodGetWithdrawals     = "get_withdrawals"
	MethodCheckSMTPConfig    = "check_smtp_config"
	MethodSendMetrics        = "send_metrics"
	MethodCreateDeposit      = "create_deposit"
	MethodCreateWithdrawal   = "create_withdrawal"
	MethodCancelDeposit      = "cancel_deposit"
	MethodCancelWithdrawal   = "cancel_withdrawal"
	MethodGetPublicConfig    = "get_public_config"
	MethodSetConfig          = "set_config"
	MethodGetConfig          = "get_config"
	MethodGetParticipantData = "get_participation_data"
	MethodAddAddress         = "add_address"
	MethodVerifyAddress      = "verify_address"
	MethodGetAddresses       = "get_addresses"

	MethodAdminCreateOrder = "admin_create_order"
	MethodAdminCancelOrder = "admin_cancel_order"
	MethodAdminDeposit     = "admin_deposit"
	MethodAdminWithdraw    = "admin_withdraw"

	OrderSideSell = "sell"
	OrderSideBuy  = "buy"

	OrderStatePending       = "p"
	OrderStateWait          = "w"
	OrderStateDone          = "d"
	OrderStateReject        = "r"
	OrderStateCancel        = "c"
	OrderStateTriggerWait   = "tw"
	OrderStateTriggerCancel = "tc"
	OrderStateTriggerReject = "tr"

	OrderTypeLimit      = "l"
	OrderTypeMarket     = "m"
	OrderTypePostOnly   = "p"
	OrderTypeFillOrKill = "f"
	OrderTypeStopLoss   = "sl"
	OrderTypeStopLimit  = "slm"
	OrderTypeTakeProfit = "tp"
	OrderTypeTakeLimit  = "tlm"

	TopicBalances      = "balances"
	TopicOrders        = "order"
	TopicTickers       = "tickers"
	TopicWithdrawals   = "withdrawals"
	TopicDeposits      = "deposits"
	TopicTrades        = "trades"
	TopicOrderbooksInc = "ob-inc"
	TopicKLines        = "kline-"

	ScopePrivate = "private"
	ScopePublic  = "public"
)

type eventDescription struct {
	description string
	bodyExample string
}

type topicDescription struct {
	name         string
	scope        string
	description  string
	reqExample   string
	eventExample string
}

var (
	events = map[string]eventDescription{
		EventBalanceUpdate:                    {"Balance Update", `[4,"bu",[["usd","0.9997","0.0003"]]]`},
		EventWithdrawalUpdate:                 {"Withdrawal Update", `[4,'wu",[{"ID":5,"UserID":"191d0965-fae4-4c2a-ab71-b433a7d226e6","SymbolID":"eth","State":"submitted","Amount":"0","TxId":null,"BlockNumber":null,"RID":"","EventID":null,"ExpiresAt":"0001-01-01T00:00:00Z","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}]]`},
		EventDepositUpdate:                    {"Deposit Update", `[4,"du",[{"ID":5,"UserID":"e8b5b5c6-7a28-44f5-9667-e6dc9c33aab4","SymbolID":"eth","State":"submitted","Amount":"0","TxId":null,"BlockNumber":null,"RID":"","EventID":null,"ExpiresAt":"0001-01-01T00:00:00Z","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}]]`},
		EventOrderCreate:                      {"Order Create", `[4,"on",["btcusd",1676,"d1b08171-5e72-467d-bf9e-8125afc954e2","buy","w","l","1","0","0.0001","0.0001","0",0,1654516051,"0.002","0.002","0"]]`},
		EventOrderCancel:                      {"Order Cancel", `[4,"oc",["btcusd",1675,"a4b35eb7-268f-4227-9a7a-14e07e02656f","buy","c","l","2","0","0.0001","0.0001","0",0,1654515716,"0.002","0.002","0"]]`},
		EventOrderUpdate:                      {"Order Update", `[4,"ou",["btcusd",1681,"790fc5ab-c344-4752-811b-30146d7425eb","sell","d","l","7","7","0","0.42","0.42",1,1654516448,"0.002","0.002","0"]]`},
		EventOrderReject:                      {"Order Reject", `[4,"or",["btcusd",0,"e65d22cd-fb5b-448b-a50a-35f800b6756c","sell","r","l","1000000000","0","100000000000000","100000000000000","0",0,1654521911,"0","0","0","(btc) balance in cache: 988, required to lock an order: 100000000000000, error: not enough funds: 988 btc"]]`},
		EventPrivateTrade:                     {"Private trade", `[4,"tr",["",45,"123.5","33.5","4445.33",0,"00000000-0000-0000-0000-000000000000","sell","sell","0","0","1654522267"]]`},
		EventTrade:                            {"Trade", `[3,"trade",["btcusd",0,"31365.35","0.00002802","0.878857107",1654517864,"sell","2"]]`},
		EventTickers:                          {"Tickers", `[3,"tickers",[["batusd",1654428483,"0.4110046233561333","0.41677723017833107","0.4159965789473684","0.41355000000000003","131491.12713629","54125.94662525234","0.41163193140135634","-0.5881247758236787"]]`},
		EventMarkets:                          {"Markets", `[3,"markets",[["avausd","spot","ava","usd","enabled",11,1,3,"0.001","0","0.1"],["balusd","spot","bal","usd","enabled",13,2,2,"0.01","0","0.01"],["batusd","spot","bat","usd","enabled",15,1,3,"0.001","0","0.1"]]]`},
		EventOrderBookIncrement:               {"Orderbook Increment", `[3,"obi",["btcusd",244098,[["32246.94","0.0001"]],[]]]`},
		EventOrderBookSnapshot:                {"Orderbook Snapshot", `[3,"obs",["btcusd",252012,[["31397.11","15.51796"],["31425.11","15.64092"],["32343.08","0.0001"]],[["31381.1","10.42636"],["31241.1","7.9297"]]]]`},
		EventRawBookIncrement:                 {"Rawbook Increment", `[3,"rbi",["btcusd",244098,[["32246.94","0.0001"]],[]]]`},
		EventRawBookSnapshot:                  {"Rawbook Snapshot", `[3,"rbs",["btcusd",252012,[["31397.11","15.51796"],["31425.11","15.64092"],["32343.08","0.0001"]],[["31381.1","10.42636"],["31241.1","7.9297"]]]]`},
		EventSystem:                           {"System", `[3,"sys",[{"hb":1654513650}]]`},
		EventConfig:                           {"Config", `[3,"config",[{"finex_custody_broker_public_key":"0x9792f5d8FCC0F7E40BC23FEB0666d5b67C3a1654","finex_version":"1.0.0"}]]`},
		"'market'." + EventKLine + "'period'": {"KLine", `[3,"btcusd.kline-15m",[1654514100,31411.907434247005,31457.383150554608,31318.440767719563,31388.286647856574,628.4550854400002]]`},
	}

	topics = map[string]topicDescription{
		TopicBalances:      {"balances", ScopePrivate, "Balances updates", "balances", getEventExample(EventBalanceUpdate)},
		TopicWithdrawals:   {"withdrawals", ScopePrivate, "Withdrawals updates", "withdrawals", getEventExample(EventWithdrawalUpdate)},
		TopicDeposits:      {"deposits", ScopePrivate, "Deposits updates", "deposits", getEventExample(EventDepositUpdate)},
		TopicTrades:        {"'market'.trades", ScopePublic, "Trades updates", "btcusd.trades", getEventExample(EventTrade)},
		TopicOrderbooksInc: {"'market'.ob-inc", ScopePublic, "Orderbook updates", "btcusd.ob-inc", getEventExample(EventOrderBookIncrement)},
		TopicOrders:        {"order", ScopePrivate, "Orders updates", "order", getEventExample(EventOrderCreate)},
		TopicTickers:       {"tickers", ScopePublic, "Tickers updates", "tickers", getEventExample(EventTickers)},
		TopicKLines:        {"'market'.kline-'period'", ScopePublic, "KLines updates", "btcusd.kline-15m", getEventExample("'market'" + EventKLine + "'period'")},
	}
)

func DescribeEvents() ([]string, [][]string) {
	keys := []string{"Event", "Description", "Example"}
	var rows [][]string
	for symbol, event := range events {
		rows = append(rows,
			[]string{
				symbol,
				event.description,
				fmt.Sprintf("`%s`", getEventExample(symbol)),
			})
	}
	return keys, rows
}

func DescribeTopics() ([]string, [][]string) {
	keys := []string{"Topic", "Scope", "Description", "Subscription", "Event example"}
	var rows [][]string
	for _, topic := range topics {
		rows = append(rows,
			[]string{
				topic.name,
				topic.scope,
				topic.description,
				fmt.Sprintf("`[1,1,\"subscribe\",[\"%s\",[\"%s\"]]]`", topic.scope, topic.reqExample),
				fmt.Sprintf("`%s`", topic.eventExample),
			})
	}
	return keys, rows
}

func getEventExample(symbol string) string {
	return "[3,\"" + symbol + "\", [" + events[symbol].bodyExample + "]]"
}
