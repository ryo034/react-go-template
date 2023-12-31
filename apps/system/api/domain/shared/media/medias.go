package media

type Medias interface {
	Size() int
	AsSlice() []Media
	IsEmpty() bool
	IsNotEmpty() bool
}

type medias struct {
	wrapped []Media
}

func NewMedias(vs []Media) Medias {
	return &medias{vs}
}

func (ms *medias) IsEmpty() bool {
	return len(ms.wrapped) == 0
}

func (ms *medias) IsNotEmpty() bool {
	return !ms.IsEmpty()
}

func (ms *medias) Size() int {
	return len(ms.wrapped)
}

func (ms *medias) AsSlice() []Media {
	return append(make([]Media, 0, ms.Size()), ms.wrapped...)
}
