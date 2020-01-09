package main

import "fmt"

/* https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
Composition with Go
William KennedySeptember 13, 2015
*/

// ======================================================

// Deska reprezentuje powierzchnie na ktorej pracujemy.
type Deska struct {
	PotrzebneGwozdzie int
	WbiteGwozdzie     int
}

// ======================================================

// WbityGwozdz reprezentuje zachowanie wbicia gwozdzia w deske.
type WbityGwozdz interface {
	WbicGwozdzia(zapasGwozdzi *int, d *Deska)
}

// WyjetyGwozdz reprezentuje zachowanie przy wyjeciu gwozdzia z Deski.
type WyjetyGwozdz interface {
	WyjmijGwozdzia(zapasGwozdzi *int, d *Deska)
}

// WbijWyjmijGwozdzia reprezentuje zachowanie przy wbiciu/wyjeciu gwozdzia z Deski.
type WbijWyjmijGwozdzia interface {
	WbityGwozdz
	WyjetyGwozdz
}

// ======================================================

// Mlotek to narzedzie to bije w gwozdzie.
type Mlotek struct{}

// WbijGwozdzia bije w gwodz w konkretna deske.
func (Mlotek) WbicGwozdzia(zapasGwozdzi *int, d *Deska) {
	// Wez gwozdzia z zapasow
	*zapasGwozdzi--

	// Uderz gwozdzia w deske
	d.WbiteGwozdzie++
	fmt.Println("Mlotek: uderzyl w gwozdz i wbil w deske.")
}

// Obcegi to narzedzie ktore wyciaga gwozdzie
type Obcegi struct{}

// Wyciagnij gwodz z konkretnej deski.
func (Obcegi) WyjmijGwozdzia(zapasGwozdzi *int, d *Deska) {
	d.WbiteGwozdzie--

	// Dodaj gwozdzia do zapasu.
	*zapasGwozdzi++

	fmt.Println("Obcegi: wyjety gwozdz z deski.")
}

// ======================================================

// Robotnik zajmuje sie zabezpieczeniem desek.
type Robotnik struct{}

// Montaz wbije gwoździe w deskę.
func (Robotnik) Montaz(wb WbityGwozdz, zapasGwozdzi *int, d *Deska) {
	for d.WbiteGwozdzie < d.PotrzebneGwozdzie {
		wb.WbicGwozdzia(zapasGwozdzi, d)
	}
}

// Demontaz wyjmie gwozdzie z deski.
func (Robotnik) Demontaz(wy WyjetyGwozdz, zapasGwozdzi *int, d *Deska) {
	for d.WbiteGwozdzie > d.PotrzebneGwozdzie {
		wy.WyjmijGwozdzia(zapasGwozdzi, d)
	}
}

// Stolarka praca na Deskach.
func (r Robotnik) Stolarka(ww WbijWyjmijGwozdzia, zapasGwozdzi *int, deski []Deska) {
	for i := range deski {
		d := &deski[i]

		fmt.Printf("robotnik: sprawdza deske #%d: %+v\n", i+1, d)

		switch {
		case d.WbiteGwozdzie < d.PotrzebneGwozdzie:
			r.Montaz(ww, zapasGwozdzi, d)

		case d.WbiteGwozdzie > d.PotrzebneGwozdzie:
			r.Demontaz(ww, zapasGwozdzi, d)
		}
	}
}

// ======================================================

// SkrzynkaNarzedziowa zawiera narzedzia.
type SkrzynkaNarzedziowa struct {
	WbityGwozdz
	WyjetyGwozdz

	gwozdzie int
}

// ======================================================

// pokazStan podaje informacje o wszystkich deskach.
func pokazStan (sn *SkrzynkaNarzedziowa, deski []Deska) {
	fmt.Printf("Skrzynia: %#v\n", sn)
	fmt.Println("Deski:")

	for _, d := range deski {
		fmt.Printf("\t%+v\n", d)
	}

	fmt.Println()
}

// main funkcja od ktorej startuje aplikacja.
func main() {

	// Skrzynia z nowymi deskami oraz ze starymi co je zamienia
	deski := []Deska{
		// Stare deski do wymiany.
		{WbiteGwozdzie: 3},
		{WbiteGwozdzie: 1},
		{WbiteGwozdzie: 6},

		// Nowa deska do montazu.
		{PotrzebneGwozdzie: 6},
		{PotrzebneGwozdzie: 9},
		{PotrzebneGwozdzie: 4},
	}

	// zapelnij skrzynke narzedziowa
	sn := SkrzynkaNarzedziowa {
		WbityGwozdz: Mlotek{},
		WyjetyGwozdz: Obcegi{},
		gwozdzie: 10,
	}

	// Najmij robotnika i kaz mu pracowac.
	var r Robotnik
	r.Stolarka(&sn, &sn.gwozdzie, deski)

	// Pokaz obecny stan skrzyni i desek.
	pokazStan(&sn, deski)
}