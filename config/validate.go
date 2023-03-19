package config

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func ValidateConfig(cfg *MainConfig) error {
	if cfg.Environment != "DEV" && cfg.Environment != "TEST" && cfg.Environment != "PROD" {
		return errors.New("Invalid Environment, must be one of 'DEV', 'TEST', or 'PROD'")
	}

	portRegex := regexp.MustCompile(`^\d{2,5}$`)
	//ipRegex := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	serviceRegex := regexp.MustCompile(`^[\w]{1,20}$`)
	fqdnRegex := regexp.MustCompile(`^((([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]))$`)

	for _, config := range cfg.Config {
		if !portRegex.MatchString(config.Port) {
			return fmt.Errorf("Invalid Port format: %s, must be numeric and length 2-5", config.Port)
		}

		sessionLifetimeMins, err := strconv.Atoi(config.SessionLifetimeMins)
		if err != nil || sessionLifetimeMins < 0 || sessionLifetimeMins > 9999 {
			return fmt.Errorf("Invalid SessionLifetimeMin value: %s, must be numeric and length 1-4", config.SessionLifetimeMins)
		}

		if config.TokenDbCheck != "Y" && config.TokenDbCheck != "N" {
			return fmt.Errorf("Invalid TokenDBCheck value: %s, must be 'Y' or 'N'", config.TokenDbCheck)
		}

		if !fqdnRegex.MatchString(config.Database.Server) {
			return fmt.Errorf("Invalid Database Server: %s, must be a valid IP address or FQDN", config.Database.Server)
		}

		if !portRegex.MatchString(config.Database.Port) {
			return fmt.Errorf("Invalid Database Port format: %s, must be numeric and length 2-5", config.Database.Port)
		}

		if !serviceRegex.MatchString(config.Database.Service) {
			return fmt.Errorf("Invalid Database Service format: %s, must be alphanumeric and max length 20", config.Database.Service)
		}

		if len(config.Database.Username) > 30 {
			return fmt.Errorf("Invalid Database Username length: %s, must be alphanumeric and max length 30", config.Database.Username)
		}

		if len(config.Database.Password) > 50 {
			return fmt.Errorf("Invalid Database Password length: %s, must be max length 50", config.Database.Password)
		}

		if !fqdnRegex.MatchString(config.LDAP.Server) {
			return fmt.Errorf("Invalid LDAP Server: %s, must be a valid IP address or FQDN", config.LDAP.Server)
		}

		if !portRegex.MatchString(config.LDAP.Port) {
			return fmt.Errorf("Invalid LDAP Port format: %s, must be numeric and length 2-5", config.LDAP.Port)
		}

		if utf8.RuneCountInString(config.LDAP.UserDN) > 30 {
			return fmt.Errorf("Invalid LDAP UserDN length: %s, must be UTF-8 with a maximum length of 30", config.LDAP.UserDN)
		}
	}

	return nil
}
