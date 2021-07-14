// Package entity provides a client implementation for the SAM Entity API
// Additional API details are published at https://open.gsa.gov/api/entity-api/
package entity

import (
	"net/url"
)

// version is the version of the SAM Entity API
const version = 2.5

// API is the URL for the SAM Entity API.  Query parameters including API key
// should be added to the RawQuery field
var API = url.URL{
	Scheme: "https",
	Host:   "api.sam.gov",
	Path:   "entity-information/v2/entities",
}

// Response is the SAM Entity API's JSON response to a request for Public API
// information
type Response struct {
	TotalRecords int      `json:"totalRecords"`
	EntityData   []Entity `json:"entityData"`
}

// Entity is a SAM Entity
type Entity struct {
	EntityRegistration Registration    `json:"entityRegistration"`
	CoreData           CoreData        `json:"coreData"`
	Assertions         Assertions      `json:"assertions"`
	RepsAndCerts       RepsAndCerts    `json:"repsAndCerts"`
	PointsOfContact    PointsOfContact `json:"pointsOfContact"`
}

// Registration holds entityRegistration section information for a SAM Entity
type Registration struct {
	SAMRegistered              string `json:"samRegistered"`
	UEISAM                     string `json:"ueiSAM"`
	UEIDUNS                    string `json:"ueiDUNS"`
	EntityEFTIndicator         string `json:"entityEFTIndicator"`
	CAGECode                   string `json:"cageCode"`
	DODAAC                     string `json:"dodaac"`
	LegalBusinessName          string `json:"legalBusinessName"`
	DBAName                    string `json:"dbaName"`
	PurposeOfRegistrationCode  string `json:"purposeOfRegistrationCode"`
	PurposeOfRegistrationDesc  string `json:"purposeOfRegistrationDesc"`
	RegistrationStatus         string `json:"registrationStatus"`
	RegistrationDate           string `json:"registrationDate"`
	LastUpdateData             string `json:"lastUpdateDate"`
	RegistrationExpirationDate string `json:"registrationExpirationDate"`
	ActivationData             string `json:"activationDate"`
	UEIStatus                  string `json:"ueiStatus"`
	UEIExpirationDate          string `json:"ueiExpirationDate"`
	UEICreationDate            string `json:"ueiCreationDate"`
	NoPublicDisplayFlag        string `json:"noPublicDisplayFlag"`
	ExclusionStatusFlag        string `json:"exclusionStatusFlag"`
	ExclusionURL               string `json:"exclusionURL"`
	DNBOpenData                string `json:"dnbOpenData"`
}

// CoreData holds coreData section information for a SAM Entity
type CoreData struct {
	EntityInformation struct {
		EntityURL              string `json:"entityURL"`
		EntityDivisionName     string `json:"entityDivisionName"`
		EntityDivisionNumber   string `json:"entityDivisionNumber"`
		EntityStartDate        string `json:"entityStartDate"`
		FiscalYearEndCloseDate string `json:"fiscalYearEndCloseDate"`
		SubmissionDate         string `json:"submissionDate"`
	} `json:"entityInformation"`
	PhysicalAddress       Address `json:"physicalAddress"`
	MailingAddress        Address `json:"mailingAddress"`
	CongressionalDistrict string  `json:"congressionalDistrict"`
	GeneralInformation    struct {
		EntityStructureCode        string `json:"entityStructureCode"`
		EntityStructureDesc        string `json:"entityStructureDesc"`
		EntityTypeCode             string `json:"entityTypeCode"`
		EntityTypeDesc             string `json:"entityTypeDesc"`
		ProfitStructureCode        string `json:"profitStructureCode"`
		ProfitStructureDesc        string `json:"profitStructureDesc"`
		OrganizationStructureCode  string `json:"organizationStructureCode"`
		OrganizationStructureDesc  string `json:"organizationStructureDesc"`
		StateOfIncorporationCode   string `json:"stateOfIncorporationCode"`
		StateOfIncorporationDesc   string `json:"stateOfIncorporationDesc"`
		CountryOfIncorporationCode string `json:"countryOfIncorporationCode"`
		CountryOfIncorporationDesc string `json:"countryOfIncorporationDesc"`
	} `json:"generalInformation"`
	BusinessTypes struct {
		BusinessTypeList []struct {
			BusinessTypeCode string `json:"businessTypeCode"`
			BusinessTypeDesc string `json:"businessTypeDesc"`
		} `json:"businessTypeList"`
		SBABusinessTypeList []struct {
			SBABusinessTypeCode    string `json:"sbaBusinessTypeCode"`
			SBABusinessTypeDesc    string `json:"sbaBusinessTypeDesc"`
			CertificationEntryDate string `json:"certificationEntryDate"`
			CertificationExitDate  string `json:"certificationExitDate"`
		} `json:"sbaBusinessTypeList"`
	} `json:"businessTypes"`
	FinancialInformation struct {
		CreditCardUsage     string `json:"creditCardUsage"`
		DebtSubjectToOffset string `json:"debtSubjectToOffset"`
	} `json:"financialInformation"`
}

