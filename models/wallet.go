package models

import (
    "bytes"
    "crypto/ed25519"
    "crypto/rand"
    "database/sql"
    "encoding/json"
    "errors"
    "fmt"
    "math/big"
    "net/http"
    "os"
    "strings"
    "time"
)

type Wallet struct {
    PublicKey  ed25519.PublicKey
    PrivateKey ed25519.PrivateKey
}

var (
    base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
)

func EnsureWallet(userID int64) *Wallet {
    if w, ok := GetWallet(userID); ok {
        return w
    }
    pub, priv, _ := ed25519.GenerateKey(rand.Reader)
    if DB() != nil {
        _, _ = DB().Exec(`INSERT INTO wallets (user_id, pub_key, priv_key) VALUES ($1, $2, $3)`,
            userID, []byte(pub), []byte(priv))
    }
    return &Wallet{PublicKey: pub, PrivateKey: priv}
}

func GetWallet(userID int64) (*Wallet, bool) {
    if DB() == nil {
        return nil, false
    }
    var pub, priv []byte
    err := DB().QueryRow(`SELECT pub_key, priv_key FROM wallets WHERE user_id=$1`, userID).Scan(&pub, &priv)
    if err == sql.ErrNoRows || err != nil {
        return nil, false
    }
    w := &Wallet{PublicKey: ed25519.PublicKey(pub), PrivateKey: ed25519.PrivateKey(priv)}
    return w, true
}

func (w *Wallet) AddressBase58() string { return Base58Encode(w.PublicKey) }

func (w *Wallet) PrivateKeyBase58() string { return Base58Encode(w.PrivateKey) }

func Base58Encode(input []byte) string {
    if len(input) == 0 {
        return ""
    }
    zeros := 0
    for zeros < len(input) && input[zeros] == 0 {
        zeros++
    }
    var intVal = new(big.Int).SetBytes(input)
    var mod = new(big.Int)
    var fiftyEight = big.NewInt(58)
    var encoded []byte
    for intVal.Sign() > 0 {
        intVal, mod = new(big.Int).DivMod(intVal, fiftyEight, mod)
        encoded = append(encoded, base58Alphabet[mod.Int64()])
    }
    for i := 0; i < zeros; i++ {
        encoded = append(encoded, base58Alphabet[0])
    }
    for i, j := 0, len(encoded)-1; i < j; i, j = i+1, j-1 {
        encoded[i], encoded[j] = encoded[j], encoded[i]
    }
    return string(encoded)
}

func (w *Wallet) GetBalance() (uint64, error) {
    rpc := os.Getenv("SOLANA_RPC_URL")
    rpc = strings.TrimSpace(strings.Trim(rpc, "\""))
    if rpc == "" {
        return 0, errors.New("SOLANA_RPC_URL is not configured")
    }
    addr := w.AddressBase58()
    // First try getBalance (most providers)
    {
        payload := map[string]any{
            "jsonrpc": "2.0",
            "id":      1,
            "method":  "getBalance",
            "params":  []any{addr},
        }
        body, _ := json.Marshal(payload)
        req, err := http.NewRequest("POST", rpc, bytes.NewReader(body))
        if err != nil {
            return 0, err
        }
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Accept", "application/json")
        client := &http.Client{Timeout: 10 * time.Second}
        resp, err := client.Do(req)
        if err != nil {
            return 0, err
        }
        defer resp.Body.Close()
        var raw bytes.Buffer
        if resp.StatusCode < 200 || resp.StatusCode >= 300 {
            _, _ = raw.ReadFrom(resp.Body)
            msg := raw.String()
            if len(msg) > 300 {
                msg = msg[:300]
            }
            return 0, fmt.Errorf("rpc http status %d: %s", resp.StatusCode, strings.TrimSpace(msg))
        }
        var res struct {
            JSONRPC string `json:"jsonrpc"`
            ID      int    `json:"id"`
            Result  struct {
                Value uint64 `json:"value"`
            } `json:"result"`
            Error *struct {
                Code    int    `json:"code"`
                Message string `json:"message"`
            } `json:"error"`
        }
        if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
            return 0, err
        }
        if res.Error == nil {
            return res.Result.Value, nil
        }
        // If method not found, try fallback
        if res.Error.Code != -32601 && !strings.Contains(strings.ToLower(res.Error.Message), "method") {
            return 0, errors.New(res.Error.Message)
        }
    }

    // Fallback to getAccountInfo and read lamports
    payload := map[string]any{
        "jsonrpc": "2.0",
        "id":      1,
        "method":  "getAccountInfo",
        "params":  []any{addr, map[string]any{"encoding": "jsonParsed", "commitment": "finalized"}},
    }
    body, _ := json.Marshal(payload)
    req, err := http.NewRequest("POST", rpc, bytes.NewReader(body))
    if err != nil {
        return 0, err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        var buf bytes.Buffer
        _, _ = buf.ReadFrom(resp.Body)
        msg := buf.String()
        if len(msg) > 300 {
            msg = msg[:300]
        }
        return 0, fmt.Errorf("rpc http status %d: %s", resp.StatusCode, strings.TrimSpace(msg))
    }
    var res2 struct {
        JSONRPC string `json:"jsonrpc"`
        ID      int    `json:"id"`
        Result  struct {
            Value *struct {
                Lamports uint64 `json:"lamports"`
            } `json:"value"`
        } `json:"result"`
        Error *struct {
            Code    int    `json:"code"`
            Message string `json:"message"`
        } `json:"error"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&res2); err != nil {
        return 0, err
    }
    if res2.Error != nil {
        return 0, errors.New(res2.Error.Message)
    }
    if res2.Result.Value == nil {
        return 0, nil
    }
    return res2.Result.Value.Lamports, nil
}

func BuildSolanaPayURI(address string, amount string) string {
    if amount == "" {
        return "solana:" + address
    }
    return "solana:" + address + "?amount=" + amount
}
