package trade

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/ramiqadoumi/ggezchain/v2/testutil/sample"
	tradesimulation "github.com/ramiqadoumi/ggezchain/v2/x/trade/simulation"
	"github.com/ramiqadoumi/ggezchain/v2/x/trade/types"
)

// avoid unused import issue
var (
	_ = tradesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTrade = "op_weight_msg_create_trade"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTrade int = 100

	opWeightMsgProcessTrade = "op_weight_msg_process_trade"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProcessTrade int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tradeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TradeIndex: types.TradeIndex{
			NextId: uint64(simState.Rand.Intn(99) + 1),
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tradeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTrade int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTrade, &weightMsgCreateTrade, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTrade = defaultWeightMsgCreateTrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTrade,
		tradesimulation.SimulateMsgCreateTrade(am.accountKeeper, am.bankKeeper, am.aclKeeper, am.keeper),
	))

	var weightMsgProcessTrade int
	simState.AppParams.GetOrGenerate(opWeightMsgProcessTrade, &weightMsgProcessTrade, nil,
		func(_ *rand.Rand) {
			weightMsgProcessTrade = defaultWeightMsgProcessTrade
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProcessTrade,
		tradesimulation.SimulateMsgProcessTrade(am.accountKeeper, am.bankKeeper, am.aclKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTrade,
			defaultWeightMsgCreateTrade,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tradesimulation.SimulateMsgCreateTrade(am.accountKeeper, am.bankKeeper, am.aclKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProcessTrade,
			defaultWeightMsgProcessTrade,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tradesimulation.SimulateMsgProcessTrade(am.accountKeeper, am.bankKeeper, am.aclKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
