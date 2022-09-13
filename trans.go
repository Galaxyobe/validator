/*
 * Copyright 2022 Galaxyobe.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package validator

import (
	"log"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Translation struct {
	Tag             string
	Translation     string
	Override        bool
	CustomRegisFunc validator.RegisterTranslationsFunc
	CustomTransFunc validator.TranslationFunc
}

func (t Translation) Register(v *validator.Validate, trans ut.Translator) (err error) {
	if t.CustomTransFunc != nil && t.CustomRegisFunc != nil {
		err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisFunc, t.CustomTransFunc)
	} else if t.CustomTransFunc != nil && t.CustomRegisFunc == nil {
		err = v.RegisterTranslation(t.Tag, trans, RegistrationFunc(t.Tag, t.Translation, t.Override), t.CustomTransFunc)
	} else if t.CustomTransFunc == nil && t.CustomRegisFunc != nil {
		err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisFunc, TranslateFunc)
	} else {
		err = v.RegisterTranslation(t.Tag, trans, RegistrationFunc(t.Tag, t.Translation, t.Override), TranslateFunc)
	}
	if err != nil {
		return
	}
	return
}

type Translations []Translation

func (list Translations) Register(v *validator.Validate, trans ut.Translator) (err error) {
	for _, t := range list {
		if err = t.Register(v, trans); err != nil {
			return err
		}
	}
	return
}

func RegistrationFunc(Tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(Tag, translation, override); err != nil {
			return
		}
		return
	}
}

func TranslateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}
	return t
}
