package types

import (
	"errors"
	"math/big"
	"math/rand/v2"
	"reflect"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
	"github.com/FastLane-Labs/atlas-sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	ErrInvalidAtlasAddress = errors.New("invalid atlas address (`to` parameter)")
)

type UserOperationsParams struct {
	From                 common.Address
	To                   common.Address // Mandatory (Atlas Address)
	Value                *big.Int
	Gas                  *big.Int
	MaxFeePerGas         *big.Int
	Nonce                *big.Int
	Deadline             *big.Int
	Dapp                 common.Address
	Control              common.Address
	CallConfig           uint32
	DappGasLimit         uint32   // From v1.5
	SolverGasLimit       uint32   // From v1.6
	BundlerSurchargeRate *big.Int // From v1.6
	SessionKey           common.Address
	Data                 []byte
	Signature            []byte
}

type UserOperation struct {
	chainId        uint64
	version        string
	versionAtLeast string

	from                 common.Address
	to                   common.Address
	value                *big.Int
	gas                  *big.Int
	maxFeePerGas         *big.Int
	nonce                *big.Int
	deadline             *big.Int
	dapp                 common.Address
	control              common.Address
	callConfig           uint32
	dappGasLimit         uint32   // From v1.5
	solverGasLimit       uint32   // From v1.6
	bundlerSurchargeRate *big.Int // From v1.6
	sessionKey           common.Address
	data                 []byte
	signature            []byte
}

func NewUserOperation(chainId uint64, params UserOperationsParams) (*UserOperation, error) {
	if params.To == (common.Address{}) {
		return nil, ErrInvalidAtlasAddress
	}

	version, err := config.GetVersionFromAtlasAddress(chainId, params.To)
	if err != nil {
		return nil, err
	}

	versionAtLeast, err := utils.VersionAtLeast(&version)
	if err != nil {
		return nil, err
	}

	return &UserOperation{
		chainId:              chainId,
		version:              version,
		versionAtLeast:       versionAtLeast,
		from:                 params.From,
		to:                   params.To,
		value:                params.Value,
		gas:                  params.Gas,
		maxFeePerGas:         params.MaxFeePerGas,
		nonce:                params.Nonce,
		deadline:             params.Deadline,
		dapp:                 params.Dapp,
		control:              params.Control,
		callConfig:           params.CallConfig,
		dappGasLimit:         params.DappGasLimit,
		solverGasLimit:       params.SolverGasLimit,
		bundlerSurchargeRate: params.BundlerSurchargeRate,
		sessionKey:           params.SessionKey,
		data:                 params.Data,
		signature:            params.Signature,
	}, nil
}

func (u *UserOperation) GetFrom() common.Address {
	return u.from
}

func (u *UserOperation) SetFrom(from common.Address) {
	u.from = from
}

func (u *UserOperation) GetTo() common.Address {
	return u.to
}

func (u *UserOperation) GetValue() *big.Int {
	if u.value == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(u.value)
}

func (u *UserOperation) SetValue(value *big.Int) {
	if value == nil {
		u.value = nil
	}
	u.value = new(big.Int).Set(value)
}

func (u *UserOperation) GetGas() *big.Int {
	if u.gas == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(u.gas)
}

func (u *UserOperation) SetGas(gas *big.Int) {
	if gas == nil {
		u.gas = nil
	}
	u.gas = new(big.Int).Set(gas)
}

func (u *UserOperation) GetMaxFeePerGas() *big.Int {
	if u.maxFeePerGas == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(u.maxFeePerGas)
}

func (u *UserOperation) SetMaxFeePerGas(maxFeePerGas *big.Int) {
	if maxFeePerGas == nil {
		u.maxFeePerGas = nil
	}
	u.maxFeePerGas = new(big.Int).Set(maxFeePerGas)
}

func (u *UserOperation) GetNonce() *big.Int {
	if u.nonce == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(u.nonce)
}

func (u *UserOperation) SetNonce(nonce *big.Int) {
	if nonce == nil {
		u.nonce = nil
	}
	u.nonce = new(big.Int).Set(nonce)
}

func (u *UserOperation) GetDeadline() *big.Int {
	if u.deadline == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(u.deadline)
}

func (u *UserOperation) SetDeadline(deadline *big.Int) {
	if deadline == nil {
		u.deadline = nil
	}
	u.deadline = new(big.Int).Set(deadline)
}

func (u *UserOperation) GetDapp() common.Address {
	return u.dapp
}

func (u *UserOperation) SetDapp(dapp common.Address) {
	u.dapp = dapp
}

func (u *UserOperation) GetControl() common.Address {
	return u.control
}

func (u *UserOperation) SetControl(control common.Address) {
	u.control = control
}

