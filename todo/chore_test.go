package todo

import (
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestAddChore(c *C) {
	// given
	expected := &Chore{Id: 1, Content: "hello world", Type: TYPE_DAILY}

	// when
	created, err := s.chores.Add("hello world", TYPE_DAILY)

	// then
	c.Assert(err, Equals, nil)

	found, _ := s.chores.FindAll()

	c.Assert(created, DeepEquals, expected)
	c.Assert(found[0], DeepEquals, expected)
}

func (s *TestSuite) TestFindAllChores(c *C) {
	// given
	expected := []*Chore{
		&Chore{Id: 1, Content: "hello world", Type: TYPE_DAILY},
		&Chore{Id: 2, Content: "hello universe", Type: TYPE_WEEKLY},
		&Chore{Id: 3, Content: "hello galaxy", Type: TYPE_MONTHLY},
	}

	s.chores.Add("hello world", TYPE_DAILY)
	s.chores.Add("hello universe", TYPE_WEEKLY)
	s.chores.Add("hello galaxy", TYPE_MONTHLY)

	// when
	found, err := s.chores.FindAll()

	// then
	c.Assert(err, Equals, nil)

	c.Assert(found, DeepEquals, expected)
}

func (s *TestSuite) TestDeleteChore(c *C) {
	// given
	chore, _ := s.chores.Add("hello world", TYPE_DAILY)

	// when
	err := s.chores.Delete(chore.Id)

	// then
	c.Assert(err, Equals, nil)

	chores, _ := s.chores.FindAll()

	c.Assert(len(chores), Equals, 0)
}
