package models

type Claims struct {
	Sub				string		`json:"sub"`
	Role			string	    `json:"role"`
	Groups			[]string	`json:"groups"`
}
