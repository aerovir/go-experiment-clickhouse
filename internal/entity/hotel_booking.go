package entity

type HotelBooking struct {
	LeadTime                    int     `column:"lead_time"`
	ArrivalDateYear             int     `column:"arrival_date_year"`
	ArrivalDateWeek             int     `column:"arrival_date_week_number"`
	ArrivalDateDayOfWeek        int     `column:"arrival_date_day_of_month"`
	StayInWeekendNights         int     `column:"stays_in_weekend_nights"`
	StayInWeekNights            int     `column:"stays_in_week_nights"`
	Adults                      int     `column:"adults"`
	Children                    int     `column:"children"`
	Babies                      int     `column:"babies"`
	PreviousBookingsNotCanceled int     `column:"previous_bookings_not_canceled"`
	DaysInWaitingList           int     `column:"days_in_waiting_list"`
	TotalOfSpecialRequests      int     `column:"total_of_special_requests"`
	Name                        string  `column:"name"`
	ArrivalDateMonth            string  `column:"arrival_date_month"`
	Meal                        string  `column:"meal"`
	Country                     string  `column:"country"`
	MarketSegment               string  `column:"market_segment"`
	DistributionChannel         string  `column:"distribution_channel"`
	ReservedRoomType            string  `column:"reserved_room_type"`
	AssignedRoomType            string  `column:"assigned_room_type"`
	BookingChanges              string  `column:"booking_changes"`
	DepositType                 string  `column:"deposit_type"`
	Agent                       string  `column:"agent"`
	Company                     string  `column:"company"`
	CustomerType                string  `column:"customer_type"`
	ReservationStatus           string  `column:"reservation_status"`
	ReservationStatusDate       string  `column:"reservation_status_date"`
	Adr                         float32 `column:"adr"`
	IsCanceled                  bool    `column:"is_canceled"`
	IsRepeatedGuest             bool    `column:"is_repeated_guest"`
	PreviousCancellations       bool    `column:"previous_cancellations"`
	RequiredCarParkingSpaces    bool    `column:"required_car_parking_spaces"`
}
