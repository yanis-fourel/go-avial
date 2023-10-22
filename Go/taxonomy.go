package av

import (
	"encoding/json"
	"fmt"
	"github.com/yanis-fourel/go-avial/hgtp"
)

type Parameter hgtp.Parameter

const (
	ParameterNull      Parameter = iota // NULL
	ParameterAvesterra                  // AVESTERRA
	ParameterFilter                     // FILTER
	ParameterUnfilter                   // UNFILTER
	ParameterLog                        // LOG
	ParameterProcess                    // PROCESS
	ParameterSize                       // SIZE
	ParameterWipe                       // WIPE
	ParameterTerminate                  // TERMINATE
	ParameterTest                       // TEST

	maxParameter = ParameterTest
)

var MapStringToParameter = func() map[string]Parameter {
	m := make(map[string]Parameter)
	for i := Parameter(0); i <= maxParameter; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Parameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Parameter) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Parameter) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Parameter) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToParameter[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Parameter: %q", s)
}

// ----

type Action hgtp.Action

const (
	ActionNull      Action = iota // NULL
	ActionAvesterra               // AVESTERRA

	maxAction = ActionAvesterra
)

var MapStringToAction = func() map[string]Action {
	m := make(map[string]Action)
	for i := Action(0); i <= maxAction; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Action) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Action) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Action) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Action) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToAction[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Action: %q", s)
}

// ----

type Aspect hgtp.Aspect

const (
	AspectNull           Aspect = iota // NULL
	AspectAvesterra                    // AVESTERRA
	AspectLabel                        // LABEL
	AspectValue                        // VALUE
	AspectProperty                     // PROPERTY
	AspectAnnotation                   // ANNOTATION
	AspectData                         // DATA
	AspectAttribute                    // ATTRIBUTE
	AspectEntity                       // ENTITY
	AspectInstance                     // INSTANCE
	AspectIndex                        // INDEX
	AspectCount                        // COUNT
	AspectReference                    // REFERENCE
	AspectMethod                       // METHOD
	AspectName                         // NAME
	AspectKey                          // KEY
	AspectAspect                       // ASPECT
	AspectContext                      // CONTEXT
	AspectCategory                     // CATEGORY
	AspectClass                        // CLASS
	AspectMode                         // MODE
	AspectPrecedence                   // PRECEDENCE
	AspectTag                          // TAG
	AspectTime                         // TIME
	AspectError                        // ERROR
	AspectEvent                        // EVENT
	AspectParameter                    // PARAMETER
	AspectResult                       // RESULT
	AspectTimeout                      // TIMEOUT
	AspectAuthorization                // AUTHORIZATION
	AspectAuthority                    // AUTHORITY
	AspectAuxiliary                    // AUXILIARY
	AspectAncillary                    // ANCILLARY
	AspectOutlet                       // OUTLET
	AspectFunction                     // FUNCTION
	AspectVariable                     // VARIABLE
	AspectTable                        // TABLE
	AspectUnit                         // UNIT
	AspectPrefix                       // PREFIX
	AspectConfidence                   // CONFIDENCE
	AspectUncertainty                  // UNCERTAINTY
	AspectRegistry                     // REGISTRY
	AspectFolder                       // FOLDER
	AspectObject                       // OBJECT
	AspectTemplate                     // TEMPLATE
	AspectCharacteristic               // CHARACTERISTIC
	AspectProperty_set                 // PROPERTY_SET
	AspectValue_set                    // VALUE_SET
	AspectProperty_name                // PROPERTY_NAME
	AspectProperty_key                 // PROPERTY_KEY
	AspectProperty_value               // PROPERTY_VALUE
	AspectFact                         // FACT
	AspectFacet                        // FACET
	AspectFactor                       // FACTOR
	AspectFeature                      // FEATURE
	AspectField                        // FIELD
	AspectFrame                        // FRAME

	maxAspect = AspectFrame
)

var MapStringToAspect = func() map[string]Aspect {
	m := make(map[string]Aspect)
	for i := Aspect(0); i <= maxAspect; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Aspect) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Aspect) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Aspect) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Aspect) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToAspect[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Aspect: %q", s)
}

// ----

type Attribute hgtp.Attribute

