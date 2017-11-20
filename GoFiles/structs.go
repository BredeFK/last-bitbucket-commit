//====================================================//
//             AUTHOR: 	Brede Fritjof Klausen         //
//         UNIVERSITY: 	NTNU in Gjøvik                //
//====================================================//

package main

import "time"

// Bitbucket struct
type Bitbucket struct {
	Pagelen int      `json:"pagelen"`
	Values  []Values `json:"values"`
}

// Values struct
type Values struct {
	Hash       string     `json:"hash"`
	Repository Repository `json:"repository"`
	Links      Links      `json:"links"`
	Author     Author     `json:"author"`
	Parents    []Parents  `json:"parents"`
	Date       time.Time  `json:"date"`
	Message    string     `json:"message"`
	Type       string     `json:"type"`
}

// Repository struct
type Repository struct {
	Links    Links  `json:"links"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	UUID     string `json:"uuid"`
}

// Links struct
type Links struct {
	Self struct {
		Href string `json:"href"`
	} `json:"self"`
	Comments struct {
		Href string `json:"href"`
	} `json:"comments"`
	Patch struct {
		Href string `json:"href"`
	} `json:"patch"`
	HTML struct {
		Href string `json:"href"`
	} `json:"html"`
	Diff struct {
		Href string `json:"href"`
	} `json:"diff"`
	Avatar struct {
		Href string `json:"href"`
	} `json:"avatar"`
	Approve struct {
		Href string `json:"href"`
	} `json:"approve"`
	Statuses struct {
		Href string `json:"href"`
	} `json:"statuses"`
}

// Author struct
type Author struct {
	Raw  string `json:"raw"`
	Type string `json:"type"`
	User User   `json:"user"`
}

// User struct
type User struct {
	UserName    string `json:"username"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	UUID        string `json:"uuid"`
	Links       Links  `json:"links"`
}

// Parents struct
type Parents struct {
	Hash  string `json:"hash"`
	Type  string `json:"type"`
	Links Links  `json:"links"`
}

// ShowInfo struct
type ShowInfo struct {
	DisplayName string `json:"displayname"`
	UserName    string `json:"userName"`
	Message     string `json:"message"`
	Diff        string `json:"diff"`
	Date        string `json:"date"`
}
