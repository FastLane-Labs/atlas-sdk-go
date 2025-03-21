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
			callConfig: 1 << IndexLegacyUserNoncesSequential,
			want:       true,
		},
		{
			name:       "UserNoncesSequentialNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UserNoncesSequentialAndDAppNoncesSequential",
			callConfig: 1<<IndexLegacyUserNoncesSequential | 1<<IndexLegacyDappNoncesSequential,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUserNoncesSequential(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyDappNoncesSequential,
			want:       true,
		},
		{
			name:       "DAppNoncesSequentialNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "DAppNoncesSequentialAndUserNoncesSequential",
			callConfig: 1<<IndexLegacyDappNoncesSequential | 1<<IndexLegacyUserNoncesSequential,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDappNoncesSequential(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyRequirePreOps,
			want:       true,
		},
		{
			name:       "RequirePreOpsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePreOpsAndTrackPreOpsReturnData",
			callConfig: 1<<IndexLegacyRequirePreOps | 1<<IndexLegacyTrackPreOpsReturnData,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePreOps(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyTrackPreOpsReturnData,
			want:       true,
		},
		{
			name:       "TrackPreOpsReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrackPreOpsReturnDataAndTrackUserReturnData",
			callConfig: 1<<IndexLegacyTrackPreOpsReturnData | 1<<IndexLegacyTrackUserReturnData,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrackPreOpsReturnData(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyTrackUserReturnData,
			want:       true,
		},
		{
			name:       "TrackUserReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrackUserReturnDataAndDelegateUser",
			callConfig: 1<<IndexLegacyTrackUserReturnData | 1<<IndexLegacyDelegateUser,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrackUserReturnData(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyDelegateUser,
			want:       true,
		},
		{
			name:       "DelegateUserNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "DelegateUserAndRequirePreSolver",
			callConfig: 1<<IndexLegacyDelegateUser | 1<<IndexLegacyRequirePreSolver,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDelegateUser(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyRequirePreSolver,
			want:       true,
		},
		{
			name:       "RequirePreSolverNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePreSolverAndRequirePostSolver",
			callConfig: 1<<IndexLegacyRequirePreSolver | 1<<IndexLegacyRequirePostSolver,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePreSolver(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyRequirePostSolver,
			want:       true,
		},
		{
			name:       "RequirePostSolverNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePostSolverAndRequirePostOps",
			callConfig: 1<<IndexLegacyRequirePostSolver | 1<<IndexLegacyRequirePostOps,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePostSolver(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyRequirePostOps,
			want:       true,
		},
		{
			name:       "RequirePostOpsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePostOpsAndZeroSolvers",
			callConfig: 1<<IndexLegacyRequirePostOps | 1<<IndexLegacyZeroSolvers,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePostOps(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyZeroSolvers,
			want:       true,
		},
		{
			name:       "ZeroSolversNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ZeroSolversAndReuseUserOp",
			callConfig: 1<<IndexLegacyZeroSolvers | 1<<IndexLegacyReuseUserOp,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagZeroSolvers(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyReuseUserOp,
			want:       true,
		},
		{
			name:       "ReuseUserOpNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ReuseUserOpAndUserAuctioneer",
			callConfig: 1<<IndexLegacyReuseUserOp | 1<<IndexLegacyUserAuctioneer,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagReuseUserOp(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyUserAuctioneer,
			want:       true,
		},
		{
			name:       "UserAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UserAuctioneerAndSolverAuctioneer",
			callConfig: 1<<IndexLegacyUserAuctioneer | 1<<IndexLegacySolverAuctioneer,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUserAuctioneer(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacySolverAuctioneer,
			want:       true,
		},
		{
			name:       "SolverAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "SolverAuctioneerAndUnknownAuctioneer",
			callConfig: 1<<IndexLegacySolverAuctioneer | 1<<IndexLegacyUnknownAuctioneer,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagSolverAuctioneer(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyUnknownAuctioneer,
			want:       true,
		},
		{
			name:       "UnknownAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UnknownAuctioneerAndUserNoncesSequential",
			callConfig: 1<<IndexLegacyUnknownAuctioneer | 1<<IndexLegacyUserNoncesSequential,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUnknownAuctioneer(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyVerifyCallChainHash,
			want:       true,
		},
		{
			name:       "VerifyCallChainHashNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "VerifyCallChainHashAndIndexInvertBidValue",
			callConfig: 1<<IndexLegacyVerifyCallChainHash | 1<<IndexLegacyInvertBidValue,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagVerifyCallChainHash(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyInvertBidValue,
			want:       true,
		},
		{
			name:       "InvertBidValueNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "InvertBidValueAndUserNoncesSequential",
			callConfig: 1<<IndexLegacyInvertBidValue | 1<<IndexLegacyUserNoncesSequential,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagInvertBidValue(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyRequireFulfillment,
			want:       true,
		},
		{
			name:       "RequireFulfillmentNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequireFulfillmentAndTrustedOpHash",
			callConfig: 1<<IndexLegacyRequireFulfillment | 1<<IndexLegacyTrustedOpHash,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequireFulfillment(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyTrustedOpHash,
			want:       true,
		},
		{
			name:       "TrustedOpHashNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrustedOpHashAndFlagForwardReturnData",
			callConfig: 1<<IndexLegacyTrustedOpHash | 1<<IndexLegacyForwardReturnData,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrustedOpHash(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyForwardReturnData,
			want:       true,
		},
		{
			name:       "ForwardReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ForwardReturnDataAndFlagRequireFulfillment",
			callConfig: 1<<IndexLegacyForwardReturnData | 1<<IndexLegacyRequireFulfillment,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagForwardReturnData(tt.callConfig, &version); got != tt.want {
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
			callConfig: 1 << IndexLegacyExPostBids,
			want:       true,
		},
		{
			name:       "ExPostBidsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ExPostBidsAndFlagTrustedOpHash",
			callConfig: 1<<IndexLegacyExPostBids | 1<<IndexLegacyTrustedOpHash,
			want:       true,
		},
	}

	version := "1.0"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagExPostBids(tt.callConfig, &version); got != tt.want {
				t.Errorf("FlagExPostBids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagVerifyCallChainHashAndInvertBidValue(t *testing.T) {
	t.Parallel()

	version := "1.0"
	callConfig := uint32(1<<IndexLegacyVerifyCallChainHash | 1<<IndexLegacyInvertBidValue)
	want := true

	if got := FlagVerifyCallChainHash(callConfig, &version); got != want {
		t.Errorf("FlagVerifyCallChainHash() = %v, want %v", got, want)
	}

	want = true

	if got := FlagInvertBidValue(callConfig, &version); got != want {
		t.Errorf("FlagInvertBidValue() = %v, want %v", got, want)
	}
}

func TestFlagTrackUserReturnDataAndDelegateUser(t *testing.T) {
	t.Parallel()

	version := "1.0"
	callConfig := uint32(1<<IndexLegacyTrackUserReturnData | 1<<IndexLegacyDelegateUser)
	want := true

	if got := FlagTrackUserReturnData(callConfig, &version); got != want {
		t.Errorf("FlagTrackUserReturnData() = %v, want %v", got, want)
	}

	want = true

	if got := FlagDelegateUser(callConfig, &version); got != want {
		t.Errorf("FlagDelegateUser() = %v, want %v", got, want)
	}
}

func TestFlagRequirePreSolverAndRequirePostSolver(t *testing.T) {
	t.Parallel()

	version := "1.0"
	callConfig := uint32(1<<IndexLegacyRequirePreSolver | 1<<IndexLegacyRequirePostSolver)
	want := true

	if got := FlagRequirePreSolver(callConfig, &version); got != want {
		t.Errorf("FlagRequirePreSolver() = %v, want %v", got, want)
	}

	want = true

	if got := FlagRequirePostSolver(callConfig, &version); got != want {
		t.Errorf("FlagRequirePostSolver() = %v, want %v", got, want)
	}
}

func TestFlagMultipleSuccessfulSolvers(t *testing.T) {
	t.Parallel()

	callConfig := uint32(532548)

	version := "1.5"
	if got := FlagMultipleSuccessfulSolvers(callConfig, &version); got != true {
		t.Errorf("FlagMultipleSuccessfulSolvers() = %v, want %v", got, true)
	}

	if got := FlagExPostBids(callConfig, &version); got != false {
		t.Errorf("FlagExPostBids() = %v, want %v", got, false)
	}
}
