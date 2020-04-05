--
-- file: migrations/0001.create-customers-schema.sql
--

-- create schema for customer data
CREATE SCHEMA customer;

-- create a method we can combine with a trigger for each table to automatically update modified_at field
CREATE OR REPLACE FUNCTION customer.update_modified() returns trigger as
     $$begin
       new.modified_at = now();
         return new;
     end;$$
     language plpgsql;
