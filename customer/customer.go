package customer

import (
	"api-cyclesoft/database"
	"context"
	"strconv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	CustomerCode              string  `json:"CustomerCode"`
	CustomerName              string  `json:"CustomerName"`
	CustomerAddress1          string  `json:"CustomerAddress1"`
	CustomerAddress2          any  `json:"CustomerAddress2"`
	CustomerPhone             any  `json:"CustomerPhone"`
	CustomerFax               any  `json:"CustomerFax"`
	CustomerSubAreaCode       any  `json:"CustomerSubAreaCode"`
	CustomerGroupCode         any  `json:"CustomerGroupCode"`
	CustomerTypeCode          any  `json:"CustomerTypeCode"`
	CustomerRegionCode        any  `json:"CustomerRegionCode"`
	CustomerDeliveryArea      any  `json:"CustomerDeliveryArea"`
	CustomerDriverCode        any  `json:"CustomerDriverCode"`
	CustomerDOAddress1        any  `json:"CustomerDOAddress1"`
	CustomerDOAddress2        any  `json:"CustomerDOAddress2"`
	CustomerInvAddress1       any  `json:"CustomerInvAddress1"`
	CustomerInvAddress2       any  `json:"CustomerInvAddress2"`
	CustomerNPPKP             any  `json:"CustomerNPPKP"`
	CustomerNPWP              any  `json:"CustomerNPWP"`
	CustomerNPWPDate          any  `json:"CustomerNPWPDate"`
	CustomerNPWPName          any  `json:"CustomerNPWPName"`
	CustomerNPWPAddress1      any  `json:"CustomerNPWPAddress1"`
	CustomerNPWPAddress2      any  `json:"CustomerNPWPAddress2"`
	CustomerNPWPAddress3      any  `json:"CustomerNPWPAddress3"`
	CustomerPKP               any    `json:"CustomerPKP"`
	CustomerPPN               any `json:"CustomerPPN"`
	CustomerCreditLimit       any `json:"CustomerCreditLimit"`
	CustomerTOP               any     `json:"CustomerTOP"`
	CustomerOutstanding       any `json:"CustomerOutstanding"`
	CustomerBankCode          any  `json:"CustomerBankCode"`
	CustomerAccNo             any  `json:"CustomerAccNo"`
	CustomerAccName           any  `json:"CustomerAccName"`
	CustomerBankBranch        any  `json:"CustomerBankBranch"`
	CustomerCollectorCode     any  `json:"CustomerCollectorCode"`
	CustomerDeleted           any    `json:"CustomerDeleted"`
	DefaultCreditLimit        any `json:"DefaultCreditLimit"`
	BlackListReason           any  `json:"BlackListReason"`
	Flag_Blocked              any  `json:"Flag_Blocked"`
	CustomerFakturis          any  `json:"CustomerFakturis"`
	CustomerInvoiceLimit      any `json:"CustomerInvoiceLimit"`
	ExternalCustomerGroupCode any  `json:"ExternalCustomerGroupCode"`
	Area                      any  `json:"area"`
	IsHoldSO                  any    `json:"IsHoldSO"`
	HoldSOReason              any  `json:"HoldSOReason"`
	CreatedDate               any  `json:"CreatedDate"`
	VirtualAccount            any  `json:"VirtualAccount"`
	StatusCustomer            any    `json:"StatusCustomer"`
	CustNonActiveDate         any  `json:"CustNonActiveDate"`
	CustomerZipcode           any  `json:"CustomerZipcode"`
	CustomerBookingCode       any  `json:"CustomerBookingCode"`
	NamaPemilik               any  `json:"namaPemilik"`
	AlamatPemilik             any  `json:"alamatPemilik"`
	NoRek                     any  `json:"noRek"`
	AtasNama                  any  `json:"atasNama"`
	Since                     any  `json:"since"`
	Longitude                 any  `json:"longitude"`
	Latitude                  any  `json:"latitude"`
	ContactPerson             any  `json:"contactPerson"`
	Keterangan                any  `json:"keterangan"`
}

