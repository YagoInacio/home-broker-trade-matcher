package transformer

import (
	"github.com/yagoinacio/home-broker-trade-matcher/internal/market/dtos"
	"github.com/yagoinacio/home-broker-trade-matcher/internal/market/entities"
)

func TransformInput(input dtos.TradeInput) *entities.Order {
	asset := entities.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entities.NewInvestor(input.InvestorID)
	order := entities.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		assetPosition := entities.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}

	return order
}

func TransformOutput(order *entities.Order) *dtos.OrderOutput {
	output := &dtos.OrderOutput{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  order.OrderType,
		Status:     order.Status,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutput []*dtos.TransactionOutput
	for _, t := range order.Transactions {
		transactionOutput := &dtos.TransactionOutput{
			TransactionID: t.ID,
			BuyerId:       t.BuyingOrder.ID,
			SellerId:      t.SellingOrder.ID,
			AssetID:       t.SellingOrder.Asset.ID,
			Price:         t.Price,
			Shares:        t.SellingOrder.Shares - t.SellingOrder.PendingShares,
		}

		transactionsOutput = append(transactionsOutput, transactionOutput)
	}

	output.TransactionOutput = transactionsOutput

	return output
}
