package av

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/assert/cmp"
)

func TestPropertyUnmarshal1(t *testing.T) {
	expected := Property{
		Name:  "The name",
		Key:   "The key",
		Value: ValueText("The value"),
	}
	jsonstr := []byte(`
	[
		"The name",
		"The key",
		{
			"TEXT": "The value"
		}
	]
	`)

	var actual Property
	err := json.Unmarshal(jsonstr, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actual, expected)(); !res.Success() {
		t.Fatal(res)
	}
}

func TestPropertyUnmarshal2(t *testing.T) {
	expected := Property{
		Name:  "Annotation Property",
		Key:   "Annotation",
		Value: Value{},
		Annotations: AnnotationMap{
			AttributeExample: ValueString("Annotation Value"),
		},
	}
	jsonstr := []byte(`
	[
		"Annotation Property",
		"Annotation",
		{
			"NULL": ""
		},
		{
			"EXAMPLE_ATTRIBUTE": {
				"STRING": "Annotation Value"
			}
		}
	]
	`)

	var actual Property
	err := json.Unmarshal(jsonstr, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actual, expected)(); !res.Success() {
		t.Fatal(res)
	}
}

func TestFactSliceUnmarshalValue(t *testing.T) {
	expected := []Fact{
		{
			Attribute: AttributeAcceleration,
			Value:     ValueDuration(12345.678900000),
		},
	}
	jsonstr := []byte(`
	[
		[
			"ACCELERATION_ATTRIBUTE",
			{"DURATION":"12345.678900000"},
			[],
			[],
			[],
			[]
		]
	]
	`)

	actual := []Fact{}
	err := json.Unmarshal(jsonstr, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actual, expected)(); !res.Success() {
		t.Fatal(res)
	}
}

func TestFactUnmarshalFacet(t *testing.T) {
	expected := Fact{
		Attribute: AttributeAcceleration,
		Value:     ValueDuration(12345.678900000),
		Facets: []Facet{
			{
				Name:  "The facet name",
				Value: ValueText("The facet value"),
				Factors: FactorList{
					Factor{
						Key:   "The factor 1 key",
						Value: ValueInteger(1),
					},
					Factor{
						Key:   "The factor 2 key",
						Value: ValueInteger(2),
					},
				},
			},
			{
				Name:  "The second facet name",
				Value: ValueString("The second facet value"),
				Factors: FactorList{
					Factor{
						Key:   "The factor 1 key",
						Value: ValueText("one"),
					},
					Factor{
						Key:   "The factor 2 key",
						Value: ValueText("two"),
					},
				},
			},
		},
	}
	jsonstr := []byte(`
	[
		"ACCELERATION_ATTRIBUTE",
		{"DURATION":"12345.678900000"},
		[
			[
				"The facet name",
				{
					"TEXT": "The facet value"
				},
				[
					[
						"The factor 1 key",
						{
							"INTEGER": "1"
						}
					],
					[
						"The factor 2 key",
						{
							"INTEGER": "2"
						}
					]
				]
			],
			[
				"The second facet name",
				{
					"STRING": "The second facet value"
				},
				[
					[
						"The factor 1 key",
						{
							"TEXT": "one"
						}
					],
					[
						"The factor 2 key",
						{
							"TEXT": "two"
						}
					]
				]
			]
		],
		[],
		[],
		[]
	]
	`)

	actual := Fact{}
	err := json.Unmarshal(jsonstr, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actual, expected)(); !res.Success() {
		t.Fatal(res)
	}
}