// Assertions holds assertions section information for a SAM Entity
type Assertions struct {
	GoodsAndServices struct {
		PrimaryNAICS string `json:"primaryNaics"`
		NAICSList    []struct {
			NAICSCode        string `json:"naicsCode"`
			NAICSDescription string `json:"naicsDescription"`
			SBASmallBusiness string `json:"sbaSmallBusiness"`
			NAICSException   string `json:"naicsException"`
		} `json:"naicsList"`
		PSCList []struct {
			PSCCode        string `json:"pscCode"`
			PSCDescription string `json:"pscDescription"`
		} `json:"pscList"`
	} `json:"goodsAndServices"`
	DisasterReliefData struct {
		DisasterReliefFlag     string `json:"disasterReliefFlag"`
		BondingFlag            string `json:"bondingFlag"`
		GeographicalAreaServed []struct {
			GeographicalAreaServedStateCode           string `json:"geographicalAreaServedStateCode"`
			GeographicaLAreaServedStateName           string `json:"geographicalAreaServedStateName"`
			GeographicalAreaServedCountyCode          string `json:"geographicalAreaServedCountyCode"`
			GeographicalAreaServedCountyName          string `json:"geographicalAreaServedCountyName"`
			GeographicalAreaServedStatisticalAreaCode string `json:"geographicalAreaServedmetropolitanStatisticalAreaCode"`
			GeographicalAreaServedStatisticalAreaName string `json:"geographicalAreaServedmetropolitanStatisticalAreaName"`
		} `json:"geographicalAreaServed"`
	} `json:"disasterReliefData"`
	EDIInformation struct {
		EDIInformationFlag string `json:"ediInformationFlag"`
	} `json:"ediInformation"`
}

// RepsAndCerts holds repsAndCerts section information for a SAM Entity
type RepsAndCerts struct {
	Certifications struct {
		FARResponses  []FARResponse   `json:"fARResponses"`
		DFARResponses []DFARSResponse `json:"dFARResponses"` // [sic]
	} `json:"certifications"`
	Qualifications struct {
		// ArchitectEngineerResponses []Response `json:"architectEngineerResponses"`
		// ArchitectEngineerResponses Response `json:"architectEngineerResponses"`
	} `json:"qualifications"`
	FinancialAssistanceCertifications struct {
		GrantsCertificationStatus string `json:"grantsCertificationStatus"`
		GrantsCertifyingResponse  string `json:"grantsCertifyingResponse"`
		CertifierFirstName        string `json:"certifierFirstName"`
		CertifierLastName         string `json:"certifierLastName"`
		CertifierMiddleInitial    string `json:"certifierMiddleInitial"`
	} `json:"financialAssistanceCertifications"`
	PDFLinks struct {
		FARPDF                               string `json:"farPDF"`
		FARAndDFARSPDF                       string `json:"farAndDfarsPDF"`
		ArchitectEngineeringPDF              string `json:"architectEngineeringPDF"`
		FinancialAssistanceCertificationsPDF string `json:"financialAssistanceCertificationsPDF"`
	} `json:"pdfLinks"`
}

// PointsOfContact holds pointsOfContact section information for a SAM Entity
type PointsOfContact struct {
	GovernmentBusinessPOC          PointOfContact `json:"governmentBusinessPOC"`
	GovernmentBusinessAlternatePOC PointOfContact `json:"governmentBusinessAlternatePOC"`
	ElectronicBusinessPOC          PointOfContact `json:"electronicBusinessPOC"`
	ElectronicBusinessAlternatePOC PointOfContact `json:"electronicBusinessAlternatePOC"`
	PastPerformancePOC             PointOfContact `json:"pastPerformancePOC"`
	PastPerformanceAlternatePOC    PointOfContact `json:"pastPerformanceAlternatePOC"`
}

