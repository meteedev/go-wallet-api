package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}



type Storer interface {
	FindAll() ([]Wallet, error)
	
	FindByWalletType(walletType string) ([]Wallet, error)
	
	FindByWalletId(walletID int) (*Wallet, error)
	
	FindByUserId(userId int) ([]Wallet, error)
	
	Create(wallet *Wallet) (*Wallet, error)
	
	CountByCriteria(criteria Wallet) (int, error)
	
	DeleteByUserId(userId string) (int64, error)
	
	UpdateByWalletId(walletId int,wallet Wallet)(int64,error)
}

type Postgres struct {
	Db *sql.DB
}

func (p *Postgres) FindByWalletType(walletType string) ([]Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type = $1", walletType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}
	return wallets, nil
}

func (p *Postgres) FindAll() ([]Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}
	return wallets, nil
}

func (p *Postgres) FindByUserId(userId int) ([]Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}
	return wallets, nil
}

func (p *Postgres) Create(w *Wallet) (*Wallet, error) {
	row := p.Db.QueryRow("INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance) values ($1, $2, $3, $4, $5) RETURNING id",
		w.UserID,
		w.UserName,
		w.WalletName, w.WalletType,
		w.Balance)

	err := row.Scan(&w.ID)
	if err != nil {
		return nil, err
	}

	return w, err
}

// func (p *Postgres) CountWalletsByCriteria(criteria *Wallet) (int, error) {
//     query := "SELECT count(id) FROM wallets WHERE "
//     var conditions []string
//     var args []interface{}

//     // Iterate over struct fields
//     v := reflect.ValueOf(criteria)
//     for i := 0; i < v.NumField(); i++ {
//         fieldName := v.Type().Field(i).Tag.Get("postgres") // Get the tag value from struct definition
//         fieldValue := v.Field(i).Interface()               // Get the field value
//         if fieldValue != reflect.Zero(v.Field(i).Type()).Interface() {
//             conditions = append(conditions, fmt.Sprintf("%s = $%d", fieldName, len(args)+1))
//             args = append(args, fieldValue)
//         }
//     }

//     query += strings.Join(conditions, " AND ")

//     rows, err := p.Db.Query(query, args...)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var wallets []Wallet
//     for rows.Next() {
//         var wallet Wallet
//         err := rows.Scan(&wallet.ID, &wallet.UserID, &wallet.UserName, &wallet.WalletName, &wallet.WalletType, &wallet.Balance, &wallet.CreatedAt)
//         if err != nil {
//             return nil, fmt.Errorf("error scanning wallet row: %v", err)
//         }
//         wallets = append(wallets, wallet)
//     }
//     if err := rows.Err(); err != nil {
//         return nil, fmt.Errorf("error iterating over wallet rows: %v", err)
//     }

//     return wallets, nil
// }

func (p *Postgres) CountByCriteria(criteria Wallet) (int, error) {
	query := "SELECT count(id) FROM user_wallet WHERE "
	var args []interface{}

	// Iterate over struct fields
	var sb strings.Builder
	v := reflect.ValueOf(criteria)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Tag.Get("postgres") // Get the tag value from struct definition
		fieldValue := v.Field(i).Interface()               // Get the field value
		if fieldValue != reflect.Zero(v.Field(i).Type()).Interface() {
			if sb.Len() > 0 {
				sb.WriteString(" AND ") // Add AND if there are multiple conditions
			}
			sb.WriteString(fmt.Sprintf("%s = $%d", fieldName, len(args)+1))
			args = append(args, fieldValue)
		}
	}

	query += sb.String()

	row := p.Db.QueryRow(query, args...)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil

}

func (p *Postgres) DeleteByUserId(userId string) (int64, error) {

	sql := "DELETE FROM user_wallet WHERE user_id= $1"
	res, err := p.Db.Exec(sql, userId)

	if err != nil {
		return 0, err
	}

	deletedRows, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return deletedRows, nil
}


func (p *Postgres) FindByWalletId(walletID int) (*Wallet, error) {
    // Prepare the SQL query with a placeholder for the wallet ID
    query := `
        SELECT id, user_id, user_name, wallet_name, wallet_type, balance, created_at 
        FROM user_wallet 
        WHERE id = $1 
        LIMIT 1
    `

    // Execute the query using the QueryRow method of the DB object
    row := p.Db.QueryRow(query, walletID)

    // Initialize a Wallet struct
    var wallet Wallet

    // Scan the values returned by the query into the fields of the wallet struct
    err := row.Scan(&wallet.ID, &wallet.UserID, &wallet.UserName, &wallet.WalletName, &wallet.WalletType, &wallet.Balance, &wallet.CreatedAt)
    if err != nil {
        // If no rows are returned, check for sql.ErrNoRows error
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("wallet not found for ID: %d", walletID)
        }
        // Otherwise, return any other error
        return nil, err
    }

    // Return the populated wallet pointer
    return &wallet, nil
}

func (p *Postgres) UpdateByWalletId(walletId int, wallet Wallet) (int64, error) {
    var updates []string
    var args []interface{}

    val := reflect.ValueOf(wallet)
    typ := reflect.TypeOf(wallet)

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        tag := typ.Field(i).Tag.Get("postgres")

        if field.IsZero() || !field.CanInterface() {
            continue
        }

        updates = append(updates, fmt.Sprintf("%s = $%d", tag, len(args)+1))
        args = append(args, field.Interface())
    }

    // Check if any updates are provided
    if len(updates) == 0 {
        return 0, fmt.Errorf("no updates provided")
    }

    // Construct the query string
    query := fmt.Sprintf("UPDATE user_wallet SET %s WHERE id = $%d", strings.Join(updates, ", "), len(args)+1)
    args = append(args, walletId)

    // Execute the query
    res, err := p.Db.Exec(query, args...)
    if err != nil {
        return 0, err
    }

    // Get the number of rows affected
    numRows, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    return numRows, nil
}