	"github.com/aymanbagabas/git-module"
	if f.OldIndex != ZeroID {
	if f.Index != ZeroID {
			fmt.Sprintf("index %s..%s", ZeroID, to.Hash()),
			fmt.Sprintf("index %s..%s", from.Hash(), ZeroID),

func toDiff(ddiff *git.Diff) *Diff {
	files := make([]*DiffFile, 0, len(ddiff.Files))
	for _, df := range ddiff.Files {
		sections := make([]*DiffSection, 0, len(df.Sections))
		for _, ds := range df.Sections {
			sections = append(sections, &DiffSection{
				DiffSection: ds,
			})
		}
		files = append(files, &DiffFile{
			DiffFile: df,
			Sections: sections,
		})
	}
	diff := &Diff{
		Diff:  ddiff,
		Files: files,
	}
	return diff
}