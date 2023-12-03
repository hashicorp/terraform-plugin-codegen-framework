package convert

type Sensitive struct {
	sensitive *bool
}

func NewSensitive(s *bool) Sensitive {
	return Sensitive{
		sensitive: s,
	}
}

func (s Sensitive) IsSensitive() bool {
	if s.sensitive == nil {
		return false
	}

	return *s.sensitive
}

func (s Sensitive) Schema() []byte {
	if s.IsSensitive() {
		return []byte("Sensitive: true,\n")
	}

	return nil
}
