package cloudstore

type CloudStore interface {
	func UploadVideoFile(localFilePath string)(string, error)
	func UploadAudioFile(localFilePath string)(string, error)
	func UploadImageFile(localFilePath string)(string, error)
	func UploadAppFile(localFilePath string)(string, error)
	func UploadFere(localFilePath string)(string, error)
}

func NewCloudStore(confing *conf.CloudStoreConf) CloudStore {
	
}