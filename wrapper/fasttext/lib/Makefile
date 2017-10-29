#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.
#

BASE = fasttext/
OBJS = $(BUILD)args.o $(BUILD)dictionary.o $(BUILD)productquantizer.o $(BUILD)matrix.o $(BUILD)qmatrix.o $(BUILD)vector.o $(BUILD)model.o $(BUILD)utils.o $(BUILD)fasttext.o $(BUILD)fsttxt_wrapper.o
BUILD = build/

CXX = c++
CXXFLAGS = -pthread -std=c++0x
INCLUDES = -I.

$(shell mkdir -p $(BUILD))

opt: CXXFLAGS += -O3 -funroll-loops
opt: build

$(BUILD)args.o: $(BASE)src/args.cc $(BASE)src/args.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/args.cc -o $(BUILD)args.o

$(BUILD)dictionary.o: $(BASE)src/dictionary.cc $(BASE)src/dictionary.h $(BASE)src/args.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/dictionary.cc -o $(BUILD)dictionary.o

$(BUILD)productquantizer.o: $(BASE)src/productquantizer.cc $(BASE)src/productquantizer.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/productquantizer.cc -o $(BUILD)productquantizer.o

$(BUILD)matrix.o: $(BASE)src/matrix.cc $(BASE)src/matrix.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/matrix.cc -o $(BUILD)matrix.o

$(BUILD)qmatrix.o: $(BASE)src/qmatrix.cc $(BASE)src/qmatrix.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/qmatrix.cc -o $(BUILD)qmatrix.o

$(BUILD)vector.o: $(BASE)src/vector.cc $(BASE)src/vector.h $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/vector.cc -o $(BUILD)vector.o

$(BUILD)model.o: $(BASE)src/model.cc $(BASE)src/model.h $(BASE)src/args.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/model.cc -o $(BUILD)model.o

$(BUILD)utils.o: $(BASE)src/utils.cc $(BASE)src/utils.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/utils.cc -o $(BUILD)utils.o

$(BUILD)fasttext.o: $(BASE)src/fasttext.cc $(BASE)src/*.h
	$(CXX) $(CXXFLAGS) -c $(BASE)src/fasttext.cc -o $(BUILD)fasttext.o

$(BUILD)fsttxt_wrapper.o: src/fsttxt_wrapper.cc
	$(CXX) $(CXXFLAGS) -c src/fsttxt_wrapper.cc -o $(BUILD)fsttxt_wrapper.o

libfasttext.a: $(OBJS)
	$(AR) rcs libfasttext.a $(OBJS)

clean:
	rm -rf $(BUILD)*.o libfasttext.a

build: libfasttext.a