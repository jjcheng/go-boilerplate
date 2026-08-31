BEGIN;
ALTER TABLE account.users DISABLE TRIGGER USER;
DELETE FROM "account"."users" WHERE "users"."id" = 1;
ALTER TABLE account.users ENABLE TRIGGER USER;
-- Reset identity sequences to the previous max id after rollback;
-- This ensures clean state after removing data;
SELECT setval(pg_get_serial_sequence('account.users', 'id'), 0, false);
COMMIT;