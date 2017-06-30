package todo

import (
	"os"
	"time"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestAddWork(c *C) {
	// given
	s.chores.Add("hello world", TYPE_DAILY)
	s.chores.Add("hello galaxy", TYPE_DAILY)
	s.chores.Add("hello universe", TYPE_DAILY)

	t1, _ := time.Parse("2006-01-02T15:04:05.000Z", "2016-01-02T15:04:05.000Z")
	t2, _ := time.Parse("2006-01-02T15:04:05.000Z", "2017-01-02T15:04:05.000Z")

	_, err1 := s.tasks.LogWork(1, time.Now())
	_, err2 := s.tasks.LogWork(2, t1)
	_, err3 := s.tasks.LogWork(2, t2)

	// when
	found, _ := s.tasks.FindAll()
	// then

	// chore order, created asc
	c.Assert(err1, Equals, nil)
	c.Assert(err2, Equals, nil)
	c.Assert(err3, Equals, nil)
	c.Assert(len(found[0].Tasks), Equals, 1)
	c.Assert(len(found[1].Tasks), Equals, 2)
	c.Assert(len(found[2].Tasks), Equals, 0)

	// test work ordering, created desc
	c.Assert(found[1].Tasks[0].Time, Equals, t2)
	c.Assert(found[1].Tasks[1].Time, Equals, t1)
}

func (s *TestSuite) TestDeleteWork(c *C) {
	// given
	s.chores.Add("hello world", TYPE_DAILY)

	work, _ := s.tasks.LogWork(1, time.Now())
	os.Exit(0)
	// when
	err := s.tasks.DeleteWork(work.Id)

	// then
	found, _ := s.tasks.FindAll()

	c.Assert(err, Equals, nil)
	c.Assert(len(found[0].Tasks), Equals, 0)
}
