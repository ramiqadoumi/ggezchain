package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/x/trade/testutil"
	"github.com/ramiqadoumi/ggezchain/x/trade/types"
)

func (suite *IntegrationTestSuite) TestCreateTrade() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	createResponse, err := suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})
	suite.Nil(err)
	suite.EqualValues(types.MsgCreateTradeResponse{
		TradeIndex: 1,
		Status:     types.Pending,
	}, *createResponse)
}

func (suite *IntegrationTestSuite) TestIfTradeSaved() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	keeper := suite.app.TradeKeeper

	suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})
	tradeIndex, found := keeper.GetTradeIndex(suite.ctx)
	suite.True(found)
	suite.EqualValues(types.TradeIndex{
		NextId: 2,
	}, tradeIndex)
	trade, found := keeper.GetStoredTrade(suite.ctx, 1)
	suite.True(found)
	suite.EqualValues(types.StoredTrade{
		TradeIndex:      1,
		Status:          types.Pending,
		CreateDate:      trade.CreateDate,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		Maker:           testutil.Mutaz,
		Checker:         "",
		ProcessDate:     trade.CreateDate,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
		Result:          types.ErrTradeCreatedSuccessfully.Error(),
	}, trade)
}

func (suite *IntegrationTestSuite) TestIfTempTradeSaved() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	keeper := suite.app.TradeKeeper

	suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})
	tradeIndex, found := keeper.GetTradeIndex(suite.ctx)
	suite.True(found)
	suite.EqualValues(types.TradeIndex{
		NextId: 2,
	}, tradeIndex)
	tempTrade, found := keeper.GetStoredTempTrade(suite.ctx, 1)
	suite.True(found)
	suite.EqualValues(types.StoredTempTrade{
		TradeIndex:     1,
		CreateDate:     tempTrade.CreateDate,
		TempTradeIndex: 1,
	}, tempTrade)
}

func (suite *IntegrationTestSuite) TestGetAllStoredTrade() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	keeper := suite.app.TradeKeeper

	suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})
	tradeIndex, found := keeper.GetTradeIndex(suite.ctx)
	suite.True(found)
	suite.EqualValues(types.TradeIndex{
		NextId: 2,
	}, tradeIndex)
	allTrades := keeper.GetAllStoredTrade(suite.ctx)
	suite.EqualValues(types.StoredTrade{
		TradeIndex:      1,
		Status:          types.Pending,
		CreateDate:      allTrades[0].CreateDate,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		Maker:           testutil.Mutaz,
		Checker:         "",
		ProcessDate:     allTrades[0].CreateDate,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
		Result:          types.ErrTradeCreatedSuccessfully.Error(),
	}, allTrades[0])
}

func (suite *IntegrationTestSuite) TestGetAllStoredTempTrade() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	keeper := suite.app.TradeKeeper

	suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})
	tradeIndex, found := keeper.GetTradeIndex(suite.ctx)
	suite.True(found)
	suite.EqualValues(types.TradeIndex{
		NextId: 2,
	}, tradeIndex)
	allTempTrades := keeper.GetAllStoredTempTrade(suite.ctx)
	suite.EqualValues(types.StoredTempTrade{
		TradeIndex:     1,
		CreateDate:     allTempTrades[0].CreateDate,
		TempTradeIndex: 1,
	}, allTempTrades[0])
}

func (suite *IntegrationTestSuite) TestCreateTradeWithInvalidMakerPermission() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	createResponse, err := suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mohd,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})

	suite.Nil(createResponse)
	suite.ErrorIs(err, types.ErrInvalidMakerPermission)
}

func (suite *IntegrationTestSuite) TestCreateTradeWithInvalidTradeData() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	createResponse, err := suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":0,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})

	suite.Nil(createResponse)
	suite.ErrorIs(err, types.ErrTradeDataRequestID)
}

func (suite *IntegrationTestSuite) TestCreate2Trades() {
	suite.SetupTestForCreateTrade()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	keeper := suite.app.TradeKeeper

	suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})
	trade, found := keeper.GetStoredTrade(suite.ctx, 1)
	suite.True(found)
	suite.EqualValues(types.StoredTrade{
		TradeIndex:      1,
		Status:          types.Pending,
		CreateDate:      trade.CreateDate,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		Maker:           testutil.Mutaz,
		Checker:         "",
		ProcessDate:     trade.CreateDate,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
		Result:          types.ErrTradeCreatedSuccessfully.Error(),
	}, trade)

	tempTrade, found := keeper.GetStoredTempTrade(suite.ctx, 1)
	suite.True(found)
	suite.EqualValues(types.StoredTempTrade{
		TradeIndex:     1,
		CreateDate:     tempTrade.CreateDate,
		TempTradeIndex: 1,
	}, tempTrade)

	suite.msgServer.CreateTrade(goCtx, &types.MsgCreateTrade{
		Creator:         testutil.Mutaz,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
	})

	trade, found = keeper.GetStoredTrade(suite.ctx, 2)
	suite.True(found)
	suite.EqualValues(types.StoredTrade{
		TradeIndex:      2,
		Status:          types.Pending,
		CreateDate:      trade.CreateDate,
		TradeType:       types.Buy,
		Coin:            types.DefaultCoinDenom,
		Price:           "0.001",
		Quantity:        "100000",
		ReceiverAddress: testutil.Mutaz,
		Maker:           testutil.Mutaz,
		Checker:         "",
		ProcessDate:     trade.CreateDate,
		TradeData:       "{\"TradeData\":{\"tradeRequestID\":123,\"assetHolderID\":456,\"assetID\":789,\"tradeType\":\"Buy\",\"tradeValue\":100.50,\"currency\":\"USD\",\"exchange\":\"NYSE\",\"fundName\":\"TechFund\",\"issuer\":\"CompanyA\",\"noShares\":\"1000\",\"price\":\"50.25\",\"quantity\":\"10\",\"segment\":\"Technology\",\"sharePrice\":\"49.50\",\"ticker\":\"TECH\",\"tradeFee\":\"5.00\",\"tradeNetPrice\":\"500.00\",\"tradeNetValue\":\"495.00\"},\"Brokerage\":{\"name\":\"XYZBrokerage\",\"type\":\"Online\",\"country\":\"USA\"}}",
		Result:          types.ErrTradeCreatedSuccessfully.Error(),
	}, trade)

	tempTrade, found = keeper.GetStoredTempTrade(suite.ctx, 2)
	suite.True(found)
	suite.EqualValues(types.StoredTempTrade{
		TradeIndex:     2,
		CreateDate:     tempTrade.CreateDate,
		TempTradeIndex: 2,
	}, tempTrade)

	// check get all trades and temp trades and next trade index
	tradeIndex, found := keeper.GetTradeIndex(suite.ctx)
	suite.True(found)
	suite.EqualValues(types.TradeIndex{
		NextId: 3,
	}, tradeIndex)
	AllTrades := keeper.GetAllStoredTrade(suite.ctx)
	suite.EqualValues(len(AllTrades), 2)

	AllTempTrades := keeper.GetAllStoredTempTrade(suite.ctx)
	suite.EqualValues(len(AllTempTrades), 2)
}
