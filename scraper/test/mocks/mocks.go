package mocks

// Generate database repository mocks
//go:generate mockgen -source=../../internal/repositories/database/measurements/interface.go -destination=./repositories/database/measurements/db_measurements_repo_mocks.go -package=dbmeasurementsmocks

// Generate openaq repository mocks
//go:generate mockgen -source=../../internal/repositories/openaq/interface.go -destination=./repositories/openaq/openaq_repo_mocks.go -package=openaqmocks
