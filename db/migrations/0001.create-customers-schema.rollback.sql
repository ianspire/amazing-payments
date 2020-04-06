--
-- file: migrations/0001.create-customers-schema.rollback.sql
--

DROP SCHEMA customer;
DROP FUNCTION customer.update_modified();