const (
	AttributeNull                  Attribute = iota // NULL_ATTRIBUTE
	AttributeAvesterra                              // AVESTERRA_ATTRIBUTE
	AttributeCategory                               // CATEGORY_ATTRIBUTE
	AttributeClass                                  // CLASS_ATTRIBUTE
	AttributeName                                   // NAME_ATTRIBUTE
	AttributeValue                                  // VALUE_ATTRIBUTE
	AttributeAuthority                              // AUTHORITY_ATTRIBUTE
	AttributeTime                                   // TIME_ATTRIBUTE
	AttributeDate                                   // DATE_ATTRIBUTE
	AttributePlace                                  // PLACE_ATTRIBUTE
	AttributePosition                               // POSITION_ATTRIBUTE
	AttributeLocation                               // LOCATION_ATTRIBUTE
	AttributeElevation                              // ELEVATION_ATTRIBUTE
	AttributeLatitude                               // LATITUDE_ATTRIBUTE
	AttributeLongitude                              // LONGITUDE_ATTRIBUTE
	AttributeDirection                              // DIRECTION_ATTRIBUTE
	AttributeArea                                   // AREA_ATTRIBUTE
	AttributeVolume                                 // VOLUME_ATTRIBUTE
	AttributeWeight                                 // WEIGHT_ATTRIBUTE
	AttributeDistance                               // DISTANCE_ATTRIBUTE
	AttributeSize                                   // SIZE_ATTRIBUTE
	AttributeLength                                 // LENGTH_ATTRIBUTE
	AttributeWidth                                  // WIDTH_ATTRIBUTE
	AttributeHeight                                 // HEIGHT_ATTRIBUTE
	AttributeGeometry                               // GEOMETRY_ATTRIBUTE
	AttributeIntensity                              // INTENSITY_ATTRIBUTE
	AttributeAddress                                // ADDRESS_ATTRIBUTE
	AttributePostal                                 // POSTAL_ATTRIBUTE
	AttributeCity                                   // CITY_ATTRIBUTE
	AttributeDistrict                               // DISTRICT_ATTRIBUTE
	AttributeState                                  // STATE_ATTRIBUTE
	AttributeCountry                                // COUNTRY_ATTRIBUTE
	AttributeRegion                                 // REGION_ATTRIBUTE
	AttributeContinent                              // CONTINENT_ATTRIBUTE
	AttributeMember                                 // MEMBER_ATTRIBUTE
	AttributeAge                                    // AGE_ATTRIBUTE
	AttributeSex                                    // SEX_ATTRIBUTE
	AttributeRace                                   // RACE_ATTRIBUTE
	AttributeEthnicity                              // ETHNICITY_ATTRIBUTE
	AttributeReligion                               // RELIGION_ATTRIBUTE
	AttributeLanguage                               // LANGUAGE_ATTRIBUTE
	AttributeEducation                              // EDUCATION_ATTRIBUTE
	AttributeEmployment                             // EMPLOYMENT_ATTRIBUTE
	AttributeOccupation                             // OCCUPATION_ATTRIBUTE
	AttributeIncome                                 // INCOME_ATTRIBUTE
	AttributeWealth                                 // WEALTH_ATTRIBUTE
	AttributeSpouse                                 // SPOUSE_ATTRIBUTE
	AttributeChild                                  // CHILD_ATTRIBUTE
	AttributeSibling                                // SIBLING_ATTRIBUTE
	AttributeParent                                 // PARENT_ATTRIBUTE
	AttributeFamily                                 // FAMILY_ATTRIBUTE
	AttributeFather                                 // FATHER_ATTRIBUTE
	AttributeMother                                 // MOTHER_ATTRIBUTE
	AttributeHome                                   // HOME_ATTRIBUTE
	AttributeHealth                                 // HEALTH_ATTRIBUTE
	AttributeGenetic                                // GENETIC_ATTRIBUTE
	AttributePopulation                             // POPULATION_ATTRIBUTE
	AttributeSociocultural                          // SOCIOCULTURAL_ATTRIBUTE
	AttributeLinguistic                             // LINGUISTIC_ATTRIBUTE
	AttributeEthnographic                           // ETHNOGRAPHIC_ATTRIBUTE
	AttributeVariable                               // VARIABLE_ATTRIBUTE
	AttributeEquation                               // EQUATION_ATTRIBUTE
	AttributeCell                                   // CELL_ATTRIBUTE
	AttributeMobile                                 // MOBILE_ATTRIBUTE
	AttributeOffice                                 // OFFICE_ATTRIBUTE
	AttributeWork                                   // WORK_ATTRIBUTE
	AttributeVersion                                // VERSION_ATTRIBUTE
	AttributeManufacturer                           // MANUFACTURER_ATTRIBUTE
	AttributeModel                                  // MODEL_ATTRIBUTE
	AttributeOwner                                  // OWNER_ATTRIBUTE
	AttributeIdentity                               // IDENTITY_ATTRIBUTE
	AttributeId                                     // ID_ATTRIBUTE
	AttributeNetwork                                // NETWORK_ATTRIBUTE
	AttributeRouting                                // ROUTING_ATTRIBUTE
	AttributeGraph                                  // GRAPH_ATTRIBUTE
	AttributeNode                                   // NODE_ATTRIBUTE
	AttributeLink                                   // LINK_ATTRIBUTE
	AttributeVertex                                 // VERTEX_ATTRIBUTE
	AttributeEdge                                   // EDGE_ATTRIBUTE
	AttributeSource                                 // SOURCE_ATTRIBUTE
	AttributeDestination                            // DESTINATION_ATTRIBUTE
	AttributeFrom                                   // FROM_ATTRIBUTE
	AttributeTo                                     // TO_ATTRIBUTE
	AttributeTarget                                 // TARGET_ATTRIBUTE
	AttributeEgresses                               // EGRESSES_ATTRIBUTE
	AttributeThreat                                 // THREAT_ATTRIBUTE
	AttributeEndpoint                               // ENDPOINT_ATTRIBUTE
	AttributeConnectome                             // CONNECTOME_ATTRIBUTE
	AttributeReference                              // REFERENCE_ATTRIBUTE
	AttributeWebsite                                // WEBSITE_ATTRIBUTE
	AttributeFeature                                // FEATURE_ATTRIBUTE
	AttributeCondition                              // CONDITION_ATTRIBUTE
	AttributeWeather                                // WEATHER_ATTRIBUTE
	AttributeNeurological                           // NEUROLOGICAL_ATTRIBUTE
	AttributeSociological                           // SOCIOLOGICAL_ATTRIBUTE
	AttributePsychological                          // PSYCHOLOGICAL_ATTRIBUTE
	AttributeRelationship                           // RELATIONSHIP_ATTRIBUTE
	AttributeAssociation                            // ASSOCIATION_ATTRIBUTE
	AttributeInteraction                            // INTERACTION_ATTRIBUTE
	AttributeWorkspace                              // WORKSPACE_ATTRIBUTE
	AttributeView                                   // VIEW_ATTRIBUTE
	AttributeFilter                                 // FILTER_ATTRIBUTE
	AttributeCoordinate                             // COORDINATE_ATTRIBUTE
	AttributeAffiliation                            // AFFILIATION_ATTRIBUTE
	AttributeMembership                             // MEMBERSHIP_ATTRIBUTE
	AttributeResidence                              // RESIDENCE_ATTRIBUTE
	AttributeNear                                   // NEAR_ATTRIBUTE
	AttributeSentiment                              // SENTIMENT_ATTRIBUTE
	AttributeConfidence                             // CONFIDENCE_ATTRIBUTE
	AttributeDistinct                               // DISTINCT_ATTRIBUTE
	AttributeAmbiguity                              // AMBIGUITY_ATTRIBUTE
	AttributeError                                  // ERROR_ATTRIBUTE
	AttributeReliability                            // RELIABILITY_ATTRIBUTE
	AttributeOther                                  // OTHER_ATTRIBUTE
	AttributeZip                                    // ZIP_ATTRIBUTE
	AttributeUrl                                    // URL_ATTRIBUTE
	AttributePhone                                  // PHONE_ATTRIBUTE
	AttributeWeb                                    // WEB_ATTRIBUTE
	AttributeCode                                   // CODE_ATTRIBUTE
	AttributeTerritory                              // TERRITORY_ATTRIBUTE
	AttributeEmployee                               // EMPLOYEE_ATTRIBUTE
	AttributeCollege                                // COLLEGE_ATTRIBUTE
	AttributeDivision                               // DIVISION_ATTRIBUTE
	AttributeDepartment                             // DEPARTMENT_ATTRIBUTE
	AttributeSchool                                 // SCHOOL_ATTRIBUTE
	AttributeCampus                                 // CAMPUS_ATTRIBUTE
	AttributeAccount                                // ACCOUNT_ATTRIBUTE
	AttributeRegistry                               // REGISTRY_ATTRIBUTE
	AttributeManager                                // MANAGER_ATTRIBUTE
	AttributeStatus                                 // STATUS_ATTRIBUTE
	AttributeEmail                                  // EMAIL_ATTRIBUTE
	AttributeTitle                                  // TITLE_ATTRIBUTE
	AttributeOrganization                           // ORGANIZATION_ATTRIBUTE
	AttributeJob                                    // JOB_ATTRIBUTE
	AttributeAgreement                              // AGREEMENT_ATTRIBUTE
	AttributeInvention                              // INVENTION_ATTRIBUTE
	AttributeCollaboration                          // COLLABORATION_ATTRIBUTE
	AttributeFinance                                // FINANCE_ATTRIBUTE
	AttributeInformation                            // INFORMATION_ATTRIBUTE
	AttributeTimezone                               // TIMEZONE_ATTRIBUTE
	AttributeCapital                                // CAPITAL_ATTRIBUTE
	AttributeNeighbor                               // NEIGHBOR_ATTRIBUTE
	AttributePlanet                                 // PLANET_ATTRIBUTE
	AttributeCurrency                               // CURRENCY_ATTRIBUTE
	AttributeServer                                 // SERVER_ATTRIBUTE
	AttributeAnalytic                               // ANALYTIC_ATTRIBUTE
	AttributeAdapter                                // ADAPTER_ATTRIBUTE
	AttributeList                                   // LIST_ATTRIBUTE
	AttributeSet                                    // SET_ATTRIBUTE
	AttributeApplication                            // APPLICATION_ATTRIBUTE
	AttributeHuman                                  // HUMAN_ATTRIBUTE
	AttributeFunction                               // FUNCTION_ATTRIBUTE
	AttributeBrain                                  // BRAIN_ATTRIBUTE
	AttributeNeuron                                 // NEURON_ATTRIBUTE
	AttributeSynapse                                // SYNAPSE_ATTRIBUTE
	AttributeLaptop                                 // LAPTOP_ATTRIBUTE
	AttributeDesktop                                // DESKTOP_ATTRIBUTE
	AttributeTelephone                              // TELEPHONE_ATTRIBUTE
	AttributeCellphone                              // CELLPHONE_ATTRIBUTE
	AttributeSmartphone                             // SMARTPHONE_ATTRIBUTE
	AttributeMainframe                              // MAINFRAME_ATTRIBUTE
	AttributeTablet                                 // TABLET_ATTRIBUTE
	AttributeMarriage                               // MARRIAGE_ATTRIBUTE
	AttributeTown                                   // TOWN_ATTRIBUTE
	AttributeVillage                                // VILLAGE_ATTRIBUTE
	AttributeCounty                                 // COUNTY_ATTRIBUTE
	AttributeProvince                               // PROVINCE_ATTRIBUTE
	AttributeRoad                                   // ROAD_ATTRIBUTE
	AttributeHighway                                // HIGHWAY_ATTRIBUTE
	AttributeRailroad                               // RAILROAD_ATTRIBUTE
	AttributeOrgan                                  // ORGAN_ATTRIBUTE
	AttributeGene                                   // GENE_ATTRIBUTE
	AttributeProtein                                // PROTEIN_ATTRIBUTE
	AttributeMolecule                               // MOLECULE_ATTRIBUTE
	AttributeDisease                                // DISEASE_ATTRIBUTE
	AttributeDrug                                   // DRUG_ATTRIBUTE
	AttributeAtlas                                  // ATLAS_ATTRIBUTE
	AttributeMap                                    // MAP_ATTRIBUTE
	AttributeComputer                               // COMPUTER_ATTRIBUTE
	AttributePassage                                // PASSAGE_ATTRIBUTE
	AttributeTract                                  // TRACT_ATTRIBUTE
	AttributeCompany                                // COMPANY_ATTRIBUTE
	AttributeFolder                                 // FOLDER_ATTRIBUTE
	AttributeGroup                                  // GROUP_ATTRIBUTE
	AttributePerson                                 // PERSON_ATTRIBUTE
	AttributeHousehold                              // HOUSEHOLD_ATTRIBUTE
	AttributeCommunity                              // COMMUNITY_ATTRIBUTE
	AttributeBusiness                               // BUSINESS_ATTRIBUTE
	AttributeProcess                                // PROCESS_ATTRIBUTE
	AttributeTask                                   // TASK_ATTRIBUTE
	AttributeProgram                                // PROGRAM_ATTRIBUTE
	AttributeMathematics                            // MATHEMATICS_ATTRIBUTE
	AttributeSubstance                              // SUBSTANCE_ATTRIBUTE
	AttributeChemical                               // CHEMICAL_ATTRIBUTE
	AttributeStructure                              // STRUCTURE_ATTRIBUTE
	AttributeInfrastructure                         // INFRASTRUCTURE_ATTRIBUTE
	AttributeFacility                               // FACILITY_ATTRIBUTE
	AttributeFarm                                   // FARM_ATTRIBUTE
	AttributePath                                   // PATH_ATTRIBUTE
	AttributeMachine                                // MACHINE_ATTRIBUTE
	AttributeDevice                                 // DEVICE_ATTRIBUTE
	AttributePlatform                               // PLATFORM_ATTRIBUTE
	AttributeVehicle                                // VEHICLE_ATTRIBUTE
	AttributeContract                               // CONTRACT_ATTRIBUTE
	AttributeHypothesis                             // HYPOTHESIS_ATTRIBUTE
	AttributeMilitary                               // MILITARY_ATTRIBUTE
	AttributeGrant                                  // GRANT_ATTRIBUTE
	AttributeProposal                               // PROPOSAL_ATTRIBUTE
	AttributeDocument                               // DOCUMENT_ATTRIBUTE
	AttributeSpecialty                              // SPECIALTY_ATTRIBUTE
	AttributeMedical                                // MEDICAL_ATTRIBUTE
	AttributeType                                   // TYPE_ATTRIBUTE
	AttributeActivity                               // ACTIVITY_ATTRIBUTE
	AttributeMoney                                  // MONEY_ATTRIBUTE
	AttributeCost_center                            // COST_CENTER_ATTRIBUTE
	AttributeProject                                // PROJECT_ATTRIBUTE
	AttributeStudy                                  // STUDY_ATTRIBUTE
	AttributeFund                                   // FUND_ATTRIBUTE
	AttributeGift                                   // GIFT_ATTRIBUTE
	AttributeAward                                  // AWARD_ATTRIBUTE
	AttributeNews                                   // NEWS_ATTRIBUTE
	AttributeInvestigator                           // INVESTIGATOR_ATTRIBUTE
	AttributeTeam                                   // TEAM_ATTRIBUTE
	AttributeClimate                                // CLIMATE_ATTRIBUTE
	AttributeSystem                                 // SYSTEM_ATTRIBUTE
	AttributeSubsystem                              // SUBSYSTEM_ATTRIBUTE
	AttributeMention                                // MENTION_ATTRIBUTE
	AttributeSubtype                                // SUBTYPE_ATTRIBUTE
	AttributeText                                   // TEXT_ATTRIBUTE
	AttributeMode                                   // MODE_ATTRIBUTE
	AttributeModality                               // MODALITY_ATTRIBUTE
	AttributeGenericity                             // GENERICITY_ATTRIBUTE
	AttributeTense                                  // TENSE_ATTRIBUTE
	AttributePolarity                               // POLARITY_ATTRIBUTE
	AttributeEntity                                 // ENTITY_ATTRIBUTE
	AttributeEvent                                  // EVENT_ATTRIBUTE
	AttributeRelation                               // RELATION_ATTRIBUTE
	AttributeGeopolitics                            // GEOPOLITICS_ATTRIBUTE
	AttributeWeapon                                 // WEAPON_ATTRIBUTE
	AttributeSponsor                                // SPONSOR_ATTRIBUTE
	AttributeCost                                   // COST_ATTRIBUTE
	AttributeSchedule                               // SCHEDULE_ATTRIBUTE
	AttributePurpose                                // PURPOSE_ATTRIBUTE
	AttributeDescription                            // DESCRIPTION_ATTRIBUTE
	AttributeSupplier                               // SUPPLIER_ATTRIBUTE
	AttributeContractor                             // CONTRACTOR_ATTRIBUTE
	AttributePayment                                // PAYMENT_ATTRIBUTE
	AttributeSubclass                               // SUBCLASS_ATTRIBUTE
	AttributeItem                                   // ITEM_ATTRIBUTE
	AttributeStock                                  // STOCK_ATTRIBUTE
	AttributeStore                                  // STORE_ATTRIBUTE
	AttributeSubsidiary                             // SUBSIDIARY_ATTRIBUTE
	AttributeLeader                                 // LEADER_ATTRIBUTE
	AttributeShareholder                            // SHAREHOLDER_ATTRIBUTE
	AttributeFounder                                // FOUNDER_ATTRIBUTE
	AttributePerpetrator                            // PERPETRATOR_ATTRIBUTE
	AttributeVictim                                 // VICTIM_ATTRIBUTE
	AttributeActor                                  // ACTOR_ATTRIBUTE
	AttributeRecipient                              // RECIPIENT_ATTRIBUTE
	AttributeBeneficiary                            // BENEFICIARY_ATTRIBUTE
	AttributeCrime                                  // CRIME_ATTRIBUTE
	AttributeSentence                               // SENTENCE_ATTRIBUTE
	AttributeProvenance                             // PROVENANCE_ATTRIBUTE
	AttributeAirport                                // AIRPORT_ATTRIBUTE
	AttributePort                                   // PORT_ATTRIBUTE
	AttributeStation                                // STATION_ATTRIBUTE
	AttributeOntology                               // ONTOLOGY_ATTRIBUTE
	AttributePower_plant                            // POWER_PLANT_ATTRIBUTE
	AttributeCapacity                               // CAPACITY_ATTRIBUTE
	AttributePower                                  // POWER_ATTRIBUTE
	AttributeFuel                                   // FUEL_ATTRIBUTE
	AttributeAirplane                               // AIRPLANE_ATTRIBUTE
	AttributeShip                                   // SHIP_ATTRIBUTE
	AttributeTrain                                  // TRAIN_ATTRIBUTE
	AttributeCar                                    // CAR_ATTRIBUTE
	AttributeInstrument                             // INSTRUMENT_ATTRIBUTE
	AttributeShipment                               // SHIPMENT_ATTRIBUTE
	AttributeObject                                 // OBJECT_ATTRIBUTE
	AttributeBuilding                               // BUILDING_ATTRIBUTE
	AttributeLaw                                    // LAW_ATTRIBUTE
	AttributeSuspect                                // SUSPECT_ATTRIBUTE
	AttributeResult                                 // RESULT_ATTRIBUTE
	AttributeOutcome                                // OUTCOME_ATTRIBUTE
	AttributeCandidate                              // CANDIDATE_ATTRIBUTE
	AttributeElection                               // ELECTION_ATTRIBUTE
	AttributeBallot                                 // BALLOT_ATTRIBUTE
	AttributeWarehouse                              // WAREHOUSE_ATTRIBUTE
	AttributePallet                                 // PALLET_ATTRIBUTE
	AttributeVendor                                 // VENDOR_ATTRIBUTE
	AttributeBuyer                                  // BUYER_ATTRIBUTE
	AttributeSeller                                 // SELLER_ATTRIBUTE
	AttributeTransaction                            // TRANSACTION_ATTRIBUTE
	AttributeArrangement                            // ARRANGEMENT_ATTRIBUTE
	AttributeDeal                                   // DEAL_ATTRIBUTE
	AttributeFax                                    // FAX_ATTRIBUTE
	AttributeNote                                   // NOTE_ATTRIBUTE
	AttributeComment                                // COMMENT_ATTRIBUTE
	AttributeCount                                  // COUNT_ATTRIBUTE
	AttributeOrigin                                 // ORIGIN_ATTRIBUTE
	AttributeHeadquarters                           // HEADQUARTERS_ATTRIBUTE
	AttributeMain                                   // MAIN_ATTRIBUTE
	AttributeSales                                  // SALES_ATTRIBUTE
	AttributeLabor                                  // LABOR_ATTRIBUTE
	AttributeScore                                  // SCORE_ATTRIBUTE
	AttributeHours                                  // HOURS_ATTRIBUTE
	AttributeUnit                                   // UNIT_ATTRIBUTE
	AttributePrice                                  // PRICE_ATTRIBUTE
	AttributeInventory                              // INVENTORY_ATTRIBUTE
	AttributeCommodity                              // COMMODITY_ATTRIBUTE
	AttributeDefendant                              // DEFENDANT_ATTRIBUTE
	AttributePlaintiff                              // PLAINTIFF_ATTRIBUTE
	AttributeReservation                            // RESERVATION_ATTRIBUTE
	AttributePlan                                   // PLAN_ATTRIBUTE
	AttributeSeason                                 // SEASON_ATTRIBUTE
	AttributePassword                               // PASSWORD_ATTRIBUTE
	AttributeCustomer                               // CUSTOMER_ATTRIBUTE
	AttributeGender                                 // GENDER_ATTRIBUTE
	AttributeNumber                                 // NUMBER_ATTRIBUTE
	AttributeRole                                   // ROLE_ATTRIBUTE
	AttributeMerchandise                            // MERCHANDISE_ATTRIBUTE
	AttributePolygon                                // POLYGON_ATTRIBUTE
	AttributeRisk                                   // RISK_ATTRIBUTE
	AttributeGoal                                   // GOAL_ATTRIBUTE
	AttributeExpense                                // EXPENSE_ATTRIBUTE
	AttributeTemplate                               // TEMPLATE_ATTRIBUTE
	AttributeIndex                                  // INDEX_ATTRIBUTE
	AttributeMetric                                 // METRIC_ATTRIBUTE
	AttributeResource                               // RESOURCE_ATTRIBUTE
	AttributeImpact                                 // IMPACT_ATTRIBUTE
	AttributeCitation                               // CITATION_ATTRIBUTE
	AttributeGain                                   // GAIN_ATTRIBUTE
	AttributeLoss                                   // LOSS_ATTRIBUTE
	AttributeComposite                              // COMPOSITE_ATTRIBUTE
	AttributeProductivity                           // PRODUCTIVITY_ATTRIBUTE
	AttributeArrival                                // ARRIVAL_ATTRIBUTE
	AttributeDeparture                              // DEPARTURE_ATTRIBUTE
	AttributeSpeed                                  // SPEED_ATTRIBUTE
	AttributeVelocity                               // VELOCITY_ATTRIBUTE
	AttributeAcceleration                           // ACCELERATION_ATTRIBUTE
	AttributeMeasurement                            // MEASUREMENT_ATTRIBUTE
	AttributeMass                                   // MASS_ATTRIBUTE
	AttributeForce                                  // FORCE_ATTRIBUTE
	AttributeMomentum                               // MOMENTUM_ATTRIBUTE
	AttributeColor                                  // COLOR_ATTRIBUTE
	AttributeTexture                                // TEXTURE_ATTRIBUTE
	AttributeShade                                  // SHADE_ATTRIBUTE
	AttributeShape                                  // SHAPE_ATTRIBUTE
	AttributeSound                                  // SOUND_ATTRIBUTE
	AttributeClient                                 // CLIENT_ATTRIBUTE
	AttributeSupervisor                             // SUPERVISOR_ATTRIBUTE
	AttributeLabel                                  // LABEL_ATTRIBUTE
	AttributeSubsubtype                             // SUBSUBTYPE_ATTRIBUTE
	AttributeVariety                                // VARIETY_ATTRIBUTE
	AttributeMake                                   // MAKE_ATTRIBUTE
	AttributeYear                                   // YEAR_ATTRIBUTE
	AttributeJustification                          // JUSTIFICATION_ATTRIBUTE
	AttributeControl                                // CONTROL_ATTRIBUTE
	AttributeSurround                               // SURROUND_ATTRIBUTE
	AttributeResponsibility                         // RESPONSIBILITY_ATTRIBUTE
	AttributeBlame                                  // BLAME_ATTRIBUTE
	AttributeOpinion                                // OPINION_ATTRIBUTE
	AttributeInfluence                              // INFLUENCE_ATTRIBUTE
	AttributeAssertion                              // ASSERTION_ATTRIBUTE
	AttributeAttribution                            // ATTRIBUTION_ATTRIBUTE
	AttributeCitizenship                            // CITIZENSHIP_ATTRIBUTE
	AttributeHeading                                // HEADING_ATTRIBUTE
	AttributeBirthday                               // BIRTHDAY_ATTRIBUTE
	AttributeFlight                                 // FLIGHT_ATTRIBUTE
	AttributeRoute                                  // ROUTE_ATTRIBUTE
	AttributeSection                                // SECTION_ATTRIBUTE
	AttributeSeat                                   // SEAT_ATTRIBUTE
	AttributeAisle                                  // AISLE_ATTRIBUTE
	AttributeRow                                    // ROW_ATTRIBUTE
	AttributeColumn                                 // COLUMN_ATTRIBUTE
	AttributeField                                  // FIELD_ATTRIBUTE
	AttributeSku                                    // SKU_ATTRIBUTE
	AttributeQuantity                               // QUANTITY_ATTRIBUTE
	AttributeQuality                                // QUALITY_ATTRIBUTE
	AttributeNet                                    // NET_ATTRIBUTE
	AttributeFile                                   // FILE_ATTRIBUTE
	AttributeMunicipality                           // MUNICIPALITY_ATTRIBUTE
	AttributeService                                // SERVICE_ATTRIBUTE
	AttributeOrder                                  // ORDER_ATTRIBUTE
	AttributePurchase                               // PURCHASE_ATTRIBUTE
	AttributeMaterial                               // MATERIAL_ATTRIBUTE
	AttributePackaging                              // PACKAGING_ATTRIBUTE
	AttributeWarranty                               // WARRANTY_ATTRIBUTE
	AttributeRevenue                                // REVENUE_ATTRIBUTE
	AttributeEarnings                               // EARNINGS_ATTRIBUTE
	AttributeDuration                               // DURATION_ATTRIBUTE
	AttributeSerial_number                          // SERIAL_NUMBER_ATTRIBUTE
	AttributeTest                                   // TEST_ATTRIBUTE
	AttributeTravel                                 // TRAVEL_ATTRIBUTE
	AttributeSymptom                                // SYMPTOM_ATTRIBUTE
	AttributeExposure                               // EXPOSURE_ATTRIBUTE
	AttributeDiagnosis                              // DIAGNOSIS_ATTRIBUTE
	AttributeContact                                // CONTACT_ATTRIBUTE
	AttributeTemperature                            // TEMPERATURE_ATTRIBUTE
	AttributeReport                                 // REPORT_ATTRIBUTE
	AttributePatient                                // PATIENT_ATTRIBUTE
	AttributeHospital                               // HOSPITAL_ATTRIBUTE
	AttributeClinic                                 // CLINIC_ATTRIBUTE
	AttributeProvider                               // PROVIDER_ATTRIBUTE
	AttributeFactory                                // FACTORY_ATTRIBUTE
	AttributeKey                                    // KEY_ATTRIBUTE
	AttributeZone                                   // ZONE_ATTRIBUTE
	AttributePerimeter                              // PERIMETER_ATTRIBUTE
	AttributeOccupancy                              // OCCUPANCY_ATTRIBUTE
	AttributeVisit                                  // VISIT_ATTRIBUTE
	AttributeVisitor                                // VISITOR_ATTRIBUTE
	AttributePrivacy                                // PRIVACY_ATTRIBUTE
	AttributeAlert                                  // ALERT_ATTRIBUTE
	AttributeAbstract                               // ABSTRACT_ATTRIBUTE
	AttributeRoom                                   // ROOM_ATTRIBUTE
	AttributeLevel                                  // LEVEL_ATTRIBUTE
	AttributeCourse                                 // COURSE_ATTRIBUTE
	AttributeStudent                                // STUDENT_ATTRIBUTE
	AttributeFaculty                                // FACULTY_ATTRIBUTE
	AttributeStaff                                  // STAFF_ATTRIBUTE
	AttributeAdvisor                                // ADVISOR_ATTRIBUTE
	AttributeInstructor                             // INSTRUCTOR_ATTRIBUTE
	AttributeAlarm                                  // ALARM_ATTRIBUTE
	AttributeHost                                   // HOST_ATTRIBUTE
	AttributeConcept                                // CONCEPT_ATTRIBUTE
	AttributeElement                                // ELEMENT_ATTRIBUTE
	AttributeInstance                               // INSTANCE_ATTRIBUTE
	AttributeProperty                               // PROPERTY_ATTRIBUTE
	AttributeAttribute                              // ATTRIBUTE_ATTRIBUTE
	AttributeKind                                   // KIND_ATTRIBUTE
	AttributeData                                   // DATA_ATTRIBUTE
	AttributeJournal                                // JOURNAL_ATTRIBUTE
	AttributeArticle                                // ARTICLE_ATTRIBUTE
	AttributePublication                            // PUBLICATION_ATTRIBUTE
	AttributeBook                                   // BOOK_ATTRIBUTE
	AttributePeriodical                             // PERIODICAL_ATTRIBUTE
	AttributeCalendar                               // CALENDAR_ATTRIBUTE
	AttributeLog                                    // LOG_ATTRIBUTE
	AttributeCollection                             // COLLECTION_ATTRIBUTE
	AttributePermission                             // PERMISSION_ATTRIBUTE
	AttributeExample                                // EXAMPLE_ATTRIBUTE
	AttributeModule                                 // MODULE_ATTRIBUTE
	AttributeDirectory                              // DIRECTORY_ATTRIBUTE
	AttributeCache                                  // CACHE_ATTRIBUTE
	AttributeBehavior                               // BEHAVIOR_ATTRIBUTE
	AttributeAgent                                  // AGENT_ATTRIBUTE
	AttributeSlot                                   // SLOT_ATTRIBUTE
	AttributeSensor                                 // SENSOR_ATTRIBUTE
	AttributeActuator                               // ACTUATOR_ATTRIBUTE
	AttributeAction                                 // ACTION_ATTRIBUTE
	AttributePercept                                // PERCEPT_ATTRIBUTE
	AttributeUtility                                // UTILITY_ATTRIBUTE
	AttributeLetter                                 // LETTER_ATTRIBUTE
	AttributeOffer                                  // OFFER_ATTRIBUTE
	AttributeMessage                                // MESSAGE_ATTRIBUTE
	AttributeTechnology                             // TECHNOLOGY_ATTRIBUTE
	AttributeBudget                                 // BUDGET_ATTRIBUTE
	AttributeFunding                                // FUNDING_ATTRIBUTE
	AttributeCollaborator                           // COLLABORATOR_ATTRIBUTE
	AttributeDeputy                                 // DEPUTY_ATTRIBUTE
	AttributeTreatment                              // TREATMENT_ATTRIBUTE
	AttributePrognosis                              // PROGNOSIS_ATTRIBUTE
	AttributeSecurity                               // SECURITY_ATTRIBUTE
	AttributeDomain                                 // DOMAIN_ATTRIBUTE
	AttributeSubmission                             // SUBMISSION_ATTRIBUTE
	AttributeRate                                   // RATE_ATTRIBUTE
	AttributeCommand                                // COMMAND_ATTRIBUTE
	AttributeMethod                                 // METHOD_ATTRIBUTE
	AttributeInterchange                            // INTERCHANGE_ATTRIBUTE
	AttributeSummary                                // SUMMARY_ATTRIBUTE
	AttributeReadiness                              // READINESS_ATTRIBUTE
	AttributeClassification                         // CLASSIFICATION_ATTRIBUTE
	AttributeSecretary                              // SECRETARY_ATTRIBUTE
	AttributePersonnel                              // PERSONNEL_ATTRIBUTE
	AttributeRepresentative                         // REPRESENTATIVE_ATTRIBUTE
	AttributeImplementation                         // IMPLEMENTATION_ATTRIBUTE
	AttributeCommercial                             // COMMERCIAL_ATTRIBUTE
	AttributeClosure                                // CLOSURE_ATTRIBUTE
	AttributeInvoice                                // INVOICE_ATTRIBUTE
	AttributeBalance                                // BALANCE_ATTRIBUTE
	AttributePeriod                                 // PERIOD_ATTRIBUTE
	AttributeVariance                               // VARIANCE_ATTRIBUTE
	AttributeAllocation                             // ALLOCATION_ATTRIBUTE
	AttributeContribution                           // CONTRIBUTION_ATTRIBUTE
	AttributeSolicitation                           // SOLICITATION_ATTRIBUTE
	AttributeHistory                                // HISTORY_ATTRIBUTE
	AttributeOutlet                                 // OUTLET_ATTRIBUTE
	AttributeCharacteristic                         // CHARACTERISTIC_ATTRIBUTE
	AttributeAspect                                 // ASPECT_ATTRIBUTE
	AttributeFacet                                  // FACET_ATTRIBUTE
	AttributePattern                                // PATTERN_ATTRIBUTE
	AttributeTrait                                  // TRAIT_ATTRIBUTE
	AttributeForm                                   // FORM_ATTRIBUTE
	AttributeImport                                 // IMPORT_ATTRIBUTE
	AttributeExport                                 // EXPORT_ATTRIBUTE
	AttributeCart                                   // CART_ATTRIBUTE
	AttributeContainer                              // CONTAINER_ATTRIBUTE
	AttributeMarket                                 // MARKET_ATTRIBUTE
	AttributeSegment                                // SEGMENT_ATTRIBUTE
	AttributeCrew                                   // CREW_ATTRIBUTE
	AttributePassenger                              // PASSENGER_ATTRIBUTE
	AttributeDatabase                               // DATABASE_ATTRIBUTE
	AttributeTable                                  // TABLE_ATTRIBUTE
	AttributeContext                                // CONTEXT_ATTRIBUTE
	AttributeFormat                                 // FORMAT_ATTRIBUTE
	AttributeCompartment                            // COMPARTMENT_ATTRIBUTE
	AttributeAuthorization                          // AUTHORIZATION_ATTRIBUTE
	AttributeToken                                  // TOKEN_ATTRIBUTE
	AttributeSubject                                // SUBJECT_ATTRIBUTE
	AttributePredicate                              // PREDICATE_ATTRIBUTE
	AttributeArgument                               // ARGUMENT_ATTRIBUTE
	AttributeNoun                                   // NOUN_ATTRIBUTE
	AttributeVerb                                   // VERB_ATTRIBUTE
	AttributeAdjective                              // ADJECTIVE_ATTRIBUTE
	AttributeAdverb                                 // ADVERB_ATTRIBUTE
	AttributeUser                                   // USER_ATTRIBUTE
	AttributeFood                                   // FOOD_ATTRIBUTE
	AttributeMeal                                   // MEAL_ATTRIBUTE
	AttributeAppetizer                              // APPETIZER_ATTRIBUTE
	AttributeEntree                                 // ENTREE_ATTRIBUTE
	AttributeSidedish                               // SIDEDISH_ATTRIBUTE
	AttributeDessert                                // DESSERT_ATTRIBUTE
	AttributeBeverage                               // BEVERAGE_ATTRIBUTE
	AttributeFruit                                  // FRUIT_ATTRIBUTE
	AttributeVegetable                              // VEGETABLE_ATTRIBUTE
	AttributeMeat                                   // MEAT_ATTRIBUTE
	AttributeSnack                                  // SNACK_ATTRIBUTE
	AttributeConfiguration                          // CONFIGURATION_ATTRIBUTE
	AttributeLayout                                 // LAYOUT_ATTRIBUTE
	AttributeCommissioner                           // COMMISSIONER_ATTRIBUTE
	AttributeAttorney                               // ATTORNEY_ATTRIBUTE
	AttributeCounselor                              // COUNSELOR_ATTRIBUTE
	AttributeDelegate                               // DELEGATE_ATTRIBUTE
	AttributeLegislator                             // LEGISLATOR_ATTRIBUTE
	AttributeCouncilor                              // COUNCILOR_ATTRIBUTE
	AttributeProxy                                  // PROXY_ATTRIBUTE
	AttributeSenator                                // SENATOR_ATTRIBUTE
	AttributeMessenger                              // MESSENGER_ATTRIBUTE
	AttributeSalesperson                            // SALESPERSON_ATTRIBUTE
	AttributeAdministrator                          // ADMINISTRATOR_ATTRIBUTE
	AttributeHatch                                  // HATCH_ATTRIBUTE
	AttributeCrane                                  // CRANE_ATTRIBUTE
	AttributeStevedore                              // STEVEDORE_ATTRIBUTE
	AttributeCargo                                  // CARGO_ATTRIBUTE
	AttributeBerth                                  // BERTH_ATTRIBUTE
	AttributeDock                                   // DOCK_ATTRIBUTE
	AttributePier                                   // PIER_ATTRIBUTE
	AttributeEquipment                              // EQUIPMENT_ATTRIBUTE
	AttributeTool                                   // TOOL_ATTRIBUTE
	AttributeGear                                   // GEAR_ATTRIBUTE
	AttributeGang                                   // GANG_ATTRIBUTE
	AttributeShift                                  // SHIFT_ATTRIBUTE
	AttributeLift                                   // LIFT_ATTRIBUTE
	AttributeShipper                                // SHIPPER_ATTRIBUTE
	AttributeSender                                 // SENDER_ATTRIBUTE
	AttributeReceiver                               // RECEIVER_ATTRIBUTE
	AttributeVerification                           // VERIFICATION_ATTRIBUTE
	AttributeValidation                             // VALIDATION_ATTRIBUTE
	AttributeAuthentication                         // AUTHENTICATION_ATTRIBUTE
	AttributeDeployment                             // DEPLOYMENT_ATTRIBUTE
	AttributeShare                                  // SHARE_ATTRIBUTE
	AttributeEmissions                              // EMISSIONS_ATTRIBUTE
	AttributeWaste                                  // WASTE_ATTRIBUTE
	AttributeWater                                  // WATER_ATTRIBUTE
	AttributeEnergy                                 // ENERGY_ATTRIBUTE
	AttributeHeat                                   // HEAT_ATTRIBUTE
	AttributeElectricity                            // ELECTRICITY_ATTRIBUTE
	AttributeRestaurant                             // RESTAURANT_ATTRIBUTE
	AttributeLobby                                  // LOBBY_ATTRIBUTE
	AttributePool                                   // POOL_ATTRIBUTE
	AttributeGarage                                 // GARAGE_ATTRIBUTE
	AttributeKitchen                                // KITCHEN_ATTRIBUTE
	AttributeLaundry                                // LAUNDRY_ATTRIBUTE
	AttributeHousekeeping                           // HOUSEKEEPING_ATTRIBUTE
	AttributeMaintenance                            // MAINTENANCE_ATTRIBUTE
	AttributeEngineering                            // ENGINEERING_ATTRIBUTE
	AttributeMarketing                              // MARKETING_ATTRIBUTE
	AttributeDining                                 // DINING_ATTRIBUTE
	AttributeDemographic                            // DEMOGRAPHIC_ATTRIBUTE
	AttributeMatch                                  // MATCH_ATTRIBUTE
	AttributeLaboratory                             // LABORATORY_ATTRIBUTE
	AttributeRestriction                            // RESTRICTION_ATTRIBUTE
	AttributePrivilege                              // PRIVILEGE_ATTRIBUTE
	AttributeEntitlement                            // ENTITLEMENT_ATTRIBUTE
	AttributeLimit                                  // LIMIT_ATTRIBUTE
	AttributeRange                                  // RANGE_ATTRIBUTE
	AttributeAlias                                  // ALIAS_ATTRIBUTE
	AttributeAlternative                            // ALTERNATIVE_ATTRIBUTE
	AttributeTrack                                  // TRACK_ATTRIBUTE
	AttributeTrend                                  // TREND_ATTRIBUTE
	AttributePolicy                                 // POLICY_ATTRIBUTE
	AttributeTransport                              // TRANSPORT_ATTRIBUTE
	AttributeCalculation                            // CALCULATION_ATTRIBUTE
	AttributeComputation                            // COMPUTATION_ATTRIBUTE
	AttributeAudit                                  // AUDIT_ATTRIBUTE
	AttributeBatch                                  // BATCH_ATTRIBUTE
	AttributeClock                                  // CLOCK_ATTRIBUTE
	AttributePublisher                              // PUBLISHER_ATTRIBUTE
	AttributeSubscriber                             // SUBSCRIBER_ATTRIBUTE
	AttributeInclusion                              // INCLUSION_ATTRIBUTE
	AttributeExclusion                              // EXCLUSION_ATTRIBUTE
	AttributeEnrichment                             // ENRICHMENT_ATTRIBUTE
	AttributeFulfillment                            // FULFILLMENT_ATTRIBUTE
	AttributeProduction                             // PRODUCTION_ATTRIBUTE
	AttributeProducer                               // PRODUCER_ATTRIBUTE
	AttributeConsumer                               // CONSUMER_ATTRIBUTE
	AttributeStorage                                // STORAGE_ATTRIBUTE
	AttributeConsumption                            // CONSUMPTION_ATTRIBUTE
	AttributeSpace                                  // SPACE_ATTRIBUTE
	AttributeScale                                  // SCALE_ATTRIBUTE
	AttributeCounter                                // COUNTER_ATTRIBUTE
	AttributePending                                // PENDING_ATTRIBUTE
	AttributeChurch                                 // CHURCH_ATTRIBUTE
	AttributeTemple                                 // TEMPLE_ATTRIBUTE
	AttributeMosque                                 // MOSQUE_ATTRIBUTE
	AttributeSynagogue                              // SYNAGOGUE_ATTRIBUTE
	AttributeCathedral                              // CATHEDRAL_ATTRIBUTE
	AttributeMonastery                              // MONASTERY_ATTRIBUTE
	AttributeMotel                                  // MOTEL_ATTRIBUTE
	AttributeBank                                   // BANK_ATTRIBUTE
	AttributeMall                                   // MALL_ATTRIBUTE
	AttributeGrocery                                // GROCERY_ATTRIBUTE
	AttributePolice                                 // POLICE_ATTRIBUTE
	AttributeFire                                   // FIRE_ATTRIBUTE
	AttributeLibrary                                // LIBRARY_ATTRIBUTE
	AttributeMuseum                                 // MUSEUM_ATTRIBUTE
	AttributeCourthouse                             // COURTHOUSE_ATTRIBUTE
	AttributePrison                                 // PRISON_ATTRIBUTE
	AttributePost_office                            // POST_OFFICE_ATTRIBUTE
	AttributeGymnasium                              // GYMNASIUM_ATTRIBUTE
	AttributeStadium                                // STADIUM_ATTRIBUTE
	AttributeArena                                  // ARENA_ATTRIBUTE
	AttributeAuditorium                             // AUDITORIUM_ATTRIBUTE
	AttributeBloodbank                              // BLOODBANK_ATTRIBUTE
	AttributePharmacy                               // PHARMACY_ATTRIBUTE
	AttributeMortuary                               // MORTUARY_ATTRIBUTE
	AttributeCemetery                               // CEMETERY_ATTRIBUTE
	AttributeCivic_center                           // CIVIC_CENTER_ATTRIBUTE
	AttributeCommunity_center                       // COMMUNITY_CENTER_ATTRIBUTE
	AttributeConference_center                      // CONFERENCE_CENTER_ATTRIBUTE
	AttributeFitness_center                         // FITNESS_CENTER_ATTRIBUTE
	AttributeDistribution_center                    // DISTRIBUTION_CENTER_ATTRIBUTE
	AttributeRecreation_center                      // RECREATION_CENTER_ATTRIBUTE
	AttributeStorage_facility                       // STORAGE_FACILITY_ATTRIBUTE
	AttributeManufacturing_plant                    // MANUFACTURING_PLANT_ATTRIBUTE
	AttributeNuclear_plant                          // NUCLEAR_PLANT_ATTRIBUTE
	AttributeWater_plant                            // WATER_PLANT_ATTRIBUTE
	AttributeSewage_plant                           // SEWAGE_PLANT_ATTRIBUTE
	AttributeFilling_station                        // FILLING_STATION_ATTRIBUTE
	AttributeCharging_station                       // CHARGING_STATION_ATTRIBUTE
	AttributeUniversity                             // UNIVERSITY_ATTRIBUTE
	AttributePreschool                              // PRESCHOOL_ATTRIBUTE
	AttributeElementary_school                      // ELEMENTARY_SCHOOL_ATTRIBUTE
	AttributeMiddle_school                          // MIDDLE_SCHOOL_ATTRIBUTE
	AttributeHigh_school                            // HIGH_SCHOOL_ATTRIBUTE
	AttributeTechnical_school                       // TECHNICAL_SCHOOL_ATTRIBUTE
	AttributeHouse                                  // HOUSE_ATTRIBUTE
	AttributeTownhouse                              // TOWNHOUSE_ATTRIBUTE
	AttributeApartment                              // APARTMENT_ATTRIBUTE
	AttributeCondominimum                           // CONDOMINIMUM_ATTRIBUTE
	AttributeHostel                                 // HOSTEL_ATTRIBUTE
	AttributeResort                                 // RESORT_ATTRIBUTE
	AttributeTrain_station                          // TRAIN_STATION_ATTRIBUTE
	AttributeBus_station                            // BUS_STATION_ATTRIBUTE
	AttributeFerry_terminal                         // FERRY_TERMINAL_ATTRIBUTE
	AttributeTaxi_stand                             // TAXI_STAND_ATTRIBUTE
	AttributeBus_stand                              // BUS_STAND_ATTRIBUTE
	AttributeBicycle_stand                          // BICYCLE_STAND_ATTRIBUTE
	AttributeTower                                  // TOWER_ATTRIBUTE
	AttributeWifi                                   // WIFI_ATTRIBUTE
	AttributeTransmitter                            // TRANSMITTER_ATTRIBUTE
	AttributeDefibrillator                          // DEFIBRILLATOR_ATTRIBUTE
	AttributeTeller_machine                         // TELLER_MACHINE_ATTRIBUTE
	AttributeVending_machine                        // VENDING_MACHINE_ATTRIBUTE
	AttributeStreetlight                            // STREETLIGHT_ATTRIBUTE
	AttributeCamera                                 // CAMERA_ATTRIBUTE
	AttributeHydrant                                // HYDRANT_ATTRIBUTE
	AttributeDirector                               // DIRECTOR_ATTRIBUTE
	AttributeComposer                               // COMPOSER_ATTRIBUTE
	AttributePerformer                              // PERFORMER_ATTRIBUTE
	AttributeStage                                  // STAGE_ATTRIBUTE
	AttributeScreen                                 // SCREEN_ATTRIBUTE
	AttributeMenu                                   // MENU_ATTRIBUTE
	AttributeHotel                                  // HOTEL_ATTRIBUTE
	AttributeLounge                                 // LOUNGE_ATTRIBUTE
	AttributeBasement                               // BASEMENT_ATTRIBUTE
	AttributeReception                              // RECEPTION_ATTRIBUTE
	AttributeAxis                                   // AXIS_ATTRIBUTE
	AttributeSeries                                 // SERIES_ATTRIBUTE
	AttributeAnnotation                             // ANNOTATION_ATTRIBUTE
	AttributeFact                                   // FACT_ATTRIBUTE
	AttributeFactor                                 // FACTOR_ATTRIBUTE
	AttributeFrame                                  // FRAME_ATTRIBUTE
	AttributeIncident                               // INCIDENT_ATTRIBUTE
	AttributeSituation                              // SITUATION_ATTRIBUTE
	AttributeStory                                  // STORY_ATTRIBUTE
	AttributePhrase                                 // PHRASE_ATTRIBUTE
	AttributeKnowledge                              // KNOWLEDGE_ATTRIBUTE
	AttributeSignal                                 // SIGNAL_ATTRIBUTE
	AttributeComponent                              // COMPONENT_ATTRIBUTE
	AttributeInjury                                 // INJURY_ATTRIBUTE
	AttributeKeyword                                // KEYWORD_ATTRIBUTE
	AttributeCluster                                // CLUSTER_ATTRIBUTE
	AttributeData_type                              // DATA_TYPE_ATTRIBUTE
	AttributeInformation_type                       // INFORMATION_TYPE_ATTRIBUTE
	AttributeKnowledge_type                         // KNOWLEDGE_TYPE_ATTRIBUTE
	AttributeData_structure                         // DATA_STRUCTURE_ATTRIBUTE
	AttributeInformation_structure                  // INFORMATION_STRUCTURE_ATTRIBUTE
	AttributeKnowledge_structure                    // KNOWLEDGE_STRUCTURE_ATTRIBUTE
	AttributeContent                                // CONTENT_ATTRIBUTE
	AttributeExtract                                // EXTRACT_ATTRIBUTE
	AttributePiece                                  // PIECE_ATTRIBUTE
	AttributeImage                                  // IMAGE_ATTRIBUTE
	AttributeIdentifier                             // IDENTIFIER_ATTRIBUTE
	AttributeRank                                   // RANK_ATTRIBUTE

	maxAttribute = AttributeRank
)