func (u *UserOperation) GetCallConfig() uint32 {
	return u.callConfig
}

func (u *UserOperation) SetCallConfig(callConfig uint32) {
	u.callConfig = callConfig
}

func (u *UserOperation) GetDappGasLimit() uint32 {
	return u.dappGasLimit
}

func (u *UserOperation) SetDappGasLimit(dappGasLimit uint32) {
	u.dappGasLimit = dappGasLimit
}

func (u *UserOperation) GetSolverGasLimit() uint32 {
	return u.solverGasLimit
}

func (u *UserOperation) SetSolverGasLimit(solverGasLimit uint32) {
	u.solverGasLimit = solverGasLimit
}

func (u *UserOperation) GetBundlerSurchargeRate() *big.Int {
	if u.bundlerSurchargeRate == nil {
		return big.NewInt(0)
	}
	return new(big.Int).Set(u.bundlerSurchargeRate)
}

func (u *UserOperation) SetBundlerSurchargeRate(bundlerSurchargeRate *big.Int) {
	if bundlerSurchargeRate == nil {
		u.bundlerSurchargeRate = nil
	}
	u.bundlerSurchargeRate = new(big.Int).Set(bundlerSurchargeRate)
}

func (u *UserOperation) GetSessionKey() common.Address {
	return u.sessionKey
}

func (u *UserOperation) SetSessionKey(sessionKey common.Address) {
	u.sessionKey = sessionKey
}

func (u *UserOperation) GetData() []byte {
	return u.data
}

func (u *UserOperation) SetData(data []byte) {
	u.data = data
}

func (u *UserOperation) GetSignature() []byte {
	return u.signature
}

func (u *UserOperation) SetSignature(signature []byte) {
	u.signature = signature
}

func (u *UserOperation) Sanitize() {
	switch u.versionAtLeast {
	case config.AtlasV_1_6:
		if u.bundlerSurchargeRate == nil {
			u.bundlerSurchargeRate = big.NewInt(0)
		}
		fallthrough

	case config.AtlasV_1_5:
		// Nothing here
		fallthrough

	default:
		if u.value == nil {
			u.value = big.NewInt(0)
		}

		if u.gas == nil {
			u.gas = big.NewInt(0)
		}

		if u.maxFeePerGas == nil {
			u.maxFeePerGas = big.NewInt(0)
		}

		if u.nonce == nil {
			u.nonce = big.NewInt(0)
		}

		if u.deadline == nil {
			u.deadline = big.NewInt(0)
		}
	}
}

func (u *UserOperation) ToParams() *UserOperationsParams {
	u.Sanitize()

	return &UserOperationsParams{
		From:                 u.from,
		To:                   u.to,
		Value:                u.value,
		Gas:                  u.gas,
		MaxFeePerGas:         u.maxFeePerGas,
		Nonce:                u.nonce,
		Deadline:             u.deadline,
		Dapp:                 u.dapp,
		Control:              u.control,
		CallConfig:           u.callConfig,
		DappGasLimit:         u.dappGasLimit,
		SolverGasLimit:       u.solverGasLimit,
		BundlerSurchargeRate: u.bundlerSurchargeRate,
		SessionKey:           u.sessionKey,
		Data:                 u.data,
		Signature:            u.signature,
	}
}

func (u *UserOperation) EncodeToRaw() *UserOperationRaw {
	u.Sanitize()

	r := &UserOperationRaw{
		From:         u.from,
		To:           u.to,
		Value:        (*hexutil.Big)(u.value),
		Gas:          (*hexutil.Big)(u.gas),
		MaxFeePerGas: (*hexutil.Big)(u.maxFeePerGas),
		Nonce:        (*hexutil.Big)(u.nonce),
		Deadline:     (*hexutil.Big)(u.deadline),
		Dapp:         u.dapp,
		Control:      u.control,
		CallConfig:   (*hexutil.Big)(big.NewInt(int64(u.callConfig))),
		SessionKey:   u.sessionKey,
		Data:         hexutil.Bytes(u.data),
		Signature:    hexutil.Bytes(u.signature),
	}

	switch u.versionAtLeast {
	case config.AtlasV_1_6:
		r.SolverGasLimit = (*hexutil.Big)(big.NewInt(int64(u.solverGasLimit)))
		r.BundlerSurchargeRate = (*hexutil.Big)(u.bundlerSurchargeRate)
		fallthrough

	case config.AtlasV_1_5:
		r.DappGasLimit = (*hexutil.Big)(big.NewInt(int64(u.dappGasLimit)))
	}

	return r
}

