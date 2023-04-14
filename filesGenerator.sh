#!/bin/bash

# Create cmd directory
mkdir cmd
touch cmd/main.go

# Create internal directory
mkdir internal

# Create app directory
mkdir internal/app
mkdir internal/app/handlers
mkdir internal/app/models
mkdir internal/app/services
touch internal/app/app.go
touch internal/app/handlers/car_handler.go
touch internal/app/models/car.go
touch internal/app/services/car_service.go
touch internal/app/services/user_service.go

# Create config directory
mkdir internal/config
touch internal/config/config.go

# Create domain directory
mkdir internal/domain
mkdir internal/domain/repositories
mkdir internal/domain/models
mkdir internal/domain/services
mkdir internal/domain/usecases
touch internal/domain/repositories/car_repository.go
touch internal/domain/repositories/user_repository.go
touch internal/domain/models/car.go
touch internal/domain/models/user.go
touch internal/domain/services/car_service.go
touch internal/domain/services/user_service.go
touch internal/domain/usecases/car_usecase.go
touch internal/domain/usecases/user_usecase.go

# Create infrastructure directory
mkdir internal/infrastructure
mkdir internal/infrastructure/database
mkdir internal/infrastructure/database/mongo
mkdir internal/infrastructure/http
mkdir internal/infrastructure/repositories
touch internal/infrastructure/database/mongo/mongo.go
touch internal/infrastructure/database/mongo/car_repository.go
touch internal/infrastructure/http/router.go
touch internal/infrastructure/repositories/user_repository.go

# Create pkg directory
mkdir internal/pkg
mkdir internal/pkg/middleware
touch internal/pkg/middleware/auth.go

# Create test directory
mkdir internal/test
mkdir internal/test/integration
mkdir internal/test/unit
touch internal/test/integration/car_repository_test.go
touch internal/test/unit/car_service_test.go
touch internal/test/unit/user_service_test.go

# Create go.mod file
go mod init git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes

