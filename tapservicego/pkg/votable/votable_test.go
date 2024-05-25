package votable

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVOTable(t *testing.T) {
	votable := &VOTable{
		Version: "1.4",
		Xmlns:   "http://www.ivoa.net/xml/VOTable/v1.4",
		Resource: Resource{
			Type: "results",
			Infos: []Info{{Name: "QUERY_STATUS", Value: "OK"}},
			Tables: []Table{
				{
					Name: "results",
					Description: "Results of the query",
					Fields: []Field{
						{Name: "RA", Datatype: "double", Unit: "deg"},
						{Name: "DEC", Datatype: "double", Unit: "deg"},
						{Name: "MAG", Datatype: "float", Unit: "mag", Description: "Magnitude"},
					},
					Data: Data{
						TableData: TableData{
							Rows: []Row{
								{Columns: []Column{{Value: "10.0"}, {Value: "20.0"}, {Value: "15.0"}}},
								{Columns: []Column{{Value: "20.0"}, {Value: "30.0"}, {Value: "16.0"}}},
								{Columns: []Column{{Value: "30.0"}, {Value: "40.0"}, {Value: "17.0"}}},
							},
						},
					},
				},
			},
		},
	}
	var xmlBuilder strings.Builder
	encoder := xml.NewEncoder(&xmlBuilder)
	encoder.Indent("", "\t")
	err := encoder.Encode(votable)
	if err != nil {
		t.Errorf("Error encoding VOTable: %v", err)
	}
	votableXML := xmlBuilder.String()
	expectedVOTableXML := `<VOTABLE version="1.4" xmlns="http://www.ivoa.net/xml/VOTable/v1.4">
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
	assert.Equal(t, expectedVOTableXML, votableXML)
}
