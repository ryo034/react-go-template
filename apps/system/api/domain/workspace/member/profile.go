package member

type Profile struct {
	displayName *DisplayName
	idNumber    *IDNumber
	bio         Bio
}

func NewProfile(displayName *DisplayName, idNumber *IDNumber, bio Bio) Profile {
	return Profile{displayName, idNumber, bio}
}

func (p *Profile) DisplayName() *DisplayName {
	return p.displayName
}

func (p *Profile) IDNumber() *IDNumber {
	return p.idNumber
}

func (p *Profile) Bio() Bio {
	return p.bio
}

func (p *Profile) HasBio() bool {
	return p.bio.v != ""
}

func (p *Profile) HasDisplayName() bool {
	return p.displayName != nil
}

func (p *Profile) HasIDNumber() bool {
	return p.idNumber != nil
}
