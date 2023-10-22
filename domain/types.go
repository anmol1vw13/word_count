package domain

type FlagSet struct {
	Char, Word, Line bool
}

type Counter struct {
	Flag  *FlagSet
	Files []string
}

type CountInfo struct {
	Identifier string
	Char       int64
	Word       int64
	Line       int64
}
