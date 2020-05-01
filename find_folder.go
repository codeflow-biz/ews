package ews

import (
    "encoding/xml"
    "errors"
)

type IndexedPageFolderView IndexedPageItemView

type FindFolderRequest struct {
    XMLName               xml.Name              `xml:"m:FindFolder"`
    Traversal             Traversal             `xml:"Traversal,attr,omitempty"`
    FolderShape           FolderShape           `xml:"m:FolderShape"`
    ParentFolderIds       ParentFolderId        `xml:"m:ParentFolderIds"`
    IndexedPageFolderView IndexedPageFolderView `xml:"m:IndexedPageFolderView"`
}
type FolderShape struct {
    BaseShape            BaseShape            `xml:"t:BaseShape"`
    AdditionalProperties AdditionalProperties `xml:"t:AdditionalProperties"`
}

type findFolderResponseEnvelop struct {
    XMLName xml.Name               `xml:"Envelope"`
    Body    findFolderResponseBody `xml:"Body"`
}
type findFolderResponseBody struct {
    FindFolderResponse FindFolderResponse `xml:"FindFolderResponse>ResponseMessages>FindFolderResponseMessage"`
}

type FindFolderResponse struct {
    Response
    RootFolder FindFolderRootfolder `xml:"RootFolder"`
}
type FindFolderRootfolder struct {
    ResultPagingInfo
    Folders        []Folder `xml:"Folders>Folder"`
    ContactsFolder []Folder `xml:"Folders>ContactsFolder"`
}

type Folder struct {
    FolderId         FolderId         `xml:"FolderId"`
    DisplayName      string           `xml:"DisplayName"`
    TotalCount       int              `xml:"TotalCount"`
    ChildFolderCount int              `xml:"ChildFolderCount"`
    ExtendedProperty ExtendedProperty `xml:"ExtendedProperty"`
}
type ExtendedProperty struct {
    ExtendedFieldURI ExtendedFieldURI `xml:"ExtendedFieldURI"`
    Value            string           `xml:"Value"`
}

func FindFolder(c Client, r *FindFolderRequest) (*FindFolderResponse, error) {

    xmlBytes, err := xml.MarshalIndent(r, "", "  ")
    if err != nil {
        return nil, err
    }

    bb, err := c.SendAndReceive(xmlBytes)
    if err != nil {
        return nil, err
    }

    var soapResp findFolderResponseEnvelop
    err = xml.Unmarshal(bb, &soapResp)
    if err != nil {
        return nil, err
    }

    if soapResp.Body.FindFolderResponse.ResponseClass == ResponseClassError {
        return nil, errors.New(soapResp.Body.FindFolderResponse.MessageText)
    }

    return &soapResp.Body.FindFolderResponse, nil
}
