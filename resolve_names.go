package ews

import (
    "encoding/xml"
    "errors"
)

type ResolveNamesRequest struct {
    XMLName               xml.Name `xml:"m:ResolveNames"`
    ReturnFullContactData bool     `xml:"ReturnFullContactData,attr,omitempty"`
    UnresolvedEntry       string   `xml:"m:UnresolvedEntry"`
}

type ResolveNamesResponseEnvelop struct {
    XMLName struct{}                 `xml:"Envelope"`
    Body    ResolveNamesResponseBody `xml:"Body"`
}
type ResolveNamesResponseBody struct {
    ResolveNamesResponseMessage ResolveNamesResponse `xml:"ResolveNamesResponse>ResponseMessages>ResolveNamesResponseMessage"`
}
type ResolveNamesResponseResponse struct {
    ResponseMessages ResolveNamesResponse `xml:"ResponseMessages"`
}
type ResolveNamesResponse struct {
    Response
    ResolutionSet ResolveNamesResolutionSetResponse `xml:"ResolutionSet"`
}
type ResolveNamesResolutionSetResponse struct {
    ResultPagingInfo
    Resolutions             []Resolution `xml:"Resolution"`
}

type Resolution struct {
    Mailbox ResolutionMailbox `xml:"Mailbox"`
    Contact Contact           `xml:"Contact"`
}
type ResolutionMailbox struct {
    Name         string `xml:"Name"`
    EmailAddress string `xml:"EmailAddress"`
    RoutingType  string `xml:"RoutingType"`
    MailboxType  string `xml:"MailboxType"`
}
type Contact struct {
    DisplayName       string                        `xml:"DisplayName"`
    GivenName         string                        `xml:"GivenName"`
    Surname           string                        `xml:"Surname"`
    Culture           string                        `xml:"Culture"`
    Initials          string                        `xml:"Initials"`
    CompanyName       string                        `xml:"CompanyName"`
    AssistantName     string                        `xml:"AssistantName"`
    ContactSource     string                        `xml:"ContactSource"`
    Department        string                        `xml:"Department"`
    JobTitle          string                        `xml:"JobTitle"`
    OfficeLocation    string                        `xml:"OfficeLocation"`
    EmailAddresses    []ContactEntry                `xml:"EmailAddresses>Entry"`
    PhoneNumbers      []ContactEntry                `xml:"PhoneNumbers>Entry"`
    PhysicalAddresses []ContactPhysicalAddressEntry `xml:"PhysicalAddresses>Entry"`
}
type ContactEntry struct {
    Key   string `xml:"Key,attr"`
    Value string `xml:",chardata"`
}
type ContactPhysicalAddressEntry struct {
    Key             string `xml:"Key,attr"`
    Street          string `xml:"Street"`
    City            string `xml:"City"`
    State           string `xml:"State"`
    CountryOrRegion string `xml:"CountryOrRegion"`
    PostalCode      string `xml:"PostalCode"`
}
type ResultPagingInfo struct {
    IndexedPagingOffset     int  `xml:"IndexedPagingOffset,attr"`
    TotalItemsInView        int  `xml:"TotalItemsInView,attr"`
    IncludesLastItemInRange bool `xml:"IncludesLastItemInRange,attr"`
}

func ResolveNames(c Client, r *ResolveNamesRequest) (*ResolveNamesResponse, error) {

    xmlBytes, err := xml.MarshalIndent(r, "", "  ")
    if err != nil {
        return nil, err
    }

    bb, err := c.SendAndReceive(xmlBytes)
    if err != nil {
        return nil, err
    }

    var soapResp ResolveNamesResponseEnvelop
    err = xml.Unmarshal(bb, &soapResp)
    if err != nil {
        return nil, err
    }

    if soapResp.Body.ResolveNamesResponseMessage.ResponseClass == ResponseClassError {
        return nil, errors.New(soapResp.Body.ResolveNamesResponseMessage.MessageText)
    }

    return &soapResp.Body.ResolveNamesResponseMessage, nil
}
