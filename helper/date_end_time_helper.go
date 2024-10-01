package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/TigorLazuardi/tanggal"
)

func StringToDate(value string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func StringToDateSepesific(value string) (time.Time, error) {
	var dateNewRow time.Time

	date, err := time.Parse("01/02/2006", value)
	if err != nil {
		return dateNewRow, err
	}
	dateNew := date.Format("2006-01-02")
	tanggals, err := time.Parse("2006-01-02", dateNew)
	if err != nil {
		return dateNewRow, err
	}
	return tanggals, nil
}

func DateToFormatIndo(value time.Time) (time.Time, error) {
	dateString := value.Format("02-01-2006")
	tanggal, err := time.Parse(dateString, "02-01-2006")
	if err != nil {
		return time.Time{}, err
	}
	return tanggal, nil
}

func DatetimeToFormatIndo(value time.Time) (time.Time, error) {
	var dateNewRow time.Time

	dateString := value.Format("2006-01-02 15:04:05")

	date, err := time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		return dateNewRow, err
	}

	dateNew := date.Format("02-01-2006 15:04:05")
	tanggals, err := time.Parse("02-01-2006 15:04:05", dateNew)
	if err != nil {
		return dateNewRow, err
	}

	return tanggals, nil
}

func StringToDateTimeIndoFormat(value string) (time.Time, error) {
	var dateNewRow time.Time

	// dateString := value.Format("2006-01-02 15:04:05")
	d := strings.Split(value, " ")
	dateString := fmt.Sprintf("%s %s", d[0], d[1])
	dtow := strings.Split(dateString, ".")
	dtowa := strings.Split(dtow[0], "{")

	date, err := time.Parse("2006-01-02 15:04:05", dtowa[1])
	if err != nil {
		return dateNewRow, err
	}
	dateNew := date.Format("02-01-2006 15:04:05")
	tanggals, _ := time.Parse("02-01-2006 15:04:05", dateNew)
	if err != nil {
		return dateNewRow, err
	}
	return tanggals, nil
}

func TempatTanggalLahirFormatIndonesia(tempat string, waktu time.Time) (string, error) {
	tanggalFormat, err := tanggal.Papar(waktu, "Jakarta", tanggal.WIB)
	if err != nil {
		return "", err
	}

	format := []tanggal.Format{
		tanggal.Hari,
		tanggal.NamaBulan,
		tanggal.Tahun,
	}
	tanggal := tanggalFormat.Format(" ", format)

	data := fmt.Sprintln(tempat, ",", tanggal)
	return data, nil
}

func IndonesiaFormat(waktu time.Time) (string, error) {
	tanggalFormat, err := tanggal.Papar(waktu, "Jakarta", tanggal.WIB)
	if err != nil {
		return "", err
	}
	format := []tanggal.Format{
		tanggal.Hari,
		tanggal.NamaBulan,
		tanggal.Tahun,
	}

	tanggal := tanggalFormat.Format(" ", format)
	return tanggal, nil
}
