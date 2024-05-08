package requests

type Header struct {
	Message				int     `json:"Message"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