func (u *UserOperation) Hash(trusted bool) (common.Hash, error) {
	eip712Domain, err := config.GetEip712Domain(u.chainId, &u.version)
	if err != nil {
		return common.Hash{}, err
	}

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       u.toTypedDataTypes(trusted),
		PrimaryType: "UserOperation",
		Domain:      *eip712Domain,
		Message:     u.toTypedDataMessage(trusted),
	})

	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(hash), nil
}

func (u *UserOperation) ValidateSignature() error {
	if len(u.signature) != 65 {
		return errors.New("invalid signature length")
	}

	userOpHash, err := u.Hash(false)
	if err != nil {
		return err
	}

	signer, err := utils.RecoverSigner(userOpHash.Bytes(), u.signature)
	if err != nil {
		return err
	}

	if signer != u.from {
		return errors.New("invalid signature")
	}

	return nil
}

func (u *UserOperation) AbiEncode() ([]byte, error) {
	u.Sanitize()

	switch u.versionAtLeast {
	case config.AtlasV_1_6:
		return userOpArgsV16.Pack(&u)

	case config.AtlasV_1_5:
		return userOpArgsV15.Pack(&u)

	default:
		return userOpArgsLegacy.Pack(&u)
	}
}

func (u *UserOperation) toTypedDataTypes(trusted bool) apitypes.Types {
	switch u.versionAtLeast {
	case config.AtlasV_1_6:
		return u.toTypedDataTypesV16(trusted)

	case config.AtlasV_1_5:
		return u.toTypedDataTypesV15(trusted)

	default:
		return u.toTypedDataTypesLegacy(trusted)
	}
}

func (u *UserOperation) toTypedDataMessage(trusted bool) apitypes.TypedDataMessage {
	u.Sanitize()

	switch u.versionAtLeast {
	case config.AtlasV_1_6:
		return u.toTypedDataMessageV16(trusted)

	case config.AtlasV_1_5:
		return u.toTypedDataMessageV15(trusted)

	default:
		return u.toTypedDataMessageLegacy(trusted)
	}
}

type UserOperationRaw struct {
	From                 common.Address `json:"from"`
	To                   common.Address `json:"to"`
	Value                *hexutil.Big   `json:"value"`
	Gas                  *hexutil.Big   `json:"gas"`
	MaxFeePerGas         *hexutil.Big   `json:"maxFeePerGas"`
	Nonce                *hexutil.Big   `json:"nonce"`
	Deadline             *hexutil.Big   `json:"deadline"`
	Dapp                 common.Address `json:"dapp"`
	Control              common.Address `json:"control"`
	CallConfig           *hexutil.Big   `json:"callConfig"`
	DappGasLimit         *hexutil.Big   `json:"dappGasLimit,omitempty"`         // From v1.5
	SolverGasLimit       *hexutil.Big   `json:"solverGasLimit,omitempty"`       // From v1.6
	BundlerSurchargeRate *hexutil.Big   `json:"bundlerSurchargeRate,omitempty"` // From v1.6
	SessionKey           common.Address `json:"sessionKey"`
	Data                 hexutil.Bytes  `json:"data"`
	Signature            hexutil.Bytes  `json:"signature"`
}

func (u *UserOperationRaw) Sanitize() {
	if u.Value == nil {
		u.Value = (*hexutil.Big)(big.NewInt(0))
	}

	if u.Gas == nil {
		u.Gas = (*hexutil.Big)(big.NewInt(0))
	}

	if u.MaxFeePerGas == nil {
		u.MaxFeePerGas = (*hexutil.Big)(big.NewInt(0))
	}

	if u.Nonce == nil {
		u.Nonce = (*hexutil.Big)(big.NewInt(0))
	}

	if u.Deadline == nil {
		u.Deadline = (*hexutil.Big)(big.NewInt(0))
	}

	if u.CallConfig == nil {
		u.CallConfig = (*hexutil.Big)(big.NewInt(0))
	}

	if u.DappGasLimit == nil {
		u.DappGasLimit = (*hexutil.Big)(big.NewInt(0))
	}

	if u.SolverGasLimit == nil {
		u.SolverGasLimit = (*hexutil.Big)(big.NewInt(0))
	}

	if u.BundlerSurchargeRate == nil {
		u.BundlerSurchargeRate = (*hexutil.Big)(big.NewInt(0))
	}
}

