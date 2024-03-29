package ebct

import (
	"errors"
	"fmt"
	"time"
)

const trackingnumbers = DefaultEndPoint + Version + "packages/tracking-numbers"
const packageasync = DefaultEndPoint + Version + "packages/async"

type TrackingNumbers struct {
	RequestId            *string       `json:"requestId,omitempty"`            // Request number generated by calling the Tracking Number Range Request method. This value is returned in the "requestId" key of this method.
	RequestDate          *string       `json:"requestDate,omitempty"`          // Date and time the request was created. It is in the pattern "YYYY-MM-DDThh:mm:ss.sss".
	DistributionModality *int          `json:"distributionModality,omitempty"` // Type of distribution in the Brazilian territory, as informed in the requisition. It must return one of the values: 33162 - PACKET STANDARD 33170 - PACKET EXPRESS 33197 - PACKET MINI
	InitialRangeCode     *string       `json:"initialRangeCode,omitempty"`     // First tracking number in the sequence.
	FinalRangeCode       *string       `json:"finalRangeCode,omitempty"`       // Last tracking number in the sequence.
	RequestedQuantity    *int          `json:"requestedQuantity,omitempty"`    // Quantity of tracking numbers requested.
	TotalUsed            *int          `json:"totalUsed,omitempty"`            // Quantity of tracking numbers, in this range, already used.
	TotalUnused          *int          `json:"totalUnused,omitempty"`          // Quantity of tracking numbers, in this range, not yet used. It is important to note that new requests must be made only when this value is close to or equal to the "requestedQuantity" field, which indicates that the requested numbers are being properly used.
	Links                []string      `json:"links,omitempty"`                // May contain a link to specific requests for queries
	ErrorMessage         *string       `json:"errorMessage,omitempty"`
	ErrorDetails         []ErrorDetail `json:"errorDetails,omitempty"`
}

func (n TrackingNumbers) Error() string {
	panic("implement me")
}

type TrackingNumbersInput struct {
	DistributionModality *int `json:"distributionModality,omitempty"` // Type of distribution in the Brazilian territory, as informed in the requisition. It must return one of the values: 33162 - PACKET STANDARD 33170 - PACKET EXPRESS 33197 - PACKET MINI
	Quantity             *int `json:"quantity,omitempty"`
}

type GetTrackingNumbers struct {
	TrackingNumber []string `json:"trackingNumber,omitempty"` // List of tracking numbers generated
}

type TrackingNumbersReturn struct {
	TrackingNumbers []string `json:"trackingNumbers,omitempty"` // List of tracking numbers generated
}

func (n TrackingNumbersReturn) Error() string {
	panic("implement me")
}

type PackList struct {
	PackageList []PackageTrack `json:"packageList,omitempty"`
}

type PackageTrack struct {
	TrackingNumber             *string  `json:"trackingNumber,omitempty"`
	SenderName                 *string  `json:"senderName,omitempty"`
	SenderAddress              *string  `json:"senderAddress,omitempty"`
	SenderAddressNumber        *string  `json:"senderAddressNumber,omitempty"`
	SenderAddressComplement    *string  `json:"senderAddressComplement,omitempty"`
	SenderZipCode              *string  `json:"senderZipCode,omitempty"`
	SenderCityName             *string  `json:"senderCityName,omitempty"`
	SenderState                *string  `json:"senderState,omitempty"`
	SenderCountryCode          *string  `json:"senderCountryCode,omitempty"`
	SenderEmail                *string  `json:"senderEmail,omitempty"`
	SenderWebsite              *string  `json:"senderWebsite,omitempty"`
	RecipientDocumentType      *string  `json:"recipientDocumentType,omitempty"`
	RecipientDocumentNumber    *string  `json:"recipientDocumentNumber,omitempty"`
	RecipientName              *string  `json:"recipientName,omitempty"`
	RecipientAddress           *string  `json:"recipientAddress,omitempty"`
	RecipientAddressNumber     *string  `json:"recipientAddressNumber,omitempty"`
	RecipientAddressComplement *string  `json:"recipientAddressComplement,omitempty"`
	RecipientZipCode           *string  `json:"recipientZipCode,omitempty"`
	RecipientCityName          *string  `json:"recipientCityName,omitempty"`
	RecipientState             *string  `json:"recipientState,omitempty"`
	RecipientEmail             *string  `json:"recipientEmail,omitempty"`
	RecipientPhoneNumber       *string  `json:"recipientPhoneNumber,omitempty"`
	Currency                   *string  `json:"currency,omitempty"`
	TotalWeight                *int     `json:"totalWeight,omitempty"`
	DistributionModality       *int     `json:"distributionModality,omitempty"`
	PackagingLength            *float64 `json:"packagingLength,omitempty"`
	PackagingWidth             *float64 `json:"packagingWidth,omitempty"`
	PackagingHeight            *float64 `json:"packagingHeight,omitempty"`
	TaxPaymentMethod           *string  `json:"taxPaymentMethod,omitempty"`
	FreightPaidValue           *float64 `json:"freightPaidValue,omitempty"`
	InsurancePaidValue         *float64 `json:"insurancePaidValue,omitempty"`
	Items                      []Item   `json:"items,omitempty"`
}

