# One Now

## Build Proto

```bash
mkdir -p frontend/src/gen backend/gen

buf lint proto
buf format proto -w
buf generate
```

## Create `localhost` certificate

Install [mkcert](https://github.com/FiloSottile/mkcert), then create `localhost` certificate.

```bash
mkdir cert && cd cert

mkcert localhost
```

## Note files

Note filenames have `<snowflake_id>_<timestamp>.md` pattern.

A note can have multiple versions. `snowflake_id` identifies a note. `timestamp` identifies a version of a note.

Commands to generate 10 dummy notes:

```bash
mkdir note && cd note

for i in {1..10}; do
    file=${i}_`date +%s`.md
    echo -e "Title $i\n\nBody $i" > $file;
done
```

## Build Backend

```bash
cd backend

echo 'ALLOWED_EMAIL=test@test.com' > .env.development.local
go run .

go build
```

## Build Frontend

```bash
cd frontend

yarn install
yarn test
yarn start

yarn build
```
