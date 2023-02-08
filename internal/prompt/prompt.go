package prompt

type Tui interface {
	Tui() (int, string, error)
}

func Start(tuis ...Tui) ([]string, error) {
	var results []string
	for _, e := range tuis {
		_, result, err := e.Tui()
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
