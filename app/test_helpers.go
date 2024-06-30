package app

import (
	"encoding/json"
	// "fmt"
	"testing"
	"time"
	// storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"github.com/GGEZLabs/ggezchain/x/trade/types"
	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"
	// codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	// cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	// "github.com/GGEZLabs/ggezchain/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
)

// DefaultConsensusParams defines the default Tendermint consensus params used in
// App testing.
var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

// func setup(withGenesis bool) (*App, GenesisState) {
// 	db := dbm.NewMemDB()
// 	// encCdc := testutil.MakeEncodingConfig()
// 	appOptions := make(simtestutil.AppOptionsMap, 0)
// 	app,_ := New(log.NewNopLogger(), db, nil, true, appOptions)
// 	// if withGenesis {
// 	// 	return app, NewDefaultGenesisState(encCdc.Marshaler)
// 	// }
// 	return app, GenesisState{}
// }

func setup(withGenesis bool, invCheckPeriod uint) (*App, GenesisState) {
	db := dbm.NewMemDB()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = invCheckPeriod

	app ,_:= New(log.NewNopLogger(), db, nil, true, appOptions)
	if withGenesis {
		return app, app.DefaultGenesis()
	}
	return app, GenesisState{}
}

// Setup initializes a new App. A Nop logger is set in App.
// func Setup(isCheckTx bool) *App {
//     sdk.GetConfig().SetBech32PrefixForAccount("ggez", "ggez")
//     sdk.GetConfig().SetBech32PrefixForValidator("ggezvaloper", "ggezvaloper")
//     sdk.GetConfig().SetBech32PrefixForConsensusNode("ggezvalcons", "ggezvalcons")
// 	app, genesisState := setup(!isCheckTx)
// 	if !isCheckTx {
// 		// init chain must be called to stop deliverState from being nil
// 		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Initialize the chain
// 		app.InitChain(
// 			&abci.RequestInitChain{
// 				Validators:      []abci.ValidatorUpdate{},
// 				ConsensusParams: DefaultConsensusParams,
// 				AppStateBytes:   stateBytes,
// 			},
// 		)
// 	}

// 	return app
// }

func Setup(t *testing.T, isCheckTx bool) *App {
	t.Helper()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)

	// create validator set with single validator
	validator := cmttypes.NewValidator(pubKey, 1)
	valSet := cmttypes.NewValidatorSet([]*cmttypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100000000000000))),
	}

	app := SetupWithGenesisValSet(t, valSet, []authtypes.GenesisAccount{acc}, balance)

	return app
}


// SetupWithGenesisValSet initializes a new App with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit (10^6) in the default token of the simapp from first genesis
// account. A Nop logger is set in App.
// func SetupWithGenesisValSet(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *App {
// 	app, genesisState := setup(true)
// 	// set genesis accounts
// 	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
// 	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

// 	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
// 	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

// 	bondAmt := sdkmath.NewInt(1000000)

// 	for _, val := range valSet.Validators {
// 		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
// 		require.NoError(t, err)
// 		pkAny, err := codectypes.NewAnyWithValue(pk)
// 		require.NoError(t, err)
// 		validator := stakingtypes.Validator{
// 			OperatorAddress:   sdk.ValAddress(val.Address).String(),
// 			ConsensusPubkey:   pkAny,
// 			Jailed:            false,
// 			Status:            stakingtypes.Bonded,
// 			Tokens:            bondAmt,
// 			DelegatorShares:   sdkmath.LegacyOneDec(),
// 			Description:       stakingtypes.Description{},
// 			UnbondingHeight:   int64(0),
// 			UnbondingTime:     time.Unix(0, 0).UTC(),
// 			Commission:        stakingtypes.NewCommission(sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec()),
// 			MinSelfDelegation: sdkmath.ZeroInt(),
// 		}
// 		validators = append(validators, validator)
// 		// Convert the validator address to a Bech32 string
// 		valAddr := sdk.ValAddress(val.Address)
// 		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress().String(), valAddr.String(), sdkmath.LegacyOneDec()))
// 	}


// 	// set validators and delegations
// 	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)

// 	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

