package sql_service

type PostgresService interface {
	Query(objecttoscan interface{}, querystring string, args ...string) error
}
