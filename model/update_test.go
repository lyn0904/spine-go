package model

import (
	"testing"

	"github.com/lyn0904/spine-go/util"
	"github.com/stretchr/testify/assert"
)

type TestUpdateData struct {
	Id           *uint `eebus:"key"`
	IsChangeable *bool `eebus:"writecheck"`
	DataItem     *int
}

type TestUpdater struct {
	// updateSelectorHashKey *string
	// deleteSelectorHashKey *string
}

func TestUpdateList_NewItem(t *testing.T) {
	existingData := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(1))}}
	newData := []TestUpdateData{{Id: util.Ptr(uint(2)), DataItem: util.Ptr(int(2))}}

	expectedResult := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(1))}, {Id: util.Ptr(uint(2)), DataItem: util.Ptr(int(2))}}

	// Act
	result, boolV := UpdateList(false, existingData, newData, nil, nil)

	assert.True(t, boolV)
	assert.Equal(t, expectedResult, result)

	expectedResult = []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(1))}}

	// Act
	result, boolV = UpdateList(true, existingData, newData, nil, nil)

	assert.False(t, boolV)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateList_ChangedItem(t *testing.T) {
	existingData := []TestUpdateData{{Id: util.Ptr(uint(1)), IsChangeable: util.Ptr(false), DataItem: util.Ptr(int(1))}}
	newData := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(2))}}

	expectedResult := []TestUpdateData{{Id: util.Ptr(uint(1)), IsChangeable: util.Ptr(false), DataItem: util.Ptr(int(2))}}

	// Act
	result, boolV := UpdateList(false, existingData, newData, nil, nil)

	assert.True(t, boolV)
	assert.Equal(t, expectedResult, result)

	expectedResult = []TestUpdateData{{Id: util.Ptr(uint(1)), IsChangeable: util.Ptr(false), DataItem: util.Ptr(int(1))}}

	// Act
	result, boolV = UpdateList(true, existingData, newData, nil, nil)

	assert.False(t, boolV)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateList_NewAndChangedItem(t *testing.T) {
	existingData := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(1))}}
	newData := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(2))}, {Id: util.Ptr(uint(3)), DataItem: util.Ptr(int(3))}}

	expectedResult := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(2))}, {Id: util.Ptr(uint(3)), DataItem: util.Ptr(int(3))}}

	// Act
	result, boolV := UpdateList(false, existingData, newData, nil, nil)

	assert.True(t, boolV)
	assert.Equal(t, expectedResult, result)

	expectedResult = []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(1))}}

	// Act
	result, boolV = UpdateList(true, existingData, newData, nil, nil)

	assert.False(t, boolV)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateList_ItemWithNoIdentifier(t *testing.T) {
	existingData := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(1))}, {Id: util.Ptr(uint(2)), DataItem: util.Ptr(int(2))}}
	newData := []TestUpdateData{{DataItem: util.Ptr(int(3))}}

	expectedResult := []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(3))}, {Id: util.Ptr(uint(2)), DataItem: util.Ptr(int(3))}}

	// Act
	result, boolV := UpdateList(false, existingData, newData, nil, nil)

	assert.True(t, boolV)
	assert.Equal(t, expectedResult, result)

	expectedResult = []TestUpdateData{{Id: util.Ptr(uint(1)), DataItem: util.Ptr(int(3))}, {Id: util.Ptr(uint(2)), DataItem: util.Ptr(int(3))}}

	// Act
	result, boolV = UpdateList(true, existingData, newData, nil, nil)

	assert.False(t, boolV)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateList_FilterDelete(t *testing.T) {
	existingData := []LoadControlLimitDataType{
		{
			LimitId: util.Ptr(LoadControlLimitIdType(0)),
			Value:   NewScaledNumberType(0),
		},
		{
			LimitId: util.Ptr(LoadControlLimitIdType(1)),
			Value:   NewScaledNumberType(0),
		},
	}
	newData := []LoadControlLimitDataType{
		{
			LimitId: util.Ptr(LoadControlLimitIdType(1)),
			Value:   NewScaledNumberType(10),
		},
	}

	expectedResult := []LoadControlLimitDataType{
		{
			LimitId: util.Ptr(LoadControlLimitIdType(1)),
			Value:   NewScaledNumberType(10),
		},
	}

	filterDelete := &FilterType{CmdControl: &CmdControlType{Delete: &ElementTagType{}}}
	filterDelete.CmdControl.Delete = new(ElementTagType)
	filterDelete.LoadControlLimitListDataSelectors = &LoadControlLimitListDataSelectorsType{
		LimitId: util.Ptr(LoadControlLimitIdType(0)),
	}

	filterPartial := NewFilterTypePartial()
	filterPartial.LoadControlLimitListDataSelectors = &LoadControlLimitListDataSelectorsType{
		LimitId: util.Ptr(LoadControlLimitIdType(1)),
	}

	// Act
	result, boolV := UpdateList(false, existingData, newData, filterPartial, filterDelete)

	assert.True(t, boolV)
	assert.Equal(t, expectedResult, result)

	newData = []LoadControlLimitDataType{
		{
			Value: NewScaledNumberType(10),
		},
	}

	expectedResult = []LoadControlLimitDataType{
		{
			LimitId: util.Ptr(LoadControlLimitIdType(0)),
			Value:   NewScaledNumberType(0),
		},
		{
			LimitId: util.Ptr(LoadControlLimitIdType(1)),
			Value:   NewScaledNumberType(0),
		},
	}

	// Act
	result, boolV = UpdateList(true, existingData, newData, filterPartial, filterDelete)

	assert.False(t, boolV)
	assert.Equal(t, expectedResult, result)
}

