package helpers

import (
	"fmt"
	"ginws/config"

	"github.com/go-ldap/ldap/v3"
)

func LdapAuth(d *config.Dependencies, username string, password string) (string, error) {
	// Set up the LDAP connection
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", d.Cfg.LDAP.Server, d.Cfg.LDAP.Port))
	if err != nil {
		return "", err
	}
	defer l.Close()

	// Search for the user's DN
	searchRequest := ldap.NewSearchRequest(
		d.Cfg.LDAP.UserDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(uid=%s)", username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return "", err
	}

	if len(sr.Entries) != 1 {
		return "", fmt.Errorf("User not found or too many entries returned: %v", len(sr.Entries))
	}

	// Bind as the user
	userDN := sr.Entries[0].DN
	err = l.Bind(userDN, password)

	if err != nil {
		return "", err
	}

	// Authentication successful!
	// fmt.Println("Authenticated successfully")
	// fmt.Println(userDN)

	return userDN, nil
}
