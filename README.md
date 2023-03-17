# Golang_Bank_APP


for the installation of migration CLI on Ubuntu follow this 

sudo wget https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz

sudo tar xvf migrate.linux-amd64.tar.gz

sudo mv migrate /usr/local/bin/migrate

Remember tha when you run this command migrate create -ext sql -dir db/migration -seq init_schema  you are not ROOT 

Because after this we need this files for SQL
