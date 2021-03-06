package d7024e

import (
	"fmt"
	"sort"
	"strings"
)

// Contact definition
// stores the KademliaID, the ip address and the distance
type Contact struct {
	ID       *KademliaID
	Address  string
	distance *KademliaID
}

// NewContact returns a new instance of a Contact
func NewContact(id *KademliaID, address string) Contact {
	return Contact{id, address, nil}
}
//Added code from us TODO actually test it
func RestoreContact(contact string) Contact {
	split := strings.Split(contact, "\"")
	id := split[1]
	addr := split[3]
	return NewContact(NewKademliaID(id), addr)
}
func (candidates *ContactCandidates) Exists(contact *Contact) bool {
	for _, ct := range candidates.contacts {
		if ct.ID.String() == contact.ID.String() {
			return true
		}
	}
	return false
}
func (candidates *ContactCandidates) Remove(id string) {
	index := -1
	for i, ct := range candidates.contacts {
		if ct.ID.String() == id {
			index = i
		}
	}
	if index != -1 {
		candidates.contacts[index] = candidates.contacts[len(candidates.contacts)-1]
		candidates.contacts = candidates.contacts[:len(candidates.contacts)-1]
	}
}
func (candidates *ContactCandidates) Sorted() bool {
	last := candidates.Len() - 1
	for i, ct := range candidates.contacts {
		if i == last {
			return true
		}
		if !(ct.Less(&candidates.contacts[i+1])) {
			return false
		}
	}
	return true
}

// CalcDistance calculates the distance to the target and 
// fills the contacts distance field
func (contact *Contact) CalcDistance(target *KademliaID) {
	contact.distance = contact.ID.CalcDistance(target)
}

// Less returns true if contact.distance < otherContact.distance
func (contact *Contact) Less(otherContact *Contact) bool {
	return contact.distance.Less(otherContact.distance)
}

// String returns a simple string representation of a Contact
func (contact *Contact) String() string {
	return fmt.Sprintf(`contact("%s", "%s")`, contact.ID, contact.Address)
}

// ContactCandidates definition
// stores an array of Contacts
type ContactCandidates struct {
	contacts []Contact
}

// Append an array of Contacts to the ContactCandidates
func (candidates *ContactCandidates) Append(contacts []Contact) {
	candidates.contacts = append(candidates.contacts, contacts...)
}

// GetContacts returns the first count number of Contacts
func (candidates *ContactCandidates) GetContacts(count int) []Contact {
	return candidates.contacts[:count]
}

// Sort the Contacts in ContactCandidates
func (candidates *ContactCandidates) Sort() {
	sort.Sort(candidates)
}

// Len returns the length of the ContactCandidates
func (candidates *ContactCandidates) Len() int {
	return len(candidates.contacts)
}

// Swap the position of the Contacts at i and j
// WARNING does not check if either i or j is within range
func (candidates *ContactCandidates) Swap(i, j int) {
	candidates.contacts[i], candidates.contacts[j] = candidates.contacts[j], candidates.contacts[i]
}

// Less returns true if the Contact at index i is smaller than 
// the Contact at index j
func (candidates *ContactCandidates) Less(i, j int) bool {
	return candidates.contacts[i].Less(&candidates.contacts[j])
}
