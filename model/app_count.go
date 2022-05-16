package model

type UserBigdataAppCount struct {
	ClusterID   int32  `json:"clusterID"`
	ClusterName string `json:"clusterName"`
	RunningNum  int32  `json:"runningNum"`
	PendingNum  int32  `json:"pendingNum"`
	FailedNum   int32  `json:"failedNum"`
	CollectedAt int64  `json:"collectedAt"`
	Partition   string `json:"partition"`
}
