package github

type Driver struct {
	parameterExtractor
}

func (driver Driver) Key() string {
	return "github"
}

func (driver Driver) Drive(url string) error {
	driver.parameterExtractor.extract(url)
	return nil
}
