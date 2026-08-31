BEGIN;
DROP SCHEMA "account" CASCADE;
DROP TABLE "account"."users";
DROP TRIGGER users_set_last_update ON account.users;
COMMIT;