BEGIN;
ALTER TABLE account.users DISABLE TRIGGER USER;
INSERT INTO "account"."users" ("entry_date","last_update","name","phone_number","email","password_hash","access_token_hash","access_token_expiry","description","type","status","id") VALUES ('2026-08-26 09:26:38.266','2026-08-26 09:39:17.531','test user','9999999999','user@example.com','$2a$10$DK5va.yiZw4CscOW/iMZvuYi6gSTCVVCNz1n.5DlhV2KKQhlD0nee',NULL,NULL,'Test admin user','ADMIN','ACTIVE',1) RETURNING "id";
ALTER TABLE account.users ENABLE TRIGGER USER;
-- Reset identity sequences after inserting records with explicit IDs;
-- This ensures auto-increment starts from the correct value after existing data;
SELECT setval(pg_get_serial_sequence('account.users', 'id'), (SELECT MAX(id) FROM account.users));

COMMIT;
ANALYZE;