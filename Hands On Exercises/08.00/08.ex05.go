// Level 8: Exercise 5
// Sort rabbanim by death, surname. Sort string teachings for each rabbi.

package main

import (
	"fmt"
	"sort"
)

type rabbi struct {
	Forname   string
	Surname   string
	Death     int
	Teachings []string
}

type ByDeath []rabbi

func (a ByDeath) Len() int           { return len(a) }
func (a ByDeath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDeath) Less(i, j int) bool { return a[i].Death < a[j].Death }

type BySurname []rabbi

func (l BySurname) Len() int           { return len(l) }
func (l BySurname) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l BySurname) Less(i, j int) bool { return l[i].Surname < l[j].Surname }

func main() {
	r1 := rabbi{
		Forname: `Hillel`,
		Surname: `haGadol`,
		Death:   10,
		Teachings: []string{
			`That which is hateful to you, do not do unto your fellow.`,
			`If I am not for myself, who will be for me? And being only for myself, what am I? And if not now, when?`,
			`Don't trust yourself until the day you die.`,
		},
	}

	r2 := rabbi{
		Forname: `Menachem`,
		Surname: `Mendel Schneerson`,
		Death:   1994,
		Teachings: []string{
			`If one does not work on themselves, what good will submitting notes, singing songs, and saying l'chaim do?`,
			`When two people meet, it should bring benefit to a third.`,
			`I have done my part, from now on you do all that you can.`,
		},
	}

	r3 := rabbi{
		Forname: `Jonathan`,
		Surname: `Sacks`,
		Death:   2020,
		Teachings: []string{
			`Rabbi Soloveitchik had challenged me to think, Rabbi Schneerson had challenged me to lead.`,
			`The consumer society was laid down by the late Steve Jobs coming down the mountain with two tablets, iPad one and iPad two, and the result is that we now have a culture of iPod, iPhone, iTune, i, i, i. When you're an individualist, egocentric culture and you only care about 'I', you don’t do terribly well.`,
			`No one creed has a monopoly on spiritual truth.`,
		},
	}

	rabbanim := []rabbi{r1, r2, r3}
	for _, r := range rabbanim {
		fmt.Println(r.Death, r.Surname, ",", r.Forname)
		sort.Strings(r.Teachings)
		for _, v := range r.Teachings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println(`-------`)

	sort.Sort(ByDeath(rabbanim))
	for _, r := range rabbanim {
		fmt.Println(r.Death, r.Surname, ",", r.Forname)
		for _, v := range r.Teachings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println(`-------`)

	sort.Sort(BySurname(rabbanim))
	for _, r := range rabbanim {
		fmt.Println(r. Death, r.Surname, ",", r.Forname)
		for _, v := range r.Teachings {
			fmt.Println("\t", v)
		}
	}

}
