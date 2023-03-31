# Golang_Bank_APP


for the installation of migration CLI on Ubuntu follow this 

sudo wget https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz

sudo tar xvf migrate.linux-amd64.tar.gz

sudo mv migrate /usr/local/bin/migrate

Remember tha when you run this command migrate create -ext sql -dir db/migration -seq init_schema  you are not ROOT 

make migrateup and migratedown when you get donw make sure that the table have this value version = 1 and dirty = false  

Because after this we need this files for SQL


go test -v -cover -short ./...
So, all together, the command go test -v -cover -short ./... will run the test suite for all packages in the current directory and its subdirectories, with verbose output and code coverage analysis enabled, and a shorter version of the tests being run.
