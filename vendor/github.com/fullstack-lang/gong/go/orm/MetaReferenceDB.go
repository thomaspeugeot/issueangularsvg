// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_MetaReference_sql sql.NullBool
var dummy_MetaReference_time time.Duration
var dummy_MetaReference_sort sort.Float64Slice

// MetaReferenceAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model metareferenceAPI
type MetaReferenceAPI struct {
	gorm.Model

	models.MetaReference_WOP

	// encoding of pointers
	MetaReferencePointersEncoding MetaReferencePointersEncoding
}

// MetaReferencePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type MetaReferencePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// MetaReferenceDB describes a metareference in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model metareferenceDB
type MetaReferenceDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field metareferenceDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	MetaReferencePointersEncoding
}

// MetaReferenceDBs arrays metareferenceDBs
// swagger:response metareferenceDBsResponse
type MetaReferenceDBs []MetaReferenceDB

// MetaReferenceDBResponse provides response
// swagger:response metareferenceDBResponse
type MetaReferenceDBResponse struct {
	MetaReferenceDB
}

// MetaReferenceWOP is a MetaReference without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type MetaReferenceWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var MetaReference_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoMetaReferenceStruct struct {
	// stores MetaReferenceDB according to their gorm ID
	Map_MetaReferenceDBID_MetaReferenceDB map[uint]*MetaReferenceDB

	// stores MetaReferenceDB ID according to MetaReference address
	Map_MetaReferencePtr_MetaReferenceDBID map[*models.MetaReference]uint

	// stores MetaReference according to their gorm ID
	Map_MetaReferenceDBID_MetaReferencePtr map[uint]*models.MetaReference

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoMetaReference *BackRepoMetaReferenceStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoMetaReference.stage
	return
}

func (backRepoMetaReference *BackRepoMetaReferenceStruct) GetDB() *gorm.DB {
	return backRepoMetaReference.db
}

// GetMetaReferenceDBFromMetaReferencePtr is a handy function to access the back repo instance from the stage instance
func (backRepoMetaReference *BackRepoMetaReferenceStruct) GetMetaReferenceDBFromMetaReferencePtr(metareference *models.MetaReference) (metareferenceDB *MetaReferenceDB) {
	id := backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference]
	metareferenceDB = backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[id]
	return
}

// BackRepoMetaReference.CommitPhaseOne commits all staged instances of MetaReference to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for metareference := range stage.MetaReferences {
		backRepoMetaReference.CommitPhaseOneInstance(metareference)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, metareference := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr {
		if _, ok := stage.MetaReferences[metareference]; !ok {
			backRepoMetaReference.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMetaReference.CommitDeleteInstance commits deletion of MetaReference to the BackRepo
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CommitDeleteInstance(id uint) (Error error) {

	metareference := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[id]

	// metareference is not staged anymore, remove metareferenceDB
	metareferenceDB := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[id]
	query := backRepoMetaReference.db.Unscoped().Delete(&metareferenceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID, metareference)
	delete(backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr, id)
	delete(backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB, id)

	return
}

// BackRepoMetaReference.CommitPhaseOneInstance commits metareference staged instances of MetaReference to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CommitPhaseOneInstance(metareference *models.MetaReference) (Error error) {

	// check if the metareference is not commited yet
	if _, ok := backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference]; ok {
		return
	}

	// initiate metareference
	var metareferenceDB MetaReferenceDB
	metareferenceDB.CopyBasicFieldsFromMetaReference(metareference)

	query := backRepoMetaReference.db.Create(&metareferenceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference] = metareferenceDB.ID
	backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID] = metareference
	backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[metareferenceDB.ID] = &metareferenceDB

	return
}

// BackRepoMetaReference.CommitPhaseTwo commits all staged instances of MetaReference to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, metareference := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr {
		backRepoMetaReference.CommitPhaseTwoInstance(backRepo, idx, metareference)
	}

	return
}

// BackRepoMetaReference.CommitPhaseTwoInstance commits {{structname }} of models.MetaReference to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, metareference *models.MetaReference) (Error error) {

	// fetch matching metareferenceDB
	if metareferenceDB, ok := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[idx]; ok {

		metareferenceDB.CopyBasicFieldsFromMetaReference(metareference)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoMetaReference.db.Save(&metareferenceDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown MetaReference intance %s", metareference.Name))
		return err
	}

	return
}

