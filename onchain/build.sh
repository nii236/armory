#!/bin/bash
set -e
solc --overwrite --abi onchain.sol -o .
abigen --abi=WithdrawalQueue.abi --pkg=onchain --type WithdrawalQueue --out=WithdrawalQueue.go
abigen --abi=Lido.abi --pkg=onchain --type Lido --out=Lido.go