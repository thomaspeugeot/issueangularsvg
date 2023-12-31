// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Countrys           map[*Country]any
	Countrys_mapString map[string]*Country

	// insertion point for slice of pointers maps
	Country_AlternateHellos_reverseMap map[*Hello]*Country

	OnAfterCountryCreateCallback OnAfterCreateInterface[Country]
	OnAfterCountryUpdateCallback OnAfterUpdateInterface[Country]
	OnAfterCountryDeleteCallback OnAfterDeleteInterface[Country]
	OnAfterCountryReadCallback   OnAfterReadInterface[Country]

	Hellos           map[*Hello]any
	Hellos_mapString map[string]*Hello

	// insertion point for slice of pointers maps

	OnAfterHelloCreateCallback OnAfterCreateInterface[Hello]
	OnAfterHelloUpdateCallback OnAfterUpdateInterface[Hello]
	OnAfterHelloDeleteCallback OnAfterDeleteInterface[Hello]
	OnAfterHelloReadCallback   OnAfterReadInterface[Hello]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "issueangularsvg/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitCountry(country *Country)
	CheckoutCountry(country *Country)
	CommitHello(hello *Hello)
	CheckoutHello(hello *Hello)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Countrys:           make(map[*Country]any),
		Countrys_mapString: make(map[string]*Country),

		Hellos:           make(map[*Hello]any),
		Hellos_mapString: make(map[string]*Hello),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Country"] = len(stage.Countrys)
	stage.Map_GongStructName_InstancesNb["Hello"] = len(stage.Hellos)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Country"] = len(stage.Countrys)
	stage.Map_GongStructName_InstancesNb["Hello"] = len(stage.Hellos)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts country to the model stage
func (country *Country) Stage(stage *StageStruct) *Country {
	stage.Countrys[country] = __member
	stage.Countrys_mapString[country.Name] = country

	return country
}

// Unstage removes country off the model stage
func (country *Country) Unstage(stage *StageStruct) *Country {
	delete(stage.Countrys, country)
	delete(stage.Countrys_mapString, country.Name)
	return country
}

// UnstageVoid removes country off the model stage
func (country *Country) UnstageVoid(stage *StageStruct) {
	delete(stage.Countrys, country)
	delete(stage.Countrys_mapString, country.Name)
}

// commit country to the back repo (if it is already staged)
func (country *Country) Commit(stage *StageStruct) *Country {
	if _, ok := stage.Countrys[country]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCountry(country)
		}
	}
	return country
}

func (country *Country) CommitVoid(stage *StageStruct) {
	country.Commit(stage)
}

// Checkout country to the back repo (if it is already staged)
func (country *Country) Checkout(stage *StageStruct) *Country {
	if _, ok := stage.Countrys[country]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCountry(country)
		}
	}
	return country
}

// for satisfaction of GongStruct interface
func (country *Country) GetName() (res string) {
	return country.Name
}

// Stage puts hello to the model stage
func (hello *Hello) Stage(stage *StageStruct) *Hello {
	stage.Hellos[hello] = __member
	stage.Hellos_mapString[hello.Name] = hello

	return hello
}

// Unstage removes hello off the model stage
func (hello *Hello) Unstage(stage *StageStruct) *Hello {
	delete(stage.Hellos, hello)
	delete(stage.Hellos_mapString, hello.Name)
	return hello
}

// UnstageVoid removes hello off the model stage
func (hello *Hello) UnstageVoid(stage *StageStruct) {
	delete(stage.Hellos, hello)
	delete(stage.Hellos_mapString, hello.Name)
}

// commit hello to the back repo (if it is already staged)
func (hello *Hello) Commit(stage *StageStruct) *Hello {
	if _, ok := stage.Hellos[hello]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitHello(hello)
		}
	}
	return hello
}

func (hello *Hello) CommitVoid(stage *StageStruct) {
	hello.Commit(stage)
}

// Checkout hello to the back repo (if it is already staged)
func (hello *Hello) Checkout(stage *StageStruct) *Hello {
	if _, ok := stage.Hellos[hello]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutHello(hello)
		}
	}
	return hello
}

// for satisfaction of GongStruct interface
func (hello *Hello) GetName() (res string) {
	return hello.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCountry(Country *Country)
	CreateORMHello(Hello *Hello)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCountry(Country *Country)
	DeleteORMHello(Hello *Hello)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Countrys = make(map[*Country]any)
	stage.Countrys_mapString = make(map[string]*Country)

	stage.Hellos = make(map[*Hello]any)
	stage.Hellos_mapString = make(map[string]*Hello)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Countrys = nil
	stage.Countrys_mapString = nil

	stage.Hellos = nil
	stage.Hellos_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for country := range stage.Countrys {
		country.Unstage(stage)
	}

	for hello := range stage.Hellos {
		hello.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Country | Hello
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*Country | *Hello
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*Country]any |
		map[*Hello]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Country |
		map[string]*Hello |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Country]any:
		return any(&stage.Countrys).(*Type)
	case map[*Hello]any:
		return any(&stage.Hellos).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Country:
		return any(&stage.Countrys_mapString).(*Type)
	case map[string]*Hello:
		return any(&stage.Hellos_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Country:
		return any(&stage.Countrys).(*map[*Type]any)
	case Hello:
		return any(&stage.Hellos).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Country:
		return any(&stage.Countrys).(*map[Type]any)
	case *Hello:
		return any(&stage.Hellos).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Country:
		return any(&stage.Countrys_mapString).(*map[string]*Type)
	case Hello:
		return any(&stage.Hellos_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Country:
		return any(&Country{
			// Initialisation of associations
			// field is initialized with an instance of Hello with the name of the field
			Hello: &Hello{Name: "Hello"},
			// field is initialized with an instance of Hello with the name of the field
			AlternateHellos: []*Hello{{Name: "AlternateHellos"}},
		}).(*Type)
	case Hello:
		return any(&Hello{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Country
	case Country:
		switch fieldname {
		// insertion point for per direct association field
		case "Hello":
			res := make(map[*Hello][]*Country)
			for country := range stage.Countrys {
				if country.Hello != nil {
					hello_ := country.Hello
					var countrys []*Country
					_, ok := res[hello_]
					if ok {
						countrys = res[hello_]
					} else {
						countrys = make([]*Country, 0)
					}
					countrys = append(countrys, country)
					res[hello_] = countrys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Hello
	case Hello:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Country
	case Country:
		switch fieldname {
		// insertion point for per direct association field
		case "AlternateHellos":
			res := make(map[*Hello]*Country)
			for country := range stage.Countrys {
				for _, hello_ := range country.AlternateHellos {
					res[hello_] = country
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Hello
	case Hello:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Country:
		res = "Country"
	case Hello:
		res = "Hello"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Country:
		res = "Country"
	case *Hello:
		res = "Hello"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Country:
		res = []string{"Name", "Hello", "AlternateHellos"}
	case Hello:
		res = []string{"Name"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Country:
		var rf ReverseField
		_ = rf
	case Hello:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Country"
		rf.Fieldname = "AlternateHellos"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Country:
		res = []string{"Name", "Hello", "AlternateHellos"}
	case *Hello:
		res = []string{"Name"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Country:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Hello":
			if inferedInstance.Hello != nil {
				res = inferedInstance.Hello.Name
			}
		case "AlternateHellos":
			for idx, __instance__ := range inferedInstance.AlternateHellos {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Hello:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Country:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Hello":
			if inferedInstance.Hello != nil {
				res = inferedInstance.Hello.Name
			}
		case "AlternateHellos":
			for idx, __instance__ := range inferedInstance.AlternateHellos {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Hello:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
