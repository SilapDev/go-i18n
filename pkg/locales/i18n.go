package locales

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/silapdev/go_i18n_test/pkg/helper"
	"golang.org/x/text/language"
)

func NewBundle() *i18n.Bundle {

	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile(fmt.Sprintf("%s/translations/en.json", helper.GetRootProjectPath()))
	bundle.MustLoadMessageFile(fmt.Sprintf("%s/translations/ru.json", helper.GetRootProjectPath()))
	bundle.MustLoadMessageFile(fmt.Sprintf("%s/translations/tm.json", helper.GetRootProjectPath()))

	return bundle

}
