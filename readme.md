The armory service would need a minimal interface

- Withdraw (for user)
- Balance (for address)
- Sign transaction (for user)
- Create wallet (for user)

With the following security considerations:

- Firewall to only allow network traffic from the bot server to the armory server
- Sign for whitelisted contracts only
- Prevent native eth transfer (unless to the user address)
- Encryption at rest for the database
