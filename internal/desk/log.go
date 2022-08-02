package desk

type Logs struct {
	List []Log
	Last int
}

type Log struct {
	String    string
	Side      string
	From      string
	To        string
	Separator string
}

func (ls Logs) String() string {
	s := ""
	for _, l := range ls.List {
		s += l.String + "\n"
	}
	return s
}

func (ls *Logs) Add(l Log) {
	ls.List = append(ls.List, l)
}

func (ls *Logs) Error(l Log) {
	if len(ls.List) > 0 {
		ls.List[0] = l
	}
}

var EmptyLog = Log{String: "___________________"}