var MapStringToAttribute = func() map[string]Attribute {
	m := make(map[string]Attribute)
	for i := Attribute(0); i <= maxAttribute; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Attribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Attribute) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Attribute) UnmarshalJSOnarsN(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Attribute) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToAttribute[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Attribute: %q", s)
}

// ----

type Category hgtp.Category

const (
	CategoryNull                Category = iota // NULL
	CategoryAvesterra                           // AVESTERRA
	CategoryServer                              // SERVER
	CategoryAnalytic                            // ANALYTIC
	CategoryAdapter                             // ADAPTER
	CategoryRegistry                            // REGISTRY
	CategoryWorkspace                           // WORKSPACE
	CategoryList                                // LIST
	CategorySet                                 // SET
	CategoryView                                // VIEW
	CategoryApplication                         // APPLICATION
	CategorySapien                              // SAPIEN
	CategoryGenetic                             // GENETIC
	CategorySociocultural                       // SOCIOCULTURAL
	CategoryLinguistic                          // LINGUISTIC
	CategoryEthnographic                        // ETHNOGRAPHIC
	CategoryFunction                            // FUNCTION
	CategoryVariable                            // VARIABLE
	CategoryBrain                               // BRAIN
	CategoryNeuron                              // NEURON
	CategorySynapse                             // SYNAPSE
	CategoryLaptop                              // LAPTOP
	CategoryDesktop                             // DESKTOP
	CategoryMobile                              // MOBILE
	CategoryTelephone                           // TELEPHONE
	CategoryCellphone                           // CELLPHONE
	CategorySmartphone                          // SMARTPHONE
	CategoryMainframe                           // MAINFRAME
	CategoryTablet                              // TABLET
	CategoryMarriage                            // MARRIAGE
	CategoryCity                                // CITY
	CategoryTown                                // TOWN
	CategoryVillage                             // VILLAGE
	CategoryState                               // STATE
	CategoryCounty                              // COUNTY
	CategoryProvince                            // PROVINCE
	CategoryCountry                             // COUNTRY
	CategoryRoad                                // ROAD
	CategoryHighway                             // HIGHWAY
	CategoryRailroad                            // RAILROAD
	CategoryOrgan                               // ORGAN
	CategoryCell                                // CELL
	CategoryGene                                // GENE
	CategoryProtein                             // PROTEIN
	CategoryMolecule                            // MOLECULE
	CategoryDisease                             // DISEASE
	CategoryDrug                                // DRUG
	CategoryAtlas                               // ATLAS
	CategoryMap                                 // MAP
	CategoryRegion                              // REGION
	CategoryComputer                            // COMPUTER
	CategoryGraph                               // GRAPH
	CategoryNode                                // NODE
	CategoryLink                                // LINK
	CategoryVertex                              // VERTEX
	CategoryEdge                                // EDGE
	CategoryDistrict                            // DISTRICT
	CategoryPassage                             // PASSAGE
	CategoryTract                               // TRACT
	CategoryIdentity                            // IDENTITY
	CategoryTerritory                           // TERRITORY
	CategoryEmployee                            // EMPLOYEE
	CategoryCollege                             // COLLEGE
	CategoryDivision                            // DIVISION
	CategorySchool                              // SCHOOL
	CategoryDepartment                          // DEPARTMENT
	CategoryCampus                              // CAMPUS
	CategoryManager                             // MANAGER
	CategoryCollaboration                       // COLLABORATION
	CategoryCompany                             // COMPANY
	CategoryTimezone                            // TIMEZONE
	CategoryContinent                           // CONTINENT
	CategoryLanguage                            // LANGUAGE
	CategoryPlanet                              // PLANET
	CategoryOrganization                        // ORGANIZATION
	CategoryFolder                              // FOLDER
	CategoryGroup                               // GROUP
	CategoryPlace                               // PLACE
	CategoryPerson                              // PERSON
	CategoryFamily                              // FAMILY
	CategoryHousehold                           // HOUSEHOLD
	CategoryCommunity                           // COMMUNITY
	CategoryBusiness                            // BUSINESS
	CategoryProcess                             // PROCESS
	CategoryTask                                // TASK
	CategoryProgram                             // PROGRAM
	CategoryMathematics                         // MATHEMATICS
	CategorySubstance                           // SUBSTANCE
	CategoryChemical                            // CHEMICAL
	CategoryStructure                           // STRUCTURE
	CategoryInfrastructure                      // INFRASTRUCTURE
	CategoryFacility                            // FACILITY
	CategoryFarm                                // FARM
	CategoryNetwork                             // NETWORK
	CategoryPath                                // PATH
	CategoryMachine                             // MACHINE
	CategoryDevice                              // DEVICE
	CategoryPlatform                            // PLATFORM
	CategoryVehicle                             // VEHICLE
	CategoryContract                            // CONTRACT
	CategoryHypothesis                          // HYPOTHESIS
	CategoryMilitary                            // MILITARY
	CategoryResidence                           // RESIDENCE
	CategoryAccount                             // ACCOUNT
	CategoryAgreement                           // AGREEMENT
	CategoryInvention                           // INVENTION
	CategoryGrant                               // GRANT
	CategoryProposal                            // PROPOSAL
	CategoryDocument                            // DOCUMENT
	CategoryActivity                            // ACTIVITY
	CategoryMoney                               // MONEY
	CategoryCost_center                         // COST_CENTER
	CategoryProject                             // PROJECT
	CategoryStudy                               // STUDY
	CategoryFund                                // FUND
	CategoryGift                                // GIFT
	CategoryAward                               // AWARD
	CategoryTeam                                // TEAM
	CategorySystem                              // SYSTEM
	CategorySubsystem                           // SUBSYSTEM
	CategoryMention                             // MENTION
	CategoryEntity                              // ENTITY
	CategoryRelation                            // RELATION
	CategoryValue                               // VALUE
	CategoryEvent                               // EVENT
	CategoryGeopolitics                         // GEOPOLITICS
	CategoryLocation                            // LOCATION
	CategoryWeapon                              // WEAPON
	CategorySponsor                             // SPONSOR
	CategoryPurpose                             // PURPOSE
	CategorySupplier                            // SUPPLIER
	CategoryContractor                          // CONTRACTOR
	CategoryPayment                             // PAYMENT
	CategorySubcontract                         // SUBCONTRACT
	CategoryStock                               // STOCK
	CategoryItem                                // ITEM
	CategoryStore                               // STORE
	CategoryFile                                // FILE
	CategoryOntology                            // ONTOLOGY
	CategoryRelationship                        // RELATIONSHIP
	CategoryAffiliation                         // AFFILIATION
	CategoryAirport                             // AIRPORT
	CategoryPort                                // PORT
	CategoryStation                             // STATION
	CategoryPower_plant                         // POWER_PLANT
	CategoryAirplane                            // AIRPLANE
	CategoryShip                                // SHIP
	CategoryTrain                               // TRAIN
	CategoryCar                                 // CAR
	CategoryInstrument                          // INSTRUMENT
	CategoryVictim                              // VICTIM
	CategoryPerpetrator                         // PERPETRATOR
	CategoryBuilding                            // BUILDING
	CategoryBorder                              // BORDER
	CategoryHelicopter                          // HELICOPTER
	CategoryPlaza                               // PLAZA
	CategoryShipment                            // SHIPMENT
	CategoryLaw                                 // LAW
	CategorySuspect                             // SUSPECT
	CategoryResult                              // RESULT
	CategoryOutcome                             // OUTCOME
	CategoryCandidate                           // CANDIDATE
	CategoryElection                            // ELECTION
	CategoryBallot                              // BALLOT
	CategoryWarehouse                           // WAREHOUSE
	CategoryPallet                              // PALLET
	CategoryVendor                              // VENDOR
	CategoryBuyer                               // BUYER
	CategorySeller                              // SELLER
	CategoryTransaction                         // TRANSACTION
	CategoryArrangement                         // ARRANGEMENT
	CategoryDeal                                // DEAL
	CategoryFilter                              // FILTER
	CategoryCategory                            // CATEGORY
	CategoryUnit                                // UNIT
	CategoryPrice                               // PRICE
	CategoryCommodity                           // COMMODITY
	CategoryCrime                               // CRIME
	CategoryReservation                         // RESERVATION
	CategoryPlan                                // PLAN
	CategorySeason                              // SEASON
	CategoryInventory                           // INVENTORY
	CategoryMerchandise                         // MERCHANDISE
	CategoryTemplate                            // TEMPLATE
	CategoryFeature                             // FEATURE
	CategoryIndex                               // INDEX
	CategoryTitle                               // TITLE
	CategorySide                                // SIDE
	CategoryTuple                               // TUPLE
	CategoryTimer                               // TIMER
	CategoryOutlet                              // OUTLET
	CategoryFlight                              // FLIGHT
	CategoryRoute                               // ROUTE
	CategorySchedule                            // SCHEDULE
	CategoryPurchase                            // PURCHASE
	CategoryTest                                // TEST
	CategoryPatient                             // PATIENT
	CategoryReport                              // REPORT
	CategoryHospital                            // HOSPITAL
	CategoryClinic                              // CLINIC
	CategoryFactory                             // FACTORY
	CategoryObject                              // OBJECT
	CategoryAbstract                            // ABSTRACT
	CategoryRoom                                // ROOM
	CategoryPerimeter                           // PERIMETER
	CategoryZone                                // ZONE
	CategoryArea                                // AREA
	CategoryVisit                               // VISIT
	CategoryVisitor                             // VISITOR
	CategoryCode                                // CODE
	CategoryCourse                              // COURSE
	CategoryStudent                             // STUDENT
	CategoryFaculty                             // FACULTY
	CategoryStaff                               // STAFF
	CategoryAdvisor                             // ADVISOR
	CategoryInstructor                          // INSTRUCTOR
	CategoryAlert                               // ALERT
	CategoryAlarm                               // ALARM
	CategoryHost                                // HOST
	CategoryConcept                             // CONCEPT
	CategoryElement                             // ELEMENT
	CategoryInstance                            // INSTANCE
	CategoryProperty                            // PROPERTY
	CategoryAttribute                           // ATTRIBUTE
	CategoryAnnotation                          // ANNOTATION
	CategoryData                                // DATA
	CategoryNote                                // NOTE
	CategoryJournal                             // JOURNAL
	CategoryArticle                             // ARTICLE
	CategoryPublication                         // PUBLICATION
	CategoryBook                                // BOOK
	CategoryPeriodical                          // PERIODICAL
	CategoryCalendar                            // CALENDAR
	CategoryLog                                 // LOG
	CategoryCollection                          // COLLECTION
	CategoryContext                             // CONTEXT
	CategoryPermission                          // PERMISSION
	CategoryExample                             // EXAMPLE
	CategoryModule                              // MODULE
	CategoryDirectory                           // DIRECTORY
	CategoryCache                               // CACHE
	CategoryBehavior                            // BEHAVIOR
	CategoryAgent                               // AGENT
	CategorySensor                              // SENSOR
	CategoryActuator                            // ACTUATOR
	CategoryPercept                             // PERCEPT
	CategoryAction                              // ACTION
	CategoryGoal                                // GOAL
	CategoryUtility                             // UTILITY
	CategoryLetter                              // LETTER
	CategoryOffer                               // OFFER
	CategoryMessage                             // MESSAGE
	CategoryTechnology                          // TECHNOLOGY
	CategoryBudget                              // BUDGET
	CategoryFunding                             // FUNDING
	CategoryCollaborator                        // COLLABORATOR
	CategoryRepesentative                       // REPESENTATIVE
	CategoryTreatment                           // TREATMENT
	CategoryPrognosis                           // PROGNOSIS
	CategorySecurity                            // SECURITY
	CategoryDomain                              // DOMAIN
	CategorySubmission                          // SUBMISSION
	CategoryInvoice                             // INVOICE
	CategorySoliciation                         // SOLICIATION
	CategoryPattern                             // PATTERN
	CategoryCharacteristic                      // CHARACTERISTIC
	CategoryForm                                // FORM
	CategoryImport                              // IMPORT
	CategoryExport                              // EXPORT
	CategoryCart                                // CART
	CategoryContainer                           // CONTAINER
	CategoryMarket                              // MARKET
	CategorySegment                             // SEGMENT
	CategoryCrew                                // CREW
	CategoryPassenger                           // PASSENGER
	CategoryDatabase                            // DATABASE
	CategoryTable                               // TABLE
	CategoryFood                                // FOOD
	CategoryCompartment                         // COMPARTMENT
	CategoryUser                                // USER
	CategorySupervisor                          // SUPERVISOR
	CategoryConfiguration                       // CONFIGURATION
	CategoryLayout                              // LAYOUT
	CategoryBeverage                            // BEVERAGE
	CategoryTraveler                            // TRAVELER
	CategoryCustomer                            // CUSTOMER
	CategoryModel                               // MODEL
	CategoryDeployment                          // DEPLOYMENT
	CategoryHotel                               // HOTEL
	CategoryLounge                              // LOUNGE
	CategoryRestaurant                          // RESTAURANT
	CategoryLobby                               // LOBBY
	CategoryPool                                // POOL
	CategoryGarage                              // GARAGE
	CategoryKitchen                             // KITCHEN
	CategoryLaundry                             // LAUNDRY
	CategoryHousekeeping                        // HOUSEKEEPING
	CategoryMaintenance                         // MAINTENANCE
	CategoryEngineering                         // ENGINEERING
	CategoryMarketing                           // MARKETING
	CategoryDining                              // DINING
	CategoryReception                           // RECEPTION
	CategoryShift                               // SHIFT
	CategoryTrack                               // TRACK
	CategoryTrend                               // TREND
	CategoryPolicy                              // POLICY
	CategoryCalculation                         // CALCULATION
	CategoryComputation                         // COMPUTATION
	CategoryAudit                               // AUDIT
	CategoryBatch                               // BATCH
	CategoryClock                               // CLOCK
	CategoryPublisher                           // PUBLISHER
	CategorySubscriber                          // SUBSCRIBER
	CategoryProducer                            // PRODUCER
	CategoryConsumer                            // CONSUMER
	CategorySpace                               // SPACE
	CategoryCounter                             // COUNTER
	CategoryChurch                              // CHURCH
	CategoryTemple                              // TEMPLE
	CategoryMosque                              // MOSQUE
	CategorySynagogue                           // SYNAGOGUE
	CategoryCathedral                           // CATHEDRAL
	CategoryMonastery                           // MONASTERY
	CategoryMotel                               // MOTEL
	CategoryBank                                // BANK
	CategoryMall                                // MALL
	CategoryGrocery                             // GROCERY
	CategoryPolice                              // POLICE
	CategoryFire                                // FIRE
	CategoryLibrary                             // LIBRARY
	CategoryMuseum                              // MUSEUM
	CategoryCourthouse                          // COURTHOUSE
	CategoryPrison                              // PRISON
	CategoryPost_office                         // POST_OFFICE
	CategoryOffice                              // OFFICE
	CategoryGymnasium                           // GYMNASIUM
	CategoryStadium                             // STADIUM
	CategoryArena                               // ARENA
	CategoryAuditorium                          // AUDITORIUM
	CategoryBloodbank                           // BLOODBANK
	CategoryPharmacy                            // PHARMACY
	CategoryMortuary                            // MORTUARY
	CategoryCemetery                            // CEMETERY
	CategoryCivic_center                        // CIVIC_CENTER
	CategoryCommunity_center                    // COMMUNITY_CENTER
	CategoryConference_center                   // CONFERENCE_CENTER
	CategoryFitness_center                      // FITNESS_CENTER
	CategoryDistribution_cent                   // DISTRIBUTION_CENT
	CategoryRecreation_center                   // RECREATION_CENTER
	CategoryStorage_facility                    // STORAGE_FACILITY
	CategoryManufacturing_plant                 // MANUFACTURING_PLANT
	CategoryNuclear_plant                       // NUCLEAR_PLANT
	CategoryWater_plant                         // WATER_PLANT
	CategorySewage_plant                        // SEWAGE_PLANT
	CategoryFilling_station                     // FILLING_STATION
	CategoryCharging_station                    // CHARGING_STATION
	CategoryUniversity                          // UNIVERSITY
	CategoryPreschool                           // PRESCHOOL
	CategoryElementary_school                   // ELEMENTARY_SCHOOL
	CategoryMiddle_school                       // MIDDLE_SCHOOL
	CategoryHigh_school                         // HIGH_SCHOOL
	CategoryTechnical_school                    // TECHNICAL_SCHOOL
	CategoryHouse                               // HOUSE
	CategoryTownhouse                           // TOWNHOUSE
	CategoryApartment                           // APARTMENT
	CategoryCondominimum                        // CONDOMINIMUM
	CategoryHostel                              // HOSTEL
	CategoryResort                              // RESORT
	CategoryTrain_station                       // TRAIN_STATION
	CategoryBus_station                         // BUS_STATION
	CategoryFerry_terminal                      // FERRY_TERMINAL
	CategoryTaxi_stand                          // TAXI_STAND
	CategoryBus_stand                           // BUS_STAND
	CategoryBicycle_stand                       // BICYCLE_STAND
	CategoryTower                               // TOWER
	CategoryWifi                                // WIFI
	CategoryTransmitter                         // TRANSMITTER
	CategoryDefibrillator                       // DEFIBRILLATOR
	CategoryTeller_machine                      // TELLER_MACHINE
	CategoryVending_machine                     // VENDING_MACHINE
	CategoryStreetlight                         // STREETLIGHT
	CategoryCamera                              // CAMERA
	CategoryHydrant                             // HYDRANT
	CategoryDirector                            // DIRECTOR
	CategoryComposer                            // COMPOSER
	CategoryPerformer                           // PERFORMER
	CategoryStage                               // STAGE
	CategoryScreen                              // SCREEN
	CategoryMenu                                // MENU

	maxCategory = CategoryMenu
)

