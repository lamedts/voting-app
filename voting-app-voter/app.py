from flask import Flask, render_template, request, make_response, g
import os
import socket
import random
import json
import grpc

import vote_pb2
import vote_pb2_grpc

option_a = os.getenv('OPTION_A', "Cats")
option_b = os.getenv('OPTION_B', "Dogs")
hostname = socket.gethostname()

app = Flask(__name__)

@app.route("/", methods=['POST','GET'])
def hello():
    voter_id = request.cookies.get('voter_id')
    if not voter_id:
        voter_id = random.randint(0, 1000)

    vote = None

    if request.method == 'POST':
        vote = request.form['vote']
        # redis = get_redis()
        # data = json.dumps({'voter_id': voter_id, 'vote': vote})
        # redis.rpush('votes', data)
        print()
        vote_data = vote_pb2.Vote(votedID=int(voter_id), vote=vote)
        response = stub.SetVote(vote_data)
        print(response)

    resp = make_response(render_template(
        'index.html',
        option_a=option_a,
        option_b=option_b,
        hostname=hostname,
        vote=vote,
    ))
    resp.set_cookie('voter_id', str(voter_id))
    return resp

@app.route("/reset", methods=['GET'])
def reset():
    resp = make_response('', 204)
    resp.delete_cookie('voter_id')
    return resp


if __name__ == "__main__":
    channel = grpc.insecure_channel('localhost:50051')
    stub = vote_pb2_grpc.VoteWorkerServiceStub(channel)
    app.run(host='0.0.0.0', port=8080, debug=True, threaded=True)


