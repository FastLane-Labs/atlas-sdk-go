TESTALL = $$(go list ./...)
GOTEST = GODEBUG=cgocheck=0 go test --count=1

test:
	$(GOTEST) --timeout 5m $(TESTALL)

contracts-bindings:
	abigen --abi ./contract/atlas/abi.json --pkg atlas --type Atlas --out ./contract/atlas/atlas.go
	abigen --abi ./contract/atlasverification/abi.json --pkg atlasverification --type AtlasVerification --out ./contract/atlasverification/atlasverification.go
	abigen --abi ./contract/dappcontrol/abi.json --pkg dappcontrol --type DAppControl --out ./contract/dappcontrol/dappcontrol.go
	abigen --abi ./contract/simulator/abi.json --pkg simulator --type Simulator --out ./contract/simulator/simulator.go
	abigen --abi ./contract/sorter/abi.json --pkg sorter --type Sorter --out ./contract/sorter/sorter.go