var MapStringToCategory = func() map[string]Category {
	m := make(map[string]Category)
	for i := Category(0); i <= maxCategory; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Category) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Category) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Category) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToCategory[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Category: %q", s)
}

// ----

type Class hgtp.Class

const (
	ClassNull                  Class = iota // NULL
	ClassAvesterra                          // AVESTERRA
	ClassServer                             // SERVER
	ClassAnalytic                           // ANALYTIC
	ClassAdapter                            // ADAPTER
	ClassRegistry                           // REGISTRY
	ClassWorkspace                          // WORKSPACE
	ClassList                               // LIST
	ClassSet                                // SET
	ClassView                               // VIEW
	ClassApplication                        // APPLICATION
	ClassSapien                             // SAPIEN
	ClassGenetic                            // GENETIC
	ClassSociocultural                      // SOCIOCULTURAL
	ClassLinguistic                         // LINGUISTIC
	ClassEthnographic                       // ETHNOGRAPHIC
	ClassFunction                           // FUNCTION
	ClassVariable                           // VARIABLE
	ClassBrain                              // BRAIN
	ClassNeuron                             // NEURON
	ClassSynapse                            // SYNAPSE
	ClassLaptop                             // LAPTOP
	ClassDesktop                            // DESKTOP
	ClassMobile                             // MOBILE
	ClassTelephone                          // TELEPHONE
	ClassCellphone                          // CELLPHONE
	ClassSmartphone                         // SMARTPHONE
	ClassMainframe                          // MAINFRAME
	ClassTablet                             // TABLET
	ClassMarriage                           // MARRIAGE
	ClassCity                               // CITY
	ClassTown                               // TOWN
	ClassVillage                            // VILLAGE
	ClassState                              // STATE
	ClassCounty                             // COUNTY
	ClassProvince                           // PROVINCE
	ClassCountry                            // COUNTRY
	ClassRoad                               // ROAD
	ClassHighway                            // HIGHWAY
	ClassRailroad                           // RAILROAD
	ClassOrgan                              // ORGAN
	ClassCell                               // CELL
	ClassGene                               // GENE
	ClassProtein                            // PROTEIN
	ClassMolecule                           // MOLECULE
	ClassDisease                            // DISEASE
	ClassDrug                               // DRUG
	ClassAtlas                              // ATLAS
	ClassMap                                // MAP
	ClassRegion                             // REGION
	ClassComputer                           // COMPUTER
	ClassGraph                              // GRAPH
	ClassNode                               // NODE
	ClassLink                               // LINK
	ClassVertex                             // VERTEX
	ClassEdge                               // EDGE
	ClassDistrict                           // DISTRICT
	ClassPassage                            // PASSAGE
	ClassTract                              // TRACT
	ClassIdentity                           // IDENTITY
	ClassTerritory                          // TERRITORY
	ClassEmployee                           // EMPLOYEE
	ClassCollege                            // COLLEGE
	ClassDivision                           // DIVISION
	ClassSchool                             // SCHOOL
	ClassDepartment                         // DEPARTMENT
	ClassCampus                             // CAMPUS
	ClassManager                            // MANAGER
	ClassCollaboration                      // COLLABORATION
	ClassCompany                            // COMPANY
	ClassTimezone                           // TIMEZONE
	ClassContinent                          // CONTINENT
	ClassLanguage                           // LANGUAGE
	ClassPlanet                             // PLANET
	ClassOrganization                       // ORGANIZATION
	ClassFolder                             // FOLDER
	ClassGroup                              // GROUP
	ClassPlace                              // PLACE
	ClassPerson                             // PERSON
	ClassFamily                             // FAMILY
	ClassHousehold                          // HOUSEHOLD
	ClassCommunity                          // COMMUNITY
	ClassBusiness                           // BUSINESS
	ClassProcess                            // PROCESS
	ClassTask                               // TASK
	ClassProgram                            // PROGRAM
	ClassMathematics                        // MATHEMATICS
	ClassSubstance                          // SUBSTANCE
	ClassChemical                           // CHEMICAL
	ClassStructure                          // STRUCTURE
	ClassInfrastructure                     // INFRASTRUCTURE
	ClassFacility                           // FACILITY
	ClassFarm                               // FARM
	ClassNetwork                            // NETWORK
	ClassPath                               // PATH
	ClassMachine                            // MACHINE
	ClassDevice                             // DEVICE
	ClassPlatform                           // PLATFORM
	ClassVehicle                            // VEHICLE
	ClassContract                           // CONTRACT
	ClassHypothesis                         // HYPOTHESIS
	ClassMilitary                           // MILITARY
	ClassResidence                          // RESIDENCE
	ClassAccount                            // ACCOUNT
	ClassAgreement                          // AGREEMENT
	ClassInvention                          // INVENTION
	ClassGrant                              // GRANT
	ClassProposal                           // PROPOSAL
	ClassDocument                           // DOCUMENT
	ClassActivity                           // ACTIVITY
	ClassMoney                              // MONEY
	ClassCost_center                        // COST_CENTER
	ClassProject                            // PROJECT
	ClassStudy                              // STUDY
	ClassFund                               // FUND
	ClassGift                               // GIFT
	ClassAward                              // AWARD
	ClassTeam                               // TEAM
	ClassSystem                             // SYSTEM
	ClassSubsystem                          // SUBSYSTEM
	ClassMention                            // MENTION
	ClassEntity                             // ENTITY
	ClassRelation                           // RELATION
	ClassValue                              // VALUE
	ClassEvent                              // EVENT
	ClassGeopolitics                        // GEOPOLITICS
	ClassLocation                           // LOCATION
	ClassWeapon                             // WEAPON
	ClassSponsor                            // SPONSOR
	ClassPurpose                            // PURPOSE
	ClassSupplier                           // SUPPLIER
	ClassContractor                         // CONTRACTOR
	ClassPayment                            // PAYMENT
	ClassSubcontract                        // SUBCONTRACT
	ClassStock                              // STOCK
	ClassItem                               // ITEM
	ClassStore                              // STORE
	ClassFile                               // FILE
	ClassOntology                           // ONTOLOGY
	ClassRelationship                       // RELATIONSHIP
	ClassAffiliation                        // AFFILIATION
	ClassAirport                            // AIRPORT
	ClassPort                               // PORT
	ClassStation                            // STATION
	ClassPower_plant                        // POWER_PLANT
	ClassAirplane                           // AIRPLANE
	ClassShip                               // SHIP
	ClassTrain                              // TRAIN
	ClassCar                                // CAR
	ClassInstrument                         // INSTRUMENT
	ClassVictim                             // VICTIM
	ClassPerpetrator                        // PERPETRATOR
	ClassBuilding                           // BUILDING
	ClassBorder                             // BORDER
	ClassHelicopter                         // HELICOPTER
	ClassPlaza                              // PLAZA
	ClassShipment                           // SHIPMENT
	ClassLaw                                // LAW
	ClassSuspect                            // SUSPECT
	ClassResult                             // RESULT
	ClassOutcome                            // OUTCOME
	ClassCandidate                          // CANDIDATE
	ClassElection                           // ELECTION
	ClassBallot                             // BALLOT
	ClassWarehouse                          // WAREHOUSE
	ClassPallet                             // PALLET
	ClassVendor                             // VENDOR
	ClassBuyer                              // BUYER
	ClassSeller                             // SELLER
	ClassTransaction                        // TRANSACTION
	ClassArrangement                        // ARRANGEMENT
	ClassDeal                               // DEAL
	ClassFilter                             // FILTER
	ClassCategory                           // CATEGORY
	ClassUnit                               // UNIT
	ClassPrice                              // PRICE
	ClassCommodity                          // COMMODITY
	ClassCrime                              // CRIME
	ClassReservation                        // RESERVATION
	ClassPlan                               // PLAN
	ClassSeason                             // SEASON
	ClassInventory                          // INVENTORY
	ClassMerchandise                        // MERCHANDISE
	ClassTemplate                           // TEMPLATE
	ClassFeature                            // FEATURE
	ClassIndex                              // INDEX
	ClassTitle                              // TITLE
	ClassSide                               // SIDE
	ClassTuple                              // TUPLE
	ClassTimer                              // TIMER
	ClassOutlet                             // OUTLET
	ClassFlight                             // FLIGHT
	ClassRoute                              // ROUTE
	ClassSchedule                           // SCHEDULE
	ClassPurchase                           // PURCHASE
	ClassTest                               // TEST
	ClassPatient                            // PATIENT
	ClassReport                             // REPORT
	ClassHospital                           // HOSPITAL
	ClassClinic                             // CLINIC
	ClassFactory                            // FACTORY
	ClassObject                             // OBJECT
	ClassAbstract                           // ABSTRACT
	ClassRoom                               // ROOM
	ClassPerimeter                          // PERIMETER
	ClassZone                               // ZONE
	ClassArea                               // AREA
	ClassVisit                              // VISIT
	ClassVisitor                            // VISITOR
	ClassCode                               // CODE
	ClassCourse                             // COURSE
	ClassStudent                            // STUDENT
	ClassFaculty                            // FACULTY
	ClassStaff                              // STAFF
	ClassAdvisor                            // ADVISOR
	ClassInstructor                         // INSTRUCTOR
	ClassAlert                              // ALERT
	ClassAlarm                              // ALARM
	ClassHost                               // HOST
	ClassConcept                            // CONCEPT
	ClassElement                            // ELEMENT
	ClassInstance                           // INSTANCE
	ClassProperty                           // PROPERTY
	ClassAttribute                          // ATTRIBUTE
	ClassAnnotation                         // ANNOTATION
	ClassData                               // DATA
	ClassNote                               // NOTE
	ClassJournal                            // JOURNAL
	ClassArticle                            // ARTICLE
	ClassPublication                        // PUBLICATION
	ClassBook                               // BOOK
	ClassPeriodical                         // PERIODICAL
	ClassCalendar                           // CALENDAR
	ClassLog                                // LOG
	ClassCollection                         // COLLECTION
	ClassContext                            // CONTEXT
	ClassPermission                         // PERMISSION
	ClassExample                            // EXAMPLE
	ClassModule                             // MODULE
	ClassDirectory                          // DIRECTORY
	ClassCache                              // CACHE
	ClassBehavior                           // BEHAVIOR
	ClassAgent                              // AGENT
	ClassSensor                             // SENSOR
	ClassActuator                           // ACTUATOR
	ClassPercept                            // PERCEPT
	ClassAction                             // ACTION
	ClassGoal                               // GOAL
	ClassUtility                            // UTILITY
	ClassLetter                             // LETTER
	ClassOffer                              // OFFER
	ClassMessage                            // MESSAGE
	ClassTechnology                         // TECHNOLOGY
	ClassBudget                             // BUDGET
	ClassFunding                            // FUNDING
	ClassCollaborator                       // COLLABORATOR
	ClassRepesentative                      // REPESENTATIVE
	ClassTreatment                          // TREATMENT
	ClassPrognosis                          // PROGNOSIS
	ClassSecurity                           // SECURITY
	ClassDomain                             // DOMAIN
	ClassSubmission                         // SUBMISSION
	ClassInvoice                            // INVOICE
	ClassSoliciation                        // SOLICIATION
	ClassPattern                            // PATTERN
	ClassCharacteristic                     // CHARACTERISTIC
	ClassForm                               // FORM
	ClassImport                             // IMPORT
	ClassExport                             // EXPORT
	ClassCart                               // CART
	ClassContainer                          // CONTAINER
	ClassMarket                             // MARKET
	ClassSegment                            // SEGMENT
	ClassCrew                               // CREW
	ClassPassenger                          // PASSENGER
	ClassDatabase                           // DATABASE
	ClassTable                              // TABLE
	ClassFood                               // FOOD
	ClassCompartment                        // COMPARTMENT
	ClassUser                               // USER
	ClassSupervisor                         // SUPERVISOR
	ClassConfiguration                      // CONFIGURATION
	ClassLayout                             // LAYOUT
	ClassBeverage                           // BEVERAGE
	ClassTraveler                           // TRAVELER
	ClassCustomer                           // CUSTOMER
	ClassModel                              // MODEL
	ClassDeployment                         // DEPLOYMENT
	ClassHotel                              // HOTEL
	ClassLounge                             // LOUNGE
	ClassRestaurant                         // RESTAURANT
	ClassLobby                              // LOBBY
	ClassPool                               // POOL
	ClassGarage                             // GARAGE
	ClassKitchen                            // KITCHEN
	ClassLaundry                            // LAUNDRY
	ClassHousekeeping                       // HOUSEKEEPING
	ClassMaintenance                        // MAINTENANCE
	ClassEngineering                        // ENGINEERING
	ClassMarketing                          // MARKETING
	ClassDining                             // DINING
	ClassReception                          // RECEPTION
	ClassShift                              // SHIFT
	ClassTrack                              // TRACK
	ClassTrend                              // TREND
	ClassPolicy                             // POLICY
	ClassCalculation                        // CALCULATION
	ClassComputation                        // COMPUTATION
	ClassAudit                              // AUDIT
	ClassBatch                              // BATCH
	ClassClock                              // CLOCK
	ClassPublisher                          // PUBLISHER
	ClassSubscriber                         // SUBSCRIBER
	ClassProducer                           // PRODUCER
	ClassConsumer                           // CONSUMER
	ClassSpace                              // SPACE
	ClassCounter                            // COUNTER
	ClassChurch                             // CHURCH
	ClassTemple                             // TEMPLE
	ClassMosque                             // MOSQUE
	ClassSynagogue                          // SYNAGOGUE
	ClassCathedral                          // CATHEDRAL
	ClassMonastery                          // MONASTERY
	ClassMotel                              // MOTEL
	ClassBank                               // BANK
	ClassMall                               // MALL
	ClassGrocery                            // GROCERY
	ClassPolice                             // POLICE
	ClassFire                               // FIRE
	ClassLibrary                            // LIBRARY
	ClassMuseum                             // MUSEUM
	ClassCourthouse                         // COURTHOUSE
	ClassPrison                             // PRISON
	ClassPost_office                        // POST_OFFICE
	ClassOffice                             // OFFICE
	ClassGymnasium                          // GYMNASIUM
	ClassStadium                            // STADIUM
	ClassArena                              // ARENA
	ClassAuditorium                         // AUDITORIUM
	ClassBloodbank                          // BLOODBANK
	ClassPharmacy                           // PHARMACY
	ClassMortuary                           // MORTUARY
	ClassCemetery                           // CEMETERY
	ClassCivic_center                       // CIVIC_CENTER
	ClassCommunity_center                   // COMMUNITY_CENTER
	ClassConference_center                  // CONFERENCE_CENTER
	ClassFitness_center                     // FITNESS_CENTER
	ClassDistribution_center                // DISTRIBUTION_CENTER
	ClassRecreation_center                  // RECREATION_CENTER
	ClassStorage_facility                   // STORAGE_FACILITY
	ClassManufacturing_plant                // MANUFACTURING_PLANT
	ClassNuclear_plant                      // NUCLEAR_PLANT
	ClassWater_plant                        // WATER_PLANT
	ClassSewage_plant                       // SEWAGE_PLANT
	ClassFilling_station                    // FILLING_STATION
	ClassCharging_station                   // CHARGING_STATION
	ClassUniversity                         // UNIVERSITY
	ClassPreschool                          // PRESCHOOL
	ClassElementary_school                  // ELEMENTARY_SCHOOL
	ClassMiddle_school                      // MIDDLE_SCHOOL
	ClassHigh_school                        // HIGH_SCHOOL
	ClassTechnical_school                   // TECHNICAL_SCHOOL
	ClassHouse                              // HOUSE
	ClassTownhouse                          // TOWNHOUSE
	ClassApartment                          // APARTMENT
	ClassCondominimum                       // CONDOMINIMUM
	ClassHostel                             // HOSTEL
	ClassResort                             // RESORT
	ClassTrain_station                      // TRAIN_STATION
	ClassBus_station                        // BUS_STATION
	ClassFerry_terminal                     // FERRY_TERMINAL
	ClassTaxi_stand                         // TAXI_STAND
	ClassBus_stand                          // BUS_STAND
	ClassBicycle_stand                      // BICYCLE_STAND
	ClassTower                              // TOWER
	ClassWifi                               // WIFI
	ClassTransmitter                        // TRANSMITTER
	ClassDefibrillator                      // DEFIBRILLATOR
	ClassTeller_machine                     // TELLER_MACHINE
	ClassVending_machine                    // VENDING_MACHINE
	ClassStreetlight                        // STREETLIGHT
	ClassCamera                             // CAMERA
	ClassHydrant                            // HYDRANT
	ClassDirector                           // DIRECTOR
	ClassComposer                           // COMPOSER
	ClassPerformer                          // PERFORMER
	ClassStage                              // STAGE
	ClassScreen                             // SCREEN
	ClassMenu                               // MENU
	ClassSubject                            // SUBJECT
	ClassPredicate                          // PREDICATE
	ClassNoun                               // NOUN
	ClassVerb                               // VERB
	ClassAdjective                          // ADJECTIVE
	ClassAdverb                             // ADVERB
	ClassIncident                           // INCIDENT
	ClassSituation                          // SITUATION
	ClassStory                              // STORY
	ClassInfluence                          // INFLUENCE
	ClassPhrase                             // PHRASE
	ClassSignal                             // SIGNAL
	ClassInformation                        // INFORMATION
	ClassKnowledge                          // KNOWLEDGE
	ClassComponent                          // COMPONENT
	ClassInjury                             // INJURY
	ClassType                               // TYPE
	ClassSubtype                            // SUBTYPE
	ClassData_type                          // DATA_TYPE
	ClassInformation_type                   // INFORMATION_TYPE
	ClassKnowledge_type                     // KNOWLEDGE_TYPE
	ClassData_structure                     // DATA_STRUCTURE
	ClassInformation_structure              // INFORMATION_STRUCTURE
	ClassKnowledge_structure                // KNOWLEDGE_STRUCTURE
	ClassContent                            // CONTENT
	ClassExtract                            // EXTRACT
	ClassPiece                              // PIECE
	ClassImage                              // IMAGE
	ClassKeyword                            // KEYWORD
	ClassSummary                            // SUMMARY
	ClassSentiment                          // SENTIMENT
	ClassCluster                            // CLUSTER
	ClassSource                             // SOURCE
	ClassIdentifier                         // IDENTIFIER
	ClassDate                               // DATE
	ClassUrl                                // URL
	ClassRank                               // RANK
	ClassDistance                           // DISTANCE

	maxClass = ClassDistance
)

