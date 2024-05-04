package db

import (
	"database/sql"

	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/application"
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

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {

	exists, err := p.exists(product.GetID())

	if err != nil {
		return nil, err
	}

	if exists {
		return p.update(product)
	}

	return p.create(product)
}

func (p *ProductDb) exists(id string) (bool, error) {
	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", id).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {

	stmt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, errExec := stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if errExec != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {

	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, errExec := stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if errExec != nil {
		return nil, err
	}

	return product, nil
}
