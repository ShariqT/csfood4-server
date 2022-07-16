package core

func GetMarkets(db DB) ([]Modelable, error) {
	markets, err := db.ListAll(MarketResource)
	if err != nil {
		return nil, err
	}
	return markets, nil
}
