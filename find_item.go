package ews

import (
    "encoding/xml"
    "errors"
)

type FindItemRequest struct {
    XMLName             xml.Name            `xml:"m:FindItem"`
    Traversal           Traversal           `xml:"Traversal,attr,omitempty"`
    ItemShape           ItemShape           `xml:"m:ItemShape"`
    ParentFolderIds     ParentFolderId      `xml:"m:ParentFolderIds"`
    IndexedPageItemView IndexedPageItemView `xml:"m:IndexedPageItemView"`
}
type ItemShape struct {
    BaseShape            BaseShape            `xml:"t:BaseShape,omitempty"`
    AdditionalProperties AdditionalProperties `xml:"t:AdditionalProperties,omitempty"`
}

type findItemResponseEnvelop struct {
    XMLName xml.Name             `xml:"Envelope"`
    Body    findItemResponseBody `xml:"Body"`
}
type findItemResponseBody struct {
    FindItemResponse FindItemResponse `xml:"FindItemResponse>ResponseMessages>FindItemResponseMessage"`
}

type FindItemResponse struct {
    Response
    RootFolder FindItemRootfolder `xml:"RootFolder"`
}
type FindItemRootfolder struct {
    ResultPagingInfo
    Contacts []Contact `xml:"Items>Contact"`
}

func FindItem(c Client, r *FindItemRequest) (*FindItemResponse, error) {

    xmlBytes, err := xml.MarshalIndent(r, "", "  ")
    if err != nil {
        return nil, err
    }

    bb, err := c.SendAndReceive(xmlBytes)
    if err != nil {
        return nil, err
    }

    var soapResp findItemResponseEnvelop
    err = xml.Unmarshal(bb, &soapResp)
    if err != nil {
        return nil, err
    }

    if soapResp.Body.FindItemResponse.ResponseClass == ResponseClassError {
        return nil, errors.New(soapResp.Body.FindItemResponse.MessageText)
    }

    return &soapResp.Body.FindItemResponse, nil
}
