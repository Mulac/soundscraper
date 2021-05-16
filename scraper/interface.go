package scraper

type DownloadManager interface {
	Download(videoID string) error
}

type DownloadFactory interface {
	SetType(dType DownloadType) DownloadFactory
	New() (DownloadManager, error)
}
