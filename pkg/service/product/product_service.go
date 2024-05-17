package product

import (
	"context"

	"github.com/eulabs/back-end/api-product/internal/config/logger"
	"github.com/eulabs/back-end/api-product/pkg/adapter/mysql"
	"github.com/eulabs/back-end/api-product/pkg/model"
)

type ProductServiceInterface interface {
	GetAll(ctx context.Context) *model.ProductList
	GetByID(ctx context.Context, ID int64) *model.Product
	GetByExternalID(ctx context.Context, ExternalID string) *model.Product
	Create(ctx context.Context, product *model.Product) (ExternalID string)
	Update(ctx context.Context, ExternalID string, product *model.Product) int64
	Delete(ctx context.Context, ExternalID string) int64
}

type product_service struct {
	dbp mysql.DatabaseInterface
}

func NewProductService(database_pool mysql.DatabaseInterface) *product_service {
	return &product_service{
		dbp: database_pool,
	}
}

func (ps *product_service) GetAll(ctx context.Context) *model.ProductList {
	rows, err := ps.dbp.GetDB().QueryContext(ctx, "SELECT id, external_id, name, quantity, price FROM tb_product LIMIT 100")
	if err != nil {
		logger.Error(err.Error(), err)
	}

	defer rows.Close()

	product_list := &model.ProductList{}

	for rows.Next() {
		p := model.Product{}
		if err := rows.Scan(&p.ID, &p.ExternalID, &p.Name, &p.Quantity, &p.Price); err != nil {
			logger.Error(err.Error(), err)
		} else {
			product_list.List = append(product_list.List, &p)
		}
	}

	return product_list
}

func (ps *product_service) GetByID(ctx context.Context, ID int64) *model.Product {
	stmt, err := ps.dbp.GetDB().PrepareContext(ctx, "SELECT id, external_id, name, quantity, price FROM tb_product WHERE external_id = ?")
	if err != nil {
		logger.Error(err.Error(), err)
	}

	defer stmt.Close()

	p := model.Product{}

	if err := stmt.QueryRowContext(ctx, ID).Scan(&p.ID, &p.ExternalID, &p.Name, &p.Quantity, &p.Price); err != nil {
		logger.Error(err.Error(), err)
	}

	return &p
}

func (ps *product_service) GetByExternalID(ctx context.Context, ExternalID string) *model.Product {
	stmt, err := ps.dbp.GetDB().PrepareContext(ctx, "SELECT id, external_id, name, quantity, price FROM tb_product WHERE external_id = ?")
	if err != nil {
		logger.Error(err.Error(), err)
	}

	defer stmt.Close()

	p := model.Product{}

	if err := stmt.QueryRowContext(ctx, ExternalID).Scan(&p.ID, &p.ExternalID, &p.Name, &p.Quantity, &p.Price); err != nil {
		logger.Error(err.Error(), err)
	}

	return &p
}

func (ps *product_service) Create(ctx context.Context, product *model.Product) (ExternalID string) {

	tx, err := ps.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error(err.Error(), err)

		return ""
	}

	query := "INSERT INTO tb_product (external_id, name, quantity, price) VALUES (?, ?, ?, ?);"

	_, err = tx.ExecContext(ctx, query, product.ExternalID, product.Name, product.Quantity, product.Price)
	if err != nil {
		logger.Error("erro to Exec SQL Query", err)

		return ""
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("erro to Commit", err)

		return ""
	} else {
		logger.Info("Insert Transaction committed")
	}

	return product.ExternalID
}

func (ps *product_service) Update(ctx context.Context, ExternalID string, product *model.Product) int64 {
	tx, err := ps.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("erro to Begin trasaction", err)

	}

	query := "UPDATE tb_product SET name = ?, quantity = ?, price = ? WHERE external_id = ?;"

	result, err := tx.ExecContext(ctx, query, product.Name, product.Quantity, product.Price, ExternalID)
	if err != nil {
		logger.Error("erro to upodate product", err)
		return 0
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("erro to Commit", err)
		tx.Rollback()
		return 0
	} else {
		logger.Info("Update Transaction committed")
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		logger.Error("erro to get rows updated", err)
		return 0
	}

	return rowsaff
}

func (ps *product_service) Delete(ctx context.Context, ExternalID string) int64 {
	tx, err := ps.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("erro to start transaction", err)

	}

	query := "DELETE FROM tb_product WHERE external_id = ?"

	result, err := tx.ExecContext(ctx, query, ExternalID)
	if err != nil {
		logger.Error("erro to delete product", err)
		return 0
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("erro to commit", err)
		tx.Rollback()
		return 0
	} else {
		logger.Info("Delete Transaction committed")
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		logger.Error("erro nome record inserted", err)
		return 0
	}

	return rowsaff
}
