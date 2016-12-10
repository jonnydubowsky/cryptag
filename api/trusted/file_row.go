// Steven Phillips / elimisteve
// 2016.07.08

package trusted

type FileRow struct {
	FilePath  string   `json:"file_path"`
	PlainTags []string `json:"plaintags"` // types.Row.PlainTags()
}

type FileRowUpdate struct {
	FilePath     string `json:"file_path"`
	OldVersionID string `json:"old_version_id_tag"`
}
