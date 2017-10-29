#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.
#

BASE = src/
OBJS = $(BUILD)emd.o
BUILD = build/

CC = cc
CCFLAGS = -fPIC
INCLUDES = -I.

$(shell mkdir -p $(BUILD))

opt: CCFLAGS += -O3 -funroll-loops
opt: build

$(BUILD)emd.o: $(BASE)emd.c $(BASE)*.h
	$(CC) $(CCFLAGS) -c $(BASE)emd.c -o $(BUILD)emd.o

libemd.a: $(OBJS)
	$(AR) rcs libemd.a $(OBJS)

clean:
	rm -rf $(BUILD)*.o libemd.a

build: libemd.a