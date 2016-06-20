package events

type ArpEntryKey struct {
	IpAddr string
}

const (
	ArpEntryLearned EventId = 1
	ArpEntryDeleted EventId = 2
	ArpEntryUpdated EventId = 3
)

var ArpdEventKeyMap KeyMap = KeyMap{
	"ArpEntry": ArpEntryKey{},
}
