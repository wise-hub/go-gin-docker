package helpers

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func LdapAuth(username string, password string) (string, error) {
	// Set up the LDAP connection
	l, err := ldap.Dial("tcp", "ldap.forumsys.com:389")
	if err != nil {
		return "", err
	}
	defer l.Close()

	// Bind with the read-only-admin credentials
	err = l.Bind("cn=read-only-admin,dc=example,dc=com", "password")
	if err != nil {
		return "", err
	}

	// Search for the user's DN
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
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
	fmt.Println("Authenticated successfully")

	return userDN, nil
}
