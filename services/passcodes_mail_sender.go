package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/datasource/models"
	"strings"
	"time"
)

var (
	activeMessage = "Hello! Dear %s %s!\r\n" +
		"The account you are using this email address on %s is trying to activate, so the email address needs to be verified.\r\n\r\n" +
		"You can use this code to verify email: %s\r\n\r\n" +
		"If you have not registered any account, please ignore this email.\r\n"
)

func SendOneTimePasscode(user models.User, email string, method int) (models.OneTimePasscode, error) {
	expired := time.Now().Add(time.Minute * 30)
	passcode := models.OneTimePasscode{
		Type:        method,
		Passcode:    strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1)[:8]),
		RefreshedAt: nil,
		ExpiredAt:   &expired,
		UserID:      user.ID,
	}

	site := viper.GetString("name")
	err := SendMail(email, fmt.Sprintf("%s Contact Method Verification", site),
		fmt.Sprintf(activeMessage, user.Details.Firstname, user.Details.Lastname, site, passcode.Passcode))

	if err != nil {
		return models.OneTimePasscode{}, err
	} else {
		err := datasource.C.Save(&passcode).Error
		return passcode, err
	}
}

func ResendOneTimePasscode(passcode models.OneTimePasscode, user models.User) (models.OneTimePasscode, error) {
	expired := time.Now().Add(time.Minute * 30)
	refreshed := time.Now()
	passcode.ExpiredAt = &expired
	passcode.RefreshedAt = &refreshed
	passcode.Passcode = strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1)[:8])

	var email string
	switch passcode.Type {
	case models.OneTimeVerifyPrimaryEmailCode:
		email = user.Details.PrimaryEmail
	case models.OneTimeVerifySecondaryEmailCode:
		email = user.Details.SecondaryEmail
	case models.OneTimeVerifyPhoneNumberCode:
		return models.OneTimePasscode{}, fmt.Errorf("unsupported passcode type")
	case models.OneTimeDangerousPasscode:
		email = user.Details.PrimaryEmail
	}

	site := viper.GetString("name")
	err := SendMail(email, fmt.Sprintf("%s Contact Method Verification", site),
		fmt.Sprintf(activeMessage, user.Details.Firstname, user.Details.Lastname, site, passcode.Passcode))

	if err != nil {
		return models.OneTimePasscode{}, err
	} else {
		err := datasource.C.Save(&passcode).Error
		return passcode, err
	}
}
