package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "issueangularsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Country": &(ref_models.Country{}),

	"ref_models.Country.AlternateHellos": (ref_models.Country{}).AlternateHellos,

	"ref_models.Country.Hello": (ref_models.Country{}).Hello,

	"ref_models.Country.Name": (ref_models.Country{}).Name,

	"ref_models.Hello": &(ref_models.Hello{}),

	"ref_models.Hello.Name": (ref_models.Hello{}).Name,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Country := (&models.GongStructShape{Name: `Default-Country`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Country := (&models.Position{Name: `Pos-Default-Country`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true

	// GongStructShape values setup
	__GongStructShape__000000_Default_Country.Name = `Default-Country`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Country]
	__GongStructShape__000000_Default_Country.Identifier = `ref_models.Country`
	__GongStructShape__000000_Default_Country.ShowNbInstances = false
	__GongStructShape__000000_Default_Country.NbInstances = 0
	__GongStructShape__000000_Default_Country.Width = 240.000000
	__GongStructShape__000000_Default_Country.Height = 63.000000
	__GongStructShape__000000_Default_Country.IsSelected = false

	// Position values setup
	__Position__000000_Pos_Default_Country.X = 16.000000
	__Position__000000_Pos_Default_Country.Y = 23.000000
	__Position__000000_Pos_Default_Country.Name = `Pos-Default-Country`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Country)
	__GongStructShape__000000_Default_Country.Position = __Position__000000_Pos_Default_Country
}


