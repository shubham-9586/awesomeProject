package Input

var InputArray = []string{
	"ADD_USER U1 Harry Male 35 Karnataka Bangalore",
	"ADD_USER U2 Ron Male 30 Karnataka Bangalore",
	//"ADD_USER U3 Albus Male 30 Karnataka Bangalore",
	//"ADD_USER U4 Draco Male 15 Karnataka Bangalore",
	//"ADD_USER U5 Dobby Male 30 Gujarat Surat",
	"ADD_VACCINATION_CENTER Karnataka Bangalore VC1",
	"ADD_VACCINATION_CENTER Karnataka Bangalore VC2",
	"ADD_VACCINATION_CENTER Karnataka Mysore VC3",
	"ADD_CAPACITY VC1 1 1",
	"ADD_CAPACITY VC2 1 3",
	"ADD_CAPACITY VC1 5 10",
	"ADD_CAPACITY VC3 3 4",
	"LIST_VACCINATION_CENTERS Bangalore",
	"BOOK_VACCINATION VC1 1 U1",
	//"LIST_ALL_BOOKINGS VC1 1",
	//"BOOK_VACCINATION VC2 1 U2",
	//"BOOK_VACCINATION VC2 1 U3",
	//"LIST_ALL_BOOKINGS VC2 1",
	//"BOOK_VACCINATION VC1 1 U5",
}

func Process(input string) (string, string) {
	return "a", "b"
}
