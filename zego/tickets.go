package zego

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type TicketArray struct {
	Count         int    `json:"count"`
	Created       string `json:"created"`
	Next_page     string `json:"next_page"`
	Previous_page string `json:"previous_page"`
	Tickets       []Ticket
}

type SingleTicket struct {
	Ticket *Ticket `json:"ticket"`
}

type TicketResult struct {
	SingleTicket *SingleTicket `json:"result"`
}

type TicketUpdate struct {
	Ticket *TicketUpdateInner `json:"ticket"`
}

type TicketUpdateInner struct {
	Id      uint64    `json:"id"`
	Status  string    `json:"status"`
	Comment *Comments `json:"comment"`
}

type Ticket struct {
	Id                    uint64      `json:"id"`
	URL                   string      `json:"url"`
	ExternalId            string      `json:"external_id"`
	CreatedAt             string      `json:"created_at"`
	UpdatedAt             string      `json:"updated_at"`
	Type                  string      `json:"type"`
	Subject               string      `json:"subject"`
	RawSubject            string      `json:"raw_subject"`
	Description           string      `json:"description"`
	Priority              string      `json:"priority"`
	Status                string      `json:"status"`
	Recipient             string      `json:"recipient"`
	RequesterId           uint32      `json:"requester_id"`
	SubmitterId           uint32      `json:"submitter_id"`
	AssigneeId            uint32      `json:"assignee_id"`
	OrganizationId        uint32      `json:"organization_id"`
	GroupId               uint32      `json:"group_id"`
	CollaboratorIds       []int32     `json:"collaborator_ids"`
	ForumTopicId          uint32      `json:"forum_topic_id"`
	ProblemId             uint32      `json:"problem_id"`
	HasIncidents          bool        `json:"has_incidents"`
	DueAt                 string      `json:"due_at"`
	Tags                  []string    `json:"tags"`
	Satisfaction_rating   string      `json:"satisfaction_rating"`
	Ticket_form_id        uint32      `json:"ticket_form_id"`
	Sharing_agreement_ids interface{} `json:"sharing_agreement_ids"`
	Via                   interface{} `json:"via"`
	Custom_Fields         interface{} `json:"custom_fields"`
	Fields                interface{} `json:"fields"`
	Comment               Comments    `json:"comment"`
}

func (a Auth) ListTickets(pag ...string) (*TicketArray, error) {

	TicketStruct := &TicketArray{}

	var path string
	if len(pag) < 1 {
		path = "/tickets.json"
	} else {
		path = pag[0]
	}
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), TicketStruct)

	return TicketStruct, nil

}

func (a Auth) GetTicket(ticket_id string) (*SingleTicket, error) {

	TicketStruct := &SingleTicket{}

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), TicketStruct)

	return TicketStruct, nil

}

func (a Auth) GetMultipleTickets(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) GetTicketComments(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + "/comments.json"
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) DeleteTicket(ticket_id string) (*Resource, error) {

	path := "/tickets/" + ticket_id + ".json"
	resource, err := api(a, "DELETE", path, "")
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) CreateTicket(data string) (*Resource, error) {

	path := "/tickets.json"
	resource, err := api(a, "POST", path, data)
	if err != nil {
		return nil, err
	}

	return resource, nil

}

func (a Auth) UpdateTicket(up TicketUpdate) (*Resource, error) {

	path := "/tickets/" + strconv.Itoa(int(up.Ticket.Id)) + ".json"
	data, err := json.Marshal(&up)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	resource, err := api(a, "PUT", path, string(data))
	if err != nil {
		return nil, err
	}

	return resource, nil

}