var MapStringToClass = func() map[string]Class {
	m := make(map[string]Class)
	for i := Class(0); i <= maxClass; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Class) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Class) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Class) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Class) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToClass[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Class: %q", s)
}

// ----

type Condition hgtp.Condition

const (
	ConditionNull      Condition = iota // NULL
	ConditionAvesterra                  // AVESTERRA
	ConditionReady                      // READY
	ConditionOnline                     // ONLINE
	ConditionArmed                      // ARMED
	ConditionEnabled                    // ENABLED
	ConditionTriggered                  // TRIGGERED
	ConditionLocked                     // LOCKED
	ConditionLoaded                     // LOADED
	ConditionLinked                     // LINKED
	ConditionWaiting                    // WAITING
	ConditionPending                    // PENDING
	ConditionHolding                    // HOLDING

	maxCondition = ConditionHolding
)

var MapStringToCondition = func() map[string]Condition {
	m := make(map[string]Condition)
	for i := Condition(0); i <= maxCondition; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Condition) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Condition) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Condition) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToCondition[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Condition: %q", s)
}

// ----

type Constant hgtp.Constant

const (
	ConstantNull                  Constant = iota // NULL
	ConstantAvesterra                             // AVESTERRA
	ConstantPi                                    // PI
	ConstantEulers_number                         // EULERS_NUMBER
	ConstantGolden_ratio                          // GOLDEN_RATIO
	ConstantImaginary_unit                        // IMAGINARY_UNIT
	ConstantSpeed_of_light                        // SPEED_OF_LIGHT
	ConstantPlanck                                // PLANCK
	ConstantGravity_of_earth                      // GRAVITY_OF_EARTH
	ConstantGravitational                         // GRAVITATIONAL
	ConstantAstronomical_unit                     // ASTRONOMICAL_UNIT
	ConstantMole                                  // MOLE
	ConstantElementa                              // ELEMENTA
	ConstantElectron_mass                         // ELECTRON_MASS
	ConstantProton_mass                           // PROTON_MASS
	ConstantNeutron_mass                          // NEUTRON_MASS
	ConstantElectric_permittivity                 // ELECTRIC_PERMITTIVITY
	ConstantMagnetic_permeability                 // MAGNETIC_PERMEABILITY

	maxConstant = ConstantMagnetic_permeability
)

