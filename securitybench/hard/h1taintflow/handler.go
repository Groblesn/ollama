package taintflow

type SnapshotRequest struct {
	Target string
}

func Handle(req SnapshotRequest) ([]byte, error) {
	return NewService().Run(req.Target)
}