func GetCustomers(c *gin.Context) {
	// Here's an array in which you can store the decoded documents
	var results []*Customer

	query := c.Request.URL.Query()
	
	page := 1

	offset := 0
	limit := 100

	for key, values := range query {
		var err error
		switch key {
		case "limit":

			cc := values[0]
			limit, err = strconv.Atoi(cc)

			if err != nil {
				panic(err)
			}
		case "page":
			dd := values[0]
			page, err = strconv.Atoi(dd)

			if err != nil {
				panic(err)
			}
		}

		
    }

	offset = (page - 1) * limit

	db := database.GetDatabaseInstance()

	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	tsql := fmt.Sprintf("SELECT CustomerCode, CustomerName, CustomerAddress1, CustomerAddress2, CustomerPhone, CustomerFax, CustomerSubAreaCode, CustomerGroupCode, CustomerTypeCode, CustomerRegionCode, CustomerDeliveryArea, CustomerDriverCode, CustomerDOAddress1, CustomerDOAddress2, CustomerInvAddress1, CustomerInvAddress2, CustomerNPPKP, CustomerNPWP, CustomerNPWPDate, CustomerNPWPName, CustomerNPWPAddress1, CustomerNPWPAddress2, CustomerNPWPAddress3, CustomerPKP, CustomerPPN, CustomerCreditLimit, CustomerTOP, CustomerOutstanding, CustomerBankCode, CustomerAccNo, CustomerAccName, CustomerBankBranch, CustomerCollectorCode, CustomerDeleted, DefaultCreditLimit, BlackListReason, Flag_Blocked, CustomerFakturis, CustomerInvoiceLimit, ExternalCustomerGroupCode,Area,  IsHoldSO, HoldSOReason, CreatedDate,  VirtualAccount,  StatusCustomer, CustNonActiveDate,CustomerZipcode, CustomerBookingCode,  NamaPemilik, AlamatPemilik, NoRek, AtasNama, Since, Longitude, Latitude, ContactPerson, Keterangan FROM %s.dbo.CustomerParameter ORDER BY CustomerCode OFFSET %d ROWS FETCH NEXT %d ROWS ONLY;", database.DB_DATABASE, offset, limit)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var customer Customer

		// Get values from row.
		err := rows.Scan(
			&customer.CustomerCode,
			&customer.CustomerName,
			&customer.CustomerAddress1,
			&customer.CustomerAddress2,
			&customer.CustomerPhone,
			&customer.CustomerFax,
			&customer.CustomerSubAreaCode,
			&customer.CustomerGroupCode,
			&customer.CustomerTypeCode,
			&customer.CustomerRegionCode,
			&customer.CustomerDeliveryArea,
			&customer.CustomerDriverCode,
			&customer.CustomerDOAddress1,
			&customer.CustomerDOAddress2,
			&customer.CustomerInvAddress1,
			&customer.CustomerInvAddress2,
			&customer.CustomerNPPKP,
			&customer.CustomerNPWP,
			&customer.CustomerNPWPDate,
			&customer.CustomerNPWPName,
			&customer.CustomerNPWPAddress1,
			&customer.CustomerNPWPAddress2,
			&customer.CustomerNPWPAddress3,
			&customer.CustomerPKP,
			&customer.CustomerPPN,
			&customer.CustomerCreditLimit,
			&customer.CustomerTOP,
			&customer.CustomerOutstanding,
			&customer.CustomerBankCode,
			&customer.CustomerAccNo,
			&customer.CustomerAccName,
			&customer.CustomerBankBranch,
			&customer.CustomerCollectorCode,
			&customer.CustomerDeleted,
			&customer.DefaultCreditLimit,
			&customer.BlackListReason,
			&customer.Flag_Blocked,
			&customer.CustomerFakturis,
			&customer.CustomerInvoiceLimit,
			&customer.ExternalCustomerGroupCode,
			&customer.Area,
			&customer.IsHoldSO,
			&customer.HoldSOReason,
			&customer.CreatedDate,
			&customer.VirtualAccount,
			&customer.StatusCustomer,
			&customer.CustNonActiveDate,
			&customer.CustomerZipcode,
			&customer.CustomerBookingCode,
			&customer.NamaPemilik,
			&customer.AlamatPemilik,
			&customer.NoRek,
			&customer.AtasNama,
			&customer.Since,
			&customer.Longitude,
			&customer.Latitude,
			&customer.ContactPerson,
			&customer.Keterangan)
		if err != nil {
			panic(err)
		}
		results = append(results, &customer)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "OK", "data": results})
}

