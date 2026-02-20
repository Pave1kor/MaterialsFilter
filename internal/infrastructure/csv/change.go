package csv

func (obj *CSVFile) ChangeHeadlines(newHeadlines []string) {
	obj.Headlines = newHeadlines
}
