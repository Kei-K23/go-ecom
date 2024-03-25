package products

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllProducts() ([]types.Product, error) {
	// Prepare SQL statement
	stmt, err := s.db.Prepare("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize slice to hold products
	var products []types.Product

	// Iterate over the rows
	for rows.Next() {
		// Create a new Product instance to hold the scanned values
		var p types.Product
		// Scan the values from the current row into the Product struct fields
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		// Append the scanned product to the slice of products
		products = append(products, p)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	// Prepare SQL statement
	stmt, err := s.db.Prepare("SELECT * FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query and scan the result into a User struct
	var p types.Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found")
	} else if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Store) CreateProduct(p types.CreateProduct) (*types.CreateProduct, error) {
	var product types.CreateProduct
	stmt, err := s.db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.Quantity)
	if err != nil {
		return nil, err
	}
	product.Name = p.Name
	product.Description = p.Description
	product.Price = p.Price
	product.Quantity = p.Quantity

	return &product, nil
}

func (s *Store) UpdateProduct(p types.CreateProduct, id int) (*types.CreateProduct, error) {
	// Prepare the SQL statement for updating the product
	stmt, err := s.db.Prepare("UPDATE products SET name=?, description=?, price=?, quantity=? WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the update statement with the provided values
	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.Quantity, id)

	if err != nil {
		return nil, err
	}

	// Return the updated product
	return &p, nil
}

func (s *Store) DeleteProduct(id uint) error {
	// Prepare the SQL statement for updating the product
	stmt, err := s.db.Prepare("DELETE TABLE products WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
