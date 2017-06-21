package todo

import (
	"time"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestAddWork(c *C) {
	// given
	s.chores.Add("hello world", TYPE_DAILY)
	s.chores.Add("hello galaxy", TYPE_DAILY)
	s.chores.Add("hello universe", TYPE_DAILY)

	s.tasks.LogWork(1, time.Now())
	s.tasks.LogWork(2, time.Now())
	s.tasks.LogWork(2, time.Now())

	// when
	found, err := s.tasks.FindAllDaily()

	// then
	c.Assert(err, Equals, nil)
	c.Assert(len(found[0].Tasks), Equals, 1)
	c.Assert(len(found[1].Tasks), Equals, 2)
	c.Assert(len(found[2].Tasks), Equals, 0)
}

func (s *TestSuite) TestDeleteWork(c *C) {
	// given
	s.chores.Add("hello world", TYPE_DAILY)

	work, _ := s.tasks.LogWork(1, time.Now())

	// when
	err := s.tasks.DeleteWork(work.Id)

	// then
	found, _ := s.tasks.FindAllDaily()

	c.Assert(err, Equals, nil)
	c.Assert(len(found[0].Tasks), Equals, 0)
}

func (s *TestSuite) TestChoreTypeFiltering(c *C) {
	// given
	s.chores.Add("hello world", TYPE_DAILY)
	s.chores.Add("hello galaxy", TYPE_WEEKLY)
	s.chores.Add("hello universe", TYPE_WEEKLY)

	s.tasks.LogWork(1, time.Now())
	s.tasks.LogWork(2, time.Now())
	s.tasks.LogWork(3, time.Now())

	// when
	daily, _ := s.tasks.FindAllDaily()
	weekly, _ := s.tasks.FindAllWeekly()
	monthly, _ := s.tasks.FindAllMonthly()

	// then
	c.Assert(len(daily), Equals, 1)
	c.Assert(len(weekly), Equals, 2)
	c.Assert(len(monthly), Equals, 0)
}