var MapStringToConstant = func() map[string]Constant {
	m := make(map[string]Constant)
	for i := Constant(0); i <= maxConstant; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Constant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Constant) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Constant) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Constant) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToConstant[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Constant: %q", s)
}

// ----

type Context hgtp.Context

const (
	ContextNull      Context = iota // NULL
	ContextAvesterra                // AVESTERRA
	ContextAtra                     // ATRA
	ContextAvert                    // AVERT
	ContextAvian                    // AVIAN
	ContextOrchestra                // ORCHESTRA
	ContextSirius                   // SIRIUS

	maxContext = ContextSirius
)

var MapStringToContext = func() map[string]Context {
	m := make(map[string]Context)
	for i := Context(0); i <= maxContext; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Context) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Context) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Context) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Context) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToContext[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Context: %q", s)
}

// ----

type Event hgtp.Event

const (
	EventNull      Event = iota // NULL
	EventAvesterra              // AVESTERRA
	EventUpdate                 // UPDATE
	EventMessage                // MESSAGE
	EventError                  // ERROR
	EventRefresh                // REFRESH
	EventTimer                  // TIMER
	EventExecute                // EXECUTE
	EventCreate                 // CREATE
	EventDelete                 // DELETE
	EventAlert                  // ALERT
	EventImport                 // IMPORT
	EventExport                 // EXPORT
	EventReset                  // RESET

	maxEvent = EventReset
)

var MapStringToEvent = func() map[string]Event {
	m := make(map[string]Event)
	for i := Event(0); i <= maxEvent; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Event) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Event) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Event) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToEvent[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Event: %q", s)
}

// ----

type Format hgtp.Format

const (
	FormatNull       Format = iota // NULL
	FormatAvesterra                // AVESTERRA
	FormatBoolean                  // BOOLEAN
	FormatString                   // STRING
	FormatNatural                  // NATURAL
	FormatInteger                  // INTEGER
	FormatFloat                    // FLOAT
	FormatTime                     // TIME
	FormatDate                     // DATE
	FormatTimestamp                // TIMESTAMP
	FormatMm_dd_yyyy               // MM_DD_YYYY
	FormatDd_mm_yyyy               // DD_MM_YYYY
	FormatYyyy_mm_dd               // YYYY_MM_DD
	FormatSsn                      // SSN
	FormatSex                      // SEX
	FormatGender                   // GENDER

	maxFormat = Format(FormatGender)
)

var MapStringToFormat = func() map[string]Format {
	m := make(map[string]Format)
	for i := Format(0); i <= maxFormat; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Format) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Format) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Format) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Format) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToFormat[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Format: %q", s)
}

// ----

type Goal hgtp.Goal

const (
	GoalNull      Goal = iota // NULL
	GoalAvesterra             // AVESTERRA

	maxGoal = GoalAvesterra
)

var MapStringToGoal = func() map[string]Goal {
	m := make(map[string]Goal)
	for i := Goal(0); i <= maxGoal; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Goal) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Goal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Goal) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToGoal[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Goal: %q", s)
}

// ----

type Installation hgtp.Installation

const (
	InstallationNull          Installation = iota // NULL
	InstallationAvesterra                         // AVESTERRA
	InstallationAvu                               // AVU
	InstallationCertificate                       // CERTIFICATE
	InstallationDocumentation                     // DOCUMENTATION
	InstallationScript                            // SCRIPT
	InstallationLinux                             // LINUX
	InstallationMac                               // MAC
	InstallationWindows                           // WINDOWS
	InstallationAda                               // ADA
	InstallationC                                 // C
	InstallationJava                              // JAVA
	InstallationPython                            // PYTHON
	InstallationRust                              // RUST
	InstallationScala                             // SCALA
	InstallationSwift                             // SWIFT
	InstallationTupleo                            // TUPLEO
	InstallationTypescript                        // TYPESCRIPT

	maxInstallation = InstallationTypescript
)

var MapStringToInstallation = func() map[string]Installation {
	m := make(map[string]Installation)
	for i := Installation(0); i <= maxInstallation; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Installation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Installation) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Installation) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Installation) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToInstallation[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Installation: %q", s)
}

// ----

type Language hgtp.Language

const (
	LanguageNull       Language = iota // NULL
	LanguageAda                        // ADA
	LanguageC                          // C
	LanguageJava                       // JAVA
	LanguagePython                     // PYTHON
	LanguageRust                       // RUST
	LanguageScala                      // SCALA
	LanguageSwift                      // SWIFT
	LanguageTupleo                     // TUPLEO
	LanguageTypescript                 // TYPESCRIPT

	maxLanguage = LanguageTypescript
)

var MapStringToLanguage = func() map[string]Language {
	m := make(map[string]Language)
	for i := Language(0); i <= maxLanguage; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Language) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Language) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Language) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Language) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToLanguage[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Language: %q", s)
}

// ----

type Match hgtp.Match

const (
	MatchNull          Match = iota // NULL
	MatchIdentifier                 // IDENTIFIER
	MatchFirst                      // FIRST
	MatchLast                       // LAST
	MatchSex                        // SEX
	MatchRace                       // RACE
	MatchDob                        // DOB
	MatchSsn                        // SSN
	MatchSoundex_first              // SOUNDEX_FIRST
	MatchSoundex_last               // SOUNDEX_LAST
	MatchOther_first                // OTHER_FIRST
	MatchOther_last                 // OTHER_LAST
	MatchPartial_first              // PARTIAL_FIRST
	MatchPartial_last               // PARTIAL_LAST
	MatchPartial_name               // PARTIAL_NAME
	MatchPartial_ssn                // PARTIAL_SSN
	MatchPartial_dob                // PARTIAL_DOB

	maxMatch = MatchPartial_dob
)

var MapStringToMatch = func() map[string]Match {
	m := make(map[string]Match)
	for i := Match(0); i <= maxMatch; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Match) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Match) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Match) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Match) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToMatch[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Match: %q", s)
}

// ----

type Method hgtp.Method

const (
	MethodNull        Method = iota // NULL
	MethodAvesterra                 // AVESTERRA
	MethodCreate                    // CREATE
	MethodDelete                    // DELETE
	MethodGet                       // GET
	MethodSet                       // SET
	MethodClear                     // CLEAR
	MethodPurge                     // PURGE
	MethodInsert                    // INSERT
	MethodRemove                    // REMOVE
	MethodReplace                   // REPLACE
	MethodMember                    // MEMBER
	MethodStart                     // START
	MethodStop                      // STOP
	MethodInclude                   // INCLUDE
	MethodExclude                   // EXCLUDE
	MethodReset                     // RESET
	MethodCount                     // COUNT
	MethodLength                    // LENGTH
	MethodAttribute                 // ATTRIBUTE
	MethodName                      // NAME
	MethodKey                       // KEY
	MethodValue                     // VALUE
	MethodFind                      // FIND
	MethodSearch                    // SEARCH
	MethodLoad                      // LOAD
	MethodSave                      // SAVE
	MethodAuthorize                 // AUTHORIZE
	MethodDeauthorize               // DEAUTHORIZE
	MethodX_lookup                  // X_LOOKUP
	MethodRead                      // READ
	MethodWrite                     // WRITE
	MethodStore                     // STORE
	MethodRetrieve                  // RETRIEVE
	MethodEcho                      // ECHO
	MethodRefresh                   // REFRESH
	MethodSort                      // SORT
	MethodIndex                     // INDEX
	MethodTest                      // TEST
	MethodExecute                   // EXECUTE
	MethodDeclare                   // DECLARE

	maxMethod = MethodDeclare
)

var MapStringToMethod = func() map[string]Method {
	m := make(map[string]Method)
	for i := Method(0); i <= maxMethod; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Method) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Method) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Method) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Method) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToMethod[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Method: %q", s)
}

// ----

type Mode hgtp.Mode

const (
	ModeNull        Mode = iota // NULL
	ModeAvesterra               // AVESTERRA
	ModeUpdate                  // UPDATE
	ModeAppend                  // APPEND
	ModeReplace                 // REPLACE
	ModeInterchange             // INTERCHANGE
	ModePush                    // PUSH
	ModePull                    // PULL
	ModePersistent              // PERSISTENT
	ModeVolatile                // VOLATILE

	maxMode = ModeVolatile
)

var MapStringToMode = func() map[string]Mode {
	m := make(map[string]Mode)
	for i := Mode(0); i <= maxMode; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Mode) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Mode) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Mode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Mode) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToMode[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Mode: %q", s)
}

// ----

type Notice hgtp.Notice

const (
	NoticeNull      Notice = iota // NULL
	NoticeAvesterra               // AVESTERRA
	NoticeSuccess                 // SUCCESS
	NoticeInform                  // INFORM
	NoticeWarning                 // WARNING
	NoticeError                   // ERROR
	NoticeFatal                   // FATAL
	NoticeDebug                   // DEBUG
	NoticeTest                    // TEST

	maxNotice = NoticeTest
)

var MapStringToNotice = func() map[string]Notice {
	m := make(map[string]Notice)
	for i := Notice(0); i <= maxNotice; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Notice) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Notice) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Notice) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Notice) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToNotice[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Notice: %q", s)
}

// ----

type Operator hgtp.Operator

const (
	OperatorNull           Operator = iota // NULL
	OperatorAvesterra                      // AVESTERRA
	OperatorNil                            // NIL
	OperatorAtomic                         // ATOMIC
	OperatorHost                           // HOST
	OperatorNetwork                        // NETWORK
	OperatorUnique                         // UNIQUE
	OperatorCast                           // CAST
	OperatorValid                          // VALID
	OperatorExtinct                        // EXTINCT
	OperatorNew                            // NEW
	OperatorCreate                         // CREATE
	OperatorDelete                         // DELETE
	OperatorInsert                         // INSERT
	OperatorRemove                         // REMOVE
	OperatorReplace                        // REPLACE
	OperatorErase                          // ERASE
	OperatorLength                         // LENGTH
	OperatorElement                        // ELEMENT
	OperatorFirst                          // FIRST
	OperatorLast                           // LAST
	OperatorFind                           // FIND
	OperatorInclude                        // INCLUDE
	OperatorExclude                        // EXCLUDE
	OperatorCopy                           // COPY
	OperatorAnd                            // AND
	OperatorOr                             // OR
	OperatorXor                            // XOR
	OperatorNand                           // NAND
	OperatorNor                            // NOR
	OperatorNxor                           // NXOR
	OperatorNot                            // NOT
	OperatorSimilar                        // SIMILAR
	OperatorDissimilar                     // DISSIMILAR
	OperatorShift                          // SHIFT
	OperatorRotate                         // ROTATE
	OperatorBit                            // BIT
	OperatorAdd                            // ADD
	OperatorSubtract                       // SUBTRACT
	OperatorMultiply                       // MULTIPLY
	OperatorDivide                         // DIVIDE
	OperatorModulo                         // MODULO
	OperatorRemainder                      // REMAINDER
	OperatorAbsolute                       // ABSOLUTE
	OperatorExponent                       // EXPONENT
	OperatorIncrement                      // INCREMENT
	OperatorDecrement                      // DECREMENT
	OperatorEqual                          // EQUAL
	OperatorUnequal                        // UNEQUAL
	OperatorDecreasing                     // DECREASING
	OperatorIncreasing                     // INCREASING
	OperatorGreater                        // GREATER
	OperatorLess                           // LESS
	OperatorMinimum                        // MINIMUM
	OperatorMaximum                        // MAXIMUM
	OperatorRandom                         // RANDOM
	OperatorFloor                          // FLOOR
	OperatorCeiling                        // CEILING
	OperatorIf                             // IF
	OperatorCase                           // CASE
	OperatorUnless                         // UNLESS
	OperatorLoop                           // LOOP
	OperatorForward                        // FORWARD
	OperatorBackward                       // BACKWARD
	OperatorRepeat                         // REPEAT
	OperatorWhile                          // WHILE
	OperatorBreak                          // BREAK
	OperatorNext                           // NEXT
	OperatorUntil                          // UNTIL
	OperatorBlock                          // BLOCK
	OperatorThread                         // THREAD
	OperatorHalt                           // HALT
	OperatorBusy                           // BUSY
	OperatorGuard                          // GUARD
	OperatorError                          // ERROR
	OperatorMessage                        // MESSAGE
	OperatorTag                            // TAG
	OperatorBoolean                        // BOOLEAN
	OperatorCharacter                      // CHARACTER
	OperatorException                      // EXCEPTION
	OperatorInteger                        // INTEGER
	OperatorReal                           // REAL
	OperatorOperator                       // OPERATOR
	OperatorString                         // STRING
	OperatorText                           // TEXT
	OperatorClock                          // CLOCK
	OperatorTuple                          // TUPLE
	OperatorFunction                       // FUNCTION
	OperatorAlter                          // ALTER
	OperatorArgument                       // ARGUMENT
	OperatorArguments                      // ARGUMENTS
	OperatorDepth                          // DEPTH
	OperatorClosure                        // CLOSURE
	OperatorConjure                        // CONJURE
	OperatorSelf                           // SELF
	OperatorReturn                         // RETURN
	OperatorEvaluate                       // EVALUATE
	OperatorPure                           // PURE
	OperatorApply                          // APPLY
	OperatorCall                           // CALL
	OperatorDeclare                        // DECLARE
	OperatorLet                            // LET
	OperatorLocal                          // LOCAL
	OperatorLocals                         // LOCALS
	OperatorLevels                         // LEVELS
	OperatorChange                         // CHANGE
	OperatorGlobal                         // GLOBAL
	OperatorGlobals                        // GLOBALS
	OperatorScope                          // SCOPE
	OperatorTop                            // TOP
	OperatorSubscribe                      // SUBSCRIBE
	OperatorUnsubscribe                    // UNSUBSCRIBE
	OperatorWait                           // WAIT
	OperatorPublish                        // PUBLISH
	OperatorPending                        // PENDING
	OperatorTimer                          // TIMER
	OperatorStart                          // START
	OperatorStop                           // STOP
	OperatorActive                         // ACTIVE
	OperatorElapsed                        // ELAPSED
	OperatorRemaining                      // REMAINING
	OperatorReset                          // RESET
	OperatorLock                           // LOCK
	OperatorUnlock                         // UNLOCK
	OperatorLocked                         // LOCKED
	OperatorSqrt                           // SQRT
	OperatorExp                            // EXP
	OperatorLog                            // LOG
	OperatorSin                            // SIN
	OperatorCos                            // COS
	OperatorTan                            // TAN
	OperatorAsin                           // ASIN
	OperatorAcos                           // ACOS
	OperatorAtan                           // ATAN
	OperatorSinh                           // SINH
	OperatorCosh                           // COSH
	OperatorTanh                           // TANH
	OperatorThis                           // THIS
	OperatorDelay                          // DELAY
	OperatorInvoke                         // INVOKE
	OperatorInquire                        // INQUIRE
	OperatorAdapt                          // ADAPT
	OperatorConnect                        // CONNECT
	OperatorDisconnect                     // DISCONNECT
	OperatorAttach                         // ATTACH
	OperatorDetach                         // DETACH
	OperatorActivate                       // ACTIVATE
	OperatorDeactivate                     // DEACTIVATE
	OperatorFlush                          // FLUSH
	OperatorSync                           // SYNC
	OperatorBlocking                       // BLOCKING
	OperatorWaiting                        // WAITING
	OperatorActivated                      // ACTIVATED
	OperatorAccess                         // ACCESS
	OperatorRendezvous                     // RENDEZVOUS
	OperatorFetch                          // FETCH
	OperatorAdapting                       // ADAPTING
	OperatorInvoking                       // INVOKING
	OperatorConnection                     // CONNECTION
	OperatorConnections                    // CONNECTIONS
	OperatorConnected                      // CONNECTED
	OperatorAttachment                     // ATTACHMENT
	OperatorAttachments                    // ATTACHMENTS
	OperatorAttached                       // ATTACHED
	OperatorSubscription                   // SUBSCRIPTION
	OperatorSubscriptions                  // SUBSCRIPTIONS
	OperatorSubscribed                     // SUBSCRIBED
	OperatorReference                      // REFERENCE
	OperatorDereference                    // DEREFERENCE
	OperatorReferences                     // REFERENCES
	OperatorArm                            // ARM
	OperatorDisarm                         // DISARM
	OperatorArmed                          // ARMED
	OperatorAuthorize                      // AUTHORIZE
	OperatorDeauthorize                    // DEAUTHORIZE
	OperatorAuthority                      // AUTHORITY
	OperatorAuthorization                  // AUTHORIZATION
	OperatorAuthorizations                 // AUTHORIZATIONS
	OperatorAuthorized                     // AUTHORIZED
	OperatorSet                            // SET
	OperatorClear                          // CLEAR
	OperatorTest                           // TEST
	OperatorCondition                      // CONDITION
	OperatorConditions                     // CONDITIONS
	OperatorEntity                         // ENTITY
	OperatorOutlet                         // OUTLET
	OperatorAuxiliary                      // AUXILIARY
	OperatorAncillary                      // ANCILLARY
	OperatorParameter                      // PARAMETER
	OperatorResultant                      // RESULTANT
	OperatorInstance                       // INSTANCE
	OperatorOffset                         // OFFSET
	OperatorIndex                          // INDEX
	OperatorCount                          // COUNT
	OperatorName                           // NAME
	OperatorKey                            // KEY
	OperatorValue                          // VALUE
	OperatorContext                        // CONTEXT
	OperatorCategory                       // CATEGORY
	OperatorClass                          // CLASS
	OperatorMethod                         // METHOD
	OperatorAttribute                      // ATTRIBUTE
	OperatorEvent                          // EVENT
	OperatorMode                           // MODE
	OperatorState                          // STATE
	OperatorPrecedence                     // PRECEDENCE
	OperatorAspect                         // ASPECT
	OperatorTime                           // TIME
	OperatorTimeout                        // TIMEOUT
	OperatorProperty                       // PROPERTY
	OperatorAnnotation                     // ANNOTATION
	OperatorFact                           // FACT
	OperatorFacet                          // FACET
	OperatorFactor                         // FACTOR
	OperatorFeature                        // FEATURE
	OperatorField                          // FIELD
	OperatorFrame                          // FRAME
	OperatorMeasurement                    // MEASUREMENT
	OperatorUnit                           // UNIT
	OperatorPrefix                         // PREFIX
	OperatorConfidence                     // CONFIDENCE
	OperatorUncertainty                    // UNCERTAINTY
	OperatorServer                         // SERVER
	OperatorGateway                        // GATEWAY
	OperatorVersion                        // VERSION
	OperatorHostname                       // HOSTNAME
	OperatorInternet                       // INTERNET
	OperatorStatus                         // STATUS
	OperatorRedirect                       // REDIRECT
	OperatorRedirection                    // REDIRECTION
	OperatorDebug                          // DEBUG

	maxOperator = OperatorDebug
)

