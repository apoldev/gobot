# Telegeam bot example

Included saving users in postgres database

# Start

Set your token in env var `TELEGRAM_BOT`

and

`docker-compose up`

# Env variables:

DB_HOST: 'db'

DB_PORT: '5432'

DB_USER: 'postgres'

DB_PASSWORD: 'example'

DB_NAME: 'gobot' // need created

TELEGRAM_BOT: 'TOKEN'

TELEGRAM_BOT_DEBUG: 'false' | 'true'


