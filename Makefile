TESTALL = $$(go list ./...)
GOTEST = GODEBUG=cgocheck=0 go test --count=1

test:
	$(GOTEST) --timeout 5m $(TESTALL)

contracts-bindings:
	abigen --abi ./contract/atlas/1.0/abi.json --pkg atlas_1_0 --type Atlas --out ./contract/atlas/1.0/atlas.go
	abigen --abi ./contract/atlas/1.1/abi.json --pkg atlas_1_1 --type Atlas --out ./contract/atlas/1.1/atlas.go
	abigen --abi ./contract/atlas/1.2/abi.json --pkg atlas_1_2 --type Atlas --out ./contract/atlas/1.2/atlas.go
	abigen --abi ./contract/atlas/1.3/abi.json --pkg atlas_1_3 --type Atlas --out ./contract/atlas/1.3/atlas.go
	abigen --abi ./contract/atlas/1.5/abi.json --pkg atlas_1_5 --type Atlas --out ./contract/atlas/1.5/atlas.go
	abigen --abi ./contract/atlas/1.5-monad/abi.json --pkg atlas_1_5_monad --type Atlas --out ./contract/atlas/1.5-monad/atlas.go
	abigen --abi ./contract/atlas/1.6/abi.json --pkg atlas_1_6 --type Atlas --out ./contract/atlas/1.6/atlas.go

	abigen --abi ./contract/atlasverification/1.0/abi.json --pkg atlasverification_1_0 --type AtlasVerification --out ./contract/atlasverification/1.0/atlasverification.go
	abigen --abi ./contract/atlasverification/1.1/abi.json --pkg atlasverification_1_1 --type AtlasVerification --out ./contract/atlasverification/1.1/atlasverification.go
	abigen --abi ./contract/atlasverification/1.2/abi.json --pkg atlasverification_1_2 --type AtlasVerification --out ./contract/atlasverification/1.2/atlasverification.go
	abigen --abi ./contract/atlasverification/1.3/abi.json --pkg atlasverification_1_3 --type AtlasVerification --out ./contract/atlasverification/1.3/atlasverification.go
	abigen --abi ./contract/atlasverification/1.5/abi.json --pkg atlasverification_1_5 --type AtlasVerification --out ./contract/atlasverification/1.5/atlasverification.go
	abigen --abi ./contract/atlasverification/1.5-monad/abi.json --pkg atlasverification_1_5_monad --type AtlasVerification --out ./contract/atlasverification/1.5-monad/atlasverification.go
	abigen --abi ./contract/atlasverification/1.6/abi.json --pkg atlasverification_1_6 --type AtlasVerification --out ./contract/atlasverification/1.6/atlasverification.go

	abigen --abi ./contract/simulator/1.0/abi.json --pkg simulator_1_0 --type Simulator --out ./contract/simulator/1.0/simulator.go
	abigen --abi ./contract/simulator/1.1/abi.json --pkg simulator_1_1 --type Simulator --out ./contract/simulator/1.1/simulator.go
	abigen --abi ./contract/simulator/1.2/abi.json --pkg simulator_1_2 --type Simulator --out ./contract/simulator/1.2/simulator.go
	abigen --abi ./contract/simulator/1.3/abi.json --pkg simulator_1_3 --type Simulator --out ./contract/simulator/1.3/simulator.go
	abigen --abi ./contract/simulator/1.5/abi.json --pkg simulator_1_5 --type Simulator --out ./contract/simulator/1.5/simulator.go
	abigen --abi ./contract/simulator/1.5-monad/abi.json --pkg simulator_1_5_monad --type Simulator --out ./contract/simulator/1.5-monad/simulator.go
	abigen --abi ./contract/simulator/1.6/abi.json --pkg simulator_1_6 --type Simulator --out ./contract/simulator/1.6/simulator.go

	abigen --abi ./contract/sorter/1.0/abi.json --pkg sorter_1_0 --type Sorter --out ./contract/sorter/1.0/sorter.go
	abigen --abi ./contract/sorter/1.1/abi.json --pkg sorter_1_1 --type Sorter --out ./contract/sorter/1.1/sorter.go
	abigen --abi ./contract/sorter/1.2/abi.json --pkg sorter_1_2 --type Sorter --out ./contract/sorter/1.2/sorter.go
	abigen --abi ./contract/sorter/1.3/abi.json --pkg sorter_1_3 --type Sorter --out ./contract/sorter/1.3/sorter.go
	abigen --abi ./contract/sorter/1.5/abi.json --pkg sorter_1_5 --type Sorter --out ./contract/sorter/1.5/sorter.go
	abigen --abi ./contract/sorter/1.5-monad/abi.json --pkg sorter_1_5_monad --type Sorter --out ./contract/sorter/1.5-monad/sorter.go
	abigen --abi ./contract/sorter/1.6/abi.json --pkg sorter_1_6 --type Sorter --out ./contract/sorter/1.6/sorter.go

	abigen --abi ./contract/dappcontrol/legacy/abi.json --pkg dappcontrol_legacy --type DAppControl --out ./contract/dappcontrol/legacy/dappcontrol.go
	abigen --abi ./contract/dappcontrol/1.5/abi.json --pkg dappcontrol_1_5 --type DAppControl --out ./contract/dappcontrol/1.5/dappcontrol.go
	abigen --abi ./contract/dappcontrol/1.6/abi.json --pkg dappcontrol_1_6 --type DAppControl --out ./contract/dappcontrol/1.6/dappcontrol.go

	abigen --abi ./contract/shmonad/abi.json --pkg shmonad --type ShMonad --out ./contract/shmonad/shmonad.go
