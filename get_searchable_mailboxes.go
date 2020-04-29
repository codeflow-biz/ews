package ews

import (
    "encoding/xml"
    "errors"
)

type GetSearchableMailboxesRequest struct {
    XMLName struct{} `xml:"m:GetSearchableMailboxes"`
}

type getSearchableMailboxesResponseEnvelop struct {
    XMLName struct{}                           `xml:"Envelope"`
    Body    getSearchableMailboxesResponseBody `xml:"Body"`
}
type getSearchableMailboxesResponseBody struct {
    SearchableMailboxesResponse GetSearchableMailboxesResponse `xml:"GetSearchableMailboxesResponse"`
}

type GetSearchableMailboxesResponse struct {
    Response
    AllSearchableMailboxes GetSearchableMailboxResponse `xml:"SearchableMailboxes"`
}

type GetSearchableMailboxResponse struct {
    SearchableMailboxes []SearchableMailbox `xml:"SearchableMailbox"`
}

type SearchableMailbox struct {
    Guid                 string `xml:"Guid"`
    PrimarySmtpAddress   string `xml:"PrimarySmtpAddress"`
    IsExternalMailbox    bool   `xml:"IsExternalMailbox"`
    ExternalEmailAddress string `xml:"ExternalEmailAddress"`
    DisplayName          string `xml:"DisplayName"`
    IsMembershipGroup    bool   `xml:"IsMembershipGroup"`
    ReferenceId          string `xml:"ReferenceId"`
}

func GetSearchableMailboxes(c Client, r *GetSearchableMailboxesRequest) (*GetSearchableMailboxesResponse, error) {

    xmlBytes, err := xml.MarshalIndent(r, "", "  ")
    if err != nil {
        return nil, err
    }

    bb, err := c.SendAndReceive(xmlBytes)
    if err != nil {
        return nil, err
    }

    var soapResp getSearchableMailboxesResponseEnvelop
    err = xml.Unmarshal(bb, &soapResp)
    if err != nil {
        return nil, err
    }

    if soapResp.Body.SearchableMailboxesResponse.ResponseClass == ResponseClassError {
        return nil, errors.New(soapResp.Body.SearchableMailboxesResponse.MessageText)
    }

    return &soapResp.Body.SearchableMailboxesResponse, nil
}