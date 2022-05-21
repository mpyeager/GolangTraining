// Level 8: Exercise 2
// Unmarshal JSON into a Go data structure.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type dooooooonut []struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	Batters struct {
		Batter []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"batter"`
	} `json:"batters"`
	Topping []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"topping"`
}

func main() {
	data := `[[
		{
			"id": "0001",
			"type": "donut",
			"name": "Cake",
			"ppu": 0.55,
			"batters":
				{
					"batter":
						[
							{ "id": "1001", "type": "Regular" },
							{ "id": "1002", "type": "Chocolate" },
							{ "id": "1003", "type": "Blueberry" },
							{ "id": "1004", "type": "Devil's Food" }
						]
				},
			"topping":
				[
					{ "id": "5001", "type": "None" },
					{ "id": "5002", "type": "Glazed" },
					{ "id": "5005", "type": "Sugar" },
					{ "id": "5007", "type": "Powdered Sugar" },
					{ "id": "5006", "type": "Chocolate with Sprinkles" },
					{ "id": "5003", "type": "Chocolate" },
					{ "id": "5004", "type": "Maple" }
				]
		},
		{
			"id": "0002",
			"type": "donut",
			"name": "Raised",
			"ppu": 0.55,
			"batters":
				{
					"batter":
						[
							{ "id": "1001", "type": "Regular" }
						]
				},
			"topping":
				[
					{ "id": "5001", "type": "None" },
					{ "id": "5002", "type": "Glazed" },
					{ "id": "5005", "type": "Sugar" },
					{ "id": "5003", "type": "Chocolate" },
					{ "id": "5004", "type": "Maple" }
				]
		},
		{
			"id": "0003",
			"type": "donut",
			"name": "Old Fashioned",
			"ppu": 0.55,
			"batters":
				{
					"batter":
						[
							{ "id": "1001", "type": "Regular" },
							{ "id": "1002", "type": "Chocolate" }
						]
				},
			"topping":
				[
					{ "id": "5001", "type": "None" },
					{ "id": "5002", "type": "Glazed" },
					{ "id": "5003", "type": "Chocolate" },
					{ "id": "5004", "type": "Maple" }
				]
		}
	]]`

	fmt.Println(data)

	err := json.NewEncoder(os.Stdout).Encode(data)
	if err != nil {
		fmt.Println("Oops!", err)
	}
}
