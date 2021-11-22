package twitter

func Tweet(content string) error {
	_, _, err := client.Statuses.Update(content, nil)
	return err
}
