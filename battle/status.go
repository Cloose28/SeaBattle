package battle

// ShotResult describe info after shot
type ShotResult struct {
	Destroy bool
	Knock bool
	End bool
}

func NewDestroyedShot(end bool) ShotResult {
	return ShotResult{true, true, end}
}

func NewKnockedShot() ShotResult {
	return ShotResult{false, true, false}
}