package consumer

type WorkType int
const (
	WorkHttpRequest = 0
	WorkSaveContentToLocalStorage = 1
)
type ConsumerInterface interface {
	 CanDo(workType WorkType)bool
	 Work()
}
