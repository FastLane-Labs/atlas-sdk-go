package utils

import "github.com/FastLane-Labs/atlas-sdk-go/config"

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

const (
	IndexV15UserNoncesSequential = iota
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
	IndexV15MultipleSuccessfulSolvers
)

func FlagUserNoncesSequential(callConfig uint32, version *string) bool {
	idx := IndexLegacyUserNoncesSequential
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15UserNoncesSequential
	}
	return callConfig&(1<<idx) != 0
}

func FlagDappNoncesSequential(callConfig uint32, version *string) bool {
	idx := IndexLegacyDappNoncesSequential
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15DappNoncesSequential
	}
	return callConfig&(1<<idx) != 0
}

func FlagRequirePreOps(callConfig uint32, version *string) bool {
	idx := IndexLegacyRequirePreOps
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15RequirePreOps
	}
	return callConfig&(1<<idx) != 0
}

func FlagTrackPreOpsReturnData(callConfig uint32, version *string) bool {
	idx := IndexLegacyTrackPreOpsReturnData
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15TrackPreOpsReturnData
	}
	return callConfig&(1<<idx) != 0
}

func FlagTrackUserReturnData(callConfig uint32, version *string) bool {
	idx := IndexLegacyTrackUserReturnData
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15TrackUserReturnData
	}
	return callConfig&(1<<idx) != 0
}

func FlagDelegateUser(callConfig uint32, version *string) bool {
	idx := IndexLegacyDelegateUser
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15DelegateUser
	}
	return callConfig&(1<<idx) != 0
}

func FlagRequirePreSolver(callConfig uint32, version *string) bool {
	idx := IndexLegacyRequirePreSolver
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15RequirePreSolver
	}
	return callConfig&(1<<idx) != 0
}

func FlagRequirePostSolver(callConfig uint32, version *string) bool {
	idx := IndexLegacyRequirePostSolver
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15RequirePostSolver
	}
	return callConfig&(1<<idx) != 0
}

func FlagRequirePostOps(callConfig uint32, version *string) bool {
	idx := IndexLegacyRequirePostOps
	return callConfig&(1<<idx) != 0
}

func FlagZeroSolvers(callConfig uint32, version *string) bool {
	idx := IndexLegacyZeroSolvers
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15ZeroSolvers
	}
	return callConfig&(1<<idx) != 0
}

func FlagReuseUserOp(callConfig uint32, version *string) bool {
	idx := IndexLegacyReuseUserOp
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15ReuseUserOp
	}
	return callConfig&(1<<idx) != 0
}

func FlagUserAuctioneer(callConfig uint32, version *string) bool {
	idx := IndexLegacyUserAuctioneer
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15UserAuctioneer
	}
	return callConfig&(1<<idx) != 0
}

func FlagSolverAuctioneer(callConfig uint32, version *string) bool {
	idx := IndexLegacySolverAuctioneer
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15SolverAuctioneer
	}
	return callConfig&(1<<idx) != 0
}

func FlagUnknownAuctioneer(callConfig uint32, version *string) bool {
	idx := IndexLegacyUnknownAuctioneer
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15UnknownAuctioneer
	}
	return callConfig&(1<<idx) != 0
}

func FlagVerifyCallChainHash(callConfig uint32, version *string) bool {
	idx := IndexLegacyVerifyCallChainHash
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15VerifyCallChainHash
	}
	return callConfig&(1<<idx) != 0
}

func FlagForwardReturnData(callConfig uint32, version *string) bool {
	idx := IndexLegacyForwardReturnData
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15ForwardReturnData
	}
	return callConfig&(1<<idx) != 0
}

func FlagRequireFulfillment(callConfig uint32, version *string) bool {
	idx := IndexLegacyRequireFulfillment
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15RequireFulfillment
	}
	return callConfig&(1<<idx) != 0
}

func FlagTrustedOpHash(callConfig uint32, version *string) bool {
	idx := IndexLegacyTrustedOpHash
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15TrustedOpHash
	}
	return callConfig&(1<<idx) != 0
}

func FlagInvertBidValue(callConfig uint32, version *string) bool {
	idx := IndexLegacyInvertBidValue
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15InvertBidValue
	}
	return callConfig&(1<<idx) != 0
}

func FlagExPostBids(callConfig uint32, version *string) bool {
	idx := IndexLegacyExPostBids
	minVersion := "1.5"
	if ok, err := config.IsVersionAtLeast(version, &minVersion); ok && err == nil {
		idx = IndexV15ExPostBids
	}
	return callConfig&(1<<idx) != 0
}

func FlagAllowAllocateValueFailure(callConfig uint32, version *string) bool {
	idx := IndexLegacyAllowAllocateValueFailure
	return callConfig&(1<<idx) != 0
}

func FlagMultipleSuccessfulSolvers(callConfig uint32, version *string) bool {
	idx := IndexV15MultipleSuccessfulSolvers
	return callConfig&(1<<idx) != 0
}
