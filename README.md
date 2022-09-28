# One Now

## Build Proto
```
mkdir -p frontend/src/gen backend/gen

buf lint proto
buf format proto -w
buf generate
```

## Migrate notes from git to files
```
cd note
python3 git_to_file.py | tee git_to_file.log
```

## Build Backend
```
cd backend

go run main.go ../note/*-*_*.md
go build
```

## Build Frontend
```
cd frontend

yarn install
yarn test
yarn start
yarn build
```