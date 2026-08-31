## Modern Go project skeleton

This project is the boilerplate Go project I use in every new Go project. Some of the features are:
1. Web framework using Gin
2. Database using Gorm + Postgres + Repository
3. File storaging using Alibaba Cloud OSS
4. Web hosting using Alibaba Cloud Function Compute
5. Logging using slog
6. API doc auto generated everytime a project is build in Develop environment
7. Auto DB migration included
8. Deploy to Function Compute with script
9. Deploy production DB with script
10. Environment variables segrated by local and live
11. Cli and Worker process included
12. Authentication with user access token


## before committing to develop branch

`make migration-files`

This generates migration SQL files in `migration/`; it does not update a remote database. The CLI creates a temporary comparison database named by `MigrationName`, applies the committed migration files, compares it with the local database, and writes new schema/data files only when changes are found.

Staging uses the same committed `migration/` history. `deploy_staging_db.sh` obtains its remote connection settings from `DB_HOST_EXTERNAL`, `DB_USER_EXTERNAL`, `DB_PASSWORD_EXTERNAL`, `DB_NAME`, `DB_PORT`, and `DB_SSLMODE` in `.env.staging` or `.env`; it never contains database credentials in source.


## grant permissions to the web db user

-- Grant usage on the schema
GRANT USAGE ON SCHEMA public TO web_user;

-- Grant select, insert, update, delete on all existing tables
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO web_user;

-- Grant usage/select on all sequences (needed for serial/identity columns)
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO web_user;

-- Ensure future tables also have these permissions
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO web_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON SEQUENCES TO web_user;

## deploy fc
if there is a change of fc_user in RAM, call 
`aliyun configure` with the latest ak and sk
use `aliyun configure list` to get list of users used

## https for localhost
brew install ngrok
ngrok config add-authtoken xxx (in .env.staging NGROK_AUTH_TOKEN)
ngrok http 9000