func TestRemoveFieldFromType(t *testing.T) {
	items := &LoadControlLimitListDataType{
		LoadControlLimitData: []LoadControlLimitDataType{
			{
				LimitId: util.Ptr(LoadControlLimitIdType(1)),
				Value:   NewScaledNumberType(16.0),
			},
		},
	}

	elements := &LoadControlLimitDataElementsType{
		Value: &ScaledNumberElementsType{},
	}

	RemoveElementFromItem(&items.LoadControlLimitData[0], elements)

	var nilValue *ScaledNumberType

	assert.Equal(t, nilValue, items.LoadControlLimitData[0].Value)
}

// TODO: Fix, as these tests won't work right now as TestUpdater doesn't use FilterProvider and its data structure
/*
func TestUpdateList_UpdateSelector(t *testing.T) {
	existingData := []TestUpdateData{{Id: util.Ptr(1), DataItem: 1}, {Id: util.Ptr(2), DataItem: 2}}
	newData := []TestUpdateData{{DataItem: 3}}

	dataProvider := &TestUpdater{
		updateSelectorHashKey: util.Ptr("1"),
	}
	expectedResult := []TestUpdateData{{Id: util.Ptr(1), DataItem: 3}, {Id: util.Ptr(2), DataItem: 2}}

	// Act
	result := UpdateList[TestUpdateData](existingData, newData, dataProvider)

	assert.Equal(t, expectedResult, result)
}

func TestUpdateList_DeleteSelector(t *testing.T) {
	existingData := []TestUpdateData{{Id: util.Ptr(1), DataItem: 1}, {Id: util.Ptr(2), DataItem: 2}}
	newData := []TestUpdateData{{Id: util.Ptr(0), DataItem: 0}}

	dataProvider := &TestUpdater{
		deleteSelectorHashKey: util.Ptr("1"),
	}
	expectedResult := []TestUpdateData{{Id: util.Ptr(2), DataItem: 2}}

	// Act
	result := UpdateList[TestUpdateData](existingData, newData, dataProvider)

	assert.Equal(t, expectedResult, result)
}
*/
