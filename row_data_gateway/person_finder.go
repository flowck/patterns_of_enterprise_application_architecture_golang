package row_data_gateway

type PersonFinder struct{}

func (p PersonFinder) Find(id string) (*PersonGateway, error) {
	return nil, nil
}