var MapStringToOperator = func() map[string]Operator {
	m := make(map[string]Operator)
	for i := Operator(0); i <= maxOperator; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Operator) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Operator) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Operator) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToOperator[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Operator: %q", s)
}

// ----

type Percept hgtp.Percept

const (
	PerceptNull      Percept = iota // NULL
	PerceptAvesterra                // AVESTERRA

	maxPercept = PerceptAvesterra
)

var MapStringToPercept = func() map[string]Percept {
	m := make(map[string]Percept)
	for i := Percept(0); i <= maxPercept; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Percept) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Percept) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Percept) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Percept) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToPercept[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Percept: %q", s)
}

// ----

type Permission hgtp.Permission

const (
	PermissionNull      Permission = iota // NULL
	PermissionAvesterra                   // AVESTERRA
	PermissionRead                        // READ
	PermissionWrite                       // WRITE
	PermissionExecute                     // EXECUTE
	PermissionDelete                      // DELETE

	maxPermission = PermissionDelete
)

var MapStringToPermission = func() map[string]Permission {
	m := make(map[string]Permission)
	for i := Permission(0); i <= maxPermission; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Permission) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Permission) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Permission) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToPermission[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Permission: %q", s)
}

type Platform hgtp.Platform

const (
	PlatformNull    Platform = iota // NULL
	PlatformLinux                   // LINUX
	PlatformMac                     // MAC
	PlatformWindows                 // WINDOWS

	maxPlatform = PlatformWindows
)

var MapStringToPlatform = func() map[string]Platform {
	m := make(map[string]Platform)
	for i := Platform(0); i <= maxPlatform; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Platform) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Platform) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Platform) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToPlatform[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Platform: %q", s)
}

type Precedence hgtp.Precedence

const (
	PrecedenceNull      Precedence = iota // NULL
	PrecedenceAvesterra                   // AVESTERRA
	PrecedenceTemplate                    // TEMPLATE
	PrecedenceTimer                       // TIMER
	PrecedenceTest                        // TEST
	PrecedenceTuple                       // TUPLE
	PrecedenceGraph                       // GRAPH
	PrecedenceSchedule                    // SCHEDULE

	maxPrecedence = PrecedenceSchedule
)

var MapStringToPrecedence = func() map[string]Precedence {
	m := make(map[string]Precedence)
	for i := Precedence(0); i <= maxPrecedence; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Precedence) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Precedence) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Precedence) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Precedence) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToPrecedence[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Precedence: %q", s)
}

type Prefix hgtp.Prefix

const (
	PrefixNull      Prefix = iota // NULL
	PrefixAvesterra               // AVESTERRA
	PrefixYotta                   // YOTTA
	PrefixZetta                   // ZETTA
	PrefixExa                     // EXA
	PrefixPeta                    // PETA
	PrefixTera                    // TERA
	PrefixGiga                    // GIGA
	PrefixMega                    // MEGA
	PrefixMyria                   // MYRIA
	PrefixKilo                    // KILO
	PrefixHecto                   // HECTO
	PrefixDeca                    // DECA
	PrefixDeci                    // DECI
	PrefixCenti                   // CENTI
	PrefixMilli                   // MILLI
	PrefixMicro                   // MICRO
	PrefixNano                    // NANO
	PrefixPico                    // PICO
	PrefixFempto                  // FEMPTO
	PrefixAtto                    // ATTO
	PrefixZopto                   // ZOPTO
	PrefixYocto                   // YOCTO
	PrefixExbi                    // EXBI
	PrefixPebi                    // PEBI
	PrefixTebi                    // TEBI
	PrefixGibi                    // GIBI
	PrefixMebi                    // MEBI
	PrefixKibi                    // KIBI

	maxPrefix = PrefixKibi
)

var MapStringToPrefix = func() map[string]Prefix {
	m := make(map[string]Prefix)
	for i := Prefix(0); i <= maxPrefix; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Prefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Prefix) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Prefix) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Prefix) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToPrefix[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Prefix: %q", s)
}

type Scheme hgtp.Scheme

const (
	SchemeNull      Scheme = iota // NULL
	SchemeAvesterra               // AVESTERRA
	SchemeAttribute               // ATTRIBUTE
	SchemeProperty                // PROPERTY
	SchemeInstance                // INSTANCE

	maxScheme = SchemeInstance
)

var MapStringToScheme = func() map[string]Scheme {
	m := make(map[string]Scheme)
	for i := Scheme(0); i <= maxScheme; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Scheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Scheme) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Scheme) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Scheme) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToScheme[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Scheme: %q", s)
}

type State hgtp.State

const (
	StateNull      State = iota // NULL
	StateAvesterra              // AVESTERRA
	StateAdvisory               // ADVISORY
	StateGood                   // GOOD
	StateCaution                // CAUTION
	StateAlert                  // ALERT
	StateWarning                // WARNING
	StateDanger                 // DANGER
	StateOffline                // OFFLINE

	maxState = StateOffline
)

var MapStringToState = func() map[string]State {
	m := make(map[string]State)
	for i := State(0); i <= maxState; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r State) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r State) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *State) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *State) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToState[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid State: %q", s)
}

type Tag hgtp.Tag

const (
	TagNull          Tag = iota // NULL
	TagAvesterra                // AVESTERRA
	TagBoolean                  // BOOLEAN
	TagCharacter                // CHARACTER
	TagString                   // STRING
	TagText                     // TEXT
	TagInteger                  // INTEGER
	TagFloat                    // FLOAT
	TagEntity                   // ENTITY
	TagTime                     // TIME
	TagWeb                      // WEB
	TagInterchange              // INTERCHANGE
	TagData                     // DATA
	TagException                // EXCEPTION
	TagOperator                 // OPERATOR
	TagFunction                 // FUNCTION
	TagMeasurement              // MEASUREMENT
	TagLocutor                  // LOCUTOR
	TagAuthorization            // AUTHORIZATION
	TagDuration                 // DURATION
	TagVariable                 // VARIABLE
	TagAggregate                // AGGREGATE
	TagArray                    // ARRAY

	maxTag = TagArray
)

var MapStringToTag = func() map[string]Tag {
	m := make(map[string]Tag)
	for i := Tag(0); i <= maxTag; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Tag) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Tag) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Tag) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToTag[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Tag: %q", s)
}

type Template hgtp.Template

const (
	TemplateNull          Template = iota // NULL
	TemplateAvesterra                     // AVESTERRA
	TemplateModel                         // MODEL
	TemplatePattern                       // PATTERN
	TemplateImport                        // IMPORT
	TemplateExport                        // EXPORT
	TemplateConfiguration                 // CONFIGURATION
	TemplateEnrichment                    // ENRICHMENT

	maxTemplate = TemplateEnrichment
)

var MapStringToTemplate = func() map[string]Template {
	m := make(map[string]Template)
	for i := Template(0); i <= maxTemplate; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Template) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Template) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Template) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Template) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToTemplate[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Template: %q", s)
}

type Unit hgtp.Unit

