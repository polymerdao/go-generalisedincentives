TMP_ABI_DIR_VIBC = .tmp_abi

.PHONY: abigen clean

CONTRACTS_DIR = ./GeneralisedIncentives/out

CONTRACT_ABI_FILES = \
	$(CONTRACTS_DIR)/vIBCEscrow.sol/IncentivizedPolymerEscrow.json \

TMP_ABI_DIR = .tmp_abi

$(TMP_ABI_DIR): $(CONTRACT_ABI_FILES)
	rm -rf $(TMP_ABI_DIR)
	mkdir -p $(TMP_ABI_DIR)
	for abi_file in $(CONTRACT_ABI_FILES); do \
		cat $$abi_file | jq .abi > $(TMP_ABI_DIR)/$$(basename $$abi_file); \
	done

abigen: $(TMP_ABI_DIR)
	for abi_file in $(wildcard $(TMP_ABI_DIR)/*.json); do \
		abi_base=$$(basename $$abi_file); \
		type=$$(basename $$abi_file .json); \
		abigen --pkg generalisedincentives --abi $$abi_file --type $$type --out $$type.go; \
	done
