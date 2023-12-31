// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Classdiagram:
		ok = stage.IsStagedClassdiagram(target)

	case *DiagramPackage:
		ok = stage.IsStagedDiagramPackage(target)

	case *Field:
		ok = stage.IsStagedField(target)

	case *GongEnumShape:
		ok = stage.IsStagedGongEnumShape(target)

	case *GongEnumValueEntry:
		ok = stage.IsStagedGongEnumValueEntry(target)

	case *GongStructShape:
		ok = stage.IsStagedGongStructShape(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteShapeLink:
		ok = stage.IsStagedNoteShapeLink(target)

	case *Position:
		ok = stage.IsStagedPosition(target)

	case *UmlState:
		ok = stage.IsStagedUmlState(target)

	case *Umlsc:
		ok = stage.IsStagedUmlsc(target)

	case *Vertice:
		ok = stage.IsStagedVertice(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedClassdiagram(classdiagram *Classdiagram) (ok bool) {

		_, ok = stage.Classdiagrams[classdiagram]
	
		return
	}

	func (stage *StageStruct) IsStagedDiagramPackage(diagrampackage *DiagramPackage) (ok bool) {

		_, ok = stage.DiagramPackages[diagrampackage]
	
		return
	}

	func (stage *StageStruct) IsStagedField(field *Field) (ok bool) {

		_, ok = stage.Fields[field]
	
		return
	}

	func (stage *StageStruct) IsStagedGongEnumShape(gongenumshape *GongEnumShape) (ok bool) {

		_, ok = stage.GongEnumShapes[gongenumshape]
	
		return
	}

	func (stage *StageStruct) IsStagedGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry) (ok bool) {

		_, ok = stage.GongEnumValueEntrys[gongenumvalueentry]
	
		return
	}

	func (stage *StageStruct) IsStagedGongStructShape(gongstructshape *GongStructShape) (ok bool) {

		_, ok = stage.GongStructShapes[gongstructshape]
	
		return
	}

	func (stage *StageStruct) IsStagedLink(link *Link) (ok bool) {

		_, ok = stage.Links[link]
	
		return
	}

	func (stage *StageStruct) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

		_, ok = stage.NoteShapes[noteshape]
	
		return
	}

	func (stage *StageStruct) IsStagedNoteShapeLink(noteshapelink *NoteShapeLink) (ok bool) {

		_, ok = stage.NoteShapeLinks[noteshapelink]
	
		return
	}

	func (stage *StageStruct) IsStagedPosition(position *Position) (ok bool) {

		_, ok = stage.Positions[position]
	
		return
	}

	func (stage *StageStruct) IsStagedUmlState(umlstate *UmlState) (ok bool) {

		_, ok = stage.UmlStates[umlstate]
	
		return
	}

	func (stage *StageStruct) IsStagedUmlsc(umlsc *Umlsc) (ok bool) {

		_, ok = stage.Umlscs[umlsc]
	
		return
	}

	func (stage *StageStruct) IsStagedVertice(vertice *Vertice) (ok bool) {

		_, ok = stage.Vertices[vertice]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Classdiagram:
		stage.StageBranchClassdiagram(target)

	case *DiagramPackage:
		stage.StageBranchDiagramPackage(target)

	case *Field:
		stage.StageBranchField(target)

	case *GongEnumShape:
		stage.StageBranchGongEnumShape(target)

	case *GongEnumValueEntry:
		stage.StageBranchGongEnumValueEntry(target)

	case *GongStructShape:
		stage.StageBranchGongStructShape(target)

	case *Link:
		stage.StageBranchLink(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteShapeLink:
		stage.StageBranchNoteShapeLink(target)

	case *Position:
		stage.StageBranchPosition(target)

	case *UmlState:
		stage.StageBranchUmlState(target)

	case *Umlsc:
		stage.StageBranchUmlsc(target)

	case *Vertice:
		stage.StageBranchVertice(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchClassdiagram(classdiagram *Classdiagram) {

	// check if instance is already staged
	if IsStaged(stage, classdiagram) {
		return
	}

	classdiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		StageBranch(stage, _gongstructshape)
	}
	for _, _gongenumshape := range classdiagram.GongEnumShapes {
		StageBranch(stage, _gongenumshape)
	}
	for _, _noteshape := range classdiagram.NoteShapes {
		StageBranch(stage, _noteshape)
	}

}

func (stage *StageStruct) StageBranchDiagramPackage(diagrampackage *DiagramPackage) {

	// check if instance is already staged
	if IsStaged(stage, diagrampackage) {
		return
	}

	diagrampackage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackage.SelectedClassdiagram != nil {
		StageBranch(stage, diagrampackage.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackage.Classdiagrams {
		StageBranch(stage, _classdiagram)
	}
	for _, _umlsc := range diagrampackage.Umlscs {
		StageBranch(stage, _umlsc)
	}

}

func (stage *StageStruct) StageBranchField(field *Field) {

	// check if instance is already staged
	if IsStaged(stage, field) {
		return
	}

	field.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongEnumShape(gongenumshape *GongEnumShape) {

	// check if instance is already staged
	if IsStaged(stage, gongenumshape) {
		return
	}

	gongenumshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongenumshape.Position != nil {
		StageBranch(stage, gongenumshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueentry := range gongenumshape.GongEnumValueEntrys {
		StageBranch(stage, _gongenumvalueentry)
	}

}

func (stage *StageStruct) StageBranchGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry) {

	// check if instance is already staged
	if IsStaged(stage, gongenumvalueentry) {
		return
	}

	gongenumvalueentry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongStructShape(gongstructshape *GongStructShape) {

	// check if instance is already staged
	if IsStaged(stage, gongstructshape) {
		return
	}

	gongstructshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongstructshape.Position != nil {
		StageBranch(stage, gongstructshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _field := range gongstructshape.Fields {
		StageBranch(stage, _field)
	}
	for _, _link := range gongstructshape.Links {
		StageBranch(stage, _link)
	}

}

func (stage *StageStruct) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Middlevertice != nil {
		StageBranch(stage, link.Middlevertice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshapelink := range noteshape.NoteShapeLinks {
		StageBranch(stage, _noteshapelink)
	}

}

func (stage *StageStruct) StageBranchNoteShapeLink(noteshapelink *NoteShapeLink) {

	// check if instance is already staged
	if IsStaged(stage, noteshapelink) {
		return
	}

	noteshapelink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPosition(position *Position) {

	// check if instance is already staged
	if IsStaged(stage, position) {
		return
	}

	position.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUmlState(umlstate *UmlState) {

	// check if instance is already staged
	if IsStaged(stage, umlstate) {
		return
	}

	umlstate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUmlsc(umlsc *Umlsc) {

	// check if instance is already staged
	if IsStaged(stage, umlsc) {
		return
	}

	umlsc.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _umlstate := range umlsc.States {
		StageBranch(stage, _umlstate)
	}

}

func (stage *StageStruct) StageBranchVertice(vertice *Vertice) {

	// check if instance is already staged
	if IsStaged(stage, vertice) {
		return
	}

	vertice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Classdiagram:
		stage.UnstageBranchClassdiagram(target)

	case *DiagramPackage:
		stage.UnstageBranchDiagramPackage(target)

	case *Field:
		stage.UnstageBranchField(target)

	case *GongEnumShape:
		stage.UnstageBranchGongEnumShape(target)

	case *GongEnumValueEntry:
		stage.UnstageBranchGongEnumValueEntry(target)

	case *GongStructShape:
		stage.UnstageBranchGongStructShape(target)

	case *Link:
		stage.UnstageBranchLink(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteShapeLink:
		stage.UnstageBranchNoteShapeLink(target)

	case *Position:
		stage.UnstageBranchPosition(target)

	case *UmlState:
		stage.UnstageBranchUmlState(target)

	case *Umlsc:
		stage.UnstageBranchUmlsc(target)

	case *Vertice:
		stage.UnstageBranchVertice(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchClassdiagram(classdiagram *Classdiagram) {

	// check if instance is already staged
	if ! IsStaged(stage, classdiagram) {
		return
	}

	classdiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		UnstageBranch(stage, _gongstructshape)
	}
	for _, _gongenumshape := range classdiagram.GongEnumShapes {
		UnstageBranch(stage, _gongenumshape)
	}
	for _, _noteshape := range classdiagram.NoteShapes {
		UnstageBranch(stage, _noteshape)
	}

}

func (stage *StageStruct) UnstageBranchDiagramPackage(diagrampackage *DiagramPackage) {

	// check if instance is already staged
	if ! IsStaged(stage, diagrampackage) {
		return
	}

	diagrampackage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackage.SelectedClassdiagram != nil {
		UnstageBranch(stage, diagrampackage.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackage.Classdiagrams {
		UnstageBranch(stage, _classdiagram)
	}
	for _, _umlsc := range diagrampackage.Umlscs {
		UnstageBranch(stage, _umlsc)
	}

}

func (stage *StageStruct) UnstageBranchField(field *Field) {

	// check if instance is already staged
	if ! IsStaged(stage, field) {
		return
	}

	field.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongEnumShape(gongenumshape *GongEnumShape) {

	// check if instance is already staged
	if ! IsStaged(stage, gongenumshape) {
		return
	}

	gongenumshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongenumshape.Position != nil {
		UnstageBranch(stage, gongenumshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueentry := range gongenumshape.GongEnumValueEntrys {
		UnstageBranch(stage, _gongenumvalueentry)
	}

}

func (stage *StageStruct) UnstageBranchGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry) {

	// check if instance is already staged
	if ! IsStaged(stage, gongenumvalueentry) {
		return
	}

	gongenumvalueentry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongStructShape(gongstructshape *GongStructShape) {

	// check if instance is already staged
	if ! IsStaged(stage, gongstructshape) {
		return
	}

	gongstructshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongstructshape.Position != nil {
		UnstageBranch(stage, gongstructshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _field := range gongstructshape.Fields {
		UnstageBranch(stage, _field)
	}
	for _, _link := range gongstructshape.Links {
		UnstageBranch(stage, _link)
	}

}

func (stage *StageStruct) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if ! IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Middlevertice != nil {
		UnstageBranch(stage, link.Middlevertice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if ! IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshapelink := range noteshape.NoteShapeLinks {
		UnstageBranch(stage, _noteshapelink)
	}

}

func (stage *StageStruct) UnstageBranchNoteShapeLink(noteshapelink *NoteShapeLink) {

	// check if instance is already staged
	if ! IsStaged(stage, noteshapelink) {
		return
	}

	noteshapelink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPosition(position *Position) {

	// check if instance is already staged
	if ! IsStaged(stage, position) {
		return
	}

	position.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUmlState(umlstate *UmlState) {

	// check if instance is already staged
	if ! IsStaged(stage, umlstate) {
		return
	}

	umlstate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUmlsc(umlsc *Umlsc) {

	// check if instance is already staged
	if ! IsStaged(stage, umlsc) {
		return
	}

	umlsc.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _umlstate := range umlsc.States {
		UnstageBranch(stage, _umlstate)
	}

}

func (stage *StageStruct) UnstageBranchVertice(vertice *Vertice) {

	// check if instance is already staged
	if ! IsStaged(stage, vertice) {
		return
	}

	vertice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

