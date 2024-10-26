package armory

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/shopspring/decimal"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

// NewDB for persistence
func NewDB(dsnURI string) (*DB, error) {
	slog.Info("connecting to db", "dsn", dsnURI)
	if dsnURI != ":memory:" {
		dir, _ := filepath.Split(dsnURI)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("create directory: %w", err)
		}
	}
	conn, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	_, err = conn.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return nil, fmt.Errorf("set journal mode: %w", err)
	}

	conn.SetMaxOpenConns(1)

	db := &DB{conn}
	err = db.CreateTable()
	if err != nil {
		return nil, fmt.Errorf("create table: %w", err)
	}
	return db, nil

}

func (d *DB) CreateTable() error {
	_, err := d.Exec(`
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    wallet_address TEXT UNIQUE NOT NULL CHECK(length(address) = 42 AND address LIKE '0x%'),
	encrypted_wallet_private_key TEXT UNIQUE NOT NULL,
    user_address TEXT UNIQUE NOT NULL CHECK(length(address) = 42 AND address LIKE '0x%'),
	created_at INT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS stats (
	total_yield TEXT NOT NULL,
	total_balance TEXT NOT NULL,
	total_balance_seconds TEXT NOT NULL,
	total_deposits TEXT NOT NULL,
	total_withdrawals_queued TEXT NOT NULL,
	total_withdrawals_finalized TEXT NOT NULL,
	total_withdrawals_completed TEXT NOT NULL,
	total_transactions TEXT NOT NULL,
	last_updated INT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS user_stats (
	user_id INTEGER PRIMARY KEY,
	total_yield TEXT NOT NULL,
	total_balance TEXT NOT NULL,
	total_balance_seconds TEXT NOT NULL,
	total_deposits TEXT NOT NULL,
	total_withdrawals_queued TEXT NOT NULL,
	total_withdrawals_finalized TEXT NOT NULL,
	total_withdrawals_completed TEXT NOT NULL,
	total_transactions TEXT NOT NULL,
	last_updated INT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (user_id)
) STRICT;

CREATE TABLE IF NOT EXISTS deposits (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    amount TEXT NOT NULL,
    tx_id TEXT NOT NULL,
    log_index INTEGER NOT NULL,
    block_number INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
) STRICT;

CREATE TABLE IF NOT EXISTS withdrawals (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
	token_id INTEGER NOT NULL,
	amount TEXT NOT NULL,
	withdrawal_requested_id INTEGER NOT NULL,
	requested_at INT NOT NULL,
	withdrawal_finalized_id INTEGER,
	finalized_at INT,
	withdrawal_completed_id INTEGER,
	completed_at INT,
	created_at INT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (user_id),
	FORIEGN KEY (withdrawal_requested_id) REFERENCES withdrawals_requested (id),
	FOREIGN KEY (withdrawal_finalized_id) REFERENCES withdrawals_finalized (id),
	FOREIGN KEY (withdrawal_completed_id) REFERENCES withdrawals_completed (id)
) STRICT;

CREATE TABLE IF NOT EXISTS withdrawals_requested (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
	token_id INTEGER NOT NULL,
    tx_id TEXT NOT NULL,
    log_index INTEGER NOT NULL,
    block_number INT NOT NULL,
	FOREIGN KEY (token_id) REFERENCES tokens (id)
) STRICT;

CREATE TABLE IF NOT EXISTS withdrawals_finalized (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	token_id INTEGER NOT NULL,
	tx_id TEXT NOT NULL,
	log_index INTEGER NOT NULL,
	block_number INT NOT NULL,
	FOREIGN KEY (token_id) REFERENCES tokens (id)
) STRICT;

CREATE TABLE IF NOT EXISTS withdrawals_completed (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
	token_id INTEGER NOT NULL,
    tx_id TEXT NOT NULL,
    log_index INTEGER NOT NULL,
    block_number INT NOT NULL,
	FOREIGN KEY (token_id) REFERENCES tokens (id)
) STRICT;

INSERT INTO stats (total_yield, total_balance, total_balance_seconds, total_deposits, total_withdrawals, total_transactions, last_updated) VALUES ("0", "0", "0", "0", "0", "0", 0) ON CONFLICT DO NOTHING;
`)
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}
	return nil
}

type User struct {
	ID                        int
	WalletAddress             common.Address
	EncryptedWalletPrivateKey string
	UserAddress               common.Address
}

type UserStats struct {
	TotalYield          decimal.Decimal
	TotalBalance        decimal.Decimal
	TotalBalanceSeconds decimal.Decimal
	TotalDeposits       decimal.Decimal
	TotalWithdrawals    decimal.Decimal
	TotalTransactions   string
	LastUpdated         int
}

func (d *DB) UserByAddress(ctx context.Context, userAddress common.Address) (*User, error) {
	result := &User{}
	q := "SELECT id, wallet_address, user_address FROM users WHERE user_address = ?;"
	err := sqlscan.Get(ctx, d, result, q, userAddress)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	return result, nil
}

