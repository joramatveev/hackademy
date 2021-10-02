package orderbook

import (
	"sort"
)

type OrderBook struct {
	Ask []*Order
	Bid []*Order
}

func New() *OrderBook {
	return &OrderBook{}
}

func (orderBook *OrderBook) Match(order *Order) ([]*Trade, *Order) {
	return orderBook.Trade(order)
}

func (orderBook *OrderBook) Trade(order *Order) ([]*Trade, *Order) {
	var tradesArray []*Trade
	var propsArr *[]*Order
	switch order.Side {
	case SideAsk:
		propsArr = &orderBook.Bid
	case SideBid:
		propsArr = &orderBook.Ask
	}
	for order.Volume != 0 {
		if len(*propsArr) == 0 {
			orderBook.OrderCancel(order)
			break
		}
		mostProp := (*propsArr)[0]
		if propApproved(order, mostProp) {
			var tradeValue uint64
			if order.Volume < mostProp.Volume {
				tradeValue = order.Volume
			} else {
				tradeValue = mostProp.Volume
			}
			trade := Trade{mostProp, order, tradeValue, mostProp.Price}
			order.Volume -= tradeValue
			mostProp.Volume -= tradeValue
			tradesArray = append(tradesArray, &trade)
			if mostProp.Volume == 0 {
				*propsArr = (*propsArr)[1:]
			}
		} else {
			orderBook.OrderCancel(order)
			break
		}
	}
	if order.Volume > 0 && order.Kind == KindMarket {
		return tradesArray, order
	}
	return tradesArray, nil
}

func propApproved(order *Order, prop *Order) bool {
	switch order.Kind {
	case KindMarket:
		return true
	case KindLimit:
		switch order.Side {
		case SideAsk:
			return order.Price <= prop.Price
		case SideBid:
			return order.Price >= prop.Price
		}
	}
	return false
}

func (orderBook *OrderBook) OrderCancel(order *Order) {
	switch order.Side {
	case SideAsk:
		orderBook.Ask = append(orderBook.Ask, order)
		sort.Slice(orderBook.Ask, func(i, j int) bool {
			return orderBook.Ask[i].Price < orderBook.Ask[j].Price
		})
	case SideBid:
		orderBook.Bid = append(orderBook.Bid, order)
		sort.Slice(orderBook.Bid, func(i, j int) bool {
			return orderBook.Bid[i].Price > orderBook.Bid[j].Price
		})
	}
}
