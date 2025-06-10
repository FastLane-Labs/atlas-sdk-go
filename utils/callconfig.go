package utils

import (
	"errors"

	"github.com/FastLane-Labs/atlas-sdk-go/config"
)

// Legacy config
const (
	IndexLegacyUserNoncesSequential = uint32(iota)
	IndexLegacyDappNoncesSequential
	IndexLegacyRequirePreOps
	IndexLegacyTrackPreOpsReturnData
	IndexLegacyTrackUserReturnData
	IndexLegacyDelegateUser
	IndexLegacyRequirePreSolver
	IndexLegacyRequirePostSolver
	IndexLegacyRequirePostOps
	IndexLegacyZeroSolvers
	IndexLegacyReuseUserOp
	IndexLegacyUserAuctioneer
	IndexLegacySolverAuctioneer
	IndexLegacyUnknownAuctioneer
	IndexLegacyVerifyCallChainHash
	IndexLegacyForwardReturnData
	IndexLegacyRequireFulfillment
	IndexLegacyTrustedOpHash
	IndexLegacyInvertBidValue
	IndexLegacyExPostBids
	IndexLegacyAllowAllocateValueFailure
)

// v1.5 config
const (
	IndexV15UserNoncesSequential = uint32(iota)
	IndexV15DappNoncesSequential
	IndexV15RequirePreOps
	IndexV15TrackPreOpsReturnData
	IndexV15TrackUserReturnData
	IndexV15DelegateUser
	IndexV15RequirePreSolver
	IndexV15RequirePostSolver
	IndexV15ZeroSolvers
	IndexV15ReuseUserOp
	IndexV15UserAuctioneer
	IndexV15SolverAuctioneer
	IndexV15UnknownAuctioneer
	IndexV15VerifyCallChainHash
	IndexV15ForwardReturnData
	IndexV15RequireFulfillment
	IndexV15TrustedOpHash
	IndexV15InvertBidValue
	IndexV15ExPostBids
)

// v1.6 config
const (
	IndexV16UserNoncesSequential = uint32(iota)
	IndexV16DappNoncesSequential
	IndexV16RequirePreOps
	IndexV16TrackPreOpsReturnData
	IndexV16TrackUserReturnData
	IndexV16DelegateUser
	IndexV16RequirePreSolver
	IndexV16RequirePostSolver
	IndexV16ZeroSolvers
	IndexV16ReuseUserOp
	IndexV16UserAuctioneer
	IndexV16SolverAuctioneer
	IndexV16UnknownAuctioneer
	IndexV16VerifyCallChainHash
	IndexV16ForwardReturnData
	IndexV16RequireFulfillment
	IndexV16TrustedOpHash
	IndexV16InvertBidValue
	IndexV16ExPostBids
	IndexV16MultipleSuccessfulSolvers
)

var (
	legacy = "legacy"
	v1_5   = config.AtlasV_1_5
	v1_6   = config.AtlasV_1_6

	ErrCallConfigNotAvailableInLtV15 = errors.New("call config not available in atlas lower than v1.5")
	ErrCallConfigNotAvailableInV15   = errors.New("call config not available in atlas v1.5")
	ErrCallConfigNotAvailableInV16   = errors.New("call config not available in atlas v1.6")
)

func VersionAtLeast(v *string) (string, error) {
	_v := config.GetVersion(v)

	gte_1_6, err := config.IsVersionAtLeast(&_v, &v1_6)
	if err != nil {
		return "", err
	}

	if gte_1_6 {
		return v1_6, nil
	}

	gte_1_5, err := config.IsVersionAtLeast(&_v, &v1_5)
	if err != nil {
		return "", err
	}

	if gte_1_5 {
		return v1_5, nil
	}

	return legacy, nil
}

