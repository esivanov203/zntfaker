package zntfaker

import "testing"

func BenchmarkFirstName(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.FirstName("m")
	}
}

func BenchmarkLastName(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.LastName("f")
	}
}

func BenchmarkPerson(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Person("")
	}
}

func BenchmarkPersonalInn(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.PersonalInn()
	}
}

func BenchmarkCompanyInn(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.CompanyInn()
	}
}

func BenchmarkCompanyOgrn(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.CompanyOgrn()
	}
}

func BenchmarkEmail(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Email()
	}
}

func BenchmarkPhoneBeauty(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.PhoneBeauty()
	}
}

func BenchmarkInt(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Int(1, 100)
	}
}

func BenchmarkString(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.String(5, 10)
	}
}

// Комплексный бенчмарк - создание полного профиля
func BenchmarkFullProfile(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Person("")
		f.PersonalInn()
		f.Email()
		f.PhoneBeauty()
	}
}

func BenchmarkMiddleName(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.MiddleName("m")
	}
}

func BenchmarkE164PhoneNumber(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.E164PhoneNumber()
	}
}

func BenchmarkPhoneBeautyNoCountry(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.PhoneBeautyNoCountry()
	}
}

func BenchmarkNumerify(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Numerify("+# (###) ###-##-##")
	}
}
