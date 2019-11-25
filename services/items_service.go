package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct{}

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
