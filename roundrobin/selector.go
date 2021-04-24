package roundrobin

type Selector struct {
	values  []string
	counter int
}

func NewSelector(values []string) *Selector {
	return &Selector{
		values:  values,
		counter: 0,
	}
}

func (s *Selector) Get() string {
	val := s.values[s.counter]

	s.counter++

	if s.counter == len(s.values) {
		s.counter = 0
	}

	return val
}
