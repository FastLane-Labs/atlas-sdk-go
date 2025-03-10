TESTALL = $$(go list ./...)
GOTEST = GODEBUG=cgocheck=0 go test --count=1

test:
	$(GOTEST) --timeout 5m $(TESTALL)

contracts-bindings:
	abigen --abi ./contract/atlas/1.0/abi.json --pkg atlas_1_0 --type Atlas --out ./contract/atlas/1.0/atlas.go
	abigen --abi ./contract/atlas/1.1/abi.json --pkg atlas_1_1 --type Atlas --out ./contract/atlas/1.1/atlas.go
	abigen --abi ./contract/atlas/1.2/abi.json --pkg atlas_1_2 --type Atlas --out ./contract/atlas/1.2/atlas.go
	abigen --abi ./contract/atlas/1.3/abi.json --pkg atlas_1_3 --type Atlas --out ./contract/atlas/1.3/atlas.go

	abigen --abi ./contract/atlasverification/1.0/abi.json --pkg atlasverification_1_0 --type AtlasVerification --out ./contract/atlasverification/1.0/atlasverification.go
	abigen --abi ./contract/atlasverification/1.1/abi.json --pkg atlasverification_1_1 --type AtlasVerification --out ./contract/atlasverification/1.1/atlasverification.go
	abigen --abi ./contract/atlasverification/1.2/abi.json --pkg atlasverification_1_2 --type AtlasVerification --out ./contract/atlasverification/1.2/atlasverification.go
	abigen --abi ./contract/atlasverification/1.3/abi.json --pkg atlasverification_1_3 --type AtlasVerification --out ./contract/atlasverification/1.3/atlasverification.go

	abigen --abi ./contract/simulator/1.0/abi.json --pkg simulator_1_0 --type Simulator --out ./contract/simulator/1.0/simulator.go
	abigen --abi ./contract/simulator/1.1/abi.json --pkg simulator_1_1 --type Simulator --out ./contract/simulator/1.1/simulator.go
	abigen --abi ./contract/simulator/1.2/abi.json --pkg simulator_1_2 --type Simulator --out ./contract/simulator/1.2/simulator.go
	abigen --abi ./contract/simulator/1.3/abi.json --pkg simulator_1_3 --type Simulator --out ./contract/simulator/1.3/simulator.go

	abigen --abi ./contract/sorter/1.0/abi.json --pkg sorter_1_0 --type Sorter --out ./contract/sorter/1.0/sorter.go
	abigen --abi ./contract/sorter/1.1/abi.json --pkg sorter_1_1 --type Sorter --out ./contract/sorter/1.1/sorter.go
	abigen --abi ./contract/sorter/1.2/abi.json --pkg sorter_1_2 --type Sorter --out ./contract/sorter/1.2/sorter.go
	abigen --abi ./contract/sorter/1.3/abi.json --pkg sorter_1_3 --type Sorter --out ./contract/sorter/1.3/sorter.go

	abigen --abi ./contract/dappcontrol/abi.json --pkg dappcontrol --type DAppControl --out ./contract/dappcontrol/dappcontrol.go
