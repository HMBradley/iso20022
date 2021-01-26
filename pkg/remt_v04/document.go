// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentRemt00100104 struct {
	RmtAdvc RemittanceAdviceV04 `xml:"RmtAdvc"`
}

func (doc DocumentRemt00100104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentRemt00100104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RmtAdvc RemittanceAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtAdvc"`
	}
	output.RmtAdvc = doc.RmtAdvc
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:remt.001.001.04")
	return e.EncodeElement(&output, start)
}
