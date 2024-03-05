package i18n

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Translator yapılandırmasını tanımlayın
type TranslatorService struct {
	bundle   *i18n.Bundle
	language string
}

var Translator *TranslatorService

// Yeni bir Translator örneği oluşturmak için bir fonksiyon
func NewTranslator(defaultLanguage language.Tag) *TranslatorService {
	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	return &TranslatorService{
		bundle: bundle,
	}
}

// Bir dil dosyası yükleme metodu
func (t *TranslatorService) LoadLanguageFile(filePath string) error {
	_, err := t.bundle.LoadMessageFile(filePath)
	return err
}

func (t *TranslatorService) GetTranslator() TranslatorService {
	return *Translator
}

// Belirli bir dil için çeviri yapma metodu
func (t *TranslatorService) Translate(messageID string, data interface{}) string {
	loc := i18n.NewLocalizer(t.bundle, t.language)
	translation, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	if err != nil {
		// Çeviri başarısız olursa, orijinal messageID döndürülebilir veya hata loglanabilir
		return messageID
	}
	return translation
}

func (t *TranslatorService) ChangeLanguage(lang string) {
	t.language = lang
}
