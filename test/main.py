import flask
import json

app = flask.Flask(__name__)

@app.route('/')
def index():
    return 'Hello, World!'

@app.route('/test', methods=['POST'])
def test():
    data = flask.request.get_json()
    return json.dumps(data)

@app.route('/test', methods=['PUT'])
def test_put():
    data = flask.request.get_json()
    return json.dumps(data)

@app.route('/test', methods=['DELETE'])
def test_delete():
    return 'Deleted'


if __name__ == '__main__':
    app.run(debug=True, port=8080)