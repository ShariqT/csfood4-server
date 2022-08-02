package core

func GetProductVariationById(productId string, db DB) (Modelable, error) {
	product, err := db.GetById(ProductResouce, productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
