#!/bin/bash
set -e
solc --overwrite --abi contracts.sol -o .
abigen --abi=WithdrawalQueue.abi --pkg=contracts --type WithdrawalQueue --out=WithdrawalQueue.go
abigen --abi=Lido.abi --pkg=contracts --type Lido --out=Lido.go