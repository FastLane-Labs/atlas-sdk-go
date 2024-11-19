TESTALL = $$(go list ./...)
GOTEST = GODEBUG=cgocheck=0 go test --count=1

test:
	$(GOTEST) --timeout 5m $(TESTALL)

contracts-bindings:
	abigen --abi ./contract/atlas/1.0.0/abi.json --pkg atlas_1_0_0 --type Atlas --out ./contract/atlas/1.0.0/atlas.go
	abigen --abi ./contract/atlas/1.0.1/abi.json --pkg atlas_1_0_1 --type Atlas --out ./contract/atlas/1.0.1/atlas.go
	abigen --abi ./contract/atlas/1.1.0/abi.json --pkg atlas_1_1_0 --type Atlas --out ./contract/atlas/1.1.0/atlas.go

	abigen --abi ./contract/atlasverification/1.0.0/abi.json --pkg atlasverification_1_0_0 --type AtlasVerification --out ./contract/atlasverification/1.0.0/atlasverification.go
	abigen --abi ./contract/atlasverification/1.0.1/abi.json --pkg atlasverification_1_0_1 --type AtlasVerification --out ./contract/atlasverification/1.0.1/atlasverification.go
	abigen --abi ./contract/atlasverification/1.1.0/abi.json --pkg atlasverification_1_1_0 --type AtlasVerification --out ./contract/atlasverification/1.1.0/atlasverification.go

	abigen --abi ./contract/simulator/1.0.0/abi.json --pkg simulator_1_0_0 --type Simulator --out ./contract/simulator/1.0.0/simulator.go
	abigen --abi ./contract/simulator/1.0.1/abi.json --pkg simulator_1_0_1 --type Simulator --out ./contract/simulator/1.0.1/simulator.go
	abigen --abi ./contract/simulator/1.1.0/abi.json --pkg simulator_1_1_0 --type Simulator --out ./contract/simulator/1.1.0/simulator.go

	abigen --abi ./contract/sorter/1.0.0/abi.json --pkg sorter_1_0_0 --type Sorter --out ./contract/sorter/1.0.0/sorter.go
	abigen --abi ./contract/sorter/1.0.1/abi.json --pkg sorter_1_0_1 --type Sorter --out ./contract/sorter/1.0.1/sorter.go
	abigen --abi ./contract/sorter/1.1.0/abi.json --pkg sorter_1_1_0 --type Sorter --out ./contract/sorter/1.1.0/sorter.go

	abigen --abi ./contract/dappcontrol/abi.json --pkg dappcontrol --type DAppControl --out ./contract/dappcontrol/dappcontrol.go
