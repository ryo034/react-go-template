package media

type Photo struct {
	id          ID
	title       string
	description string
	filePath    Path
	width       int
	height      int
	order       uint32
}

func NewPhoto(id ID, title string, description string, filePath Path, width int, height int, order uint32) *Photo {
	return &Photo{id, title, description, filePath, width, height, order}
}

func (p *Photo) ID() ID {
	return p.id
}

func (p *Photo) Title() string {
	return p.title
}

func (p *Photo) Description() string {
	return p.description
}

func (p *Photo) FilePath() Path {
	return p.filePath
}

func (p *Photo) Width() int {
	return p.width
}

func (p *Photo) Height() int {
	return p.height
}

func (p *Photo) IsVideo() bool {
	return false
}

func (p *Photo) IsImage() bool {
	return true
}

func (p *Photo) Order() uint32 {
	return p.order
}
