# Usage

docker build -t mifomm/golang-dev:1.14.7 .

docker push mifomm/golang-dev:1.14.7

docker run --rm -it --name golang-dev -v /vagrant/go/src:/home/rh/go/src mifomm/golang-dev:1.14.7

## inside container

source .profile