func (n PackageTrack) Error() string {
	panic("implement me")
}

type PackageAsyncReturn struct {
	RequestId     *string `json:"requestId,omitempty"`     // Value generated to identify the request. This data must be passed as the value of the key "requestid" of the GET method "Get unit async request info", used to obtain the unit number.
	RequestStatus *string `json:"requestStatus,omitempty"` // Contains the status of the request. Pending: Pending processing; Processing: In processing; Success: Processing was completed successfully. Error: An error has occurred in the processing of the information. The error message is displayed in the “errorMessage” field.
}

func (n PackageAsyncReturn) Error() string {
	panic("implement me")
}

type TrackingNumbersInputDate struct {
	InitialRequestDate *string `json:"initialRequestDate,omitempty"` // Start date in pattern YYYY-MM-DD
	FinalRequestDate   *string `json:"finalRequestDate,omitempty"`   // End date in pattern YYYY-MM-DD
}

func (c *Client) GetTrackingNumbers() (*TrackingNumbers, error) {

	newTrackingNumbers := &TrackingNumbers{}
	err := c.Get(trackingnumbers, nil, nil, newTrackingNumbers)
	return newTrackingNumbers, err
}

func (c *Client) PostTrackingNumbers(createtracks *TrackingNumbersInput) (*TrackingNumbersReturn, error) {

	newTrackingNumbersReturn := &TrackingNumbersReturn{}
	err := c.Post(trackingnumbers, createtracks, nil, newTrackingNumbersReturn)
	return newTrackingNumbersReturn, err
}

func (c *Client) PutTrackingNumbers(createtracks *PackageTrack, tracknumber string) (*PackageTrack, error) {

	newTrackingNumbersReturn := &PackageTrack{}
	err := c.Put(packagesRequest+"/"+tracknumber, createtracks, nil, newTrackingNumbersReturn)
	return newTrackingNumbersReturn, err
}

func (c *Client) PutPackageAsync(createtracks *PackList) (*PackageAsyncReturn, error) {

	newTrackingNumbersReturn := &PackageAsyncReturn{}
	err := c.Put(packageasync, createtracks, nil, newTrackingNumbersReturn)
	return newTrackingNumbersReturn, err
}

func (c *Client) GetPackageAsync(requestId string) (*PackageAsyncReturn, error) {

	sendRequest := &Cn38RequestGetAsync{}
	var err error

	if !IsUUID(requestId) {
		err = errors.New("the requestId must be a valid UUID")
		return nil, err
	}
	sendRequest.RequestId = StringPtr(requestId)

	newTrackingNumbersReturn := &PackageAsyncReturn{}
	err = c.Get(packageasync, sendRequest, nil, newTrackingNumbersReturn)
	return newTrackingNumbersReturn, err

}

func (c *Client) GetPackageTrackingNumbersId(tracknumbers ...string) (*PackList, error) {

	newTrackingNumbers := &PackList{}
	sendTrackingNumbers := &GetTrackingNumbers{}

	if len(tracknumbers) == 0 {
		err := errors.New("1 or more tracking number must be sent")
		return nil, err
	}

	for _, sTracking := range tracknumbers {
		sendTrackingNumbers.TrackingNumber = append(sendTrackingNumbers.TrackingNumber, sTracking)
	}

	err := c.Get(packagesRequest, sendTrackingNumbers, nil, newTrackingNumbers)
	return newTrackingNumbers, err
}

func (c *Client) GetPackageTrackingNumbersList(InitialRequestDate, finalRequestDate time.Time) ([]TrackingNumbers, error) {

	const formatdate = "2006-01-02"
	const formatdate2 = "02/01/2006" // colocado isso aqui pq o 2006-01-02 retorna erro... e essa data assim aceita...
	sendRequest := &TrackingNumbersInputDate{}
	sendRequest.InitialRequestDate = StringPtr(fmt.Sprint(InitialRequestDate.Format(formatdate2)))
	sendRequest.FinalRequestDate = StringPtr(fmt.Sprint(finalRequestDate.Format(formatdate2)))

	var newTrackingNumbersReturn []TrackingNumbers
	err := c.Get(trackingnumbers, sendRequest, nil, newTrackingNumbersReturn)
	return newTrackingNumbersReturn, err

}