func TestFactUnmarshalFeature(t *testing.T) {
	expected := Fact{
		Attribute: AttributeAcceleration,
		Value:     ValueDuration(12345.678900000),
		Features: []Feature{
			{
				Name:  "First feature name",
				Key:   "First feature key",
				Value: ValueText("First feature value"),
			},
			{
				Name:  "Second feature name",
				Value: ValueInteger(42),
			},
			{
				Name:  "Third feature name",
				Value: ValueInteger(84),
			},
		},
	}
	jsonstr := []byte(`
	[
		"ACCELERATION_ATTRIBUTE",
		{"DURATION":"12345.678900000"},
		[],
		[
			[
				"First feature name",
				"First feature key",
				{
					"TEXT": "First feature value"
				}
			],
			[
				"Second feature name",
				"",
				{
					"INTEGER": "42"
				}
			],
			[
				"Third feature name",
				"",
				{
					"INTEGER": "84"
				}
			]
		],
		[],
		[]
	]
	`)

	actual := Fact{}
	err := json.Unmarshal(jsonstr, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actual, expected)(); !res.Success() {
		t.Fatal(res)
	}
}

func TestFactMapUnmarshalFieldsFrames(t *testing.T) {
	expected := Fact{
		Attribute: AttributeAcceleration,
		Value:     ValueDuration(12345.678900000),
		Fields: []Field{
			{
				Name:         "First field name",
				DefaultValue: ValueText("First field default value"),
			},
			{
				Name:         "Second field name",
				DefaultValue: ValueInteger(22),
			},
		},
		Frames: []Frame{
			{
				Key: "First frame key",
				Values: []Value{
					ValueText("First frame value 1"),
					ValueText("First frame value 2"),
				},
			},
			{
				Key: "Second frame key",
				Values: []Value{
					ValueInteger(1),
					ValueInteger(2),
				},
			},
		},
	}
	jsonstr := []byte(`
	[
		"ACCELERATION_ATTRIBUTE",
		{"DURATION":"12345.678900000"},
		[],
		[],
		[
			[
				"First field name",
				{
					"TEXT": "First field default value"
				}
			],
			[
				"Second field name",
				{
					"INTEGER": "22"
				}
			]
		],
		[
			[
				"First frame key",
				[
					{
						"TEXT": "First frame value 1"
					},
					{
						"TEXT": "First frame value 2"
					}
				]
			],
			[
				"Second frame key",
				[
					{
						"INTEGER": "1"
					},
					{
						"INTEGER": "2"
					}
				]
			]
		]
	]
	`)

	actual := Fact{}
	err := json.Unmarshal(jsonstr, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actual, expected)(); !res.Success() {
		t.Fatal(res)
	}
}