const (
	UnitNull                        Unit = iota // NULL
	UnitAvesterra                               // AVESTERRA
	UnitRadian                                  // RADIAN
	UnitDegree                                  // DEGREE
	UnitCircle                                  // CIRCLE
	UnitTurn                                    // TURN
	UnitRevolution                              // REVOLUTION
	UnitGrade                                   // GRADE
	UnitArcdegree                               // ARCDEGREE
	UnitArcminute                               // ARCMINUTE
	UnitCcs                                     // CCS
	UnitArsecond                                // ARSECOND
	UnitSteradian                               // STERADIAN
	UnitSphere                                  // SPHERE
	UnitSecond                                  // SECOND
	UnitMinute                                  // MINUTE
	UnitHour                                    // HOUR
	UnitDay                                     // DAY
	UnitWeek                                    // WEEK
	UnitYear                                    // YEAR
	UnitMonth                                   // MONTH
	UnitKilogram                                // KILOGRAM
	UnitGram                                    // GRAM
	UnitMetricton                               // METRICTON
	UnitPound                                   // POUND
	UnitPound_force                             // POUND_FORCE
	UnitOunce                                   // OUNCE
	UnitDram                                    // DRAM
	UnitGrain                                   // GRAIN
	UnitTon                                     // TON
	UnitLongton                                 // LONGTON
	UnitScruple                                 // SCRUPLE
	UnitApdram                                  // APDRAM
	UnitTroy_ounce                              // TROY_OUNCE
	UnitTroy_pound                              // TROY_POUND
	UnitPennyweight                             // PENNYWEIGHT
	UnitMeter                                   // METER
	UnitMicron                                  // MICRON
	UnitAngstrom                                // ANGSTROM
	UnitBohr_radious                            // BOHR_RADIOUS
	UnitInch                                    // INCH
	UnitFoot                                    // FOOT
	UnitYard                                    // YARD
	UnitRod                                     // ROD
	UnitMile                                    // MILE
	UnitStatue_mile                             // STATUE_MILE
	UnitNautical_mile                           // NAUTICAL_MILE
	UnitAcre                                    // ACRE
	UnitCubic_centimeter                        // CUBIC_CENTIMETER
	UnitLiter                                   // LITER
	UnitMilliliter                              // MILLILITER
	UnitGallon                                  // GALLON
	UnitImperial                                // IMPERIAL
	UnitQuart                                   // QUART
	UnitPint                                    // PINT
	UnitFluid_ounce                             // FLUID_OUNCE
	UnitFluid_dram                              // FLUID_DRAM
	UnitShot                                    // SHOT
	UnitDry                                     // DRY
	UnitPeck                                    // PECK
	UnitBushel                                  // BUSHEL
	UnitChaldron                                // CHALDRON
	UnitBritish_gallon                          // BRITISH_GALLON
	UnitBritish_quart                           // BRITISH_QUART
	UnitBritish_pint                            // BRITISH_PINT
	UnitBritish_fluid_ounce                     // BRITISH_FLUID_OUNCE
	UnitBritish_peck                            // BRITISH_PECK
	UnitBritish_hundred_weight                  // BRITISH_HUNDRED_WEIGHT
	UnitBottle                                  // BOTTLE
	UnitMiniature                               // MINIATURE
	UnitSplit                                   // SPLIT
	UnitHalf                                    // HALF
	UnitMagnum                                  // MAGNUM
	UnitJeroboam                                // JEROBOAM
	UnitRehoboam                                // REHOBOAM
	UnitMethuselah                              // METHUSELAH
	UnitSalmanazar                              // SALMANAZAR
	UnitBalthazar                               // BALTHAZAR
	UnitNebuchadnezzar                          // NEBUCHADNEZZAR
	UnitSovereign                               // SOVEREIGN
	UnitPony                                    // PONY
	UnitFillette                                // FILLETTE
	UnitTappit_hen                              // TAPPIT_HEN
	UnitRheoboam                                // RHEOBOAM
	UnitShalmaneser                             // SHALMANESER
	UnitBerkovets                               // BERKOVETS
	UnitPood                                    // POOD
	UnitFunt                                    // FUNT
	UnitLot                                     // LOT
	UnitZolotnik                                // ZOLOTNIK
	UnitDolya                                   // DOLYA
	UnitRumile                                  // RUMILE
	UnitMezhevayaverst                          // MEZHEVAYAVERST
	UnitVerst                                   // VERST
	UnitSazhen                                  // SAZHEN
	UnitKosayasazhen                            // KOSAYASAZHEN
	UnitArshin                                  // ARSHIN
	UnitRuell                                   // RUELL
	UnitLiniya                                  // LINIYA
	UnitVershok                                 // VERSHOK
	UnitPyad                                    // PYAD
	UnitVedro                                   // VEDRO
	UnitShtoff                                  // SHTOFF
	UnitVinebottle                              // VINEBOTTLE
	UnitVodkabottle                             // VODKABOTTLE
	UnitCharka                                  // CHARKA
	UnitShkalik                                 // SHKALIK
	UnitDesyatina_state                         // DESYATINA_STATE
	UnitDesyatina_farmery                       // DESYATINA_FARMERY
	UnitSqverst                                 // SQVERST
	UnitSqarshin_unit                           // SQARSHIN_UNIT
	UnitSqfoot                                  // SQFOOT
	UnitNewton                                  // NEWTON
	UnitJoule                                   // JOULE
	UnitCalorie                                 // CALORIE
	UnitNewton_meter                            // NEWTON_METER
	UnitNewton_centimeter                       // NEWTON_CENTIMETER
	UnitPound_foot                              // POUND_FOOT
	UnitPound_inch                              // POUND_INCH
	UnitKelvin                                  // KELVIN
	UnitCelcius                                 // CELCIUS
	UnitFarenheit                               // FARENHEIT
	UnitRankine                                 // RANKINE
	UnitReaumur                                 // REAUMUR
	UnitCoulomb                                 // COULOMB
	UnitAmpere                                  // AMPERE
	UnitWatt                                    // WATT
	UnitVolt                                    // VOLT
	UnitOhm                                     // OHM
	UnitMho                                     // MHO
	UnitFarad                                   // FARAD
	UnitHenry                                   // HENRY
	UnitWeber                                   // WEBER
	UnitCandela                                 // CANDELA
	UnitLumen                                   // LUMEN
	UnitLux                                     // LUX
	UnitEuro                                    // EURO
	UnitAustriaschilling                        // AUSTRIASCHILLING
	UnitBelgiumfranc                            // BELGIUMFRANC
	UnitFinlandmarkka                           // FINLANDMARKKA
	UnitFrancefranc                             // FRANCEFRANC
	UnitGermanymark                             // GERMANYMARK
	UnitGreecedrachma                           // GREECEDRACHMA
	UnitIrelandpunt                             // IRELANDPUNT
	UnitItalylira                               // ITALYLIRA
	UnitLuxembourgfranc                         // LUXEMBOURGFRANC
	UnitNetherlandsguilder                      // NETHERLANDSGUILDER
	UnitPortugalescudo                          // PORTUGALESCUDO
	UnitSpainpeseta                             // SPAINPESETA
	UnitSloveniantolar                          // SLOVENIANTOLAR
	UnitCypriotpound                            // CYPRIOTPOUND
	UnitMalteselira                             // MALTESELIRA
	UnitSlovakkoruna                            // SLOVAKKORUNA
	UnitBosniaherzegovinamark                   // BOSNIAHERZEGOVINAMARK
	UnitBulgarianlev                            // BULGARIANLEV
	UnitCapeverdeanescudo                       // CAPEVERDEANESCUDO
	UnitCentralafricancfafranc                  // CENTRALAFRICANCFAFRANC
	UnitComorosfranc                            // COMOROSFRANC
	UnitDanishkrone                             // DANISHKRONE
	UnitEstoniakroon                            // ESTONIAKROON
	UnitLithuanianlitas                         // LITHUANIANLITAS
	UnitPacificfrancexchange                    // PACIFICFRANCEXCHANGE
	UnitWestafricancfafranc                     // WESTAFRICANCFAFRANC
	UnitUs_dollar                               // US_DOLLAR
	UnitDollar                                  // DOLLAR
	UnitArubanflorin                            // ARUBANFLORIN
	UnitBahamiandollar                          // BAHAMIANDOLLAR
	UnitBahrainidinar                           // BAHRAINIDINAR
	UnitBarbadiandollar                         // BARBADIANDOLLAR
	UnitBelizedollar                            // BELIZEDOLLAR
	UnitBelarusianruble                         // BELARUSIANRUBLE
	UnitBermudiandollar                         // BERMUDIANDOLLAR
	UnitCaymanislandsdollar                     // CAYMANISLANDSDOLLAR
	UnitCubanconvertiblepeso                    // CUBANCONVERTIBLEPESO
	UnitDjiboutianfranc                         // DJIBOUTIANFRANC
	UnitEastcaribbeandollar                     // EASTCARIBBEANDOLLAR
	UnitEritreannakfa                           // ERITREANNAKFA
	UnitHongkongdollar                          // HONGKONGDOLLAR
	UnitMacanesepatacas                         // MACANESEPATACAS
	UnitJordaniandinar                          // JORDANIANDINAR
	UnitLebanesepound                           // LEBANESEPOUND
	UnitMaldivianrufiyaa                        // MALDIVIANRUFIYAA
	UnitNetherlandsantilleanguilder             // NETHERLANDSANTILLEANGUILDER
	UnitOmanirial                               // OMANIRIAL
	UnitQataririyal                             // QATARIRIYAL
	UnitSaudiriyal                              // SAUDIRIYAL
	UnitUnitedarabemiratesdirham                // UNITEDARABEMIRATESDIRHAM
	UnitBosniamark                              // BOSNIAMARK
	UnitCfafranc                                // CFAFRANC
	UnitCfpfranc                                // CFPFRANC
	UnitDrachma                                 // DRACHMA
	UnitEscudo                                  // ESCUDO
	UnitFranc                                   // FRANC
	UnitGuilder                                 // GUILDER
	UnitHerzegovinamark                         // HERZEGOVINAMARK
	UnitHollandguilder                          // HOLLANDGUILDER
	UnitLira                                    // LIRA
	UnitMark                                    // MARK
	UnitMarkka                                  // MARKKA
	UnitPeseta                                  // PESETA
	UnitRand                                    // RAND
	UnitErlang                                  // ERLANG
	UnitBit                                     // BIT
	UnitBaud                                    // BAUD
	UnitNibble                                  // NIBBLE
	UnitNybble                                  // NYBBLE
	UnitByte                                    // BYTE
	UnitWord                                    // WORD
	UnitBlock                                   // BLOCK
	UnitKbyte                                   // KBYTE
	UnitMegabyte                                // MEGABYTE
	UnitGigabyte                                // GIGABYTE
	UnitTerabyte                                // TERABYTE
	UnitPetabyte                                // PETABYTE
	UnitExabyte                                 // EXABYTE
	UnitZettabyte                               // ZETTABYTE
	UnitYottabyte                               // YOTTABYTE
	UnitKilobyte                                // KILOBYTE
	UnitPercent                                 // PERCENT
	UnitAbampere                                // ABAMPERE
	UnitAdmiraltyknot                           // ADMIRALTYKNOT
	UnitApostilb                                // APOSTILB
	UnitAre                                     // ARE
	UnitAsb                                     // ASB
	UnitArpentcan                               // ARPENTCAN
	UnitArpentlin                               // ARPENTLIN
	UnitAtmosphere                              // ATMOSPHERE
	UnitAtomic_mass_unit                        // ATOMIC_MASS_UNIT
	UnitSeconds_in_a_year                       // SECONDS_IN_A_YEAR
	UnitAmu                                     // AMU
	UnitBag                                     // BAG
	UnitBakersdozen                             // BAKERSDOZEN
	UnitBar                                     // BAR
	UnitBarie                                   // BARIE
	UnitBarleycorn                              // BARLEYCORN
	UnitBarn                                    // BARN
	UnitOilbarrel                               // OILBARREL
	UnitBarrel                                  // BARREL
	UnitBarye                                   // BARYE
	UnitBaryl                                   // BARYL
	UnitBev                                     // BEV
	UnitBiot                                    // BIOT
	UnitBlondel                                 // BLONDEL
	UnitBoardfoot                               // BOARDFOOT
	UnitBolt                                    // BOLT
	UnitBottommeasure                           // BOTTOMMEASURE
	UnitBritishthermalunit                      // BRITISHTHERMALUNIT
	UnitBtu                                     // BTU
	UnitRefrigeration                           // REFRIGERATION
	UnitBuck                                    // BUCK
	UnitCable                                   // CABLE
	UnitCaliber                                 // CALIBER
	UnitCarat                                   // CARAT
	UnitCaratgold                               // CARATGOLD
	UnitCent                                    // CENT
	UnitCental                                  // CENTAL
	UnitCentesimalminute                        // CENTESIMALMINUTE
	UnitCentesimalsecond                        // CENTESIMALSECOND
	UnitCentury                                 // CENTURY
	UnitCfs                                     // CFS
	UnitChain                                   // CHAIN
	UnitCircularinch                            // CIRCULARINCH
	UnitCircularmil                             // CIRCULARMIL
	UnitClusec                                  // CLUSEC
	UnitCoomb                                   // COOMB
	UnitCord                                    // CORD
	UnitCordfoot                                // CORDFOOT
	UnitCrith                                   // CRITH
	UnitCubit                                   // CUBIT
	UnitCup                                     // CUP
	UnitCurie                                   // CURIE
	UnitDalton                                  // DALTON
	UnitDecade                                  // DECADE
	UnitDioptre                                 // DIOPTRE
	UnitDisplacementton                         // DISPLACEMENTTON
	UnitDoppelzentner                           // DOPPELZENTNER
	UnitDozen                                   // DOZEN
	UnitDrop                                    // DROP
	UnitDyne                                    // DYNE
	UnitElectronvolt                            // ELECTRONVOLT
	UnitEll                                     // ELL
	UnitEngineerschain                          // ENGINEERSCHAIN
	UnitEngineerslink                           // ENGINEERSLINK
	UnitEquivalentfootcandle                    // EQUIVALENTFOOTCANDLE
	UnitEquivalentlux                           // EQUIVALENTLUX
	UnitEquivalentphot                          // EQUIVALENTPHOT
	UnitErg                                     // ERG
	UnitElectron_volt                           // ELECTRON_VOLT
	UnitFaraday                                 // FARADAY
	UnitFathom                                  // FATHOM
	UnitFermi                                   // FERMI
	UnitFifth                                   // FIFTH
	UnitFin                                     // FIN
	UnitFinger                                  // FINGER
	UnitFirkin                                  // FIRKIN
	UnitFootcandle                              // FOOTCANDLE
	UnitFootlambert                             // FOOTLAMBERT
	UnitFortnight                               // FORTNIGHT
	UnitFranklin                                // FRANKLIN
	UnitFrigorie                                // FRIGORIE
	UnitFurlong                                 // FURLONG
	UnitGalileo                                 // GALILEO
	UnitGamma                                   // GAMMA
	UnitGauss                                   // GAUSS
	UnitGeodeticfoot                            // GEODETICFOOT
	UnitGeographicalmile                        // GEOGRAPHICALMILE
	UnitGilbert                                 // GILBERT
	UnitGill                                    // GILL
	UnitGross                                   // GROSS
	UnitGunterschain                            // GUNTERSCHAIN
	UnitHand                                    // HAND
	UnitHartree                                 // HARTREE
	UnitHead                                    // HEAD
	UnitHectare                                 // HECTARE
	UnitHefnercandle                            // HEFNERCANDLE
	UnitHertz                                   // HERTZ
	UnitHogshead                                // HOGSHEAD
	UnitHomestead                               // HOMESTEAD
	UnitHorsepower                              // HORSEPOWER
	UnitHubble                                  // HUBBLE
	UnitHyl                                     // HYL
	UnitImaginarycubicfoot                      // IMAGINARYCUBICFOOT
	UnitJansky                                  // JANSKY
	UnitKarat                                   // KARAT
	UnitKayser                                  // KAYSER
	UnitKey                                     // KEY
	UnitKilderkin                               // KILDERKIN
	UnitKnot                                    // KNOT
	UnitLambert                                 // LAMBERT
	UnitLangley                                 // LANGLEY
	UnitLast                                    // LAST
	UnitLeague                                  // LEAGUE
	UnitLightyear                               // LIGHTYEAR
	UnitLine                                    // LINE
	UnitLink                                    // LINK
	UnitLonghundredweight                       // LONGHUNDREDWEIGHT
	UnitLongquarter                             // LONGQUARTER
	UnitLusec                                   // LUSEC
	UnitMass_earth                              // MASS_EARTH
	UnitMass_moon                               // MASS_MOON
	UnitMass_sun                                // MASS_SUN
	UnitMach                                    // MACH
	UnitMarineleague                            // MARINELEAGUE
	UnitMaxwell                                 // MAXWELL
	UnitMetriccarat                             // METRICCARAT
	UnitMgd                                     // MGD
	UnitMillihenry                              // MILLIHENRY
	UnitMil                                     // MIL
	UnitMillennium                              // MILLENNIUM
	UnitMinersinch                              // MINERSINCH
	UnitMinim                                   // MINIM
	UnitMiles_per_gallon                        // MILES_PER_GALLON
	UnitMiles_per_hour                          // MILES_PER_HOUR
	UnitNail                                    // NAIL
	UnitNmile                                   // NMILE
	UnitNit                                     // NIT
	UnitNoggin                                  // NOGGIN
	UnitNox                                     // NOX
	UnitOersted                                 // OERSTED
	UnitPace                                    // PACE
	UnitPair                                    // PAIR
	UnitPalm                                    // PALM
	UnitParasang                                // PARASANG
	UnitParsec                                  // PARSEC
	UnitPascal                                  // PASCAL
	UnitPerch                                   // PERCH
	UnitPhot                                    // PHOT
	UnitPica                                    // PICA
	UnitPieze                                   // PIEZE
	UnitPipe                                    // PIPE
	UnitPoint                                   // POINT
	UnitPoise                                   // POISE
	UnitPole                                    // POLE
	UnitPond                                    // POND
	UnitPoundal                                 // POUNDAL
	UnitPdl                                     // PDL
	UnitProof                                   // PROOF
	UnitPsi                                     // PSI
	UnitQuarter                                 // QUARTER
	UnitQuartersection                          // QUARTERSECTION
	UnitQuintal                                 // QUINTAL
	UnitQuire                                   // QUIRE
	UnitRadious_earth                           // RADIOUS_EARTH
	UnitAdious_moon                             // ADIOUS_MOON
	UnitAdious_sun                              // ADIOUS_SUN
	UnitRackunit                                // RACKUNIT
	UnitRad                                     // RAD
	UnitReam                                    // REAM
	UnitRegisterton                             // REGISTERTON
	UnitRem                                     // REM
	UnitRhe                                     // RHE
	UnitRontgen                                 // RONTGEN
	UnitRoentgen                                // ROENTGEN
	UnitRood                                    // ROOD
	UnitRope                                    // ROPE
	UnitRutherford                              // RUTHERFORD
	UnitRydbergconst                            // RYDBERGCONST
	UnitRydberg                                 // RYDBERG
	UnitSabin                                   // SABIN
	UnitSack                                    // SACK
	UnitScore                                   // SCORE
	UnitSeam                                    // SEAM
	UnitSection                                 // SECTION
	UnitShed                                    // SHED
	UnitShippington                             // SHIPPINGTON
	UnitShorthundredweight                      // SHORTHUNDREDWEIGHT
	UnitShortquarter                            // SHORTQUARTER
	UnitSiemens                                 // SIEMENS
	UnitSigma                                   // SIGMA
	UnitSkein                                   // SKEIN
	UnitSkot                                    // SKOT
	UnitSlug                                    // SLUG
	UnitSmoot                                   // SMOOT
	UnitSpan                                    // SPAN
	UnitSpat                                    // SPAT
	UnitSpindle                                 // SPINDLE
	UnitSquare                                  // SQUARE
	UnitStere                                   // STERE
	UnitSthene                                  // STHENE
	UnitStilb                                   // STILB
	UnitSb                                      // SB
	UnitStoke                                   // STOKE
	UnitStone                                   // STONE
	UnitStrike                                  // STRIKE
	UnitSurveyfoot                              // SURVEYFOOT
	UnitSurveyyard                              // SURVEYYARD
	UnitSurveyorschain                          // SURVEYORSCHAIN
	UnitSurveyorslink                           // SURVEYORSLINK
	UnitTablespoon                              // TABLESPOON
	UnitTeaspoon                                // TEASPOON
	UnitTesla                                   // TESLA
	UnitTherm                                   // THERM
	UnitThermie                                 // THERMIE
	UnitTimberfoot                              // TIMBERFOOT
	UnitTnt                                     // TNT
	UnitTonne                                   // TONNE
	UnitTorr                                    // TORR
	UnitTownship                                // TOWNSHIP
	UnitTsp                                     // TSP
	UnitTun                                     // TUN
	UnitWater                                   // WATER
	UnitWey                                     // WEY
	UnitWeymass                                 // WEYMASS
	UnitX                                       // X
	UnitK                                       // K
	UnitDegc                                    // DEGC
	UnitBrewster                                // BREWSTER
	UnitDrachm                                  // DRACHM
	UnitPoncelet                                // PONCELET
	UnitDenier                                  // DENIER
	UnitTex                                     // TEX
	UnitEnglishell                              // ENGLISHELL
	UnitScottishell                             // SCOTTISHELL
	UnitFlemishell                              // FLEMISHELL
	UnitBohrradius                              // BOHRRADIUS
	UnitBecquerel                               // BECQUEREL
	UnitFresnel                                 // FRESNEL
	UnitStatcoul                                // STATCOUL
	UnitStatamp                                 // STATAMP
	UnitStatvolt                                // STATVOLT
	UnitStatcoulomb                             // STATCOULOMB
	UnitStatampere                              // STATAMPERE
	UnitDebye                                   // DEBYE
	UnitPulsatance                              // PULSATANCE
	UnitRpm                                     // RPM
	UnitRps                                     // RPS
	UnitKilohm                                  // KILOHM
	UnitMegohm                                  // MEGOHM
	UnitSiderealyear                            // SIDEREALYEAR
	UnitSiderealday                             // SIDEREALDAY
	UnitSiderealhour                            // SIDEREALHOUR
	UnitLunarmonth                              // LUNARMONTH
	UnitSynodicmonth                            // SYNODICMONTH
	UnitSiderealmonth                           // SIDEREALMONTH
	UnitTropicalyear                            // TROPICALYEAR
	UnitSolaryear                               // SOLARYEAR
	UnitLunaryear                               // LUNARYEAR
	UnitCran                                    // CRAN
	UnitKip                                     // KIP
	UnitFrenchfoot                              // FRENCHFOOT
	UnitToise                                   // TOISE
	UnitSievert                                 // SIEVERT
	UnitCandle                                  // CANDLE
	UnitMilitarypace                            // MILITARYPACE
	UnitIudiptheria                             // IUDIPTHERIA
	UnitIupenicillin                            // IUPENICILLIN
	UnitIuinsulin                               // IUINSULIN
	UnitCottonyarncount                         // COTTONYARNCOUNT
	UnitLinenyarncount                          // LINENYARNCOUNT
	UnitWorstedyarncount                        // WORSTEDYARNCOUNT
	UnitMetricyarncount                         // METRICYARNCOUNT
	UnitJewlerspoint                            // JEWLERSPOINT

	maxUnit = UnitJewlerspoint
)

var MapStringToUnit = func() map[string]Unit {
	m := make(map[string]Unit)
	for i := Unit(0); i <= maxUnit; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Unit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Unit) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Unit) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}

func (r *Unit) UnmarshalText(text []byte) error {
	s := string(text)
	if v, ok := MapStringToUnit[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Unit: %q", s)
}
