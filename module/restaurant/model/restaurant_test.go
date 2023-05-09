package restaurantmodel

import "testing"

type testData struct {
	Input  RestaurantCreate
	Expect error
}

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTable := []testData{
		{Input: RestaurantCreate{Name: ""}, Expect: ErrNameISEmpty},
		{Input: RestaurantCreate{Name: "Some Name"}, Expect: nil},
	}

	for _, item := range dataTable {
		err := item.Input.Validate()

		if err != item.Expect {
			t.Errorf("Validate restaurant. Input name: %v, Expect: %v, Ouput: %v", item.Input, item.Expect, err)
		}

	}
}
