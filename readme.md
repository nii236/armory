# Armory Package

This package provides a service to manage user withdrawals, balances, and yield for an Ethereum-based game wallet system. It includes functions for managing user collateral and facilitating controlled balance withdrawals to user wallets, while tracking yield and collateral growth.

## Key Functions

- **RequestWithdrawal**: Initiates a withdrawal request from the user's balance to their wallet. Only non-yield funds are withdrawn.
- **Balance**: Retrieves the user's current game wallet balance.
- **BalanceSeconds**: Calculates the user's balance in "eth seconds" for tracking collateral value.
- **Yield**: Fetches the yield accumulated on the user’s collateral.
- **Wallet**: Returns an existing wallet or creates a new one if it doesn't exist.

### Notes

This package is currently under development, and some features may not be fully functional.
