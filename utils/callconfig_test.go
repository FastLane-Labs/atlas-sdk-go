package utils

import (
	"testing"
)

func TestFlagUserNoncesSequential(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "UserNoncesSequential",
			callConfig: 1 << IndexUserNoncesSequential,
			want:       true,
		},
		{
			name:       "UserNoncesSequentialNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UserNoncesSequentialAndDAppNoncesSequential",
			callConfig: 1<<IndexUserNoncesSequential | 1<<IndexDappNoncesSequential,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUserNoncesSequential(tt.callConfig); got != tt.want {
				t.Errorf("FlagUserNoncesSequential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagDAppNoncesSequential(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "DAppNoncesSequential",
			callConfig: 1 << IndexDappNoncesSequential,
			want:       true,
		},
		{
			name:       "DAppNoncesSequentialNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "DAppNoncesSequentialAndUserNoncesSequential",
			callConfig: 1<<IndexDappNoncesSequential | 1<<IndexUserNoncesSequential,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDappNoncesSequential(tt.callConfig); got != tt.want {
				t.Errorf("FlagDAppNoncesSequential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRequirePreOps(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "RequirePreOps",
			callConfig: 1 << IndexRequirePreOps,
			want:       true,
		},
		{
			name:       "RequirePreOpsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePreOpsAndTrackPreOpsReturnData",
			callConfig: 1<<IndexRequirePreOps | 1<<IndexTrackPreOpsReturnData,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePreOps(tt.callConfig); got != tt.want {
				t.Errorf("FlagRequirePreOps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagTrackPreOpsReturnData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "TrackPreOpsReturnData",
			callConfig: 1 << IndexTrackPreOpsReturnData,
			want:       true,
		},
		{
			name:       "TrackPreOpsReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrackPreOpsReturnDataAndTrackUserReturnData",
			callConfig: 1<<IndexTrackPreOpsReturnData | 1<<IndexTrackUserReturnData,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrackPreOpsReturnData(tt.callConfig); got != tt.want {
				t.Errorf("FlagTrackPreOpsReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagTrackUserReturnData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "TrackUserReturnData",
			callConfig: 1 << IndexTrackUserReturnData,
			want:       true,
		},
		{
			name:       "TrackUserReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrackUserReturnDataAndDelegateUser",
			callConfig: 1<<IndexTrackUserReturnData | 1<<IndexDelegateUser,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrackUserReturnData(tt.callConfig); got != tt.want {
				t.Errorf("FlagTrackUserReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagDelegateUser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "DelegateUser",
			callConfig: 1 << IndexDelegateUser,
			want:       true,
		},
		{
			name:       "DelegateUserNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "DelegateUserAndRequirePreSolver",
			callConfig: 1<<IndexDelegateUser | 1<<IndexRequirePreSolver,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDelegateUser(tt.callConfig); got != tt.want {
				t.Errorf("FlagDelegateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRequirePreSolver(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "RequirePreSolver",
			callConfig: 1 << IndexRequirePreSolver,
			want:       true,
		},
		{
			name:       "RequirePreSolverNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePreSolverAndRequirePostSolver",
			callConfig: 1<<IndexRequirePreSolver | 1<<IndexRequirePostSolver,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePreSolver(tt.callConfig); got != tt.want {
				t.Errorf("FlagRequirePreSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRequirePostSolver(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "RequirePostSolver",
			callConfig: 1 << IndexRequirePostSolver,
			want:       true,
		},
		{
			name:       "RequirePostSolverNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePostSolverAndRequirePostOps",
			callConfig: 1<<IndexRequirePostSolver | 1<<IndexRequirePostOps,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePostSolver(tt.callConfig); got != tt.want {
				t.Errorf("FlagRequirePostSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRequirePostOps(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "RequirePostOps",
			callConfig: 1 << IndexRequirePostOps,
			want:       true,
		},
		{
			name:       "RequirePostOpsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePostOpsAndZeroSolvers",
			callConfig: 1<<IndexRequirePostOps | 1<<IndexZeroSolvers,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePostOps(tt.callConfig); got != tt.want {
				t.Errorf("FlagRequirePostOps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagZeroSolvers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ZeroSolvers",
			callConfig: 1 << IndexZeroSolvers,
			want:       true,
		},
		{
			name:       "ZeroSolversNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ZeroSolversAndReuseUserOp",
			callConfig: 1<<IndexZeroSolvers | 1<<IndexReuseUserOp,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagZeroSolvers(tt.callConfig); got != tt.want {
				t.Errorf("FlagZeroSolvers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagReuseUserOp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ReuseUserOp",
			callConfig: 1 << IndexReuseUserOp,
			want:       true,
		},
		{
			name:       "ReuseUserOpNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ReuseUserOpAndUserAuctioneer",
			callConfig: 1<<IndexReuseUserOp | 1<<IndexUserAuctioneer,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagReuseUserOp(tt.callConfig); got != tt.want {
				t.Errorf("FlagReuseUserOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagUserAuctioneer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "UserAuctioneer",
			callConfig: 1 << IndexUserAuctioneer,
			want:       true,
		},
		{
			name:       "UserAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UserAuctioneerAndSolverAuctioneer",
			callConfig: 1<<IndexUserAuctioneer | 1<<IndexSolverAuctioneer,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUserAuctioneer(tt.callConfig); got != tt.want {
				t.Errorf("FlagUserAuctioneer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagSolverAuctioneer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "SolverAuctioneer",
			callConfig: 1 << IndexSolverAuctioneer,
			want:       true,
		},
		{
			name:       "SolverAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "SolverAuctioneerAndUnknownAuctioneer",
			callConfig: 1<<IndexSolverAuctioneer | 1<<IndexUnknownAuctioneer,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagSolverAuctioneer(tt.callConfig); got != tt.want {
				t.Errorf("FlagSolverAuctioneer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagUnknownAuctioneer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "UnknownAuctioneer",
			callConfig: 1 << IndexUnknownAuctioneer,
			want:       true,
		},
		{
			name:       "UnknownAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UnknownAuctioneerAndUserNoncesSequential",
			callConfig: 1<<IndexUnknownAuctioneer | 1<<IndexUserNoncesSequential,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUnknownAuctioneer(tt.callConfig); got != tt.want {
				t.Errorf("FlagUnknownAuctioneer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagVerifyCallChainHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "VerifyCallChainHash",
			callConfig: 1 << IndexVerifyCallChainHash,
			want:       true,
		},
		{
			name:       "VerifyCallChainHashNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "VerifyCallChainHashAndIndexInvertBidValue",
			callConfig: 1<<IndexVerifyCallChainHash | 1<<IndexInvertBidValue,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagVerifyCallChainHash(tt.callConfig); got != tt.want {
				t.Errorf("FlagVerifyCallChainHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagInvertBidValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "InvertBidValue",
			callConfig: 1 << IndexInvertBidValue,
			want:       true,
		},
		{
			name:       "InvertBidValueNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "InvertBidValueAndUserNoncesSequential",
			callConfig: 1<<IndexInvertBidValue | 1<<IndexUserNoncesSequential,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagInvertBidValue(tt.callConfig); got != tt.want {
				t.Errorf("FlagInvertBidValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRequireFulfillment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "RequireFulfillment",
			callConfig: 1 << IndexRequireFulfillment,
			want:       true,
		},
		{
			name:       "RequireFulfillmentNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequireFulfillmentAndTrustedOpHash",
			callConfig: 1<<IndexRequireFulfillment | 1<<IndexTrustedOpHash,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequireFulfillment(tt.callConfig); got != tt.want {
				t.Errorf("FlagRequireFulfillment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagTrustedOpHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "TrustedOpHash",
			callConfig: 1 << IndexTrustedOpHash,
			want:       true,
		},
		{
			name:       "TrustedOpHashNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrustedOpHashAndFlagForwardReturnData",
			callConfig: 1<<IndexTrustedOpHash | 1<<IndexForwardReturnData,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrustedOpHash(tt.callConfig); got != tt.want {
				t.Errorf("FlagTrustedOpHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagForwardReturnData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ForwardReturnData",
			callConfig: 1 << IndexForwardReturnData,
			want:       true,
		},
		{
			name:       "ForwardReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ForwardReturnDataAndFlagRequireFulfillment",
			callConfig: 1<<IndexForwardReturnData | 1<<IndexRequireFulfillment,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagForwardReturnData(tt.callConfig); got != tt.want {
				t.Errorf("FlagForwardReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagExPostBids(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ExPostBids",
			callConfig: 1 << IndexExPostBids,
			want:       true,
		},
		{
			name:       "ExPostBidsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ExPostBidsAndFlagTrustedOpHash",
			callConfig: 1<<IndexExPostBids | 1<<IndexTrustedOpHash,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagExPostBids(tt.callConfig); got != tt.want {
				t.Errorf("FlagExPostBids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagVerifyCallChainHashAndInvertBidValue(t *testing.T) {
	t.Parallel()

	callConfig := uint32(1<<IndexVerifyCallChainHash | 1<<IndexInvertBidValue)
	want := true

	if got := FlagVerifyCallChainHash(callConfig); got != want {
		t.Errorf("FlagVerifyCallChainHash() = %v, want %v", got, want)
	}

	want = true

	if got := FlagInvertBidValue(callConfig); got != want {
		t.Errorf("FlagInvertBidValue() = %v, want %v", got, want)
	}
}

func TestFlagTrackUserReturnDataAndDelegateUser(t *testing.T) {
	t.Parallel()

	callConfig := uint32(1<<IndexTrackUserReturnData | 1<<IndexDelegateUser)
	want := true

	if got := FlagTrackUserReturnData(callConfig); got != want {
		t.Errorf("FlagTrackUserReturnData() = %v, want %v", got, want)
	}

	want = true

	if got := FlagDelegateUser(callConfig); got != want {
		t.Errorf("FlagDelegateUser() = %v, want %v", got, want)
	}
}

func TestFlagRequirePreSolverAndRequirePostSolver(t *testing.T) {
	t.Parallel()

	callConfig := uint32(1<<IndexRequirePreSolver | 1<<IndexRequirePostSolver)
	want := true

	if got := FlagRequirePreSolver(callConfig); got != want {
		t.Errorf("FlagRequirePreSolver() = %v, want %v", got, want)
	}

	want = true

	if got := FlagRequirePostSolver(callConfig); got != want {
		t.Errorf("FlagRequirePostSolver() = %v, want %v", got, want)
	}
}