// BackRepoMetaReference.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CheckoutPhaseOne() (Error error) {

	metareferenceDBArray := make([]MetaReferenceDB, 0)
	query := backRepoMetaReference.db.Find(&metareferenceDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	metareferenceInstancesToBeRemovedFromTheStage := make(map[*models.MetaReference]any)
	for key, value := range backRepoMetaReference.stage.MetaReferences {
		metareferenceInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, metareferenceDB := range metareferenceDBArray {
		backRepoMetaReference.CheckoutPhaseOneInstance(&metareferenceDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		metareference, ok := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID]
		if ok {
			delete(metareferenceInstancesToBeRemovedFromTheStage, metareference)
		}
	}

	// remove from stage and back repo's 3 maps all metareferences that are not in the checkout
	for metareference := range metareferenceInstancesToBeRemovedFromTheStage {
		metareference.Unstage(backRepoMetaReference.GetStage())

		// remove instance from the back repo 3 maps
		metareferenceID := backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference]
		delete(backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID, metareference)
		delete(backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB, metareferenceID)
		delete(backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr, metareferenceID)
	}

	return
}

// CheckoutPhaseOneInstance takes a metareferenceDB that has been found in the DB, updates the backRepo and stages the
// models version of the metareferenceDB
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CheckoutPhaseOneInstance(metareferenceDB *MetaReferenceDB) (Error error) {

	metareference, ok := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID]
	if !ok {
		metareference = new(models.MetaReference)

		backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID] = metareference
		backRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference] = metareferenceDB.ID

		// append model store with the new element
		metareference.Name = metareferenceDB.Name_Data.String
		metareference.Stage(backRepoMetaReference.GetStage())
	}
	metareferenceDB.CopyBasicFieldsToMetaReference(metareference)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	metareference.Stage(backRepoMetaReference.GetStage())

	// preserve pointer to metareferenceDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_MetaReferenceDBID_MetaReferenceDB)[metareferenceDB hold variable pointers
	metareferenceDB_Data := *metareferenceDB
	preservedPtrToMetaReference := &metareferenceDB_Data
	backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[metareferenceDB.ID] = preservedPtrToMetaReference

	return
}

// BackRepoMetaReference.CheckoutPhaseTwo Checkouts all staged instances of MetaReference to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, metareferenceDB := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB {
		backRepoMetaReference.CheckoutPhaseTwoInstance(backRepo, metareferenceDB)
	}
	return
}

// BackRepoMetaReference.CheckoutPhaseTwoInstance Checkouts staged instances of MetaReference to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMetaReference *BackRepoMetaReferenceStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, metareferenceDB *MetaReferenceDB) (Error error) {

	metareference := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID]

	metareferenceDB.DecodePointers(backRepo, metareference)

	return
}

