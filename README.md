# Key Management System (KMS)
KMS is a Key Management System that stores generated keys and is able to retrieve keys for for future use. 

## Instructions
```bash
# Perform Database Migrations
./goose -dir=./migrations mysql "username:password@tcp(localhost:3306)/kms?parseTime=true" up;

# Build application
go build -o kms;

# Run application
./kms -dsn="username:password@tcp(localhost:3306)/kms?parseTime=true";

# Create a user
./kms -dsn="username:password@tcp(localhost:3306)/kms?parseTime=true" -createUser -username="EPOINT" -password="Oc3UmUZsK0" -programID=1;
```

## API Documentation
API documentation is available on Postman (Collection: KMS).