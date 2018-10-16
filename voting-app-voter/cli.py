import grpc
# from . import vote_pb2
# from . import vote_pb2_grpc
import vote_pb2
import vote_pb2_grpc
channel = grpc.insecure_channel('localhost:50051')
stub = vote_pb2_grpc.VoteWorkerServiceStub(channel)

# to vote
vote_data = vote_pb2.Vote(votedID=54, vote="dog")
response = stub.SetVote(vote_data)
print(response.status)

# to get vote
worker_request = vote_pb2.Vote(votedID=54)
response = stub.GetVote(vote_data)
print(response)

#to get result
worker_request = vote_pb2.WorkerRequest(query="")
for response in stub.GetVotesResults(worker_request):
    print(response)
