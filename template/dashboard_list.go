package template

type DashboardList struct {
	Id          int      `json:"id"`
	Uid         string   `json:"uid"`
	Title       string   `json:"title`
	Url         string   `json:"url"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	FolderId    int      `json:"folderId"`
	FolderUid   string   `json:"folderUid"`
	FolderTitle string   `"folderTitle"`
	FolderUrl   string   `json:"folderUrl"`
}
