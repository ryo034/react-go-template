package media

type Media interface {
	ID() ID
	Title() string
	Description() string
	FilePath() Path
	Width() int
	Height() int
	IsVideo() bool
	IsImage() bool
	Order() uint32
}