func (metareferenceDB *MetaReferenceDB) DecodePointers(backRepo *BackRepoStruct, metareference *models.MetaReference) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitMetaReference allows commit of a single metareference (if already staged)
func (backRepo *BackRepoStruct) CommitMetaReference(metareference *models.MetaReference) {
	backRepo.BackRepoMetaReference.CommitPhaseOneInstance(metareference)
	if id, ok := backRepo.BackRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference]; ok {
		backRepo.BackRepoMetaReference.CommitPhaseTwoInstance(backRepo, id, metareference)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitMetaReference allows checkout of a single metareference (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMetaReference(metareference *models.MetaReference) {
	// check if the metareference is staged
	if _, ok := backRepo.BackRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference]; ok {

		if id, ok := backRepo.BackRepoMetaReference.Map_MetaReferencePtr_MetaReferenceDBID[metareference]; ok {
			var metareferenceDB MetaReferenceDB
			metareferenceDB.ID = id

			if err := backRepo.BackRepoMetaReference.db.First(&metareferenceDB, id).Error; err != nil {
				log.Fatalln("CheckoutMetaReference : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMetaReference.CheckoutPhaseOneInstance(&metareferenceDB)
			backRepo.BackRepoMetaReference.CheckoutPhaseTwoInstance(backRepo, &metareferenceDB)
		}
	}
}

// CopyBasicFieldsFromMetaReference
func (metareferenceDB *MetaReferenceDB) CopyBasicFieldsFromMetaReference(metareference *models.MetaReference) {
	// insertion point for fields commit

	metareferenceDB.Name_Data.String = metareference.Name
	metareferenceDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromMetaReference_WOP
func (metareferenceDB *MetaReferenceDB) CopyBasicFieldsFromMetaReference_WOP(metareference *models.MetaReference_WOP) {
	// insertion point for fields commit

	metareferenceDB.Name_Data.String = metareference.Name
	metareferenceDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromMetaReferenceWOP
func (metareferenceDB *MetaReferenceDB) CopyBasicFieldsFromMetaReferenceWOP(metareference *MetaReferenceWOP) {
	// insertion point for fields commit

	metareferenceDB.Name_Data.String = metareference.Name
	metareferenceDB.Name_Data.Valid = true
}

// CopyBasicFieldsToMetaReference
func (metareferenceDB *MetaReferenceDB) CopyBasicFieldsToMetaReference(metareference *models.MetaReference) {
	// insertion point for checkout of basic fields (back repo to stage)
	metareference.Name = metareferenceDB.Name_Data.String
}

// CopyBasicFieldsToMetaReference_WOP
func (metareferenceDB *MetaReferenceDB) CopyBasicFieldsToMetaReference_WOP(metareference *models.MetaReference_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	metareference.Name = metareferenceDB.Name_Data.String
}

// CopyBasicFieldsToMetaReferenceWOP
func (metareferenceDB *MetaReferenceDB) CopyBasicFieldsToMetaReferenceWOP(metareference *MetaReferenceWOP) {
	metareference.ID = int(metareferenceDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	metareference.Name = metareferenceDB.Name_Data.String
}

// Backup generates a json file from a slice of all MetaReferenceDB instances in the backrepo
func (backRepoMetaReference *BackRepoMetaReferenceStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "MetaReferenceDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MetaReferenceDB, 0)
	for _, metareferenceDB := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB {
		forBackup = append(forBackup, metareferenceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json MetaReference ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json MetaReference file", err.Error())
	}
}

// Backup generates a json file from a slice of all MetaReferenceDB instances in the backrepo
func (backRepoMetaReference *BackRepoMetaReferenceStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MetaReferenceDB, 0)
	for _, metareferenceDB := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB {
		forBackup = append(forBackup, metareferenceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("MetaReference")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&MetaReference_Fields, -1)
	for _, metareferenceDB := range forBackup {

		var metareferenceWOP MetaReferenceWOP
		metareferenceDB.CopyBasicFieldsToMetaReferenceWOP(&metareferenceWOP)

		row := sh.AddRow()
		row.WriteStruct(&metareferenceWOP, -1)
	}
}

// RestoreXL from the "MetaReference" sheet all MetaReferenceDB instances
func (backRepoMetaReference *BackRepoMetaReferenceStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoMetaReferenceid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["MetaReference"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoMetaReference.rowVisitorMetaReference)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoMetaReference *BackRepoMetaReferenceStruct) rowVisitorMetaReference(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var metareferenceWOP MetaReferenceWOP
		row.ReadStruct(&metareferenceWOP)

		// add the unmarshalled struct to the stage
		metareferenceDB := new(MetaReferenceDB)
		metareferenceDB.CopyBasicFieldsFromMetaReferenceWOP(&metareferenceWOP)

		metareferenceDB_ID_atBackupTime := metareferenceDB.ID
		metareferenceDB.ID = 0
		query := backRepoMetaReference.db.Create(metareferenceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[metareferenceDB.ID] = metareferenceDB
		BackRepoMetaReferenceid_atBckpTime_newID[metareferenceDB_ID_atBackupTime] = metareferenceDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "MetaReferenceDB.json" in dirPath that stores an array
// of MetaReferenceDB and stores it in the database
// the map BackRepoMetaReferenceid_atBckpTime_newID is updated accordingly
func (backRepoMetaReference *BackRepoMetaReferenceStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMetaReferenceid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "MetaReferenceDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json MetaReference file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*MetaReferenceDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_MetaReferenceDBID_MetaReferenceDB
	for _, metareferenceDB := range forRestore {

		metareferenceDB_ID_atBackupTime := metareferenceDB.ID
		metareferenceDB.ID = 0
		query := backRepoMetaReference.db.Create(metareferenceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[metareferenceDB.ID] = metareferenceDB
		BackRepoMetaReferenceid_atBckpTime_newID[metareferenceDB_ID_atBackupTime] = metareferenceDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json MetaReference file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<MetaReference>id_atBckpTime_newID
// to compute new index
func (backRepoMetaReference *BackRepoMetaReferenceStruct) RestorePhaseTwo() {

	for _, metareferenceDB := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB {

		// next line of code is to avert unused variable compilation error
		_ = metareferenceDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoMetaReference.db.Model(metareferenceDB).Updates(*metareferenceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoMetaReference.ResetReversePointers commits all staged instances of MetaReference to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMetaReference *BackRepoMetaReferenceStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, metareference := range backRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr {
		backRepoMetaReference.ResetReversePointersInstance(backRepo, idx, metareference)
	}

	return
}

func (backRepoMetaReference *BackRepoMetaReferenceStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, metareference *models.MetaReference) (Error error) {

	// fetch matching metareferenceDB
	if metareferenceDB, ok := backRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB[idx]; ok {
		_ = metareferenceDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMetaReferenceid_atBckpTime_newID map[uint]uint
