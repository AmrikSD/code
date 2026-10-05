// Package picker holds the selection state behind the repo picker: owners that
// expand into repos, where ticking an owner ticks every repo under it.
package picker

// Repo is one repository under an owner.
type Repo struct {
	Name     string
	SSHURL   string
	Cloned   bool
	Selected bool
}

// Owner is a GitHub user or organisation and its repositories.
type Owner struct {
	Login    string
	Repos    []Repo
	Expanded bool
	Loading  bool
	Err      error
}

// State says how much of an owner is selected.
type State int

const (
	None State = iota
	Some
	All
)

// State reports how many of the owner's repos are selected. Repos that are
// already cloned cannot be selected, so they do not count either way.
func (o *Owner) State() State {
	selected, selectable := 0, 0
	for _, r := range o.Repos {
		if r.Cloned {
			continue
		}
		selectable++
		if r.Selected {
			selected++
		}
	}
	switch {
	case selected == 0:
		return None
	case selected == selectable:
		return All
	default:
		return Some
	}
}

// SetAll selects or clears every repo that is not already cloned.
func (o *Owner) SetAll(selected bool) {
	for i := range o.Repos {
		if !o.Repos[i].Cloned {
			o.Repos[i].Selected = selected
		}
	}
}

// Count returns how many repos are selected and how many are already cloned.
func (o *Owner) Count() (selected, cloned int) {
	for _, r := range o.Repos {
		if r.Cloned {
			cloned++
		} else if r.Selected {
			selected++
		}
	}
	return selected, cloned
}

// Row is one visible line of the picker. Repo is -1 on an owner's own line.
type Row struct {
	Owner int
	Repo  int
}

// Selection is a repo chosen for cloning.
type Selection struct {
	Owner  string
	Name   string
	SSHURL string
}

// Picker is the list of owners plus the cursor position within the visible rows.
type Picker struct {
	Owners []*Owner
	Cursor int
}

// Rows lists the visible lines: every owner, and the repos of expanded owners.
func (p *Picker) Rows() []Row {
	var rows []Row
	for i, o := range p.Owners {
		rows = append(rows, Row{Owner: i, Repo: -1})
		if !o.Expanded {
			continue
		}
		for j := range o.Repos {
			rows = append(rows, Row{Owner: i, Repo: j})
		}
	}
	return rows
}

// Move shifts the cursor by delta rows, stopping at either end.
func (p *Picker) Move(delta int) {
	last := len(p.Rows()) - 1
	p.Cursor += delta
	if p.Cursor > last {
		p.Cursor = last
	}
	if p.Cursor < 0 {
		p.Cursor = 0
	}
}

// Toggle flips the row under the cursor. On an owner it selects every repo
// under it, or clears them if they were all selected already.
func (p *Picker) Toggle() {
	row, ok := p.current()
	if !ok {
		return
	}
	owner := p.Owners[row.Owner]
	if row.Repo == -1 {
		owner.SetAll(owner.State() != All)
		return
	}
	repo := &owner.Repos[row.Repo]
	if !repo.Cloned {
		repo.Selected = !repo.Selected
	}
}

// Expand shows the repos of the owner under the cursor.
func (p *Picker) Expand() {
	if row, ok := p.current(); ok {
		p.Owners[row.Owner].Expanded = true
	}
}

// Collapse hides the repos of the owner under the cursor. From a repo row it
// collapses the parent and moves the cursor up to it.
func (p *Picker) Collapse() {
	row, ok := p.current()
	if !ok {
		return
	}
	p.Owners[row.Owner].Expanded = false
	for i, r := range p.Rows() {
		if r.Owner == row.Owner {
			p.Cursor = i
			return
		}
	}
}

// SetAll selects or clears every repo of every owner.
func (p *Picker) SetAll(selected bool) {
	for _, o := range p.Owners {
		o.SetAll(selected)
	}
}

// Selected lists the chosen repos in display order.
func (p *Picker) Selected() []Selection {
	var out []Selection
	for _, o := range p.Owners {
		for _, r := range o.Repos {
			if r.Selected && !r.Cloned {
				out = append(out, Selection{Owner: o.Login, Name: r.Name, SSHURL: r.SSHURL})
			}
		}
	}
	return out
}

func (p *Picker) current() (Row, bool) {
	rows := p.Rows()
	if p.Cursor < 0 || p.Cursor >= len(rows) {
		return Row{}, false
	}
	return rows[p.Cursor], true
}
