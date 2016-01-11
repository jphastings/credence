apt-get install -y \
  pkg-config \
  libsodium-dev \
  autoconf \
  automake \
  libtool \
  curl \
  unzip \
  libssl-dev \
  uuid-dev \
  gettext-base

cd /tmp
curl -sSL https://github.com/google/protobuf/archive/master.zip > protobuf.zip
unzip protobuf.zip
cd protobuf-master
./autogen.sh
./configure
make
make install && ldconfig

go get github.com/golang/protobuf/protoc-gen-go

ln -s /usr/bin/libtoolize /usr/bin/libtool

# install libzmq
cd ../
curl -sSL https://github.com/zeromq/libzmq/archive/master.zip > zmq.zip
unzip zmq.zip
cd libzmq-master
./autogen.sh
./configure
make -j 4
make install && ldconfig

# install libczmq
cd ../
curl -sSL https://github.com/zeromq/czmq/archive/v3.0.2.zip > czmq.zip
unzip czmq.zip
cd czmq-3.0.2
./autogen.sh
./configure
make -j 4
make install && ldconfig
