package p

//go:generate group-interface -f example.go

type IPv4 [4]byte
type IPv6 [16]byte

type IP interface {
	GroupTypes(IPv4, IPv6)
}

type NoGroup interface {
	GroupTypes(IPv4, IPv6) int
}

type TypeWithNoName interface {
	GroupTypes(int, IPv4)
}
