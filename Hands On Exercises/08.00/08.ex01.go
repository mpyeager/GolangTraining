// Level 8: Exercise x1
// Marshal []user to JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type rabbi struct {
	Forname string
	Surname string
	House   string
}

func main() {

	bh1 := rabbi{
		Forname: `Hillel`,
		Surname: `haGadol`,
		House:   `Beit Hillel`,
	}
	bh2 := rabbi{
		Forname: `Shimon`,
		Surname: `ben Gamliel`,
		House:   `Beit Hillel`,
	}

	bs1 := rabbi{
		Forname: `Shammai`,
		Surname: ``,
		House:   `Beit Shammai`,
	}

	bs2 := rabbi{
		Forname: `Yohanan`,
		Surname: `ben Zakkai`,
		House:   `Beit Shammai`,
	}

	rabbanim := []rabbi{bh1, bh2, bs1, bs2}

	fmt.Println(rabbanim)

	bs, err := json.Marshal(rabbanim)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
