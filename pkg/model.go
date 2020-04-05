package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type PGDB struct {
	pg       *sqlx.DB
	dbLogger *zap.SugaredLogger
}

func NewDB(c *Config, appLogger *zap.SugaredLogger) (*PGDB, error) {
	dbConnString := fmt.Sprintf(
		"host=%s port=%v user=%s dbname=%s password=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBName,
		c.DBPassword,
	)

	sqlxPGConn, err := sqlx.Open("postgres", dbConnString)
	if err != nil {
		log.Printf("dbConnString: %v", dbConnString)
		log.Fatalf("failed to load PGDB connection: %s", err.Error())
		return nil, err
	}

	return &PGDB{pg: sqlxPGConn, dbLogger: appLogger}, nil
}

type CustomerRecord struct {
	CustomerID int64    `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	StripeCustomerKey string `db:"stripe_customer_key"`
	StripeChargeDate string `db:"stripe_charge_date"`
	Created    time.Time `db:"created_at"`
	Modified   time.Time `db:"modified_at"`
}

// HealthCheck verifies that the underlying datastore is working properly
func (db *PGDB) HealthCheck() error {
	return db.pg.Ping()
}

func (db *PGDB) InsertCustomer(ctx context.Context, name, email, stripeChargeDate, customerKey string) (
	*CustomerRecord, error) {

	var customer *CustomerRecord

	// columns inserted on record creation
	customerColumns := []string{
		"name",
		"email",
		"stripe_customer_key",
		"stripe_charge_date",
	}

	// generating SQL statement structure
	query := squirrel.Insert("customer.customer").
		Columns(customerColumns...).
		Values(name, email, customerKey, stripeChargeDate).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(customer)
	// Use prepared statement for input sanitization
	if err != nil {
		db.dbLogger.Errorw("failed to prepare SQL string",
			"columns", customerColumns,
		)
		return nil, err
	}

	return customer, err
}

//
func (db *PGDB) GetCustomer(ctx context.Context, customerID int64) (*CustomerRecord, error) {

	var cr CustomerRecord

	if customerID == 0 {
		return nil, errors.New("no customer can be retrieved with ID 0")
	}

	selectColumns := []string{
		"id",
		"name",
		"email",
		"stripe_customer_id",
		"stripe_charge_date",
	}

	selectSQL, _, err := squirrel.Select(selectColumns...).
		From("customer.customer").
		Where(squirrel.Eq{"id": customerID}).ToSql()
	if err != nil {
		db.dbLogger.Errorw("failed to create SQL string",
			"id", customerID,
		)
	}

	err = db.pg.SelectContext(ctx, &cr, selectSQL, customerID)
	if err != nil {
		db.dbLogger.Errorw("failed to select customer",
			"id", customerID,
			)
	}

	return &cr, err
}

func (db *PGDB) UpdateCustomer(ctx context.Context, name, email, stripeChargeDate *string) (
	*CustomerRecord, error) {

	var cr CustomerRecord

	if name == nil && email == nil && stripeChargeDate == nil {
		return nil, errors.New("no values to update")
	}

	updateSQL := squirrel.Update("customer.customer")
	if name != nil {
		updateSQL.Set("name", &name)
	}
	if email != nil {
		updateSQL.Set("email", &email)
	}
	if stripeChargeDate != nil {
		updateSQL.Set("stripe_charge_date", &stripeChargeDate)
	}

	updateStmt, _, err := updateSQL.ToSql()
	if err != nil {
		db.dbLogger.Errorw("failed to create SQL string",
			"name", name,
			"email", email,
			"stripe_charge_date", stripeChargeDate,
		)
		return nil, err
	}

	err = db.pg.QueryRowContext(ctx, updateStmt).Scan(&cr)
	if err != nil {
		db.dbLogger.Errorw("failed to update DB record",
			"name", name,
			"email", email,
			"stripe_charge_date", stripeChargeDate,
		)
		return nil, err
	}

	return &cr, err
}