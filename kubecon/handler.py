def handle(req):
    msg = req
    if len(req) == 0:
        msg = "pass a value to the function"

    return str.format("You said: {}", msg)
