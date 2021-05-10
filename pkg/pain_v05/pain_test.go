// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestDocumentPain00900105(t *testing.T) {
	sample := DocumentPain00900105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentPain00900105{
		MndtInitnReq: MandateInitiationRequestV05{
			GrpHdr: GroupHeader47{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"MndtInitnReq":{"GrpHdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentPain00900105><MndtInitnReq><GrpHdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></MndtInitnReq></DocumentPain00900105>`,
		string(buf))
}

func TestDocumentPain01000105(t *testing.T) {
	sample := DocumentPain01000105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentPain01000105{
		MndtAmdmntReq: MandateAmendmentRequestV05{
			GrpHdr: GroupHeader47{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"MndtAmdmntReq":{"GrpHdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentPain01000105><MndtAmdmntReq><GrpHdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></MndtAmdmntReq></DocumentPain01000105>`,
		string(buf))
}

func TestDocumentPain01200105(t *testing.T) {
	sample := DocumentPain01200105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentPain01200105{
		MndtAccptncRpt: MandateAcceptanceReportV05{
			GrpHdr: GroupHeader47{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"MndtAccptncRpt":{"GrpHdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentPain01200105><MndtAccptncRpt><GrpHdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></MndtAccptncRpt></DocumentPain01200105>`,
		string(buf))
}

func TestDocumentPain01100105(t *testing.T) {
	sample := DocumentPain01100105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentPain01100105{
		MndtCxlReq: MandateCancellationRequestV05{
			GrpHdr: GroupHeader47{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"MndtCxlReq":{"GrpHdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentPain01100105><MndtCxlReq><GrpHdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></MndtCxlReq></DocumentPain01100105>`,
		string(buf))
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AuthenticationChannel1Choice{}.Validate())
	assert.NotNil(t, Authorisation1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, CashAccount24{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, DatePeriodDetails1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, Frequency37Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader47{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, Mandate10{}.Validate())
	assert.Nil(t, MandateAdjustment1{}.Validate())
	assert.Nil(t, MandateAuthentication1{}.Validate())
	assert.NotNil(t, MandateClassification1Choice{}.Validate())
	assert.NotNil(t, MandateInitiationRequestV05{}.Validate())
	assert.NotNil(t, MandateOccurrences4{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, MandateTypeInformation2{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, ReferredMandateDocument1{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, Mandate8{}.Validate())
	assert.NotNil(t, Mandate9{}.Validate())
	assert.NotNil(t, MandateAmendment5{}.Validate())
	assert.NotNil(t, MandateAmendmentReason1{}.Validate())
	assert.NotNil(t, MandateAmendmentRequestV05{}.Validate())
	assert.NotNil(t, MandateReason1Choice{}.Validate())
	assert.NotNil(t, OriginalMandate4Choice{}.Validate())
	assert.NotNil(t, OriginalMessageInformation1{}.Validate())
	assert.NotNil(t, MandateCancellation5{}.Validate())
	assert.NotNil(t, MandateCancellationRequestV05{}.Validate())
	assert.NotNil(t, PaymentCancellationReason1{}.Validate())
	assert.Nil(t, AcceptanceResult6{}.Validate())
	assert.Nil(t, Mandate11{}.Validate())
	assert.Nil(t, MandateAcceptance5{}.Validate())
	assert.NotNil(t, MandateAcceptanceReportV05{}.Validate())
	assert.NotNil(t, OriginalMandate5Choice{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalCashAccountType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalCategoryPurpose1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalAuthenticationChannel1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type10 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type12 ExternalLocalInstrument1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalMandateSetupReason1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalPersonIdentification1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type17 ExternalMandateReason1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 ExternalServiceLevel1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type24 Frequency10Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "NEVR"
	assert.Nil(t, type24.Validate())

	var type29 DocumentType6Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MSIN"
	assert.Nil(t, type29.Validate())

	var type30 Frequency6Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "YEAR"
	assert.Nil(t, type30.Validate())

	var type31 SequenceType2Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "RCUR"
	assert.Nil(t, type31.Validate())
}
