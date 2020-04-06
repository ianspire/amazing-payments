package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Datastore interface {
	HealthCheck() error
	InsertCustomer(ctx context.Context, name, email, stripeChargeDate, customerKey string) (*CustomerRecord, error)
	GetCustomer(ctx context.Context, customerID int64) (*CustomerRecord, error)
}

type pgdb struct {
	pgx      *sqlx.DB
	dbLogger *zap.SugaredLogger
}

func NewDB(c *Config, appLogger *zap.SugaredLogger) (Datastore, error) {
	dbConnString := fmt.Sprintf(
		"host=%s port=%v user=%s dbname=%s password=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBName,
		c.DBPassword,
	)

	sqlxPGConn, err := sqlx.Connect("postgres", dbConnString)
	if err != nil {
		log.Printf("dbConnString: %v", dbConnString)
		log.Fatalf("failed to load pgdb connection: %s", err.Error())
		return nil, err
	}

	return &pgdb{pgx: sqlxPGConn, dbLogger: appLogger}, nil
}

type CustomerRecord struct {
	CustomerID        int64     `db:"id"`
	Name              string    `db:"name"`
	Email             string    `db:"email"`
	StripeCustomerKey string    `db:"stripe_customer_key"`
	StripeChargeDate  string    `db:"stripe_charge_date"`
	Created           time.Time `db:"created_at"`
	Modified          time.Time `db:"modified_at"`
}

// HealthCheck verifies that the underlying datastore is working properly
func (db *pgdb) HealthCheck() error {
	return db.pgx.Ping()
}

func (db *pgdb) InsertCustomer(ctx context.Context, name, email, stripeChargeDate, customerKey string) (
	*CustomerRecord, error) {

	var customer CustomerRecord

	// columns inserted on record creation
	customerColumns := []string{
		"name",
		"email",
		"stripe_customer_key",
		"stripe_charge_date",
	}

	// generating SQL statement structure
	query, args, err := squirrel.Insert("customer.customer").
		Columns(customerColumns...).
		Values(name, email, customerKey, stripeChargeDate).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	err = db.pgx.QueryRowxContext(ctx, query, args...).StructScan(&customer)
	if err != nil {
		db.dbLogger.Errorw("failed to scan row",
			"columns", customerColumns,
			"statement", query,
			"error", err,
		)
		return nil, err
	}

	return &customer, err
}

//
func (db *pgdb) GetCustomer(ctx context.Context, customerID int64) (*CustomerRecord, error) {

	var cr CustomerRecord

	if customerID == 0 {
		return nil, errors.New("no customer can be retrieved with ID 0")
	}

	selectColumns := []string{
		"id",
		"name",
		"email",
		"stripe_customer_key",
		"stripe_charge_date",
	}

	selectSQL, args, err := squirrel.Select(selectColumns...).
		From("customer.customer").
		Where(squirrel.Eq{"id": customerID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		db.dbLogger.Errorw("failed to create SQL string",
			"id", customerID,
		)
	}

	db.dbLogger.Infow("selectSQL",
		"statement", selectSQL)

	err = db.pgx.GetContext(ctx, &cr, selectSQL, args...)
	if err != nil {
		db.dbLogger.Errorw("failed to select customer",
			"id", customerID,
			"error", err,
		)
	}

	return &cr, err
}

func (db *pgdb) UpdateCustomer(ctx context.Context, name, email, stripeChargeDate *string) (
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

	err = db.pgx.QueryRowContext(ctx, updateStmt).Scan(&cr)
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
