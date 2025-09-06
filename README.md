# Telegram Dice Roller + Solana Wallet Bot

A Telegram bot that:
- Rolls dice with commands like `/roll 2d6+1` (works in groups and DMs).
- Manages a per-user Solana wallet in direct messages (DMs): address, balance, and exporting the private key.

## Features
- Group chats: `/roll`, `/help`.
- Direct messages (DM): `/roll`, `/wallet_address`, `/wallet_balance`, `/wallet_private_key`, `/help`.
- One wallet per Telegram user ID, persisted in PostgreSQL.
- Solana balance fetched via JSON-RPC.

## Requirements
- Go 1.22+
- Docker + Docker Compose

## Environment
Configure via `.env` in the project root. Docker Compose will load it automatically.

Required variables:
- `BOT_TOKEN`: Telegram bot token (from BotFather)
- `SOLANA_RPC_URL`: Solana JSON-RPC endpoint (e.g. `https://api.mainnet-beta.solana.com` or `https://api.devnet.solana.com`)
- `DATABASE_URL`: Postgres DSN (Compose example uses `postgres` service hostname)

Example `.env` (do not commit secrets):
```
BOT_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
SOLANA_RPC_URL=https://solana.drpc.org
DATABASE_URL=postgres://diceroller:password@postgres:5432/diceroller?sslmode=disable
```

## Run with Docker Compose
```
docker compose up -d --build
```
- The app connects to Postgres and auto-migrates the `wallets` table.
- The bot starts polling and responds to messages.

Stop and clean:
```
docker compose down
```

## Local Development
- Ensure `BOT_TOKEN` is set in your shell, and optionally `DATABASE_URL` & `SOLANA_RPC_URL`.
- Run:
```
go run ./...
```

Build binary:
```
go build -o bin/diceroller
```

## Commands
- Group chats:
  - `/roll [NdX[+M]]` – Roll dice (e.g. `/roll 2d6+1`).
  - `/help` – Show available commands.
- Direct messages (DM):
  - `/roll [NdX[+M]]` – Same as above.
  - `/wallet_address` – Show your Solana wallet address (created on first use).
  - `/wallet_balance` – Show balance (lamports and SOL; requires valid `SOLANA_RPC_URL`).
  - `/wallet_private_key` – Show private key (base58). Be careful — this is sensitive.
  - `/help` – Show full help.

## Notes
- Security: exposing a private key is dangerous. Consider disabling `/wallet_private_key` in production or adding confirmation/limits.
- RPC: live balances require a working Solana RPC endpoint. The app validates the URL and HTTP status; errors are returned as messages.
- Persistence: wallets are stored in Postgres (`wallets` table) keyed by Telegram `user_id`.

## Project Structure
- `main.go` – app entry: loads env, DB init/migration, starts bot.
- `controllers/controller.go` – message routing and command handling.
- `models/` – domain logic (dice, wallet, DB helpers).
- `views/` – response formatting and messaging helpers.
- `Dockerfile`, `docker-compose.yaml` – containerization.

## License
Non-Commercial License. You may use, modify, and share this project for personal and non-commercial purposes under the terms in `LICENSE`. Commercial use is not permitted without prior written consent.
