package services

type LdapConfig struct {
	Url        string
	ServerName string
	Tls        bool
}

type LoginConfig struct {
	UserName string
	Password string
}

type SvEn struct {
	Sv string `json:"sv"`
	En string `json:"en"`
}

type FKITSuperGroup struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"prettyName"`
	Type       string `json:"type"`
	Email      string `json:"email"`
}

type Post struct {
	ID          string `json:"id"`
	Sv          string `json:"sv"`
	En          string `json:"en"`
	EmailPrefix string `json:"emailPrefix"`
}

type FKITGroupDTO struct {
	ID              string         `json:"id"`
	BecomesActive   int64          `json:"becomesActive"`
	BecomesInactive int64          `json:"becomesInactive"`
	Description     SvEn           `json:"description"`
	Email           string         `json:"email"`
	Function        SvEn           `json:"function"`
	Name            string         `json:"name"`
	PrettyName      string         `json:"prettyName"`
	AvatarURL       interface{}    `json:"avatarURL"`
	SuperGroup      FKITSuperGroup `json:"superGroup"`
	Active          bool           `json:"active"`
}

type FKITUser struct {
	Post                  Post         `json:"post"`
	FkitGroupDTO          FKITGroupDTO `json:"fkitGroupDTO"`
	UnofficialPostName    string       `json:"unofficialPostName"`
	ID                    string       `json:"id"`
	Cid                   string       `json:"cid"`
	Nick                  string       `json:"nick"`
	FirstName             string       `json:"firstName"`
	LastName              string       `json:"lastName"`
	Email                 string       `json:"email"`
	Phone                 interface{}  `json:"phone"`
	Language              string       `json:"language"`
	AvatarURL             string       `json:"avatarUrl"`
	Gdpr                  bool         `json:"gdpr"`
	UserAgreement         bool         `json:"userAgreement"`
	AccountLocked         bool         `json:"accountLocked"`
	AcceptanceYear        int          `json:"acceptanceYear"`
	Authorities           interface{}  `json:"authorities"`
	Activated             bool         `json:"activated"`
	Username              string       `json:"username"`
	AccountNonExpired     bool         `json:"accountNonExpired"`
	AccountNonLocked      bool         `json:"accountNonLocked"`
	CredentialsNonExpired bool         `json:"credentialsNonExpired"`
	Enabled               bool         `json:"enabled"`
}

type FKITGroup struct {
	ID               string         `json:"id"`
	BecomesActive    int64          `json:"becomesActive"`
	BecomesInactive  int64          `json:"becomesInactive"`
	Description      SvEn           `json:"description"`
	Email            string         `json:"email"`
	Function         SvEn           `json:"function"`
	Name             string         `json:"name"`
	PrettyName       string         `json:"prettyName"`
	AvatarURL        interface{}    `json:"avatarURL"`
	SuperGroup       FKITSuperGroup `json:"superGroup"`
	Active           bool           `json:"active"`
	GroupMembers     []FKITUser     `json:"groupMembers"`
	NoAccountMembers []interface{}  `json:"noAccountMembers"`
}

type ITUser struct {
	ID                    string      `json:"id"`
	Cid                   string      `json:"cid"`
	Nick                  string      `json:"nick"`
	FirstName             string      `json:"firstName"`
	LastName              string      `json:"lastName"`
	Email                 string      `json:"email"`
	Phone                 string      `json:"phone"`
	Language              string      `json:"language"`
	AvatarURL             string      `json:"avatarUrl"`
	Gdpr                  bool        `json:"gdpr"`
	UserAgreement         bool        `json:"userAgreement"`
	AccountLocked         bool        `json:"accountLocked"`
	AcceptanceYear        int         `json:"acceptanceYear"`
	Authorities           []string    `json:"authorities"`
	Activated             bool        `json:"activated"`
	Enabled               bool        `json:"enabled"`
	Username              string      `json:"username"`
	AccountNonLocked      bool        `json:"accountNonLocked"`
	AccountNonExpired     bool        `json:"accountNonExpired"`
	CredentialsNonExpired bool        `json:"credentialsNonExpired"`
	Groups                []FKITGroup `json:"groups"`
	WebsiteURLs           interface{} `json:"websiteURLs"`
}