// Address contains geographic information
type Address struct {
	AddressLine1        string `json:"addressLine1,omitempty"`
	AddressLine2        string `json:"addressLine2,omitempty"`
	City                string `json:"city,omitempty"`
	StateOrProvinceCode string `json:"stateOrProvinceCode,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	ZipCode             string `json:"zipCode,omitempty"`
	ZipCodePlus4        string `json:"zipCodePlus4,omitempty"`
}

// PointOfContact is an Address with additional name and title fields
type PointOfContact struct {
	Address
	FirstName     string `json:"firstName"`
	MiddleInitial string `json:"middleInitial"`
	LastName      string `json:"lastName"`
	Title         string `json:"title"`
}

// FARResponse is an entity's response to a FAR provision
type FARResponse struct {
	ProvisionID   string   `json:"provisionId"`
	ListOfAnswers []Answer `json:"listOfAnswers"`
}

// DFARSResponse is an entity's response to a DFARS provision
type DFARSResponse = FARResponse

// Answer is an item-level answer within a FARResponse
type Answer struct {
	Section      string `json:"section"`
	QuestionText string `json:"questionText"`
	AnswerID     string `json:"answerId"`
	AnswerText   string `json:"answerText"`
	Country      string `json:"country"`
	Company      struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		TIN             string `json:"tin"`
		DUNS            string `json:"duns"`
		YearEstablished string `json:"yearEstablished"`
	} `json:"company"`
	HighestLevelOwnerCage struct {
		CAGECode          string `json:"cageCode"`
		NCAGECode         string `json:"nCageCode"`
		LegalBusinessName string `json:"legalBusinessName"`
		HasOwner          string `json:"hasOwner"`
		ID                string `json:"id"`
	} `json:"highestLevelOwnerCage"`
	ImmediateOwnerCage struct {
		CAGECode          string `json:"cageCode"`
		NCAGECode         string `json:"nCageCode"`
		LegalBusinessName string `json:"legalBusinessName"`
		HasOwner          string `json:"hasOwner"`
		ID                string `json:"id"`
	} `json:"immediateOwnerCage"`
	PersonDetails struct {
		FirstName     string `json:"firstName"`
		MiddleInitial string `json:"middleInitial"`
		LastName      string `json:"lastName"`
		Title         string `json:"title"`
	} `json:"personDetails"`
	PointOfContact struct {
		ID                  string `json:"id"`
		FirstName           string `json:"firstName"`
		MiddleInitial       string `json:"middleInitial"`
		LastName            string `json:"lastName"`
		Title               string `json:"title"`
		TelephoneNumber     string `json:"telephoneNumber"`
		Extension           string `json:"extension"`
		InternationalNumber string `json:"internationalNumber"`
	} `json:"pointOfContact"`
	ArchitectExperiencesList []struct {
		ID                          string `json:"id"`
		ExperienceCode              string `json:"experienceCode"`
		ExperienceDescription       string `json:"experienceDescription"`
		AnnualAvgRevenueCode        string `json:"annualAvgRevenueCode"`
		AnnualAvgRevenueDescription string `json:"annualAvgRevenueDescription"`
	} `json:"architectExperiencesList"`
	DisciplineInfoList []struct {
		ID                    string `json:"id"`
		DisciplineID          string `json:"disciplineID"`
		FirmNumOfEmployees    string `json:"firmNumOfEmployees"`
		BranchNumOfEmployees  string `json:"branchNumOfEmployees"`
		DisciplineDescription string `json:"disciplineDescription"`
	} `json:"disciplineInfoList"`
	// EndProductsList struct{} `json:"endProductsList"`
	// ForeignGovtEntitiesList struct{} `json:"foreignGovtEntitiesList"`
	// FormerFirmsList struct{} `json:"formerFirmsList"`
	// FSCInfoList struct{} `json:"fscInfoList"`
	// JointVentureCompaniesList struct{} `json:"jointVentureCompaniesList"`
	// LaborSurplusConcernsList struct{} `json:"laborSurplusConcernsList"`
	// NAICSList struct{} `json:"naicsList"`
	// PredecessorList struct{} `json:"predecessorsList"`
	// SAMFacilitiesList struct{} `json:"samFacilitiesList"`
	// SAMPointsOfContactList struct{} `json:"samPointsOfContactList"`
	// ServicesRevenueList struct{} `json:"servicesRevenuesList"`
	// SoftwareList struct{} `json:"softwareList"`
	// URLList struct{} `json:"urlList"`
}
