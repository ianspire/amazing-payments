--
-- file: migrations/20200331_01_Eksxi-create-customer-table.sql
--

-- create customer table for storing primary customer information
CREATE TABLE customer.customer (
    id serial PRIMARY KEY,
    name text,
    email text NOT NULL,
    stripe_customer_key text,
    stripe_charge_date text,
    created_at timestamp DEFAULT NOW(),
    modified_at timestamp DEFAULT NOW()
);

-- add our update_modified trigger created in a prior step
CREATE TRIGGER customer_update BEFORE UPDATE ON customer.customer
   FOR EACH ROW EXECUTE PROCEDURE customer.update_modified()
;
