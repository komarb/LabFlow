package utils

func FormatQueryDate(date string) string {
	return date[6:10]+"-"+date[3:5]+"-"+date[0:2] //05.01.2020
}