func _perform3waytest(t *testing.T, expected AvialModel, jsonstr []byte) {
	var actualCreated AvialModel
	err := json.Unmarshal(jsonstr, &actualCreated)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actualCreated, expected)(); !res.Success() {
		t.Fatal(res)
	}

	token := setup(t)
	defer teardown(token)

	rootToken := MustParseToken("636d384c-54d5-4bab-99b0-9b29236cff59")
	token = rootToken

	e, err := CreateEntity(ArgsCreateEntity{
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Created! ", e)

	err = e.ConnectMethod(ArgsConnectMethod{
		Outlet:        NewEntity(0, 0, 14),
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = e.Invoke(ArgsInvokeEntity{
		Method:        MethodCreate,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Initialized general adapter for ", e)
	interchange_json, err := json.Marshal(actualCreated)
	fmt.Printf("Storing interchange (%d bytes)...\n", len(interchange_json))
	err = e.Store(actualCreated, ModeNull, token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Done! Retrieving the interchange again...")
	actualRetrieved, err := e.Retrieve(token)
	if err != nil {
		t.Fatal(err)
	}

	if res := cmp.DeepEqual(actualCreated, actualRetrieved)(); !res.Success() {
		t.Fatal(res)
	}
	fmt.Printf("OK")
}

func TestInterchangeMarshal3way(t *testing.T) {
	expected := AvialModel{
		Name: "Test interchange",
		Key:  "test-interchange",
		Properties: []Property{
			{
				Name:  "First property name",
				Key:   "first-property-key",
				Value: ValueText("First property value"),
			},
			{
				Name:  "Annotation Property",
				Key:   "Annotation",
				Value: Value{},
				Annotations: AnnotationMap{
					AttributeExample: ValueString("Annotation Value"),
				},
			},
		},
		Facts: []Fact{
			{
				Attribute: AttributeAcceleration,
				Value:     ValueDuration(12345.678900000),
				Facets: []Facet{
					{
						Name:  "The facet name",
						Value: ValueText("The facet value"),
						Factors: FactorList{
							Factor{
								Key:   "The factor 1 key",
								Value: ValueInteger(1),
							},
							Factor{
								Key:   "The factor 2 key",
								Value: ValueInteger(2),
							},
						},
					},
					{
						Name:  "The second facet name",
						Value: ValueString("The second facet value"),
						Factors: FactorList{
							Factor{
								Key:   "The factor 1 key",
								Value: ValueText("one"),
							},
							Factor{
								Key:   "The factor 2 key",
								Value: ValueText("two"),
							},
						},
					},
				},
				Features: []Feature{
					{
						Name:  "First feature name",
						Key:   "First feature key", // TODO: If key is empty, storing fails
						Value: ValueText("First feature value"),
					},
					{
						Name:  "Second feature name",
						Key:   "Second feature key",
						Value: ValueInteger(42),
					},
					{
						Name:  "Third feature name",
						Key:   "Third feature key",
						Value: ValueInteger(84),
					},
				},
				Fields: []Field{
					{
						Name:         "First field name",
						DefaultValue: ValueText("First field default value"),
					},
					{
						Name:         "Second field name",
						DefaultValue: ValueInteger(22),
					},
				},
				Frames: []Frame{
					{
						Key: "First frame key",
						Values: []Value{
							ValueText("First frame value 1"),
							ValueText("First frame value 2"),
						},
					},
					{
						Key: "Second frame key",
						Values: []Value{
							ValueInteger(1),
							ValueInteger(2),
						},
					},
				},
			},
		},
	}

	jsonstr := []byte(`
	{
		"Name": "Test interchange",
		"Key": "test-interchange",
		"Properties": [
			[
				"First property name",
				"first-property-key",
				{
					"TEXT": "First property value"
				}
			],
			[
				"Annotation Property",
				"Annotation",
				{
					"NULL": ""
				},
				{
					"EXAMPLE_ATTRIBUTE": {
						"STRING": "Annotation Value"
					}
				}
			]
		],
		"Facts": [
			[
				"ACCELERATION_ATTRIBUTE",
				{"DURATION":"12345.678900000"},
				[
					[
						"The facet name",
						{
							"TEXT": "The facet value"
						},
						[
							[
								"The factor 1 key",
								{
									"INTEGER": "1"
								}
							],
							[
								"The factor 2 key",
								{
									"INTEGER": "2"
								}
							]
						]
					],
					[
						"The second facet name",
						{
							"STRING": "The second facet value"
						},
						[
							[
								"The factor 1 key",
								{
									"TEXT": "one"
								}
							],
							[
								"The factor 2 key",
								{
									"TEXT": "two"
								}
							]
						]
					]
				],
				[
					[
						"First feature name",
						"First feature key",
						{
							"TEXT": "First feature value"
						}
					],
					[
						"Second feature name",
						"Second feature key",
						{
							"INTEGER": "42"
						}
					],
					[
						"Third feature name",
						"Third feature key",
						{
							"INTEGER": "84"
						}
					]
				],
				[
					[
						"First field name",
						{
							"TEXT": "First field default value"
						}
					],
					[
						"Second field name",
						{
							"INTEGER": "22"
						}
					]
				],
				[
					[
						"First frame key",
						[
							{
								"TEXT": "First frame value 1"
							},
							{
								"TEXT": "First frame value 2"
							}
						]
					],
					[
						"Second frame key",
						[
							{
								"INTEGER": "1"
							},
							{
								"INTEGER": "2"
							}
						]
					]
				]
			]
		]
	}
	`)

	_perform3waytest(t, expected, jsonstr)
}

// Try multiple null attribute, properties with null keys, etc.
