package pb

import (
	"voting-app/voting-app-worker/datastore"
	"voting-app/voting-app-worker/utils/logger"

	"github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

var appLogger *logrus.Entry = logger.GetLogger("app")

type VoteServer struct{}

func (s *VoteServer) GetVotesResults(req *WorkerRequest, stream VoteWorkerService_GetVotesResultsServer) error {
	results := datastore.PgDBInstance.GetVoteResults()
	appLogger.Infof("%+v", results)
	for _, result := range results {
		appLogger.Infof("%+v", result)
		tmpResop := &VoteResults{
			Vote:  result.Vote,
			Count: result.Count,
		}
		if err := stream.Send(tmpResop); err != nil {
			appLogger.Errorf("%+v", err)
			return err
		}
	}
	return nil
}

func (s *VoteServer) GetVotes(req *WorkerRequest, stream VoteWorkerService_GetVotesServer) error {
	// results := datastore.PgDBInstance.GetAllVotes()
	// for result := range results {
	// 	appLogger.Infof("%+v", result)
	// 	if err := stream.Send(nil); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

// TODO: next
func (s *VoteServer) SetVote(context.Context, *Vote) (*VoteStatus, error) {
	var voteStatus VoteStatus
	voteStatus.Status = "ok"
	return &voteStatus, nil
}
