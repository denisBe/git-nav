#!/bin/sh

#
# Install libgit2 to git2go in dynamic mode on Travis
#

set -ex

cd "${HOME}"
wget -O libgit2-0.23.1.tar.gz https://github.com/libgit2/libgit2/archive/v0.23.1.tar.gz
tar -xzvf libgit2-0.23.1.tar.gz
cd libgit2-0.23.1 && mkdir build && cd build
cmake -DTHREADSAFE=ON -DBUILD_CLAR=OFF -DCMAKE_BUILD_TYPE="RelWithDebInfo" .. && make && sudo make install
sudo ldconfig
cd "${TRAVIS_BUILD_DIR}"
