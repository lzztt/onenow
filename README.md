# One Now

## Build Proto
```
buf lint proto
buf format proto -w
buf generate
```

## Build Frontend
```
cd frontend

yarn install
yarn test
yarn start
yarn build
```

## Build Backend
```
cd backend
mkdir -p gen

rsync -av ../note gen/
rsync -av ../frontend/build gen/

go build
```