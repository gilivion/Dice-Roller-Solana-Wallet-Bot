---
title: Dice Roller + Solana Wallet Bot
layout: default
---

# Dice Roller + Solana Wallet Bot

Telegram bot for rolling dice in chats and providing simple Solana wallet utilities in DM.

Links and CA:
- Site: https://gilivion.github.io/Dice-Roller-Solana-Wallet-Bot/
- Telegram: https://t.me/Dices_sol_bot (@Dices_sol_bot)
- X Community: https://twitter.com/i/communities/1964419206939890018
- CA token (Solana): `9GPcQGzqRsxHAc1J7teBCZrJxvNAcfMevYz1vqtZ28Ra`

## What it does
- Group chats: `/roll`, `/help`.
- Direct messages: `/roll`, `/wallet_address`, `/wallet_balance`, `/wallet_private_key`, `/help`.
- One Solana wallet per Telegram user, stored in PostgreSQL.

## Roadmap (Fun Only)
This Solana project is built purely for fun. Planned features:
- `/game` — simple mini-games with SOL tokens and stakes.
- `/wallet_qr` — generate a QR code for your wallet address.

> Note: exposing private keys is dangerous. Consider using this bot for testing/devnet only.

## Run It
See the README for full instructions. In short:

```
docker compose up -d --build
```

Configure environment in `.env`:

```
BOT_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
SOLANA_RPC_URL=https://solana.drpc.org
DATABASE_URL=postgres://diceroller:password@postgres:5432/diceroller?sslmode=disable
```

## License
Non-Commercial License — free for personal and non-commercial use. See the LICENSE file in the repository for details.

## Links
- Source code: Refer to the repository README.
