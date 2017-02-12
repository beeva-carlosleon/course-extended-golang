export GOPATH=~/gopath

PATH=$PATH:$GOPATH/bin

sudo docker pull mysql

sudo docker run --name mysql-docker -e MYSQL_ROOT_PASSWORD=curso -p 3306:3306 -d mysql:latest

mysql -h localhost --protocol=tcp -u root --password

create database curso;

use curso;

create table users (ID int not null, NAME varchar(255), SURNAME varchar(255),AGE int, PRIMARY KEY (ID));

docker run --name mongo-docker -p 27017:27017 -d mongo --storageEngine wiredTiger

sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927

echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list

sudo apt-get update

sudo apt-get install -y mongodb-org

use curso

db.curso.insert({id:0,name:"test",surname:"test"})
db.curso.findOne()

go get github.com/coreos/etcd
go get github.com/coreos/etcd/clientv3