func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	var customer Customer

	db := database.GetDatabaseInstance()

	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	tsql := fmt.Sprintf("SELECT CustomerCode, CustomerName, CustomerAddress1, CustomerAddress2, CustomerPhone, CustomerFax, CustomerSubAreaCode, CustomerGroupCode, CustomerTypeCode, CustomerRegionCode, CustomerDeliveryArea, CustomerDriverCode, CustomerDOAddress1, CustomerDOAddress2, CustomerInvAddress1, CustomerInvAddress2, CustomerNPPKP, CustomerNPWP, CustomerNPWPDate, CustomerNPWPName, CustomerNPWPAddress1, CustomerNPWPAddress2, CustomerNPWPAddress3, CustomerPKP, CustomerPPN, CustomerCreditLimit, CustomerTOP, CustomerOutstanding, CustomerBankCode, CustomerAccNo, CustomerAccName, CustomerBankBranch, CustomerCollectorCode, CustomerDeleted, DefaultCreditLimit, BlackListReason, Flag_Blocked, CustomerFakturis, CustomerInvoiceLimit, ExternalCustomerGroupCode,Area,  IsHoldSO, HoldSOReason, CreatedDate,  VirtualAccount,  StatusCustomer, CustNonActiveDate,CustomerZipcode, CustomerBookingCode,  NamaPemilik, AlamatPemilik, NoRek, AtasNama, Since, Longitude, Latitude, ContactPerson, Keterangan  FROM %s.dbo.CustomerParameter WHERE CustomerCode='%s'", database.DB_DATABASE, id)

	row := db.QueryRowContext(ctx, tsql)

	row.Scan(&customer.CustomerCode,
		&customer.CustomerName,
		&customer.CustomerAddress1,
		&customer.CustomerAddress2,
		&customer.CustomerPhone,
		&customer.CustomerFax,
		&customer.CustomerSubAreaCode,
		&customer.CustomerGroupCode,
		&customer.CustomerTypeCode,
		&customer.CustomerRegionCode,
		&customer.CustomerDeliveryArea,
		&customer.CustomerDriverCode,
		&customer.CustomerDOAddress1,
		&customer.CustomerDOAddress2,
		&customer.CustomerInvAddress1,
		&customer.CustomerInvAddress2,
		&customer.CustomerNPPKP,
		&customer.CustomerNPWP,
		&customer.CustomerNPWPDate,
		&customer.CustomerNPWPName,
		&customer.CustomerNPWPAddress1,
		&customer.CustomerNPWPAddress2,
		&customer.CustomerNPWPAddress3,
		&customer.CustomerPKP,
		&customer.CustomerPPN,
		&customer.CustomerCreditLimit,
		&customer.CustomerTOP,
		&customer.CustomerOutstanding,
		&customer.CustomerBankCode,
		&customer.CustomerAccNo,
		&customer.CustomerAccName,
		&customer.CustomerBankBranch,
		&customer.CustomerCollectorCode,
		&customer.CustomerDeleted,
		&customer.DefaultCreditLimit,
		&customer.BlackListReason,
		&customer.Flag_Blocked,
		&customer.CustomerFakturis,
		&customer.CustomerInvoiceLimit,
		&customer.ExternalCustomerGroupCode,
		&customer.Area,
		&customer.IsHoldSO,
		&customer.HoldSOReason,
		&customer.CreatedDate,
		&customer.VirtualAccount,
		&customer.StatusCustomer,
		&customer.CustNonActiveDate,
		&customer.CustomerZipcode,
		&customer.CustomerBookingCode,
		&customer.NamaPemilik,
		&customer.AlamatPemilik,
		&customer.NoRek,
		&customer.AtasNama,
		&customer.Since,
		&customer.Longitude,
		&customer.Latitude,
		&customer.ContactPerson,
		&customer.Keterangan)

	if row.Err() != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Customer not found"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "OK", "data": customer})
	}
}

func CreateCustomer(c *gin.Context) {
	// var newAlbum Album

	// // Call BindJSON to bind the received JSON to newAlbum.
	// if err := c.BindJSON(&newAlbum); err != nil {
	// 	return
	// }

	// // Add the new album to the slice.
	// albums = append(albums, newAlbum)
	// c.IndentedJSON(http.StatusCreated, newAlbum)
}

func UpdateCustomer(c *gin.Context) {
	// var newAlbum Album

	// // Call BindJSON to bind the received JSON to newAlbum.
	// if err := c.BindJSON(&newAlbum); err != nil {
	// 	return
	// }

	// // Add the new album to the slice.
	// albums = append(albums, newAlbum)
	// c.IndentedJSON(http.StatusCreated, newAlbum)
}
