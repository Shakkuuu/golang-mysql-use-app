package entity

// type User struct
type User struct {
	ID       int
	Name     string
	Password string
	Item_ID  int
}

// type Coin struct
type Coin struct {
	User_ID int
	Qty     int
}

// type Item struct
type Buki struct {
	ID   int
	Name string
	Atk  int
}

// var ki Buki = Buki{ID: 1, Name: "木の剣", Atk: 2}
// var isi Buki = Buki{ID: 2, Name: "石の剣", Atk: 4}
// var tetu Buki = Buki{ID: 1, Name: "鉄の剣", Atk: 6}
