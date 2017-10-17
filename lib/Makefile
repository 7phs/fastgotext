#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.
#

CXX = c++
CXXFLAGS = -pthread -std=c++0x
OBJS = args.o dictionary.o productquantizer.o matrix.o qmatrix.o vector.o model.o utils.o fasttext.o wrapper.o
INCLUDES = -I.
BASE = fasttext/

opt: CXXFLAGS += -O3 -funroll-loops
opt: build

args.o: $(BASE)src/args.cc $(BASE)src/args.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/args.cc

dictionary.o: $(BASE)src/dictionary.cc $(BASE)src/dictionary.h $(BASE)src/args.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/dictionary.cc

productquantizer.o: $(BASE)src/productquantizer.cc $(BASE)src/productquantizer.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/productquantizer.cc

matrix.o: $(BASE)src/matrix.cc $(BASE)src/matrix.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/matrix.cc

qmatrix.o: $(BASE)src/qmatrix.cc $(BASE)src/qmatrix.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/qmatrix.cc

vector.o: $(BASE)src/vector.cc $(BASE)src/vector.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/vector.cc

model.o: $(BASE)src/model.cc $(BASE)src/model.h $(BASE)src/args.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/model.cc

utils.o: $(BASE)src/utils.cc $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/utils.cc

fasttext.o: $(BASE)src/fasttext.cc $(BASE)src/*.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/fasttext.cc

wrapper.o: src/wrapper.cc
	$(CXX) $(CXXFLAGS) -c src/wrapper.cc

libfasttext.a: $(OBJS)
	$(AR) rcs libfasttext.a $(OBJS)

clean:
	rm -rf *.o libfasttext.a

build: libfasttext.a