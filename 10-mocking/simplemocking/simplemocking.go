package simplemocking

type Matematik interface {
	MockTopla([]int) (int, error)
}

type islem struct{}

func (*islem) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam, nil
}
