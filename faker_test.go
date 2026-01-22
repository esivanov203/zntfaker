package zntfaker

import (
	"regexp"
	"strings"
	"testing"
)

// ------------------ FIO TESTS -------------------

func TestPerson(t *testing.T) {
	f := New()

	for i := 0; i < 10; i++ {
		fullName, lastName, firstName, middleName := f.Person("")
		if fullName == "" || lastName == "" || firstName == "" || middleName == "" {
			t.Errorf("Generated FIO has empty field: %s %s %s %s", fullName, lastName, firstName, middleName)
		}
	}

	// Проверка пола
	for _, gender := range []string{"m", "f"} {
		for i := 0; i < 10; i++ {
			_, _, firstName, middleName := f.Person(gender)
			if gender == "m" {
				if !contains(maleFirstNames, firstName) || !contains(maleMiddleNames, middleName) {
					t.Errorf("Expected male FIO for gender 'm', got %s %s", firstName, middleName)
				}
			} else {
				if !contains(femaleFirstNames, firstName) || !contains(femaleMiddleNames, middleName) {
					t.Errorf("Expected female FIO for gender 'f', got %s %s", firstName, middleName)
				}
			}
		}
	}
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// ------------------ EMAIL TESTS -------------------

func TestEmail(t *testing.T) {
	f := New()
	re := regexp.MustCompile(`^[a-zA-Z]{3,7}unface[a-zA-Z]{3,7}@[a-zA-Z]{3,7}unface[a-zA-Z]{3,7}\.(ru|com|org|net|zenit\.ru)$`)

	for i := 0; i < 10; i++ {
		email := f.Email()
		if !re.MatchString(email) {
			t.Errorf("Generated email does not match pattern: %s", email)
		}
	}
}

// ------------------ PHONE TESTS -------------------

func TestPhone(t *testing.T) {
	f := New()

	phones := []string{
		f.E164PhoneNumber(),
		f.PhoneBeauty(),
		f.PhoneBeautyNoCountry(),
	}

	for _, p := range phones {
		if strings.Count(p, "#") != 0 {
			t.Errorf("Phone number contains #: %s", p)
		}
	}
}

// ------------------ NUMERIC & STRING TESTS -------------------

func TestNumerify(t *testing.T) {
	f := New()
	for i := 0; i < 10; i++ {
		num := f.Numerify("##-##")
		if !regexp.MustCompile(`^\d{2}-\d{2}$`).MatchString(num) {
			t.Errorf("Numerify failed: %s", num)
		}
	}
}

func TestStringRange(t *testing.T) {
	f := New()
	for i := 0; i < 10; i++ {
		s := f.String(3, 7)
		if len(s) < 3 || len(s) > 7 {
			t.Errorf("String length out of range: %s", s)
		}
	}
}

func TestStringLength(t *testing.T) {
	f := New()
	word := "Образцовые текст. Text"
	for i := 0; i < 10; i++ {
		s := f.StringLength(word)
		if len([]rune(s)) != len([]rune(word)) {
			t.Errorf("String length out of range: %s", s)
		}
	}
}

// ------------------ INN & OGRN TESTS -------------------

func TestPersonalInn(t *testing.T) {
	f := New()
	for i := 0; i < 10; i++ {
		inn := f.PersonalInn()
		if len(inn) != 12 {
			t.Errorf("PersonalInn length incorrect: %s", inn)
		}
	}
}

func TestCompanyInnOgrn(t *testing.T) {
	f := New()
	for i := 0; i < 10; i++ {
		inn := f.CompanyInn()
		if len(inn) != 10 {
			t.Errorf("CompanyInn length incorrect: %s", inn)
		}
		ogrn := f.CompanyOgrn()
		if len(ogrn) != 13 {
			t.Errorf("CompanyOgrn length incorrect: %s", ogrn)
		}
	}
}

func TestBase64(t *testing.T) {
	f := New() // если нужен какой-то объект, оставим, иначе можно убрать
	cases := []struct {
		input, expected int
	}{
		{0, 0},
		{1, 4},
		{4, 4},
		{16, 16},
		{99, 100},
	}

	// регулярка для проверки URL-safe Base64
	re := regexp.MustCompile(`^[A-Za-z0-9\-_]*$`)

	for _, tc := range cases {
		s := f.Base64(tc.input)
		if len(s) != tc.expected {
			t.Errorf("input=%d: got length %d, want %d", tc.input, len(s), tc.expected)
		}
		if !re.MatchString(s) {
			t.Errorf("input=%d: string contains invalid characters: %s", tc.input, s)
		}

		s2 := f.Base64UrlSafe(tc.input)
		if len(s2) < tc.input {
			t.Errorf("Base64UrlSafe input=%d: got length %d, want ≥ %d", tc.input, len(s2), tc.input)
		}
		reUrlSafe := regexp.MustCompile(`^[A-Za-z0-9\-_]+=*$`)
		if len(s2) > 0 && !reUrlSafe.MatchString(s2) {
			t.Errorf("Base64UrlSafe input=%d: string contains invalid characters: %s", tc.input, s2)
		}
	}
}
