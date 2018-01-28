The Golang wrapper of FaceBook [Fasttext](https://github.com/facebookresearch/fastText)
and [EMD idea and source code](http://ai.stanford.edu/%7Erubner/emd/default.htm).

The primary purpose of the project is an experiment for a bridge between Golang and native.

# Init submodule

First - init and update submodule.

```
git submodule update --init --recursive

git submodule foreach git pull
```

# Make test model

Make fasttext:

```
cd wrapper/lib/fasttext
make
```

Build an unsupervised model:

```
wrapper/fasttext/lib/fasttext/fasttext skipgram -input unsupervised_text.txt -output test-data/unsupervised_model
```

Build a supervised model:

```
wrapper/fasttext/lib/fasttext/fasttext supervised -input supervised_text.txt -output test-data/supervised_model
```

# Build fasttext wrapper

```
go generate ./...
go build ./...
```

# Benchmark for Emd

```
iMac-Aleksej:emd alexey$ go test -bench=.
goos: darwin
goarch: amd64
pkg: bitbucket.org/7phs/fastgotext/wrapper/emd
BenchmarkEmd/Emd/10-8              50000             22396 ns/op
BenchmarkEmd/Emd/20-8              20000             71225 ns/op
BenchmarkEmd/Emd/30-8              10000            162125 ns/op
BenchmarkEmd/Emd/40-8               5000            332387 ns/op
BenchmarkEmd/Emd/50-8               3000            486507 ns/op
BenchmarkEmd/Emd/60-8               2000            657931 ns/op
BenchmarkEmd/Emd/70-8               2000            920516 ns/op
BenchmarkEmd/Emd/80-8               1000           1432798 ns/op
BenchmarkEmd/Emd/90-8               1000           2279080 ns/op
BenchmarkEmd/Emd/100-8              1000           1963907 ns/op
PASS
ok      bitbucket.org/7phs/fastgotext/wrapper/emd       18.945s
```

Test as just marshaling.

```
iMac-Aleksej:emd alexey$ go test -bench=.
goos: darwin
goarch: amd64
pkg: bitbucket.org/7phs/fastgotext/wrapper/emd
BenchmarkEmd/Emd/10-8             100000             18632 ns/op
BenchmarkEmd/Emd/20-8              20000             57830 ns/op
BenchmarkEmd/Emd/30-8              10000            128437 ns/op
BenchmarkEmd/Emd/40-8              10000            215187 ns/op
BenchmarkEmd/Emd/50-8               5000            329511 ns/op
BenchmarkEmd/Emd/60-8               3000            467195 ns/op
BenchmarkEmd/Emd/70-8               2000            628742 ns/op
BenchmarkEmd/Emd/80-8               2000            818129 ns/op
BenchmarkEmd/Emd/90-8               2000           1030549 ns/op
BenchmarkEmd/Emd/100-8              1000           1288463 ns/op
PASS
ok      bitbucket.org/7phs/fastgotext/wrapper/emd       17.083s
```

Test with native array.

```
go test -bench=.
goos: darwin
goarch: amd64
pkg: bitbucket.org/7phs/fastgotext/wrapper/emd
BenchmarkEmd/Emd/10-8         	  300000	      3768 ns/op
BenchmarkEmd/Emd/20-8         	  100000	     11023 ns/op
BenchmarkEmd/Emd/30-8         	   50000	     57476 ns/op
BenchmarkEmd/Emd/40-8         	   20000	     86323 ns/op
BenchmarkEmd/Emd/50-8         	   10000	    105847 ns/op
BenchmarkEmd/Emd/60-8         	   10000	    274646 ns/op
BenchmarkEmd/Emd/70-8         	    3000	    575674 ns/op
BenchmarkEmd/Emd/80-8         	    3000	    529006 ns/op
BenchmarkEmd/Emd/90-8         	    2000	    657821 ns/op
BenchmarkEmd/Emd/100-8        	    2000	    695739 ns/op
PASS
ok  	bitbucket.org/7phs/fastgotext/wrapper/emd	19.171s
```

Tes as just native marshaling.

```
go test -tags="dumb" -bench=.
goos: darwin
goarch: amd64
pkg: bitbucket.org/7phs/fastgotext/wrapper/emd
BenchmarkEmd/Emd/10-8         	 1000000	      1354 ns/op
BenchmarkEmd/Emd/20-8         	 1000000	      1376 ns/op
BenchmarkEmd/Emd/30-8         	 1000000	      1407 ns/op
BenchmarkEmd/Emd/40-8         	 1000000	      1448 ns/op
BenchmarkEmd/Emd/50-8         	 1000000	      1492 ns/op
BenchmarkEmd/Emd/60-8         	 1000000	      1526 ns/op
BenchmarkEmd/Emd/70-8         	 1000000	      1576 ns/op
BenchmarkEmd/Emd/80-8         	 1000000	      1629 ns/op
BenchmarkEmd/Emd/90-8         	 1000000	      1613 ns/op
BenchmarkEmd/Emd/100-8        	 1000000	      1704 ns/op
PASS
ok  	bitbucket.org/7phs/fastgotext/wrapper/emd	15.308s
```

Fasttext links:

* [Train and Test Supervised Text Classifier using fasttext](https://www.tutorialkart.com/fasttext/train-and-test-supervised-text-classifier-using-fasttext/)
