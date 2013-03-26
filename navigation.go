// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

type NavigationIterator struct {
	depth  []navpoint
	navMap []navpoint
	index  int
}

func newNavigationIterator(navMap []navpoint) *NavigationIterator {
	var nav NavigationIterator
	nav.navMap = navMap
	nav.index = 0
	return &nav
}

func (nav NavigationIterator) Title() string {
	return nav.item().Title()
}

func (nav NavigationIterator) Url() string {
	return nav.item().Url()
}

func (nav NavigationIterator) HasChildren() bool {
	return nav.item().Children() != nil
}

func (nav NavigationIterator) HasParents() bool {
	return nav.depth != nil
}

func (nav NavigationIterator) IsFirst() bool {
	return nav.index == 0
}

func (nav NavigationIterator) IsLast() bool {
	return nav.index == len(nav.navMap)
}

func (nav NavigationIterator) item() *navpoint {
	return &nav.navMap[nav.index]
}