func (d *DB) UserStatsByAddress(ctx context.Context, userAddress common.Address) (*UserStats, error) {
	result := &UserStats{}
	q := "SELECT total_yield, total_balance, total_balance_seconds, total_deposits, total_withdrawals, total_transactions, last_updated FROM user_stats WHERE user_id = (SELECT id FROM users WHERE user_address = ?);"
	err := sqlscan.Get(ctx, d, result, q, userAddress)
	if err != nil {
		return nil, fmt.Errorf("user stats: %w", err)
	}
	return result, nil
}

func (d *DB) InsertUser(ctx context.Context, userID int, walletAddress, userAddress common.Address, privateKey *ecdsa.PrivateKey) error {
	tx, err := d.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO users (id, wallet_address, encrypted_wallet_private_key, user_address, created_at) VALUES (?, ?, ?, ?, ?);", userID, walletAddress.Hex(), hexutil.EncodeBig(privateKey.D), userAddress.Hex(), 0)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	_, err = tx.Exec("INSERT INTO user_stats (user_id, total_yield, total_balance, total_balance_seconds, total_deposits, total_withdrawals, total_transactions, last_updated) VALUES (?, ?, ?, ?, ?, ?, ?, ?);", userID, "0", "0", "0", "0", "0", "0", 0)
	if err != nil {
		return fmt.Errorf("insert user stats: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *DB) CalculateUserStats(ctx context.Context, userID int) (*UserStats, error) {
	result := &UserStats{}

	// Single query to calculate all required stats
	query := `
		SELECT 
			IFNULL((SELECT SUM(CAST(amount AS REAL)) FROM deposits WHERE user_id = ?), '0') AS total_deposits,
			IFNULL((SELECT SUM(CAST(amount AS REAL)) FROM withdrawals WHERE user_id = ?), '0') AS total_withdrawals,
			IFNULL((SELECT COUNT(*) FROM transactions WHERE user_id = ?), '0') AS total_transactions,
			IFNULL((SELECT SUM(CAST(amount AS REAL)) FROM deposits WHERE user_id = ?) 
				- IFNULL((SELECT SUM(CAST(amount AS REAL)) FROM withdrawals WHERE user_id = ?), 0), '0') AS total_balance,
			IFNULL((SELECT SUM(CAST(balance_seconds AS REAL)) FROM balances WHERE user_id = ?), '0') AS total_balance_seconds,
			IFNULL((SELECT SUM(CAST(total_yield AS REAL)) FROM yields WHERE user_id = ?), '0') AS total_yield,
			? AS last_updated
	`

	err := sqlscan.Get(ctx, d, result, query, userID, userID, userID, userID, userID, userID, userID, time.Now().Unix())
	if err != nil {
		return nil, fmt.Errorf("calculate user stats: %w", err)
	}

	return result, nil
}

func (d *DB) InsertDeposit(ctx context.Context, userID int, amount decimal.Decimal, txID common.Hash, logIndex int, blockNumber int) error {
	return ErrNotImplemented
}

func (d *DB) InsertWithdrawal(ctx context.Context, userID int, amount decimal.Decimal, txID common.Hash, logIndex int, blockNumber int) error {
	return ErrNotImplemented
}

func (d *DB) InsertTransaction(ctx context.Context, userID int, calldataHex string, signedDataHex string, nonce int, createdAt int) error {
	_, err := d.Exec("INSERT INTO transactions (user_id, calldata_hex, signed_data_hex, nonce, created_at) VALUES (?, ?, ?, ?, ?);", userID, calldataHex, signedDataHex, nonce, createdAt)
	if err != nil {
		return fmt.Errorf("insert transaction: %w", err)
	}
	return nil
}

func (d *DB) UpdateUserStats(ctx context.Context, tx *sql.Tx, userID int, totalYield string, totalBalance string, totalBalanceSeconds string, totalDeposits string, totalWithdrawals string, totalTransactions string) error {
	lastUpdated := time.Now().Unix()
	_, err := tx.Exec("UPDATE user_stats SET total_yield = ?, total_balance = ?, total_balance_seconds = ?, total_deposits = ?, total_withdrawals = ?, total_transactions = ?, last_updated = ? WHERE user_id = ?;", totalYield, totalBalance, totalBalanceSeconds, totalDeposits, totalWithdrawals, totalTransactions, lastUpdated, userID)
	if err != nil {
		return fmt.Errorf("upsert user stats: %w", err)
	}
	return nil
}

func (d *DB) UpdateStats(ctx context.Context, totalYield string, totalBalance string, totalBalanceSeconds string, totalDeposits string, totalWithdrawals string, totalTransactions string) error {
	lastUpdated := time.Now().Unix()
	_, err := d.Exec("UPDATE stats SET total_yield = ?, total_balance = ?, total_balance_seconds = ?, total_deposits = ?, total_withdrawals = ?, total_transactions = ?, last_updated = ?;", totalYield, totalBalance, totalBalanceSeconds, totalDeposits, totalWithdrawals, totalTransactions, lastUpdated)
	if err != nil {
		return fmt.Errorf("upsert stats: %w", err)
	}
	return nil
}
