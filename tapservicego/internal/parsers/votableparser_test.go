package parsers

import (
	"ataps/pkg/votable"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVOTable(t *testing.T) {
	data := []map[string]interface{}{
		{"a": 1, "b": 2},
		{"a": 3, "b": 4},
	}
	votable, err := CreateVOTable(data)
	if err != nil {
		t.Errorf("CreateVOTable() error = %v", err)
		return
	}
	assert.Equal(t, "1.4", votable.Version)
	assert.Equal(t, "http://www.ivoa.net/xml/VOTable/v1.4", votable.Xmlns)
	assert.Equal(t, "results", votable.Resource.Type)
	assert.Equal(t, "QUERY_STATUS", votable.Resource.Infos[0].Name)
	assert.Equal(t, "OK", votable.Resource.Infos[0].Value)
	assert.Equal(t, "results", votable.Resource.Tables[0].Name)
	assert.Equal(t, "Results of the query", votable.Resource.Tables[0].Description)
	assert.Equal(t, "a", votable.Resource.Tables[0].Fields[0].Name)
	assert.Equal(t, "b", votable.Resource.Tables[0].Fields[1].Name)
	assert.Equal(t, "1", votable.Resource.Tables[0].Data.TableData.Rows[0].Columns[0].Value)
	assert.Equal(t, "2", votable.Resource.Tables[0].Data.TableData.Rows[0].Columns[1].Value)
	assert.Equal(t, "3", votable.Resource.Tables[0].Data.TableData.Rows[1].Columns[0].Value)
	assert.Equal(t, "4", votable.Resource.Tables[0].Data.TableData.Rows[1].Columns[1].Value)
}

func TestCreateVOTableEmptyData(t *testing.T) {
	data := []map[string]interface{}{}
	votable, err := CreateVOTable(data)
	if err != nil {
		t.Errorf("CreateVOTable() error = %v", err)
		return
	}
	assert.Equal(t, "1.4", votable.Version)
	assert.Equal(t, "http://www.ivoa.net/xml/VOTable/v1.4", votable.Xmlns)
	assert.Equal(t, "results", votable.Resource.Type)
	assert.Equal(t, "QUERY_STATUS", votable.Resource.Infos[0].Name)
	assert.Equal(t, "OK", votable.Resource.Infos[0].Value)
	assert.Equal(t, "results", votable.Resource.Tables[0].Name)
	assert.Equal(t, "Results of the query", votable.Resource.Tables[0].Description)
	assert.Equal(t, 0, len(votable.Resource.Tables[0].Fields))
	assert.Equal(t, 0, len(votable.Resource.Tables[0].Data.TableData.Rows))
}

func TestVOTableToXML(t *testing.T) {
	votable := votable.VOTable{
		Version: "1.4",
		Xmlns:   "http://www.ivoa.net/xml/VOTable/v1.4",
		Resource: votable.Resource{
			Type:  "results",
			Infos: []votable.Info{{Name: "QUERY_STATUS", Value: "OK"}},
			Tables: []votable.Table{
				{
					Name:        "results",
					Description: "Results of the query",
					Fields: []votable.Field{
						{Name: "RA", Datatype: "double", Unit: "deg"},
						{Name: "DEC", Datatype: "double", Unit: "deg"},
						{Name: "MAG", Datatype: "float", Unit: "mag", Description: "Magnitude"},
					},
					Data: votable.Data{
						TableData: votable.TableData{
							Rows: []votable.Row{
								{Columns: []votable.Column{{Value: "10.0"}, {Value: "20.0"}, {Value: "15.0"}}},
								{Columns: []votable.Column{{Value: "20.0"}, {Value: "30.0"}, {Value: "16.0"}}},
								{Columns: []votable.Column{{Value: "30.0"}, {Value: "40.0"}, {Value: "17.0"}}},
							},
						},
					},
				},
			},
		},
	}
	result, err := VOTableToXML(votable)
	if err != nil {
		t.Errorf("VOTableToXML() error = %v", err)
		return
	}
	expectedResult := `<?xml version="1.0" encoding="UTF-8"?>
<VOTABLE version="1.4" xmlns="http://www.ivoa.net/xml/VOTable/v1.4">
	<RESOURCE type="results">
		<INFO name="QUERY_STATUS" value="OK"></INFO>
		<TABLE name="results">
			<DESCRIPTION>Results of the query</DESCRIPTION>
			<FIELD name="RA" datatype="double" unit="deg"></FIELD>
			<FIELD name="DEC" datatype="double" unit="deg"></FIELD>
			<FIELD name="MAG" datatype="float" unit="mag">
				<DESCRIPTION>Magnitude</DESCRIPTION>
			</FIELD>
			<DATA>
				<TABLEDATA>
					<TR>
						<TD>10.0</TD>
						<TD>20.0</TD>
						<TD>15.0</TD>
					</TR>
					<TR>
						<TD>20.0</TD>
						<TD>30.0</TD>
						<TD>16.0</TD>
					</TR>
					<TR>
						<TD>30.0</TD>
						<TD>40.0</TD>
						<TD>17.0</TD>
					</TR>
				</TABLEDATA>
			</DATA>
		</TABLE>
	</RESOURCE>
</VOTABLE>`
	assert.Equal(t, expectedResult, result)
}
