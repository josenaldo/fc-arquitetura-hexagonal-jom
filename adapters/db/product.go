package db

import (
	"database/sql"

	"github.com/josenaldo/fc-arquitetura-hexagobal-jom/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	smtp, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer smtp.Close()

	err = smtp.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, application.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil

}
