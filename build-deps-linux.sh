echo "deb http://httpredir.debian.org/debian stretch main" >> /etc/apt/sources.list

apt-get update
apt-get install -y pkg-config libsodium-dev libczmq-dev autoconf automake libtool curl unzip libssl-dev

cd /tmp
curl -sSL https://github.com/google/protobuf/archive/master.zip > probuf.zip
unzip protobuf.zip
cd protobuf-master
./autogen.sh
./configure
make
make install
