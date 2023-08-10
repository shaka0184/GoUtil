package zoomUtil

type RequestParticipantJoined struct {
	Event   string `json:"event"`
	EventTs string `json:"event_ts"`
	Payload struct {
		AccountId string `json:"account_id"`
		Object    struct {
			Id          string `json:"id"`
			Uuid        string `json:"uuid"`
			HostId      string `json:"host_id"`
			Topic       string `json:"topic"`
			Type        string `json:"type"`
			StartTime   string `json:"start_time"`
			Timezone    string `json:"timezone"`
			Duration    string `json:"duration"`
			Participant struct {
				UserId            string `json:"user_id"`
				UserName          string `json:"user_name"`
				Id                string `json:"id"`
				DateTime          string `json:"date_time"`
				Email             string `json:"email"`
				RegistrantId      string `json:"registrant_id"`
				ParticipantUserId string `json:"participant_user_id"`
				CustomerKey       string `json:"customer_key"`
			} `json:"participant"`
		} `json:"object"`
	} `json:"payload"`
}
