package section2

import (
	"encoding/json"
	"os"
)

type Student struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Class struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

func ReadJson() ([]Class, error) {
	jsonString := `
	[
    {
      "name": "Class A",
      "students":  [
        {
          "name": "A1",
          "age": "22"
        },
        {
          "name": "A2",
          "age": "21"
        }
      ]
    },
    {
      "name": "Class B",
      "students":  [
        {
          "name": "B1",
          "age": "20"
        },
        {
          "name": "B2",
          "age": "21"
        }
      ]
    }
  ]
`

	var classes []Class
	err := json.Unmarshal([]byte(jsonString), &classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func WriteFile(classes []Class) error {
	data, err := json.Marshal(classes)
	if err != nil {
		return err
	}
	err = os.WriteFile("destination.json", data, 0644)
	return err
}
