package convert

type WriteOnly struct {
	write_only *bool
}

func NewWriteOnly(wo *bool) WriteOnly {
	return WriteOnly{
		write_only: wo,
	}
}

func (wo WriteOnly) Equal(other WriteOnly) bool {
	if wo.write_only == nil && other.write_only == nil {
		return true
	}

	if wo.write_only == nil || other.write_only == nil {
		return false
	}

	return *wo.write_only == *other.write_only
}

func (wo WriteOnly) IsWriteOnly() bool {
	if wo.write_only == nil {
		return false
	}

	return *wo.write_only
}

func (wo WriteOnly) Schema() []byte {
	if wo.IsWriteOnly() {
		return []byte("WriteOnly: true,\n")
	}

	return nil
}