func FlagUserNoncesSequential(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16UserNoncesSequential
	case v1_5:
		idx = IndexV15UserNoncesSequential
	default:
		idx = IndexLegacyUserNoncesSequential
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagDappNoncesSequential(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16DappNoncesSequential
	case v1_5:
		idx = IndexV15DappNoncesSequential
	default:
		idx = IndexLegacyDappNoncesSequential
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagRequirePreOps(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16RequirePreOps
	case v1_5:
		idx = IndexV15RequirePreOps
	default:
		idx = IndexLegacyRequirePreOps
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagTrackPreOpsReturnData(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16TrackPreOpsReturnData
	case v1_5:
		idx = IndexV15TrackPreOpsReturnData
	default:
		idx = IndexLegacyTrackPreOpsReturnData
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagTrackUserReturnData(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16TrackUserReturnData
	case v1_5:
		idx = IndexV15TrackUserReturnData
	default:
		idx = IndexLegacyTrackUserReturnData
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagDelegateUser(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16DelegateUser
	case v1_5:
		idx = IndexV15DelegateUser
	default:
		idx = IndexLegacyDelegateUser
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagRequirePreSolver(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16RequirePreSolver
	case v1_5:
		idx = IndexV15RequirePreSolver
	default:
		idx = IndexLegacyRequirePreSolver
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagRequirePostSolver(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16RequirePostSolver
	case v1_5:
		idx = IndexV15RequirePostSolver
	default:
		idx = IndexLegacyRequirePostSolver
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagRequirePostOps(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		return false, ErrCallConfigNotAvailableInV16
	case v1_5:
		return false, ErrCallConfigNotAvailableInV15
	default:
		idx = IndexLegacyRequirePostOps
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagZeroSolvers(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16ZeroSolvers
	case v1_5:
		idx = IndexV15ZeroSolvers
	default:
		idx = IndexLegacyZeroSolvers
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagReuseUserOp(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16ReuseUserOp
	case v1_5:
		idx = IndexV15ReuseUserOp
	default:
		idx = IndexLegacyReuseUserOp
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagUserAuctioneer(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16UserAuctioneer
	case v1_5:
		idx = IndexV15UserAuctioneer
	default:
		idx = IndexLegacyUserAuctioneer
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagSolverAuctioneer(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16SolverAuctioneer
	case v1_5:
		idx = IndexV15SolverAuctioneer
	default:
		idx = IndexLegacySolverAuctioneer
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagUnknownAuctioneer(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16UnknownAuctioneer
	case v1_5:
		idx = IndexV15UnknownAuctioneer
	default:
		idx = IndexLegacyUnknownAuctioneer
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagVerifyCallChainHash(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16VerifyCallChainHash
	case v1_5:
		idx = IndexV15VerifyCallChainHash
	default:
		idx = IndexLegacyVerifyCallChainHash
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagForwardReturnData(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16ForwardReturnData
	case v1_5:
		idx = IndexV15ForwardReturnData
	default:
		idx = IndexLegacyForwardReturnData
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagRequireFulfillment(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16RequireFulfillment
	case v1_5:
		idx = IndexV15RequireFulfillment
	default:
		idx = IndexLegacyRequireFulfillment
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagTrustedOpHash(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16TrustedOpHash
	case v1_5:
		idx = IndexV15TrustedOpHash
	default:
		idx = IndexLegacyTrustedOpHash
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagInvertBidValue(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16InvertBidValue
	case v1_5:
		idx = IndexV15InvertBidValue
	default:
		idx = IndexLegacyInvertBidValue
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagExPostBids(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16ExPostBids
	case v1_5:
		idx = IndexV15ExPostBids
	default:
		idx = IndexLegacyExPostBids
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagAllowAllocateValueFailure(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		return false, ErrCallConfigNotAvailableInV16
	case v1_5:
		return false, ErrCallConfigNotAvailableInV15
	default:
		idx = IndexLegacyAllowAllocateValueFailure
	}

	return callConfig&(1<<idx) != 0, nil
}

func FlagMultipleSuccessfulSolvers(callConfig uint32, version *string) (bool, error) {
	_versionAtLeast, err := VersionAtLeast(version)
	if err != nil {
		return false, err
	}

	var idx uint32

	switch _versionAtLeast {
	case v1_6:
		idx = IndexV16MultipleSuccessfulSolvers
	case v1_5:
		return false, ErrCallConfigNotAvailableInV15
	default:
		return false, ErrCallConfigNotAvailableInLtV15
	}

	return callConfig&(1<<idx) != 0, nil
}
