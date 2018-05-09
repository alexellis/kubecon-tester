def handle(req):
    msg = req
    if len(req) == 0:
        return "pass a value to the function"

    return str.format("Anisha said: {}", msg)
