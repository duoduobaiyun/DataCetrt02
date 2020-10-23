package models

import (
	"DataCertProject_Me/db_mysql"
	"DataCertProject_Me/util"
	"fmt"
)

type UploadRecord struct {
	Id            int
	FileName      string
	FileSize      int64
	FileCert      string //认证号
	FileTitle     string
	CertTime      int64 //整型
	FormatCertTime string//格式化时间格式,该字段只在前端展示时使用
	Phone         string //对应的用户的phone
}

/*
*保存上传记录到数据库中
 */
func (u UploadRecord) SaveRecord() (int64, error) {
	result, err := db_mysql.Db.Exec("insert into upload_record(file_name,file_size,file_cert,file_title,cert_time,phone)"+
		"values(?,?,?,?,?,?)",
		u.FileName,
		u.FileSize,
		u.FileCert,
		u.FileTitle,
		u.CertTime,
		u.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return id, nil
}

//读取数据库中phone用户对应的所有认证数据
func QueryRecordBYPhone(phone string) ([]UploadRecord, error) {
	rs, err := db_mysql.Db.Query("select id,file_name,file_size,file_cert,file_title,cert_time,phone from upload_record where phone=?", phone)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	records := make([]UploadRecord, 0)
	for rs.Next() {
		var record UploadRecord
		err := rs.Scan(&record.Id, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime, &record.Phone)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		//时间转换record.CertTime
		record.FormatCertTime=util.TimeFormat(record.CertTime,0,util.TIME_FORMAT_THREE)
		records = append(records, record)
	}
	return records, nil
}
