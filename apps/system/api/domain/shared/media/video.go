package media

type Video struct {
	id          ID
	title       string
	description string
	filePath    Path
	width       int
	height      int
	timeLength  int
	order       uint32
}

func NewVideo(id ID, title string, description string, filePath Path, width int, height int, timeLength int, order uint32) *Video {
	return &Video{id, title, description, filePath, width, height, timeLength, order}
}

func (v *Video) ID() ID {
	return v.id
}

func (v *Video) Title() string {
	return v.title
}

func (v *Video) Description() string {
	return v.description
}

func (v *Video) FilePath() Path {
	return v.filePath
}

func (v *Video) Width() int {
	return v.width
}

func (v *Video) Height() int {
	return v.height
}

func (v *Video) TimeLength() int {
	return v.timeLength
}

func (v *Video) IsVideo() bool {
	return false
}

func (v *Video) IsImage() bool {
	return true
}

func (v *Video) Order() uint32 {
	return v.order
}
