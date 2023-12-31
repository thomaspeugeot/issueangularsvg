// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Animate:
		ok = stage.IsStagedAnimate(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *Ellipse:
		ok = stage.IsStagedEllipse(target)

	case *Layer:
		ok = stage.IsStagedLayer(target)

	case *Line:
		ok = stage.IsStagedLine(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *LinkAnchoredText:
		ok = stage.IsStagedLinkAnchoredText(target)

	case *Path:
		ok = stage.IsStagedPath(target)

	case *Point:
		ok = stage.IsStagedPoint(target)

	case *Polygone:
		ok = stage.IsStagedPolygone(target)

	case *Polyline:
		ok = stage.IsStagedPolyline(target)

	case *Rect:
		ok = stage.IsStagedRect(target)

	case *RectAnchoredRect:
		ok = stage.IsStagedRectAnchoredRect(target)

	case *RectAnchoredText:
		ok = stage.IsStagedRectAnchoredText(target)

	case *RectLinkLink:
		ok = stage.IsStagedRectLinkLink(target)

	case *SVG:
		ok = stage.IsStagedSVG(target)

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedAnimate(animate *Animate) (ok bool) {

		_, ok = stage.Animates[animate]
	
		return
	}

	func (stage *StageStruct) IsStagedCircle(circle *Circle) (ok bool) {

		_, ok = stage.Circles[circle]
	
		return
	}

	func (stage *StageStruct) IsStagedEllipse(ellipse *Ellipse) (ok bool) {

		_, ok = stage.Ellipses[ellipse]
	
		return
	}

	func (stage *StageStruct) IsStagedLayer(layer *Layer) (ok bool) {

		_, ok = stage.Layers[layer]
	
		return
	}

	func (stage *StageStruct) IsStagedLine(line *Line) (ok bool) {

		_, ok = stage.Lines[line]
	
		return
	}

	func (stage *StageStruct) IsStagedLink(link *Link) (ok bool) {

		_, ok = stage.Links[link]
	
		return
	}

	func (stage *StageStruct) IsStagedLinkAnchoredText(linkanchoredtext *LinkAnchoredText) (ok bool) {

		_, ok = stage.LinkAnchoredTexts[linkanchoredtext]
	
		return
	}

	func (stage *StageStruct) IsStagedPath(path *Path) (ok bool) {

		_, ok = stage.Paths[path]
	
		return
	}

	func (stage *StageStruct) IsStagedPoint(point *Point) (ok bool) {

		_, ok = stage.Points[point]
	
		return
	}

	func (stage *StageStruct) IsStagedPolygone(polygone *Polygone) (ok bool) {

		_, ok = stage.Polygones[polygone]
	
		return
	}

	func (stage *StageStruct) IsStagedPolyline(polyline *Polyline) (ok bool) {

		_, ok = stage.Polylines[polyline]
	
		return
	}

	func (stage *StageStruct) IsStagedRect(rect *Rect) (ok bool) {

		_, ok = stage.Rects[rect]
	
		return
	}

	func (stage *StageStruct) IsStagedRectAnchoredRect(rectanchoredrect *RectAnchoredRect) (ok bool) {

		_, ok = stage.RectAnchoredRects[rectanchoredrect]
	
		return
	}

	func (stage *StageStruct) IsStagedRectAnchoredText(rectanchoredtext *RectAnchoredText) (ok bool) {

		_, ok = stage.RectAnchoredTexts[rectanchoredtext]
	
		return
	}

	func (stage *StageStruct) IsStagedRectLinkLink(rectlinklink *RectLinkLink) (ok bool) {

		_, ok = stage.RectLinkLinks[rectlinklink]
	
		return
	}

	func (stage *StageStruct) IsStagedSVG(svg *SVG) (ok bool) {

		_, ok = stage.SVGs[svg]
	
		return
	}

	func (stage *StageStruct) IsStagedText(text *Text) (ok bool) {

		_, ok = stage.Texts[text]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Animate:
		stage.StageBranchAnimate(target)

	case *Circle:
		stage.StageBranchCircle(target)

	case *Ellipse:
		stage.StageBranchEllipse(target)

	case *Layer:
		stage.StageBranchLayer(target)

	case *Line:
		stage.StageBranchLine(target)

	case *Link:
		stage.StageBranchLink(target)

	case *LinkAnchoredText:
		stage.StageBranchLinkAnchoredText(target)

	case *Path:
		stage.StageBranchPath(target)

	case *Point:
		stage.StageBranchPoint(target)

	case *Polygone:
		stage.StageBranchPolygone(target)

	case *Polyline:
		stage.StageBranchPolyline(target)

	case *Rect:
		stage.StageBranchRect(target)

	case *RectAnchoredRect:
		stage.StageBranchRectAnchoredRect(target)

	case *RectAnchoredText:
		stage.StageBranchRectAnchoredText(target)

	case *RectLinkLink:
		stage.StageBranchRectLinkLink(target)

	case *SVG:
		stage.StageBranchSVG(target)

	case *Text:
		stage.StageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if IsStaged(stage, animate) {
		return
	}

	animate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if IsStaged(stage, circle) {
		return
	}

	circle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circle.Animations {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchEllipse(ellipse *Ellipse) {

	// check if instance is already staged
	if IsStaged(stage, ellipse) {
		return
	}

	ellipse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipse.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchLayer(layer *Layer) {

	// check if instance is already staged
	if IsStaged(stage, layer) {
		return
	}

	layer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layer.Rects {
		StageBranch(stage, _rect)
	}
	for _, _text := range layer.Texts {
		StageBranch(stage, _text)
	}
	for _, _circle := range layer.Circles {
		StageBranch(stage, _circle)
	}
	for _, _line := range layer.Lines {
		StageBranch(stage, _line)
	}
	for _, _ellipse := range layer.Ellipses {
		StageBranch(stage, _ellipse)
	}
	for _, _polyline := range layer.Polylines {
		StageBranch(stage, _polyline)
	}
	for _, _polygone := range layer.Polygones {
		StageBranch(stage, _polygone)
	}
	for _, _path := range layer.Paths {
		StageBranch(stage, _path)
	}
	for _, _link := range layer.Links {
		StageBranch(stage, _link)
	}
	for _, _rectlinklink := range layer.RectLinkLinks {
		StageBranch(stage, _rectlinklink)
	}

}

func (stage *StageStruct) StageBranchLine(line *Line) {

	// check if instance is already staged
	if IsStaged(stage, line) {
		return
	}

	line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range line.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Start != nil {
		StageBranch(stage, link.Start)
	}
	if link.End != nil {
		StageBranch(stage, link.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		StageBranch(stage, _linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		StageBranch(stage, _linkanchoredtext)
	}
	for _, _point := range link.ControlPoints {
		StageBranch(stage, _point)
	}

}

func (stage *StageStruct) StageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

	// check if instance is already staged
	if IsStaged(stage, linkanchoredtext) {
		return
	}

	linkanchoredtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtext.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchPath(path *Path) {

	// check if instance is already staged
	if IsStaged(stage, path) {
		return
	}

	path.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range path.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchPoint(point *Point) {

	// check if instance is already staged
	if IsStaged(stage, point) {
		return
	}

	point.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPolygone(polygone *Polygone) {

	// check if instance is already staged
	if IsStaged(stage, polygone) {
		return
	}

	polygone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygone.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchPolyline(polyline *Polyline) {

	// check if instance is already staged
	if IsStaged(stage, polyline) {
		return
	}

	polyline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polyline.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchRect(rect *Rect) {

	// check if instance is already staged
	if IsStaged(stage, rect) {
		return
	}

	rect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rect.Animations {
		StageBranch(stage, _animate)
	}
	for _, _rectanchoredtext := range rect.RectAnchoredTexts {
		StageBranch(stage, _rectanchoredtext)
	}
	for _, _rectanchoredrect := range rect.RectAnchoredRects {
		StageBranch(stage, _rectanchoredrect)
	}

}

func (stage *StageStruct) StageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredrect) {
		return
	}

	rectanchoredrect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredtext) {
		return
	}

	rectanchoredtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtext.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *StageStruct) StageBranchRectLinkLink(rectlinklink *RectLinkLink) {

	// check if instance is already staged
	if IsStaged(stage, rectlinklink) {
		return
	}

	rectlinklink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklink.Start != nil {
		StageBranch(stage, rectlinklink.Start)
	}
	if rectlinklink.End != nil {
		StageBranch(stage, rectlinklink.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSVG(svg *SVG) {

	// check if instance is already staged
	if IsStaged(stage, svg) {
		return
	}

	svg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if svg.StartRect != nil {
		StageBranch(stage, svg.StartRect)
	}
	if svg.EndRect != nil {
		StageBranch(stage, svg.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svg.Layers {
		StageBranch(stage, _layer)
	}

}

func (stage *StageStruct) StageBranchText(text *Text) {

	// check if instance is already staged
	if IsStaged(stage, text) {
		return
	}

	text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range text.Animates {
		StageBranch(stage, _animate)
	}

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Animate:
		stage.UnstageBranchAnimate(target)

	case *Circle:
		stage.UnstageBranchCircle(target)

	case *Ellipse:
		stage.UnstageBranchEllipse(target)

	case *Layer:
		stage.UnstageBranchLayer(target)

	case *Line:
		stage.UnstageBranchLine(target)

	case *Link:
		stage.UnstageBranchLink(target)

	case *LinkAnchoredText:
		stage.UnstageBranchLinkAnchoredText(target)

	case *Path:
		stage.UnstageBranchPath(target)

	case *Point:
		stage.UnstageBranchPoint(target)

	case *Polygone:
		stage.UnstageBranchPolygone(target)

	case *Polyline:
		stage.UnstageBranchPolyline(target)

	case *Rect:
		stage.UnstageBranchRect(target)

	case *RectAnchoredRect:
		stage.UnstageBranchRectAnchoredRect(target)

	case *RectAnchoredText:
		stage.UnstageBranchRectAnchoredText(target)

	case *RectLinkLink:
		stage.UnstageBranchRectLinkLink(target)

	case *SVG:
		stage.UnstageBranchSVG(target)

	case *Text:
		stage.UnstageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if ! IsStaged(stage, animate) {
		return
	}

	animate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if ! IsStaged(stage, circle) {
		return
	}

	circle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circle.Animations {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchEllipse(ellipse *Ellipse) {

	// check if instance is already staged
	if ! IsStaged(stage, ellipse) {
		return
	}

	ellipse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipse.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchLayer(layer *Layer) {

	// check if instance is already staged
	if ! IsStaged(stage, layer) {
		return
	}

	layer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layer.Rects {
		UnstageBranch(stage, _rect)
	}
	for _, _text := range layer.Texts {
		UnstageBranch(stage, _text)
	}
	for _, _circle := range layer.Circles {
		UnstageBranch(stage, _circle)
	}
	for _, _line := range layer.Lines {
		UnstageBranch(stage, _line)
	}
	for _, _ellipse := range layer.Ellipses {
		UnstageBranch(stage, _ellipse)
	}
	for _, _polyline := range layer.Polylines {
		UnstageBranch(stage, _polyline)
	}
	for _, _polygone := range layer.Polygones {
		UnstageBranch(stage, _polygone)
	}
	for _, _path := range layer.Paths {
		UnstageBranch(stage, _path)
	}
	for _, _link := range layer.Links {
		UnstageBranch(stage, _link)
	}
	for _, _rectlinklink := range layer.RectLinkLinks {
		UnstageBranch(stage, _rectlinklink)
	}

}

func (stage *StageStruct) UnstageBranchLine(line *Line) {

	// check if instance is already staged
	if ! IsStaged(stage, line) {
		return
	}

	line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range line.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if ! IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Start != nil {
		UnstageBranch(stage, link.Start)
	}
	if link.End != nil {
		UnstageBranch(stage, link.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		UnstageBranch(stage, _linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		UnstageBranch(stage, _linkanchoredtext)
	}
	for _, _point := range link.ControlPoints {
		UnstageBranch(stage, _point)
	}

}

func (stage *StageStruct) UnstageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

	// check if instance is already staged
	if ! IsStaged(stage, linkanchoredtext) {
		return
	}

	linkanchoredtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtext.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchPath(path *Path) {

	// check if instance is already staged
	if ! IsStaged(stage, path) {
		return
	}

	path.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range path.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchPoint(point *Point) {

	// check if instance is already staged
	if ! IsStaged(stage, point) {
		return
	}

	point.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPolygone(polygone *Polygone) {

	// check if instance is already staged
	if ! IsStaged(stage, polygone) {
		return
	}

	polygone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygone.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchPolyline(polyline *Polyline) {

	// check if instance is already staged
	if ! IsStaged(stage, polyline) {
		return
	}

	polyline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polyline.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchRect(rect *Rect) {

	// check if instance is already staged
	if ! IsStaged(stage, rect) {
		return
	}

	rect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rect.Animations {
		UnstageBranch(stage, _animate)
	}
	for _, _rectanchoredtext := range rect.RectAnchoredTexts {
		UnstageBranch(stage, _rectanchoredtext)
	}
	for _, _rectanchoredrect := range rect.RectAnchoredRects {
		UnstageBranch(stage, _rectanchoredrect)
	}

}

func (stage *StageStruct) UnstageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if ! IsStaged(stage, rectanchoredrect) {
		return
	}

	rectanchoredrect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

	// check if instance is already staged
	if ! IsStaged(stage, rectanchoredtext) {
		return
	}

	rectanchoredtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtext.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *StageStruct) UnstageBranchRectLinkLink(rectlinklink *RectLinkLink) {

	// check if instance is already staged
	if ! IsStaged(stage, rectlinklink) {
		return
	}

	rectlinklink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklink.Start != nil {
		UnstageBranch(stage, rectlinklink.Start)
	}
	if rectlinklink.End != nil {
		UnstageBranch(stage, rectlinklink.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSVG(svg *SVG) {

	// check if instance is already staged
	if ! IsStaged(stage, svg) {
		return
	}

	svg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if svg.StartRect != nil {
		UnstageBranch(stage, svg.StartRect)
	}
	if svg.EndRect != nil {
		UnstageBranch(stage, svg.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svg.Layers {
		UnstageBranch(stage, _layer)
	}

}

func (stage *StageStruct) UnstageBranchText(text *Text) {

	// check if instance is already staged
	if ! IsStaged(stage, text) {
		return
	}

	text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range text.Animates {
		UnstageBranch(stage, _animate)
	}

}

