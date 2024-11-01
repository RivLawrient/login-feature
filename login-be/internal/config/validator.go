// package config

// import (
// 	"github.com/go-playground/locales/en"
// 	ut "github.com/go-playground/universal-translator"
// 	"github.com/go-playground/validator/v10"
// 	enTranslations "github.com/go-playground/validator/v10/translations/en"
// 	"github.com/spf13/viper"
// )

// func NewValidator(viper *viper.Viper) *validator.Validate {
// 	validate := validator.New()
// 	return validate
// }

// func ValidateStruct(s interface{}) []string {
// 	validate := validator.New()
// 	var translatedErrors []string

// 	eng := en.New()
// 	uni := ut.New(eng, eng)
// 	trans, _ := uni.GetTranslator("en")
// 	enTranslations.RegisterDefaultTranslations(validate, trans)

// 	err := validate.Struct(s)
// 	if err != nil {
// 		if valErr, ok := err.(validator.ValidationErrors); ok {
// 			for _, e := range valErr {
// 				translatedErrors = append(translatedErrors, e.Translate(trans))
// 			}
// 		}
// 	}

// 	return translatedErrors
// }

package config

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/spf13/viper"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func NewValidator(viper *viper.Viper) *validator.Validate {
	// Inisialisasi validator hanya sekali
	validate = validator.New()

	// Inisialisasi translator
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ = uni.GetTranslator("en")

	// Daftarkan terjemahan default untuk bahasa Inggris
	enTranslations.RegisterDefaultTranslations(validate, trans)

	// Return validator untuk digunakan di tempat lain (jika diperlukan)
	return validate
}

// Fungsi untuk memvalidasi dan menerjemahkan error
func ValidateStruct(s interface{}) []string {
	var translatedErrors []string

	// Lakukan validasi pada struct
	err := validate.Struct(s)
	if err != nil {
		// Cek apakah error adalah ValidationErrors
		if valErr, ok := err.(validator.ValidationErrors); ok {
			// Iterasi dan terjemahkan semua pesan error
			for _, e := range valErr {
				translatedErrors = append(translatedErrors, e.Translate(trans))
			}
		}
	}

	return translatedErrors
}
