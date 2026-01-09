package game

func (t *try) printGallows(l int) {
	if l == 1 {
		t.imageGallow[4][10] = "O"
	}
	if l == 2 {
		t.imageGallow[5][10] = "|"
	}
	if l == 3 {
		t.imageGallow[5][9] = "/"
	}
	if l == 4 {
		t.imageGallow[5][11] = "\\"
	}
	if l == 5 {
		t.imageGallow[6][9] = "/"
	}
	if l == 6 {
		t.imageGallow[6][11] = "\\"
	}

}
