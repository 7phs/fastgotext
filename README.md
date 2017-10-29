# Init submodule

First - update submodule 

```
git submodule update
```

# Make test model

Make fasttext:

```
cd wrapper/lib/fasttext
make
```

Build a model:
```
wrapper/lib/fasttext/fasttext skipgram -input README.md -output test/model
```

# Build fasttext wrapper

```
go generate ./...
go build ./...
```