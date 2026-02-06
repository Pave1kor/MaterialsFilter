package path

type Resolver struct {
	BaseDir string
}

func New(baseDir string) *Resolver {
	return &Resolver{BaseDir: baseDir}
}
