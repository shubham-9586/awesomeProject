package appLogic

import (
	"awesomeProject/entity"
	"fmt"
)

var CurrrentBookings = []entity.Booking{}

func GetNewBooking(CentreID string, Day int, UserID string, Active bool) entity.Booking {
	return entity.Booking{
		BookingID: string(len(CurrrentBookings)),
		CentreID:  CentreID,
		Day:       Day,
		UserID:    UserID,
		Active:    Active,
	}
}

func AddNewBooking(booking entity.Booking) {
	fmt.Println("Booking has been added successfully")
	CurrrentBookings = append(CurrrentBookings, booking)
}

func ListAllBookings(centreID string) {
	for _, booking := range CurrrentBookings {
		condition := booking.Active && booking.CentreID == centreID
		if condition {
			fmt.Println(booking)
		}
	}
}

func ListBookings() {
	for _, booking := range CurrrentBookings {
		if booking.Active {
			fmt.Println(booking)
		}
	}
}

func CancelBooking(bookingID string) {
	for _, booking := range CurrrentBookings {
		if booking.BookingID == bookingID {
			booking.Active = false
			break
		}
	}
}
