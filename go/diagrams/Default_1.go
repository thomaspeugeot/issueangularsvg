package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "issueangularsvg/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default_1 models.StageStruct
var ___dummy__Time_Default_1 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default_1 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default_1 map[string]any = map[string]any{
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
// 	InjectionGateway["Default_1"] = Default_1Injection
// }

// Default_1Injection will stage objects of database "Default_1"
func Default_1Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default_1 := (&models.Classdiagram{Name: `Default_1`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default_1.Name = `Default_1`
	__Classdiagram__000000_Default_1.IsInDrawMode = false

	// Setup of pointers
}


