package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weigle/http/domain"
	accountHeader "github.com/weigle/http/handle/account"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := mux.NewRouter()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         255,                                                                           // default size for string fields
		DisableDatetimePrecision:  true,                                                                          // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                          // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                          // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                         // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.PropertyValue{})
	accountHeader.CreateHandle(r)
	r.Use(commonMiddleware)
	http.ListenAndServe(":8080", r)

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
