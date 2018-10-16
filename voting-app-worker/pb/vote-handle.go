package pb

import (
	"voting-app/voting-app-worker/datastore"
	types "voting-app/voting-app-worker/types/datastore"
	"voting-app/voting-app-worker/utils/logger"

	"github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

var voteLogger *logrus.Entry = logger.GetLogger("vote")

type VoteServer struct{}

func (s *VoteServer) GetVotesResults(req *WorkerRequest, stream VoteWorkerService_GetVotesResultsServer) error {
	results := datastore.PgDBInstance.GetVoteResults()
	voteLogger.Infof("%+v", results)
	for _, result := range results {
		voteLogger.Infof("%+v", result)
		tmpResop := &VoteResults{
			Vote:  result.Vote,
			Count: result.Count,
		}
		if err := stream.Send(tmpResop); err != nil {
			voteLogger.Errorf("%+v", err)
			return err
		}
	}
	return nil
}

func (s *VoteServer) GetVotes(req *WorkerRequest, stream VoteWorkerService_GetVotesServer) error {
	// results := datastore.PgDBInstance.GetAllVotes()
	// for result := range results {
	// 	voteLogger.Infof("%+v", result)
	// 	if err := stream.Send(nil); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (s *VoteServer) SetVote(ctx context.Context, vote *Vote) (*VoteStatus, error) {
	var voteStatus VoteStatus
	insertData := types.Vote{
		Vote:    vote.Vote,
		VoterID: vote.VotedID,
	}
	if err := datastore.PgDBInstance.UpsertVote(insertData); err != nil {
		voteStatus.Status = "fail"
	} else {
		voteStatus.Status = "ok"
	}
	voteLogger.Infof("SetVote %+v", voteStatus)
	return &voteStatus, nil
}

func (s *VoteServer) GetVote(ctx context.Context, vote *Vote) (*Vote, error) {
	option, err := datastore.PgDBInstance.GetVote(vote.VotedID)
	if err != nil {
		return nil, err
	}
	vote.Vote = *option
	voteLogger.Infof("GetVote %+v", vote)
	return vote, nil
}
