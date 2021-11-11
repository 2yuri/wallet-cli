# 💸 wallet-cli
🪙 ETH wallet CLI

## Installation
```
$ go get -u github.com/hyperyuri/wallet-cli
```

## Accepted Currencies

- ETH/ETH
- ETH/USDT
- BSC/BNB
- BSC/USDT

## Usage:

1. Create your wallets
```
$ wallet-cli create (--pass or -p) waller_password
```

2. Add a new address to your wallet
```
$ wallet-cli add (--walet or -w) wallet_uuid (--pass or -p) wallet_password
```
3. List your addresses by your wallets and currencies
```
$ wallet-cli list (--walet or -w) wallet_uuid (--pass or -p) wallet_password (--network or -n) ETH (--currency or -c) USDT
```

4. Send a transaction 
```
$ wallet-cli list (--walet or -w) wallet_uuid (--pass or -p) wallet_password (--network or -n) ETH (--currency or -c) USDT (--from or -f) address (--to or -t) to_address (--amount or -a) 1.3
```

- After send a transaction you need to confirm the fee amount and select your option
- 1 - Accept discounting fee from balance.
- 2 - Accept discounting fee from amount.
- 3 - Cancel transaction.

5. List your transactions
```
$ wallet-cli transactions (--walet or -w) wallet_uuid (--pass or -p) wallet_password (--count or -c) 10 (last 10 transaction) (--status or -s) (all, pending or done)
```

## WIP
- Add transaction for non native currencies (USDT)
- Improve readme with images
- Add more coins
