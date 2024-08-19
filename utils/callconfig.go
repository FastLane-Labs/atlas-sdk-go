package utils

const (
	IndexUserNoncesSequential = uint32(iota)
	IndexDappNoncesSequential
	IndexRequirePreOps
	IndexTrackPreOpsReturnData
	IndexTrackUserReturnData
	IndexDelegateUser
	IndexRequirePreSolver
	IndexRequirePostSolver
	IndexRequirePostOps
	IndexZeroSolvers
	IndexReuseUserOp
	IndexUserAuctioneer
	IndexSolverAuctioneer
	IndexUnknownAuctioneer
	IndexVerifyCallChainHash
	IndexForwardReturnData
	IndexRequireFulfillment
	IndexTrustedOpHash
	IndexInvertBidValue
	IndexExPostBids
	IndexAllowAllocateValueFailure
)

func FlagUserNoncesSequential(callConfig uint32) bool {
	return callConfig&(1<<IndexUserNoncesSequential) != 0
}

func FlagDappNoncesSequential(callConfig uint32) bool {
	return callConfig&(1<<IndexDappNoncesSequential) != 0
}

func FlagRequirePreOps(callConfig uint32) bool {
	return callConfig&(1<<IndexRequirePreOps) != 0
}

func FlagTrackPreOpsReturnData(callConfig uint32) bool {
	return callConfig&(1<<IndexTrackPreOpsReturnData) != 0
}

func FlagTrackUserReturnData(callConfig uint32) bool {
	return callConfig&(1<<IndexTrackUserReturnData) != 0
}

func FlagDelegateUser(callConfig uint32) bool {
	return callConfig&(1<<IndexDelegateUser) != 0
}

func FlagRequirePreSolver(callConfig uint32) bool {
	return callConfig&(1<<IndexRequirePreSolver) != 0
}

func FlagRequirePostSolver(callConfig uint32) bool {
	return callConfig&(1<<IndexRequirePostSolver) != 0
}

func FlagRequirePostOps(callConfig uint32) bool {
	return callConfig&(1<<IndexRequirePostOps) != 0
}

func FlagZeroSolvers(callConfig uint32) bool {
	return callConfig&(1<<IndexZeroSolvers) != 0
}

func FlagReuseUserOp(callConfig uint32) bool {
	return callConfig&(1<<IndexReuseUserOp) != 0
}

func FlagUserAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<IndexUserAuctioneer) != 0
}

func FlagSolverAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<IndexSolverAuctioneer) != 0
}

func FlagUnknownAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<IndexUnknownAuctioneer) != 0
}

func FlagVerifyCallChainHash(callConfig uint32) bool {
	return callConfig&(1<<IndexVerifyCallChainHash) != 0
}

func FlagForwardReturnData(callConfig uint32) bool {
	return callConfig&(1<<IndexForwardReturnData) != 0
}

func FlagRequireFulfillment(callConfig uint32) bool {
	return callConfig&(1<<IndexRequireFulfillment) != 0
}

func FlagTrustedOpHash(callConfig uint32) bool {
	return callConfig&(1<<IndexTrustedOpHash) != 0
}

func FlagInvertBidValue(callConfig uint32) bool {
	return callConfig&(1<<IndexInvertBidValue) != 0
}

func FlagExPostBids(callConfig uint32) bool {
	return callConfig&(1<<IndexExPostBids) != 0
}

func FlagAllowAllocateValueFailure(callConfig uint32) bool {
	return callConfig&(1<<IndexAllowAllocateValueFailure) != 0
}
