package data

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"finy.be/api/structs"
	_ "github.com/mattn/go-sqlite3"
)


var TransactionsTableDef string = `
CREATE TABLE Transactions (
	Id TEXT NOT NULL PRIMARY KEY,
	Name TEXT NOT NULL,
	Amount INT NOT NULL,
	Date TEXT NOT NULL,
	CreatedAt INTEGER NOT NULL, -- unix epoch in ms
	UpdatedAt INTEGER NOT NULL, -- unix epoch in ms
	DeletedAt INTEGER -- unix epoch in ms
);`

func ConnectDatabase(name string) (*sql.DB, error) {
	connStr := fmt.Sprintf("file:%s.db?cache=shared", name)
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)

	return db, nil
}

func CreateTable(db *sql.DB) error {
	result, err := db.Exec(TransactionsTableDef)
	if err != nil {
		return err
	}

	_ = result

	return nil
}

func EnsureDbExists(name string) error {
	_, err := os.Stat(fmt.Sprintf("%s.db", name))
	if errors.Is(err, os.ErrNotExist) {
		db, err := ConnectDatabase(name)
		if err != nil {
			return err
		}
		defer db.Close()

		return CreateTable(db)
	}

	return err
}


func SelectQuery(db *sql.DB, query string, transactions *[]structs.Transaction) error {
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate over the result rows.
    for rows.Next() {
		var t structs.Transaction
		err := rows.Scan(&t.Id, &t.Name, &t.Amount, &t.Date,  &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
        if err != nil {
            return err
        }
        *transactions = append(*transactions, t)
    }
	
	if err := rows.Err(); err != nil {
        return err
    }

	return nil
}



func InsertMany(db *sql.DB, transactons *[]structs.Transaction) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	insertSQL := `INSERT INTO Transactions (Id, Name, Amount, Date, CreatedAt, UpdatedAt, DeletedAt) VALUES (?, ?, ?, ?, ?, ?, ?)`
    stmt, err := tx.Prepare(insertSQL)
    if err != nil {
		return err 
    }
    defer stmt.Close()

    for _, t := range *transactons {
        _, err = stmt.Exec(t.Id, t.Name, t.Amount, t.Date, t.CreatedAt, t.UpdatedAt, t.DeletedAt)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    err = tx.Commit()
	return err
}

func DeleteTransaction(db *sql.DB, id string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	statement := "DELETE FROM Transactions WHERE id = ?;"
	stmt, err := tx.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	
	err = tx.Commit()
	return err
}

// func UpdateMany(db *sql.DB, fileDefs *[]models.FileDef) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
//
// 	updateSQL := `UPDATE FileDef SET Path = ?, Hash = ?, Kind = ?, Size = ?, CreatedAt = ?, LastUpdate = ? WHERE Id = ?`
// 	stmt, err := tx.Prepare(updateSQL)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	defer stmt.Close()
//
// 	for _, fileDef := range *fileDefs {
// 		_, err = stmt.Exec(fileDef.Path, fileDef.Hash, fileDef.Kind, fileDef.Size, fileDef.CreatedAt, fileDef.LastUpdate, fileDef.Id)
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}
//
// 	err = tx.Commit()
// 	return err
// }
//
//
// func GetHashesById(db *sql.DB, ids []string, fileDefMap *map[string]string) error {
// 	selectSQL := "SELECT id, hash FROM FileDef WHERE id IN (?" + strings.Repeat(", ?", len(ids)-1) + ")"
// 	stmt, err := db.Prepare(selectSQL)
// 	if err != nil {
// 		return err
// 	}
//
// 	defer stmt.Close()
//
// 	// Convert []string to []interface{}.
//     var idInterfaces []interface{}
//     for _, id := range ids {
//         idInterfaces = append(idInterfaces, id)
//     }
//
// 	rows, err := stmt.Query(idInterfaces...)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
//
// 	// Iterate over the result rows.
//     for rows.Next() {
//         var id, hash string
//         if err := rows.Scan(&id, &hash); err != nil {
//             return err
//         }
//         (*fileDefMap)[id] = hash
//     }
// 	
// 	if err := rows.Err(); err != nil {
//         return err
//     }
//
//     return nil
// }
//
//
// func SearchFiles(db *sql.DB, filter models.SearchFilter, fileDefs *[]models.FileDef) error {
// 	searchSQL := "SELECT * FROM FileDef WHERE path LIKE ? LIMIT ?;"
// 	stmt, err := db.Prepare(searchSQL)
// 	if err != nil {
// 		return err
// 	}
//
// 	defer stmt.Close()
//
// 	rows, err := stmt.Query("%" + filter.Name + "%", cap(*fileDefs))
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
//
// 	// Iterate over the result rows.
//     for rows.Next() {
// 		var fileDef models.FileDef
// 		err := rows.Scan(&fileDef.Id, &fileDef.Hash, &fileDef.Path, &fileDef.Kind, &fileDef.Size, &fileDef.CreatedAt, &fileDef.LastUpdate)
//         if err != nil {
//             return err
//         }
//         *fileDefs = append(*fileDefs, fileDef)
//     }
// 	
// 	if err := rows.Err(); err != nil {
//         return err
//     }
//
// 	return nil
// }
//