func (u *UserOperationRaw) Decode(chainId uint64) (*UserOperation, error) {
	var (
		version           = u.GetAtlasVersion(chainId)
		versionAtLeast, _ = utils.VersionAtLeast(&version)
	)

	u.Sanitize()

	params := UserOperationsParams{
		From:         u.From,
		To:           u.To,
		Value:        u.Value.ToInt(),
		Gas:          u.Gas.ToInt(),
		MaxFeePerGas: u.MaxFeePerGas.ToInt(),
		Nonce:        u.Nonce.ToInt(),
		Deadline:     u.Deadline.ToInt(),
		Dapp:         u.Dapp,
		Control:      u.Control,
		CallConfig:   uint32(u.CallConfig.ToInt().Uint64()),
		SessionKey:   u.SessionKey,
		Data:         u.Data,
		Signature:    u.Signature,
	}

	switch versionAtLeast {
	case config.AtlasV_1_6:
		params.SolverGasLimit = uint32(u.SolverGasLimit.ToInt().Uint64())
		params.BundlerSurchargeRate = u.BundlerSurchargeRate.ToInt()
		fallthrough

	case config.AtlasV_1_5:
		params.DappGasLimit = uint32(u.DappGasLimit.ToInt().Uint64())
	}

	return NewUserOperation(chainId, params)
}

func (u *UserOperationRaw) GetAtlasVersion(chainId uint64) string {
	v, err := config.GetVersionFromAtlasAddress(chainId, u.To)
	if err == nil {
		return v
	}

	if u.SolverGasLimit != nil || u.BundlerSurchargeRate != nil {
		return config.AtlasV_1_6
	}

	if u.DappGasLimit != nil {
		return config.AtlasV_1_5
	}

	// Any value below 1.5 is fine
	return config.AtlasV_1_3
}

type UserOperationPartialRaw struct {
	ChainId *hexutil.Big `json:"chainId"`

	UserOpHash   common.Hash    `json:"userOpHash"`
	To           common.Address `json:"to"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`

	//Exactly one of 1. Hints  2. (Value, Data, From) must be set
	Hints map[string]interface{} `json:"hints,omitempty"`

	Value *hexutil.Big   `json:"value,omitempty"`
	Data  hexutil.Bytes  `json:"data,omitempty"`
	From  common.Address `json:"from,omitempty"`
}

func NewUserOperationPartialRaw(userOp *UserOperation, hints map[string]interface{}) (*UserOperationPartialRaw, error) {
	trustedOpHash, err := utils.FlagTrustedOpHash(userOp.GetCallConfig(), &userOp.version)
	if err != nil {
		return nil, err
	}

	userOpHash, err := userOp.Hash(trustedOpHash)
	if err != nil {
		return nil, err
	}

	userOpPartial := &UserOperationPartialRaw{
		ChainId:      (*hexutil.Big)(big.NewInt(int64(userOp.chainId))),
		UserOpHash:   userOpHash,
		To:           userOp.GetTo(),
		Gas:          (*hexutil.Big)(userOp.GetGas()),
		MaxFeePerGas: (*hexutil.Big)(userOp.GetMaxFeePerGas()),
		Deadline:     (*hexutil.Big)(userOp.GetDeadline()),
		Dapp:         userOp.GetDapp(),
		Control:      userOp.GetControl(),
	}

	if len(hints) > 0 {
		// Randomize slices/arrays
		for key, hint := range hints {
			v := reflect.ValueOf(hint)
			if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
				randomizeSlice(v)
				hints[key] = v.Interface()
			}
		}
		userOpPartial.Hints = hints
	} else {
		userOpPartial.Data = hexutil.Bytes(userOp.GetData())
		userOpPartial.From = userOp.GetFrom()
		userOpPartial.Value = (*hexutil.Big)(userOp.GetValue())
	}

	return userOpPartial, nil
}

type UserOperationWithHintsRaw struct {
	ChainId       *hexutil.Big           `json:"chainId"`
	UserOperation *UserOperationRaw      `json:"userOperation"`
	Hints         map[string]interface{} `json:"hints"`
}

func NewUserOperationWithHintsRaw(userOp *UserOperation, hints map[string]interface{}) *UserOperationWithHintsRaw {
	return &UserOperationWithHintsRaw{
		ChainId:       (*hexutil.Big)(big.NewInt(int64(userOp.chainId))),
		UserOperation: userOp.EncodeToRaw(),
		Hints:         hints,
	}
}

func (uop *UserOperationWithHintsRaw) Decode() (uint64, *UserOperation, map[string]interface{}, error) {
	userOp, err := uop.UserOperation.Decode(uop.ChainId.ToInt().Uint64())
	if err != nil {
		return 0, nil, nil, err
	}

	return uop.ChainId.ToInt().Uint64(), userOp, uop.Hints, nil
}

func randomizeSlice(v reflect.Value) {
	if v.Len() > 1 {
		rand.Shuffle(v.Len(), func(i, j int) {
			vi, vj := v.Index(i), v.Index(j)
			tmp := reflect.New(vi.Type()).Elem()
			tmp.Set(vi)
			vi.Set(vj)
			vj.Set(tmp)
		})
	}
}
