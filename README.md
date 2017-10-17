# Init submodule

First - update submodule 

```
git submodule update
```

# Make test model

Make fasttext:

```
cd lib/fasttext
make
```

Build a model:
```
lib/fasttext/fasttext skipgram -input README.md -output test/model
```

# Build fasttext wrapper

```
go generate ./...
go build
```