// 	totalSupply := sdk.NewCoins()
// 	for _, b := range balances {
// 		// add genesis acc tokens to total supply
// 		totalSupply = totalSupply.Add(b.Coins...)
// 	}

// 	// add bonded amount to bonded pool module account
// 	balances = append(balances, banktypes.Balance{
// 		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
// 		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt)},
// 	})

// 	// add bonded amount to total supply
// 	totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))

// 	// update total supply
// 	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
// 	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

// 	// trade genesis
// 	tradeGenesis := types.DefaultGenesis()
// 	genesisState[types.ModuleName] = app.AppCodec().MustMarshalJSON(tradeGenesis)

// 	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
// 	require.NoError(t, err)


// 	// init chain will set the validator set and initialize the genesis accounts
// 	app.InitChain(&abci.RequestInitChain{
// 		Validators:      []abci.ValidatorUpdate{},  // Ensure validators are included
// 		ConsensusParams: DefaultConsensusParams,
// 		AppStateBytes:   stateBytes,
// 	})

// 	// commit genesis changes
// 	app.Commit()
// 	app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: app.LastBlockHeight() + 1})
// 	ctx := app.BaseApp.NewContext(false)

// 	// Retrieve and print validators
// 	store := ctx.KVStore(app.GetKey(stakingtypes.StoreKey))
// 	iterator := storetypes.KVStorePrefixIterator(store, stakingtypes.ValidatorsKey)
// 	defer iterator.Close()
// 	for ; iterator.Valid(); iterator.Next() {
// 		var validator stakingtypes.Validator
// 		app.AppCodec().MustUnmarshal(iterator.Value(), &validator)
// 		fmt.Printf("Validator: %v\n", validator)
// 	}

// 	// Retrieve and print delegations
// 	iterator = storetypes.KVStorePrefixIterator(store, stakingtypes.DelegationKey)
// 	defer iterator.Close()
// 	for ; iterator.Valid(); iterator.Next() {
// 		var delegation stakingtypes.Delegation
// 		app.AppCodec().MustUnmarshal(iterator.Value(), &delegation)
// 		fmt.Printf("Delegation: %v\n", delegation)
// 	}

// 	// Retrieve and print balances
// 	bankStore := ctx.KVStore(app.GetKey(banktypes.StoreKey))
// 	iterator = storetypes.KVStorePrefixIterator(bankStore, banktypes.BalancesPrefix)
// 	defer iterator.Close()
// 	for ; iterator.Valid(); iterator.Next() {
// 		var balance banktypes.Balance
// 		app.AppCodec().MustUnmarshal(iterator.Value(), &balance)
// 		fmt.Printf("Balance: %v\n", balance)
// 	}
// 	return app
// }

func SetupWithGenesisValSet(t *testing.T, valSet *cmttypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *App {
	t.Helper()

	app, genesisState := setup(true, 5)
	genesisState, err := simtestutil.GenesisStateWithValSet(app.AppCodec(), genesisState, valSet, genAccs, balances...)
	require.NoError(t, err)
	tradeGenesis := types.DefaultGenesis()
 	genesisState[types.ModuleName] = app.AppCodec().MustMarshalJSON(tradeGenesis)
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	_, err = app.InitChain(&abci.RequestInitChain{
		Validators:      []abci.ValidatorUpdate{},
		ConsensusParams: simtestutil.DefaultConsensusParams,
		AppStateBytes:   stateBytes,
	},
	)
	require.NoError(t, err)

	require.NoError(t, err)
	_, err = app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height:             app.LastBlockHeight() + 1,
		Hash:               app.LastCommitID().Hash,
		NextValidatorsHash: valSet.Hash(),
	})
	require.NoError(t, err)

	return app
}

// SetupWithGenesisAccounts initializes a new App with the provided genesis
// accounts and possible balances.
func SetupWithGenesisAccounts(genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *App {
	app, genesisState := setup(true,5)
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		totalSupply = totalSupply.Add(b.Coins...)
	}

	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{},[]banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	app.InitChain(
		&abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	app.Commit()
	app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: app.LastBlockHeight() + 1})

	return app
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}
