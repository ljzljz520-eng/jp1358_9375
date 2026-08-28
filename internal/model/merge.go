package model

func MergeNotes(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return a + "; " + b
}
func AddSpecialty(p Profile, v string) Profile {
	for _, x := range p.Specialties {
		if x == v {
			return p
		}
	}
	p.Specialties = append(p.Specialties, v)
	return p
}
func AddAvailability(p Profile, v string) Profile {
	p.Availability = append(p.Availability, v)
	return p
}
