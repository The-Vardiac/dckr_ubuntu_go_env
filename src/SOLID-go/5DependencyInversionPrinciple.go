package main

/* depend on interfaces/abstraction, not concrete implementation */

/* bad example */
type OrderServiceA struct {
	repo *MysqlOrderRepository
}


/* good example */
type OrderRepository interface {
	Save(order string) error
}

type mysqlOrderRepository struct {
	db *mysql.DB
}
func NewMysqlOrderRepository(db *mysql.DB) OrderRepository {
	return &mysqlOrderRepository{db: db}
}
func (r *mysqlOrderRepository) Save(order string) error {
	return nil
}

type OrderServiceImpl struct {
	repo OrderRepository
}
func NewOrderService(r OrderRepository) OrderServiceB {
	return &OrderServiceImpl{
		repo: r
	}
}