package prompt

import "fmt"

type Tui interface {
	Tui() (int, string, error)
}

func Start(tuis ...Tui) ([]string, error) {
	results := make([]string, 0, len(tuis))

	for _, e := range tuis {
		_, result, err := e.Tui()
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		results = append(results, result)
	}

	return results, nil
}
